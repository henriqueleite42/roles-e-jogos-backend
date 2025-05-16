package collection_delivery_http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	collection_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/collection"
)

func (self *collectionController) CollectionCollective(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Collection").
		Str("mtd", r.Method).
		Str("route", "CollectionCollective").
		Str("reqId", reqId).
		Logger()

	if r.Method == http.MethodGet {
		getCollectiveCollectionInput := &collection_usecase.GetCollectiveCollectionInput{
			Pagination: models.GetDefaultPaginationInputString(),
		}

		query := r.URL.Query()

		afterQuery := query.Get("after")
		if afterQuery != "" {
			getCollectiveCollectionInput.Pagination.After = &afterQuery
		}
		limitQuery := query.Get("limit")
		if limitQuery != "" {
			limitInt, err := strconv.Atoi(limitQuery)
			if err == nil {
				getCollectiveCollectionInput.Pagination.Limit = limitInt
			} else {
				logger.Info().Err(err).Msg("fail to convert limit")
			}
		}
		accountIdQuery := query.Get("accountId")
		if accountIdQuery != "" {
			accountIdInt, err := strconv.Atoi(accountIdQuery)
			if err == nil {
				getCollectiveCollectionInput.AccountId = &accountIdInt
			} else {
				logger.Info().Err(err).Msg("fail to convert accountId")
			}
		}
		gameNameQuery := query.Get("gameName")
		if gameNameQuery != "" {
			getCollectiveCollectionInput.GameName = &gameNameQuery
		}
		kindQuery := query.Get("kind")
		if kindQuery != "" {
			getCollectiveCollectionInput.Kind = models.Kind(kindQuery)
		}
		maxAmountOfPlayersQuery := query.Get("maxAmountOfPlayers")
		if maxAmountOfPlayersQuery != "" {
			maxAmountOfPlayersInt, err := strconv.Atoi(maxAmountOfPlayersQuery)
			if err == nil {
				getCollectiveCollectionInput.MaxAmountOfPlayers = &maxAmountOfPlayersInt
			} else {
				logger.Info().Err(err).Msg("fail to convert maxAmountOfPlayers")
			}
		}

		logger.Trace().Msg("validate getCollectiveCollectionInput")
		err := self.validator.Validate(getCollectiveCollectionInput)
		if err != nil {
			logger.Info().Err(err).Msg("invalid getCollectiveCollectionInput")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Trace().Msg("create reqCtx")
		reqCtx := context.WithValue(context.Background(), "logger", logger)

		logger.Trace().Msg("call usecase")
		getCollectiveCollectionOutput, err := self.collectionUsecase.GetCollectiveCollection(reqCtx, getCollectiveCollectionInput)
		if err != nil {
			// If there are any errors that should be handled, add them here
			logger.Warn().Err(err).Msg("usecase err")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonOutput, err := json.Marshal(getCollectiveCollectionOutput)
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
