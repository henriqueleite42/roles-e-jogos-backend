package event_usecase

import (
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type EventUsecaseImplementation struct {
	Logger *zerolog.Logger

	Db *pgxpool.Pool

	SecretsAdapter *adapters.Secrets
}
