package account_delivery_http

import (
	"context"
	"encoding/json"
	"net/http"

	account_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/account"
)

func (self *accountController) ProfileListByHandle(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Account").
		Str("mtd", "ProfileListByHandle").
		Str("reqId", reqId).
		Logger()

	if r.Method == http.MethodGet {
		getProfilesListByHandleInput := &account_usecase.GetProfilesListByHandleInput{}

		query := r.URL.Query()

		handleQuery := query.Get("handle")
		if handleQuery != "" {
			getProfilesListByHandleInput.Handle = handleQuery
		}

		logger.Trace().Msg("validate getProfilesListByHandleInput")
		err := self.validator.Validate(getProfilesListByHandleInput)
		if err != nil {
			logger.Info().Err(err).Msg("invalid getProfilesListByHandleInput")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Trace().Msg("create reqCtx")
		reqCtx := context.WithValue(context.Background(), "logger", logger)

		logger.Trace().Msg("call usecase")
		getProfilesListByHandleOutput, err := self.accountUsecase.GetProfilesListByHandle(reqCtx, getProfilesListByHandleInput)
		if err != nil {
			// If there are any errors that should be handled, add them here
			logger.Warn().Err(err).Msg("usecase err")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonOutput, err := json.Marshal(getProfilesListByHandleOutput)
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
