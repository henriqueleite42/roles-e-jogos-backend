package account_delivery_http

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	account_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/account"
)

func (self *accountController) ProfileHandle(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Account").
		Str("mtd", r.Method).
		Str("route", "ProfileHandle").
		Str("reqId", reqId).
		Logger()

	if r.Method == http.MethodPut {
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

		editHandleInput := &account_usecase.EditHandleInput{}
		err = json.Unmarshal(body, editHandleInput)
		if err != nil {
			logger.Info().Err(err).Msg("error unmarshalling body")
			http.Error(w, "error unmarshalling body", http.StatusBadRequest)
			return
		}

		editHandleInput.AccountId = session.AccountId

		logger.Trace().Msg("validate editHandleInput")
		err = self.validator.Validate(editHandleInput)
		if err != nil {
			logger.Info().Err(err).Msg("invalid editHandleInput")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Trace().Msg("create reqCtx")
		reqCtx := context.WithValue(context.Background(), "logger", logger)

		logger.Trace().Msg("call usecase")
		err = self.accountUsecase.EditHandle(reqCtx, editHandleInput)
		if err != nil {
			// If there are any errors that should be handled, add them here

			if err.Error() == "conflict" {
				http.Error(w, "conflict", http.StatusConflict)
				return
			}

			logger.Warn().Err(err).Msg("usecase err")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Trace().Msg("send response")
		w.WriteHeader(http.StatusOK)
		logger.Trace().Msg("finish")
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
