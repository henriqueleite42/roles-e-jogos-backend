package xid

import (
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/rs/zerolog"
)

type xidAdapter struct {
	logger *zerolog.Logger
}

func NewXid(logger *zerolog.Logger) (adapters.Id, error) {
	newLogger := logger.With().Str("adapter", "XidAdapter").Logger()

	return &xidAdapter{
		logger: &newLogger,
	}, nil
}
