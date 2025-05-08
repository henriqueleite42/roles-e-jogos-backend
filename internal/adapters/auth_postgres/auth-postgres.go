package auth_postgres

import (
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	"github.com/rs/zerolog"
)

type authPostgresAdapter struct {
	logger *zerolog.Logger

	accountRepository account_repository.AccountRepository
}

type NewAuthPostgresInput struct {
	Logger *zerolog.Logger

	AccountRepository account_repository.AccountRepository
}

const SESSION_COOKIE_NAME = "rolesejogos-session"

func NewAuthPostgres(i *NewAuthPostgresInput) (adapters.Auth, error) {
	newLogger := i.Logger.With().Str("adapter", "AuthPostgresAdapter").Logger()

	return &authPostgresAdapter{
		logger:            &newLogger,
		accountRepository: i.AccountRepository,
	}, nil
}
