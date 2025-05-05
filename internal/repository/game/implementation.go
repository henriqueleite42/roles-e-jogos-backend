package game_repository

import (
	"context"
	"errors"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type gameRepositoryImplementation struct {
	logger  *zerolog.Logger
	queries *queries.Queries
}

type NewGameRepositoryInput struct {
	Logger  *zerolog.Logger
	Queries *queries.Queries
}

// Applies the transaction if it's needed
func (self *gameRepositoryImplementation) getSlcQueries(ctx context.Context) (*queries.Queries, error) {
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

func NewGameRepository(i *NewGameRepositoryInput) (GameRepository, error) {
	if i == nil {
		return nil, errors.New("NewGameRepository: input must not be nil")
	}

	return &gameRepositoryImplementation{
		logger:  i.Logger,
		queries: i.Queries,
	}, nil
}
