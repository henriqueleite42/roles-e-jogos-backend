package collection_repository

import (
	"context"
	"errors"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type collectionRepositoryImplementation struct {
	logger  *zerolog.Logger
	queries *queries.Queries
}

type NewCollectionRepositoryInput struct {
	Logger  *zerolog.Logger
	Queries *queries.Queries
}

// Applies the transaction if it's needed
func (self *collectionRepositoryImplementation) getSlcQueries(ctx context.Context) (*queries.Queries, error) {
	txAny := ctx.Value("tx")
	if txAny == nil {
		return self.queries, nil
	}
	tx, ok := txAny.(*pgxpool.Tx)
	if !ok {
		return self.queries, nil
	}

	return self.queries.WithTx(tx), nil
}

func NewCollectionRepository(i *NewCollectionRepositoryInput) (CollectionRepository, error) {
	if i == nil {
		return nil, errors.New("NewCollectionRepository: input must not be nil")
	}

	return &collectionRepositoryImplementation{
		logger:  i.Logger,
		queries: i.Queries,
	}, nil
}
