package collection_delivery_http

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	collection_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/collection"
)

func (self *collectionController) CollectionPersonal(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Account").
		Str("mtd", r.Method).
		Str("route", "AddToPersonalCollection").
		Str("reqId", reqId).
		Logger()

	if r.Method == http.MethodGet {
		session, err := self.authAdapter.HasValidSession(&adapters.HasValidSessionInput{
			Req: r,
		})
		if err != nil {
			logger.Warn().Err(err).Msg("invalid cookie")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			logger.Info().Err(err).Msg("error reading request body")
			http.Error(w, "error reading request body", http.StatusInternalServerError)
			return
		}

		addToPersonalCollectionInput := &collection_usecase.AddToPersonalCollectionInput{}
		err = json.Unmarshal(body, addToPersonalCollectionInput)
		if err != nil {
			logger.Info().Err(err).Msg("error unmarshalling body")
			http.Error(w, "error unmarshalling body", http.StatusBadRequest)
			return
		}

		addToPersonalCollectionInput.AccountId = session.AccountId

		logger.Trace().Msg("validate addToPersonalCollectionInput")
		err = self.validator.Validate(addToPersonalCollectionInput)
		if err != nil {
			logger.Info().Err(err).Msg("invalid addToPersonalCollectionInput")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Trace().Msg("create reqCtx")
		reqCtx := context.WithValue(context.Background(), "logger", logger)

		logger.Trace().Msg("call usecase")
		err = self.collectionUsecase.AddToPersonalCollection(reqCtx, addToPersonalCollectionInput)
		if err != nil {
			// If there are any errors that should be handled, add them here
			logger.Warn().Err(err).Msg("usecase err")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
