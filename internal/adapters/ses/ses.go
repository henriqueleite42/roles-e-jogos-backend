package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/rs/zerolog"
)

type sesAdapter struct {
	logger *zerolog.Logger
	client *ses.Client
}

func NewSes(logger *zerolog.Logger) (adapters.Email, error) {
	newLogger := logger.With().Str("adapter", "SesAdapter").Logger()

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	sesClient := ses.NewFromConfig(cfg)

	return &sesAdapter{
		logger: &newLogger,
		client: sesClient,
	}, nil
}
