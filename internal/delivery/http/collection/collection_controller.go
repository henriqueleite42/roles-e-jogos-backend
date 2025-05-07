package collection_delivery_http

import (
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	collection_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/collection"
	"github.com/rs/zerolog"
)

type collectionController struct {
	logger            *zerolog.Logger
	validator         adapters.Validator
	authAdapter       adapters.Auth
	secretsAdapter    *adapters.Secrets
	idAdapter         adapters.Id
	collectionUsecase collection_usecase.CollectionUsecase
}

type AddCollectionControllerInput struct {
	Logger *zerolog.Logger

	Validator      adapters.Validator
	AuthAdapter    adapters.Auth
	SecretsAdapter *adapters.Secrets
	IdAdapter      adapters.Id

	CollectionUsecase collection_usecase.CollectionUsecase
}

func AddCollectionController(i *AddCollectionControllerInput) {
	collectionController := &collectionController{
		logger:            i.Logger,
		validator:         i.Validator,
		authAdapter:       i.AuthAdapter,
		secretsAdapter:    i.SecretsAdapter,
		idAdapter:         i.IdAdapter,
		collectionUsecase: i.CollectionUsecase,
	}

	http.HandleFunc("/collection/personal", collectionController.AddToPersonalCollection)
	http.HandleFunc("/collection/collective", collectionController.GetCollectiveCollection)
}
