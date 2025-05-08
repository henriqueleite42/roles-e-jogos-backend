package auth_postgres

import (
	"net/http"
	"os"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

const ONE_YEAR_IN_HOURS = 8766 * time.Hour

func (self *authPostgresAdapter) SetSessionOnRes(i *adapters.SetSessionOnResInput) {
	isProd := os.Getenv("ENV") == "prod"

	var domain string
	if isProd {
		domain = "rolesejogos.com.br"
	}

	http.SetCookie(i.Res, &http.Cookie{
		Name:     SESSION_COOKIE_NAME,
		Value:    i.SessionId,
		Path:     "/",
		HttpOnly: true,
		Domain:   domain,
		Secure:   isProd,
		Expires:  time.Now().Add(10 * ONE_YEAR_IN_HOURS),
		SameSite: http.SameSiteStrictMode,
	})
}
