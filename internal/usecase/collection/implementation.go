package collection_usecase

import (
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	collection_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/collection"
	game_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/game"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type CollectionUsecaseImplementation struct {
	Logger *zerolog.Logger

	Db *pgxpool.Pool

	LudopediaAdapter adapters.Ludopedia
	IdAdapter        adapters.Id
	EmailAdapter     adapters.Email
	SecretsAdapter   *adapters.Secrets
	StorageAdapter   adapters.Storage
	MessagingAdapter adapters.Messaging

	AccountRepository    account_repository.AccountRepository
	GameRepository       game_repository.GameRepository
	CollectionRepository collection_repository.CollectionRepository
}
