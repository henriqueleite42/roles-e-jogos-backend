package collection_repository

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5/pgtype"
)

func (self *collectionRepositoryImplementation) AddToPersonalCollection(ctx context.Context, i *AddToPersonalCollectionInput) error {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return err
	}

	var paid pgtype.Int4
	if i.Paid != nil {
		paid = pgtype.Int4{
			Valid: true,
			Int32: int32(*i.Paid),
		}
	}
	var acquiredAt pgtype.Timestamptz
	if i.AcquiredAt != nil {
		acquiredAt = pgtype.Timestamptz{
			Valid: true,
			Time:  *i.AcquiredAt,
		}
	}

	err = db.AddToPersonalCollection(ctx, queries.AddToPersonalCollectionParams{
		AccountID:  int32(i.AccountId),
		GameID:     int32(i.GameId),
		Paid:       paid,
		AcquiredAt: acquiredAt,
	})
	if err != nil {
		return err
	}

	return nil
}
