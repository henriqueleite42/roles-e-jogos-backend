package account_repository

import (
	"context"
	"errors"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type accountRepositoryImplementation struct {
	logger  *zerolog.Logger
	queries *queries.Queries
}

type NewAccountRepositoryInput struct {
	Logger  *zerolog.Logger
	Queries *queries.Queries
}

// Applies the transaction if it's needed
func (self *accountRepositoryImplementation) getSlcQueries(ctx context.Context) (*queries.Queries, error) {
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

func NewAccountRepository(i *NewAccountRepositoryInput) (AccountRepository, error) {
	if i == nil {
		return nil, errors.New("NewAccountRepository: input must not be nil")
	}

	return &accountRepositoryImplementation{
		logger:  i.Logger,
		queries: i.Queries,
	}, nil
}
