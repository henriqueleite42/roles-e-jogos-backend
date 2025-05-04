package account_repository

import (
	"context"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

func (self *accountRepositoryImplementation) GetAccountDataByHandle(ctx context.Context, i *GetAccountDataByHandleInput) (*models.AccountData, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	row, err := db.GetAccountDataByHandle(ctx, i.Handle)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("not found")
		}

		return nil, err
	}

	return &models.AccountData{
		AccountId: int(row.ID),
		IsAdmin:   row.IsAdmin,
	}, nil
}
