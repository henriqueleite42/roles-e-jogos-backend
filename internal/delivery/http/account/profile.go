package account_delivery_http

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	account_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/account"
)

func (self *accountController) Profile(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Account").
		Str("mtd", r.Method).
		Str("route", "EditProfile").
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

		editProfileInput := &account_usecase.EditProfileInput{}
		err = json.Unmarshal(body, editProfileInput)
		if err != nil {
			logger.Info().Err(err).Msg("error unmarshalling body")
			http.Error(w, "error unmarshalling body", http.StatusBadRequest)
			return
		}

		editProfileInput.AccountId = session.AccountId

		logger.Trace().Msg("validate editProfileInput")
		err = self.validator.Validate(editProfileInput)
		if err != nil {
			logger.Info().Err(err).Msg("invalid editProfileInput")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Trace().Msg("create reqCtx")
		reqCtx := context.WithValue(context.Background(), "logger", logger)

		logger.Trace().Msg("call usecase")
		err = self.accountUsecase.EditProfile(reqCtx, editProfileInput)
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
