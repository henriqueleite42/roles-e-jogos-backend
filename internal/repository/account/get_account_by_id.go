package account_repository

import (
	"context"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

func (self *accountRepositoryImplementation) GetAccountById(ctx context.Context, i *GetAccountByIdInput) (*models.Account, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	row, err := db.GetAccountById(ctx, int32(i.AccountId))
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("not found")
		}

		return nil, err
	}

	return &models.Account{
		Id:         int(row.ID),
		IsAdmin:    row.IsAdmin,
		AvatarPath: &row.AvatarPath.String,
		CreatedAt:  row.CreatedAt.Time,
		Handle:     row.Handle,
		Name:       &row.Name.String,
	}, nil
}
