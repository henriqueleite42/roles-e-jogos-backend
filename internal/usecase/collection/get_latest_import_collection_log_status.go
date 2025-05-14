package collection_usecase

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	collection_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/collection"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *CollectionUsecaseImplementation) GetLatestImportCollectionLogStatus(ctx context.Context, i *GetLatestImportCollectionLogStatusInput) (*GetLatestImportCollectionLogStatusOutput, error) {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return nil, err
	}

	collection, err := self.CollectionRepository.GetLatestImportCollectionLogStatus(ctx, &collection_repository.GetLatestImportCollectionLogStatusInput{
		AccountId:  i.AccountId,
		ExternalId: i.ExternalId,
		Provider:   i.Provider,
	})
	if err != nil {
		if err.Error() == "not found" {
			return &GetLatestImportCollectionLogStatusOutput{
				Status: models.CollectionImportStatus_NotYetStarted,
			}, nil
		}
		self.Logger.Warn().Err(err).Msg("fail to get specific ongoing import collection log")
		return nil, err
	}

	return &GetLatestImportCollectionLogStatusOutput{
		Status: collection.Status,
	}, nil
}
