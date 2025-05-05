package account_repository

import (
	"context"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

func (self *accountRepositoryImplementation) GetAccountDataBySessionId(ctx context.Context, i *GetAccountDataBySessionIdInput) (*models.AccountData, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	row, err := db.GetAccountDataBySession(ctx, i.SessionId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("not found")
		}

		return nil, err
	}

	return &models.AccountData{
		AccountId: int(row.ID.Int32),
		IsAdmin:   row.IsAdmin.Bool,
	}, nil
}
