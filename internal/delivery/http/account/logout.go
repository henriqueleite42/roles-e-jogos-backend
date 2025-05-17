package account_delivery_http

import (
	"context"
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	account_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/account"
)

func (self *accountController) Logout(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Account").
		Str("mtd", "Logout").
		Str("reqId", reqId).
		Logger()

	if r.Method == http.MethodPost {
		_, err := self.authAdapter.HasValidSession(&adapters.HasValidSessionInput{
			Req: r,
		})
		if err != nil {
			logger.Warn().Err(err).Msg("invalid cookie")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		sessionId, err := self.authAdapter.GetSessionId(&adapters.GetSessionIdInput{
			Req: r,
		})
		if err != nil {
			logger.Warn().Err(err).Msg("usecase err")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		deleteSessionInput := &account_usecase.DeleteSessionInput{
			SessionId: sessionId,
		}

		logger.Trace().Msg("validate deleteSessionInput")
		err = self.validator.Validate(deleteSessionInput)
		if err != nil {
			logger.Info().Err(err).Msg("invalid deleteSessionInput")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Trace().Msg("create reqCtx")
		reqCtx := context.WithValue(context.Background(), "logger", logger)

		logger.Trace().Msg("call usecase")
		err = self.accountUsecase.DeleteSession(reqCtx, deleteSessionInput)
		if err != nil {
			// If there are any errors that should be handled, add them here
			logger.Warn().Err(err).Msg("usecase err")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Trace().Msg("send response")
		self.authAdapter.DeleteSessionFromRes(&adapters.DeleteSessionFromResInput{
			Res: w,
		})
		logger.Trace().Msg("finish")
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
