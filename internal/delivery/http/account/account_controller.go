package account_delivery_http

import (
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	account_usecase "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/account"
	"github.com/rs/zerolog"
)

type accountController struct {
	logger         *zerolog.Logger
	validator      adapters.Validator
	authAdapter    adapters.Auth
	secretsAdapter *adapters.Secrets
	idAdapter      adapters.Id
	accountUsecase account_usecase.AccountUsecase
}

type AddAccountControllerInput struct {
	Mux    *http.ServeMux
	Logger *zerolog.Logger

	Validator      adapters.Validator
	AuthAdapter    adapters.Auth
	SecretsAdapter *adapters.Secrets
	IdAdapter      adapters.Id

	AccountUsecase account_usecase.AccountUsecase
}

func AddAccountController(i *AddAccountControllerInput) {
	accountController := &accountController{
		logger:         i.Logger,
		validator:      i.Validator,
		authAdapter:    i.AuthAdapter,
		secretsAdapter: i.SecretsAdapter,
		idAdapter:      i.IdAdapter,
		accountUsecase: i.AccountUsecase,
	}

	i.Mux.HandleFunc("/auth/google", accountController.AuthGoogle)
	i.Mux.HandleFunc("/auth/link/ludopedia", accountController.AuthLinkLudopedia)
	i.Mux.HandleFunc("/logout", accountController.Logout)

	i.Mux.HandleFunc("/profile/handle", accountController.ProfileHandle)
	i.Mux.HandleFunc("/profile/handle/check", accountController.ProfileHandleCheck)
	i.Mux.HandleFunc("/profile/me", accountController.ProfileMe)
	i.Mux.HandleFunc("/profile/list/by-handle", accountController.ProfileListByHandle)
}
