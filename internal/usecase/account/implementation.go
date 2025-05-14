package account_usecase

import (
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type AccountUsecaseImplementation struct {
	Logger *zerolog.Logger

	Db *pgxpool.Pool

	AccountRepository account_repository.AccountRepository

	GoogleAdapter    adapters.SignInProvider
	LudopediaAdapter adapters.SignInProvider
	IdAdapter        adapters.Id
	EmailAdapter     adapters.Email
	SecretsAdapter   *adapters.Secrets
	StorageAdapter   adapters.Storage
	MessagingAdapter adapters.Messaging
}
