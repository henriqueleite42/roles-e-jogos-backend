package collection_repository

import (
	"errors"

	"github.com/rs/zerolog"
)

type collectionRepositoryImplementation struct {
	logger *zerolog.Logger
}

type NewCollectionRepositoryInput struct {
	Logger *zerolog.Logger
}

func NewCollectionRepository(i *NewCollectionRepositoryInput) (CollectionRepository, error) {
	if i == nil {
		return nil, errors.New("NewCollectionRepository: input must not be nil")
	}

	return &collectionRepositoryImplementation{
		logger: i.Logger,
	}, nil
}
