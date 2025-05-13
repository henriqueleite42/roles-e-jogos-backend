package collection_delivery_http

import (
	"context"
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	collection_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/collection"
)

func (self *collectionController) CollectionPersonalImportLudopedia(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Collection").
		Str("mtd", "CollectionPersonalImportLudopedia").
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

		importPersonalCollectionFromLudopediaInput := &collection_usecase.ImportPersonalCollectionFromLudopediaInput{
			AccountId: session.AccountId,
		}

		logger.Trace().Msg("validate importPersonalCollectionFromLudopediaInput")
		err = self.validator.Validate(importPersonalCollectionFromLudopediaInput)
		if err != nil {
			logger.Info().Err(err).Msg("invalid importPersonalCollectionFromLudopediaInput")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Trace().Msg("create reqCtx")
		reqCtx := context.WithValue(context.Background(), "logger", logger)

		logger.Trace().Msg("call usecase")
		err = self.collectionUsecase.ImportPersonalCollectionFromLudopedia(reqCtx, importPersonalCollectionFromLudopediaInput)
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
