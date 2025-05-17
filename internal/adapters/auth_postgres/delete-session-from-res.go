package auth_postgres

import (
	"net/http"
	"os"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

func (self *authPostgresAdapter) DeleteSessionFromRes(i *adapters.DeleteSessionFromResInput) {
	isProd := os.Getenv("ENV") == "prod"

	var domain string
	if isProd {
		domain = "rolesejogos.com.br"
	}

	http.SetCookie(i.Res, &http.Cookie{
		Name:     SESSION_COOKIE_NAME,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Domain:   domain,
		Secure:   isProd,
		Expires:  time.Now().Add(-10 * time.Minute),
		SameSite: http.SameSiteLaxMode,
	})
}
