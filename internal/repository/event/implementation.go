package event_repository

import (
	"context"
	"errors"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type eventRepositoryImplementation struct {
	logger  *zerolog.Logger
	queries *queries.Queries

	secretsAdapter *adapters.Secrets
}

type NewEventRepositoryInput struct {
	Logger  *zerolog.Logger
	Queries *queries.Queries

	SecretsAdapter *adapters.Secrets
}

// Applies the transaction if it's needed
func (self *eventRepositoryImplementation) getSlcQueries(ctx context.Context) (*queries.Queries, error) {
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

func NewEventRepository(i *NewEventRepositoryInput) (EventRepository, error) {
	if i == nil {
		return nil, errors.New("NewEventRepository: input must not be nil")
	}

	return &eventRepositoryImplementation{
		logger:         i.Logger,
		queries:        i.Queries,
		secretsAdapter: i.SecretsAdapter,
	}, nil
}
