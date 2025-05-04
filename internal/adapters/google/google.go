package google

import (
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/rs/zerolog"
)

type googleAdapter struct {
	logger     *zerolog.Logger
	httpClient *http.Client
	secrets    *adapters.Secrets
}

func NewGoogle(logger *zerolog.Logger, secretsAdapter *adapters.Secrets) (adapters.SignInProvider, error) {
	newLogger := logger.With().Str("adapter", "GoogleAdapter").Logger()

	return &googleAdapter{
		logger:     &newLogger,
		httpClient: &http.Client{},
		secrets:    secretsAdapter,
	}, nil
}
