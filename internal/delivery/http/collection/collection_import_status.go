package collection_delivery_http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	collection_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/collection"
)

func (self *collectionController) CollectionImportStatus(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Collection").
		Str("mtd", "CollectionImportStatus").
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

		query := r.URL.Query()
		externalId := query.Get("externalId")
		provider := query.Get("provider")

		getLatestImportCollectionLogStatusInput := &collection_usecase.GetLatestImportCollectionLogStatusInput{
			AccountId:  session.AccountId,
			ExternalId: externalId,
			Provider:   models.Provider(provider),
		}

		logger.Trace().Msg("validate getLatestImportCollectionLogStatusInput")
		err = self.validator.Validate(getLatestImportCollectionLogStatusInput)
		if err != nil {
			logger.Info().Err(err).Msg("invalid getLatestImportCollectionLogStatusInput")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Trace().Msg("create reqCtx")
		reqCtx := context.WithValue(context.Background(), "logger", logger)

		logger.Trace().Msg("call usecase")
		getLatestImportCollectionLogStatusOutput, err := self.collectionUsecase.GetLatestImportCollectionLogStatus(reqCtx, getLatestImportCollectionLogStatusInput)
		if err != nil {
			// If there are any errors that should be handled, add them here
			logger.Warn().Err(err).Msg("usecase err")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonOutput, err := json.Marshal(getLatestImportCollectionLogStatusOutput)
		if err != nil {
			logger.Error().Err(err).Msg("failed to marshal JSON")
			http.Error(w, "failed to marshal JSON", http.StatusInternalServerError)
			return
		}

		logger.Trace().Msg("send response")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonOutput)
		logger.Trace().Msg("finish")
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
