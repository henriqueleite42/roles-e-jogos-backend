package account_repository

import (
	"context"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5"
)

func (self *accountRepositoryImplementation) GetAccountDataByConnection(ctx context.Context, i *GetAccountDataByConnectionInput) (*models.AccountData, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	row, err := db.GetAccountDataByConnection(ctx, queries.GetAccountDataByConnectionParams{
		Provider:   queries.ProviderEnum(i.Provider),
		ExternalID: i.ExternalId,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("not found")
		}

		return nil, err
	}

	return &models.AccountData{
		AccountId: int(row.AccountID),
		IsAdmin:   row.IsAdmin.Bool,
	}, nil
}
