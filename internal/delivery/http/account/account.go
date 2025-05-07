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

	http.HandleFunc("//profile/handle/check", accountController.CheckHandle)
}
