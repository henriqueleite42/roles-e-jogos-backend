package account_repository

import (
	"context"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

func (self *accountRepositoryImplementation) GetAccountDataBySessionId(ctx context.Context, i *GetAccountDataBySessionIdInput) (*models.AccountDataDb, error) {
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

	var avatarPath *string
	if row.AvatarPath.Valid {
		avatarPath = &row.AvatarPath.String
	}

	return &models.AccountDataDb{
		AccountId:  int(row.ID.Int32),
		IsAdmin:    row.IsAdmin.Bool,
		AvatarPath: avatarPath,
		Handle:     row.Handle.String,
	}, nil
}
