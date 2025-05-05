package collection_usecase

import (
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	collection_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/collection"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type CollectionUsecaseImplementation struct {
	Logger *zerolog.Logger

	Db *pgxpool.Pool

	CollectionRepository collection_repository.CollectionRepository

	LudopediaAdapter adapters.SignInProvider
	IdAdapter        adapters.Id
	EmailAdapter     adapters.Email
	SecretsAdapter   *adapters.Secrets
	StorageAdapter   adapters.Storage
}
