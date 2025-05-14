package collection_usecase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	collection_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/collection"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *CollectionUsecaseImplementation) RequestImportPersonalCollectionFromLudopedia(ctx context.Context, i *RequestImportPersonalCollectionFromLudopediaInput) error {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return err
	}

	connection, err := self.AccountRepository.GetConnection(ctx, &account_repository.GetConnectionInput{
		ExternalId: strconv.Itoa(i.LudopediaId),
		Provider:   models.Provider_Ludopedia,
	})
	if err != nil {
		if err.Error() == "not found" {
			return fmt.Errorf("connection not found")
		}

		return err
	}
	if connection.AccountId != i.AccountId {
		return fmt.Errorf("connection not found")
	}

	ongoingImport, err := self.CollectionRepository.GetOngoingImportCollectionLog(ctx, &collection_repository.GetOngoingImportCollectionLogInput{
		ExternalIds: []string{connection.ExternalId},
		Provider:    models.Provider_Ludopedia,
	})
	if err != nil {
		if err.Error() == "not found" {
			return fmt.Errorf("import already in progress")
		}

		return err
	}
	if ongoingImport != nil && len(ongoingImport.Data) > 0 {
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
