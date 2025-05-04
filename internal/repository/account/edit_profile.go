package account_repository

import (
	"context"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (self *accountRepositoryImplementation) EditProfile(ctx context.Context, i *EditProfileInput) error {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return err
	}

	var name pgtype.Text
	if i.Name != nil {
		name = pgtype.Text{
			Valid:  true,
			String: *i.Name,
		}
	}

	err = db.UpdateAccountProfile(ctx, queries.UpdateAccountProfileParams{
		ID:   int32(i.AccountId),
		Name: name,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("not found")
		}

		return err
	}

	return nil
}
