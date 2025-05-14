package collection_delivery_http

import (
	"context"
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	collection_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/collection"
)

func (self *collectionController) CollectionImportLudopedia(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Collection").
		Str("mtd", "CollectionImportLudopedia").
		Str("reqId", reqId).
		Logger()

	if r.Method == http.MethodPost {
		session, err := self.authAdapter.HasValidSession(&adapters.HasValidSessionInput{
			Req: r,
		})
		if err != nil {
			logger.Warn().Err(err).Msg("invalid cookie")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		requestImportPersonalCollectionFromLudopediaInput := &collection_usecase.RequestImportPersonalCollectionFromLudopediaInput{
			AccountId: session.AccountId,
		}

		logger.Trace().Msg("validate requestImportPersonalCollectionFromLudopediaInput")
		err = self.validator.Validate(requestImportPersonalCollectionFromLudopediaInput)
		if err != nil {
			logger.Info().Err(err).Msg("invalid requestImportPersonalCollectionFromLudopediaInput")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Trace().Msg("create reqCtx")
		reqCtx := context.WithValue(context.Background(), "logger", logger)

		logger.Trace().Msg("call usecase")
		err = self.collectionUsecase.RequestImportPersonalCollectionFromLudopedia(reqCtx, requestImportPersonalCollectionFromLudopediaInput)
		if err != nil {
			// If there are any errors that should be handled, add them here
			logger.Warn().Err(err).Msg("usecase err")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
