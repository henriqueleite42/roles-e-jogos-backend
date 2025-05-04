package game_repository

import (
	"errors"

	"github.com/rs/zerolog"
)

type gameRepositoryImplementation struct {
	logger *zerolog.Logger
}

type NewGameRepositoryInput struct {
	Logger *zerolog.Logger
}

func NewGameRepository(i *NewGameRepositoryInput) (GameRepository, error) {
	if i == nil {
		return nil, errors.New("NewGameRepository: input must not be nil")
	}

	return &gameRepositoryImplementation{
		logger: i.Logger,
	}, nil
}
