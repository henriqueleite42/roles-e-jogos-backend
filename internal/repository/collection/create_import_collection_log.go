package collection_repository

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
)

func (self *collectionRepositoryImplementation) CreateImportCollectionLog(ctx context.Context, i *CreateImportCollectionLogInput) (*CreateImportCollectionLogOutput, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	id, err := db.CreateImportCollectionLog(ctx, queries.CreateImportCollectionLogParams{
		AccountID:  int32(i.AccountId),
		ExternalID: i.ExternalId,
		Trigger:    queries.CollectionImportTriggerEnum(i.Trigger),
		Provider:   queries.ProviderEnum(i.Provider),
		Status:     queries.CollectionImportStatusEnum(i.Status),
	})
	if err != nil {
		return nil, err
	}

	return &CreateImportCollectionLogOutput{
		Id: int(id),
	}, nil
}
