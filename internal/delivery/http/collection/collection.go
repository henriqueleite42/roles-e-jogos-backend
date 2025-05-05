package collection_delivery_http

import (
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	collection_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/collection"
	"github.com/rs/zerolog"
)

type collectionController struct {
	logger            *zerolog.Logger
	validator         adapters.Validator
	idAdapter         adapters.Id
	collectionUsecase collection_usecase.CollectionUsecase
}

type AddCollectionControllerInput struct {
	Logger *zerolog.Logger

	Validator adapters.Validator
	IdAdapter adapters.Id

	CollectionUsecase collection_usecase.CollectionUsecase
}

func AddCollectionController(i *AddCollectionControllerInput) {
	collectionController := &collectionController{
		logger:            i.Logger,
		validator:         i.Validator,
		idAdapter:         i.IdAdapter,
		collectionUsecase: i.CollectionUsecase,
	}

	// Add routes here. Ex:
	// http.HandleFunc("/", self.Handler)
}
