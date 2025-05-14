package account_delivery_http

import (
	"context"
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	account_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/account"
)

func (self *accountController) AuthLinkLudopedia(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Account").
		Str("mtd", r.Method).
		Str("route", "AuthLinkLudopedia").
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
		code := query.Get("code")

		linkLudopediaProviderInput := &account_usecase.LinkLudopediaProviderInput{
			AccountId: session.AccountId,
			Code:      code,
		}

		logger.Trace().Msg("validate linkLudopediaProviderInput")
		err = self.validator.Validate(linkLudopediaProviderInput)
		if err != nil {
			logger.Info().Err(err).Msg("invalid linkLudopediaProviderInput")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Trace().Msg("create reqCtx")
		reqCtx := context.WithValue(context.Background(), "logger", logger)

		logger.Trace().Msg("call usecase")
		err = self.accountUsecase.LinkLudopediaProvider(reqCtx, linkLudopediaProviderInput)
		if err != nil {
			// If there are any errors that should be handled, add them here
			logger.Warn().Err(err).Msg("usecase err")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Trace().Msg("send response")
		http.Redirect(w, r, self.secretsAdapter.WebsiteUrl+"/conta", http.StatusSeeOther)
		logger.Trace().Msg("finish")
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
