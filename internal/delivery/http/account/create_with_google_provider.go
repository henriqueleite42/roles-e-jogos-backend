package account_delivery_http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	account_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/account"
)

func (self *accountController) CreateWithGoogleProvider(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Account").
		Str("mtd", r.Method).
		Str("route", "CreateWithGoogleProvider").
		Str("reqId", reqId).
		Logger()

	if r.Method == http.MethodGet {
		session, err := self.authAdapter.HasValidSession(&adapters.HasValidSessionInput{
			Req: r,
		})
		if session != nil && err == nil {
			logger.Warn().Msg("user already logged")
			http.Error(w, "user already logged", http.StatusUnauthorized)
			return
		}

		query := r.URL.Query()
		code := query.Get("code")

		createWithGoogleProviderInput := &account_usecase.CreateWithGoogleProviderInput{
			Code: code,
		}

		logger.Trace().Msg("validate createWithGoogleProviderInput")
		err = self.validator.Validate(createWithGoogleProviderInput)
		if err != nil {
			logger.Info().Err(err).Msg("invalid createWithGoogleProviderInput")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Trace().Msg("create reqCtx")
		reqCtx := context.WithValue(context.Background(), "logger", logger)

		logger.Trace().Msg("call usecase")
		createWithGoogleProviderOutput, err := self.accountUsecase.CreateWithGoogleProvider(reqCtx, createWithGoogleProviderInput)
		if err != nil {
			// If there are any errors that should be handled, add them here
			logger.Warn().Err(err).Msg("usecase err")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonOutput, err := json.Marshal(createWithGoogleProviderOutput)
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
