package viacep

import (
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/rs/zerolog"
)

type viacepAdapter struct {
	logger *zerolog.Logger
}

func NewViaCepAdapter(logger *zerolog.Logger) (adapters.Address, error) {
	newLogger := logger.With().Str("adapter", "ViacepAdapter").Logger()

	return &viacepAdapter{
		logger: &newLogger,
	}, nil
}
