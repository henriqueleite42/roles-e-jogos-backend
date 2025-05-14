package account_usecase

import (
	"context"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *AccountUsecaseImplementation) LinkLudopediaProvider(ctx context.Context, i *LinkLudopediaProviderInput) error {
	logger := utils.GetLoggerFromCtx(ctx, self.Logger)

	exchangeResult, err := self.LudopediaAdapter.ExchangeCode(&adapters.ExchangeCodeInput{
		Code: i.Code,
	})
	if err != nil {
		return err
	}

	err = self.LudopediaAdapter.CheckRequiredScopes(exchangeResult.Scopes)
	if err != nil {
		return err
	}

	externalUserData, err := self.LudopediaAdapter.GetUserData(exchangeResult.AccessToken)
	if err != nil {
		return err
	}

	tx, ctx, err := utils.SetTxInCtx(self.Db, ctx)
	if err != nil {
		return err
	}

	connection, err := self.AccountRepository.GetConnection(ctx, &account_repository.GetConnectionInput{
		ExternalId: externalUserData.Id,
		Provider:   models.Provider_Ludopedia,
	})
	if err != nil && err.Error() != "not found" {
		tx.Rollback(ctx)
		return err
	}

	if connection != nil {
		if connection.AccountId == i.AccountId {
			tx.Commit(ctx)
			return nil
		} else {
			tx.Rollback(ctx)
			return fmt.Errorf("connection already linked with another account")
		}
	}

	logger.Trace().Msg("try to link connection with account")
	err = self.AccountRepository.LinkConnectionWithAccount(ctx, &account_repository.LinkConnectionWithAccountInput{
		AccountId:      i.AccountId,
		Provider:       models.Provider_Ludopedia,
		ExternalHandle: externalUserData.Handle,
		ExternalId:     externalUserData.Id,
		AccessToken:    &exchangeResult.AccessToken,
	})
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	logger.Trace().Msg("account successfully linked with connection")

	logger.Trace().Msg("try to send queue message to import ludopedia collection")
	err = self.MessagingAdapter.SendPrivateEvent(&adapters.SendEventInput{
		ListenerId: self.SecretsAdapter.CollectionImportPersonalCollectionFromLudopediaQueueId,
		EventName:  "import-collection-from-ludopedia",
		Event: models.ImportCollectionEvent{
			AccountId:  i.AccountId,
			ExternalId: externalUserData.Id,
			Trigger:    models.CollectionImportTrigger_AccountLink,
		},
	})
	if err != nil {
		logger.Error().Err(err).Msg("fail to send queue message to import collection from ludopedia")
	} else {
		logger.Trace().Msg("queue message to import ludopedia collection successfully sent")
	}

	tx.Commit(ctx)
	return nil
}
