package account_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5"
)

func (self *accountRepositoryImplementation) EditAccountHandle(ctx context.Context, i *EditAccountHandleInput) error {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return err
	}

	err = db.UpdateAccountHandle(ctx, queries.UpdateAccountHandleParams{
		ID:     int32(i.AccountId),
		Handle: i.Handle,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("not found")
		}
		if strings.Contains(err.Error(), "violates unique constraint") {
			return fmt.Errorf("conflict")
		}

		return err
	}

	return nil
}
