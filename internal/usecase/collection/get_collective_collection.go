package collection_usecase

import (
	"context"

	collection_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/collection"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *CollectionUsecaseImplementation) GetCollectiveCollection(ctx context.Context, i *GetCollectiveCollectionInput) (*GetCollectiveCollectionOutput, error) {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return nil, err
	}

	collection, err := self.CollectionRepository.GetCollectiveCollection(ctx, &collection_repository.GetCollectiveCollectionInput{
		AccountId:          i.AccountId,
		GameName:           i.GameName,
		Kind:               i.Kind,
		MaxAmountOfPlayers: i.MaxAmountOfPlayers,
		Pagination:         i.Pagination,
	})
	if err != nil {
		self.Logger.Warn().Err(err).Msg("fail to get collective collection")
		return nil, err
	}

	return &GetCollectiveCollectionOutput{
		Data:       collection.Data,
		Pagination: collection.Pagination,
	}, nil
}
