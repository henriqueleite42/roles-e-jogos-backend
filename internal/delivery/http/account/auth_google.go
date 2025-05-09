package account_delivery_http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	account_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/account"
)

func (self *accountController) AuthGoogle(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Account").
		Str("mtd", r.Method).
		Str("route", "AuthGoogle").
		Str("reqId", reqId).
		Logger()

	if r.Method == http.MethodGet {
		session, err := self.authAdapter.HasValidSession(&adapters.HasValidSessionInput{
			Req: r,
		})
		if session != nil && err == nil {
			logger.Warn().Msg("user already logged")
			fmt.Println(self.secretsAdapter.WebsiteUrl + "/conta")
			http.Redirect(w, r, self.secretsAdapter.WebsiteUrl+"/conta", http.StatusSeeOther)
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

		self.authAdapter.SetSessionOnRes(&adapters.SetSessionOnResInput{
			Res:       w,
			SessionId: createWithGoogleProviderOutput.SessionId,
		})
		fmt.Println(self.secretsAdapter.WebsiteUrl + "/conta")
		http.Redirect(w, r, self.secretsAdapter.WebsiteUrl+"/conta", http.StatusSeeOther)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
