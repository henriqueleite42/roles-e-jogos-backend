package account_repository

import (
	"context"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5"
)

func (self *accountRepositoryImplementation) GetConnection(ctx context.Context, i *GetConnectionInput) (*models.Connection, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	row, err := db.GetConnection(ctx, queries.GetConnectionParams{
		ExternalID: i.ExternalId,
		Provider:   queries.ProviderEnum(i.Provider),
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("not found")
		}

		return nil, err
	}

	var externalHandle *string
	if row.ExternalHandle.Valid {
		externalHandle = &row.ExternalHandle.String
	}

	var refreshToken *string
	if row.RefreshToken.Valid {
		refreshToken = &row.RefreshToken.String
	}

	return &models.Connection{
		AccountId:      int(row.AccountID),
		CreatedAt:      row.CreatedAt.Time,
		ExternalHandle: externalHandle,
		ExternalId:     row.ExternalID,
		Provider:       models.Provider(row.Provider),
		RefreshToken:   refreshToken,
	}, nil
}
