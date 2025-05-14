package collection_repository

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
)

func (self *collectionRepositoryImplementation) GetOngoingImportCollectionLog(ctx context.Context, i *GetOngoingImportCollectionLogInput) (*GetOngoingImportCollectionLogOutput, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	result, err := db.GetOngoingImportCollectionLog(ctx, queries.GetOngoingImportCollectionLogParams{
		Column1:  i.ExternalIds,
		Provider: queries.ProviderEnum(i.Provider),
	})
	if err != nil {
		return nil, err
	}

	importLogs := make([]*models.ImportCollectionLog, len(result))
	for k, v := range result {
		importLogs[k] = &models.ImportCollectionLog{
			AccountId:  int(v.AccountID),
			CreatedAt:  v.CreatedAt.Time,
			ExternalId: v.ExternalID,
			Id:         int(v.ID),
			Provider:   models.Provider(v.Provider),
			Status:     models.CollectionImportStatus(v.Status),
			Trigger:    models.CollectionImportTrigger(v.Trigger),
		}
	}

	return &GetOngoingImportCollectionLogOutput{
		Data: importLogs,
	}, nil
}
