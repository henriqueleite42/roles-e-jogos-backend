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
	exchangeResult, err := self.GoogleAdapter.ExchangeCode(&adapters.ExchangeCodeInput{
		Code: i.Code,
	})
	if err != nil {
		return err
	}

	err = self.GoogleAdapter.CheckRequiredScopes(exchangeResult.Scopes)
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

	err = self.AccountRepository.LinkConnectionWithAccount(ctx, &account_repository.LinkConnectionWithAccountInput{
		AccountId:      i.AccountId,
		Provider:       models.Provider_Ludopedia,
		ExternalHandle: externalUserData.Handle,
		ExternalId:     externalUserData.Id,
	})
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	tx.Commit(ctx)
	return nil
}
