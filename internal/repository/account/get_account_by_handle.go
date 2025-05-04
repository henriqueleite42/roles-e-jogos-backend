package account_repository

import (
	"context"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

func (self *accountRepositoryImplementation) GetAccountByHandle(ctx context.Context, i *GetAccountByHandleInput) (*models.Account, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	row, err := db.GetAccountByHandle(ctx, i.Handle)
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
