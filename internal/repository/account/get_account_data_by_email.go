package account_repository

import (
	"context"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

func (self *accountRepositoryImplementation) GetAccountDataByEmail(ctx context.Context, i *GetAccountDataByEmailInput) (*models.AccountData, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	row, err := db.GetAccountDataByEmail(ctx, i.Email)
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
