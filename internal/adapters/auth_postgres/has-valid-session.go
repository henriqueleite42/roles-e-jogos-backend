package auth_postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
)

func (self *authPostgresAdapter) HasValidSession(i *adapters.HasValidSessionInput) (*models.AccountDataDb, error) {
	cookie, err := i.Req.Cookie(SESSION_COOKIE_NAME)
	if err != nil {
		return nil, err
	}

	// Applies extra security on prod
	if os.Getenv("ENV") == "prod" && !cookie.HttpOnly {
		return nil, fmt.Errorf("invalid auth cookie")
	}

	sessionId := cookie.Value

	account, err := self.accountRepository.GetAccountDataBySessionId(context.TODO(), &account_repository.GetAccountDataBySessionIdInput{
		SessionId: sessionId,
	})
	if err != nil {
		return nil, err
	}

	return account, nil
}
