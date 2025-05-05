package secretmanager_paramstore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/rs/zerolog"
)

type secretmanagerParamstore struct {
	logger *zerolog.Logger

	secrets *adapters.Secrets
}

func NewSecretManagerParamStore(logger *zerolog.Logger) (*adapters.Secrets, error) {
	instance := &secretmanagerParamstore{
		logger:  logger,
		secrets: &adapters.Secrets{},
	}

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	logger.Trace().Msg("load secrets")
	err = instance.loadSecrets(cfg)
	if err != nil {
		logger.Error().Err(err).Msg("fail to load secrets")
		return nil, err
	}

	logger.Trace().Msg("load variables")
	err = instance.loadVariables(cfg)
	if err != nil {
		logger.Error().Err(err).Msg("fail to load variables")
		return nil, err
	}

	return instance.secrets, nil
}
