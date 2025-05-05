package account_delivery_http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	account_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/account"
)

func (self *accountController) CheckHandle(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Account").
		Str("mtd", "GetProfileByHandle").
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

		query := r.URL.Query()
		handle := query.Get("handle")

		checkHandleInput := &account_usecase.CheckHandleInput{
			Handle: handle,
		}

		logger.Trace().Msg("validate checkHandleInput")
		err = self.validator.Validate(checkHandleInput)
		if err != nil {
			logger.Info().Err(err).Msg("invalid checkHandleInput")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Trace().Msg("create reqCtx")
		reqCtx := context.WithValue(context.Background(), "logger", logger)

		logger.Trace().Msg("call usecase")
		checkHandleOutput, err := self.accountUsecase.CheckHandle(reqCtx, checkHandleInput)
		if err != nil {
			// If there are any errors that should be handled, add them here
			logger.Warn().Err(err).Msg("usecase err")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonOutput, err := json.Marshal(checkHandleOutput)
		if err != nil {
			logger.Error().Err(err).Msg("failed to marshal JSON")
			http.Error(w, "failed to marshal JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonOutput)
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
