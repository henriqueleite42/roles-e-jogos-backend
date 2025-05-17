package auth_postgres

import (
	"fmt"
	"os"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

func (self *authPostgresAdapter) GetSessionId(i *adapters.GetSessionIdInput) (string, error) {
	cookie, err := i.Req.Cookie(SESSION_COOKIE_NAME)
	if err != nil {
		return "", err
	}

	// Applies extra security on prod
	if os.Getenv("ENV") == "prod" && !cookie.HttpOnly {
		return "", fmt.Errorf("invalid auth cookie")
	}

	return cookie.Value, nil
}
