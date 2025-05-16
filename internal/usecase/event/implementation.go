package event_usecase

import (
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	event_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/event"
	game_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/game"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type EventUsecaseImplementation struct {
	Logger *zerolog.Logger

	Db *pgxpool.Pool

	SecretsAdapter *adapters.Secrets

	EventRepository event_repository.EventRepository
	GameRepository  game_repository.GameRepository
}
