package collection_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5"
)

func (self *collectionRepositoryImplementation) GetLatestImportCollectionLogStatus(ctx context.Context, i *GetLatestImportCollectionLogStatusInput) (*models.ImportCollectionLog, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	row, err := db.GetLatestImportCollectionLog(ctx, queries.GetLatestImportCollectionLogParams{
		AccountID:  int32(i.AccountId),
		ExternalID: i.ExternalId,
		Provider:   queries.ProviderEnum(i.Provider),
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("not found")
		}

		return nil, err
	}

	var endedAt *time.Time
	if row.EndedAt.Valid {
		endedAt = &row.EndedAt.Time
	}

	return &models.ImportCollectionLog{
		AccountId:  int(row.AccountID),
		CreatedAt:  row.CreatedAt.Time,
		EndedAt:    endedAt,
		ExternalId: row.ExternalID,
		Id:         int(row.ID),
		Provider:   models.Provider(row.Provider),
		Status:     models.CollectionImportStatus(row.Status),
		Trigger:    models.CollectionImportTrigger(row.Trigger),
	}, nil
}
