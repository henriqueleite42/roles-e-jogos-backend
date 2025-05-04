package account_repository

import (
	"context"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

func (self *accountRepositoryImplementation) GetAccountDataById(ctx context.Context, i *GetAccountDataByIdInput) (*models.AccountData, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	row, err := db.GetAccountDataById(ctx, int32(i.AccountId))
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
