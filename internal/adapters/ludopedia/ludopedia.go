package ludopedia

import (
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/rs/zerolog"
)

type ludopediaAdapter struct {
	logger     *zerolog.Logger
	httpClient *http.Client
	secrets    *adapters.Secrets
}

func NewLudopedia(logger *zerolog.Logger, secretsAdapter *adapters.Secrets) (adapters.SignInProvider, adapters.Ludopedia, error) {
	newLogger := logger.With().Str("adapter", "LudopediaAdapter").Logger()

	adapter := &ludopediaAdapter{
		logger:     &newLogger,
		httpClient: &http.Client{},
		secrets:    secretsAdapter,
	}

	return adapter, adapter, nil
}
