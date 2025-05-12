package collection_delivery_http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	collection_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/collection"
)

func (self *collectionController) CollectionCollective(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Account").
		Str("mtd", r.Method).
		Str("route", "CollectionCollective").
		Str("reqId", reqId).
		Logger()

	if r.Method == http.MethodGet {
		_, err := self.authAdapter.HasValidSession(&adapters.HasValidSessionInput{
			Req: r,
		})
		if err != nil {
			logger.Warn().Err(err).Msg("invalid cookie")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

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
			if err != nil {
				getCollectiveCollectionInput.Pagination.Limit = &limitInt
			}
		}
		accountIdQuery := query.Get("accountId")
		if accountIdQuery != "" {
			accountIdInt, err := strconv.Atoi(accountIdQuery)
			if err != nil {
				getCollectiveCollectionInput.AccountId = &accountIdInt
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
			if err != nil {
				getCollectiveCollectionInput.MaxAmountOfPlayers = &maxAmountOfPlayersInt
			}
		}

		logger.Trace().Msg("validate getCollectiveCollectionInput")
		err = self.validator.Validate(getCollectiveCollectionInput)
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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonOutput)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
