package collection_usecase

import (
	"context"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	collection_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/collection"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *CollectionUsecaseImplementation) RequestImportPersonalCollectionFromLudopedia(ctx context.Context, i *RequestImportPersonalCollectionFromLudopediaInput) error {
	logger := utils.GetLoggerFromCtx(ctx, self.Logger)

	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return err
	}

	connection, err := self.AccountRepository.GetConnection(ctx, &account_repository.GetConnectionInput{
		ExternalId: i.ExternalId,
		Provider:   models.Provider_Ludopedia,
	})
	if err != nil {
		if err.Error() == "not found" {
			logger.Trace().
				Int("AccountId", i.AccountId).
				Str("ExternalId", i.ExternalId).
				Any("Provider", models.Provider_Ludopedia).
				Msg("connection not found on database")
			return fmt.Errorf("connection not found")
		}

		return err
	}
	if connection.AccountId != i.AccountId {
		logger.Trace().
			Int("connection.AccountId", connection.AccountId).
			Int("i.AccountId", i.AccountId).
			Str("ExternalId", i.ExternalId).
			Any("Provider", models.Provider_Ludopedia).
			Msg("connection have a different account id")
		return fmt.Errorf("connection not found")
	}

	ongoingImport, err := self.CollectionRepository.GetOngoingImportCollectionLog(ctx, &collection_repository.GetOngoingImportCollectionLogInput{
		ExternalIds: []string{connection.ExternalId},
		Provider:    models.Provider_Ludopedia,
	})
	if err != nil && err.Error() != "not found" {
		return err
	}
	if ongoingImport != nil && len(ongoingImport.Data) > 0 {
		logger.Trace().
			Msg("import already in progress")
		return fmt.Errorf("import already in progress")
	}

	err = self.MessagingAdapter.SendPrivateEvent(&adapters.SendEventInput{
		ListenerId: self.SecretsAdapter.CollectionImportPersonalCollectionFromLudopediaQueueId,
		EventName:  "import-collection-from-ludopedia",
		Event: models.ImportCollectionEvent{
			AccountId:  i.AccountId,
			ExternalId: connection.ExternalId,
			Trigger:    models.CollectionImportTrigger_ManualByUser,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
