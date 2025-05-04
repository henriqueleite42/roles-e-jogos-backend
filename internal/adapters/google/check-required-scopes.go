package google

import (
	"fmt"
	"strings"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (adp *googleAdapter) CheckRequiredScopes(scopes []string) error {
	requiredScopes := []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"openid",
		"https://www.googleapis.com/auth/userinfo.email",
	}

	diff := utils.Diff(requiredScopes, scopes)

	if len(diff) > 0 {
		return fmt.Errorf("missing required scopes: %s", strings.Join(diff, ", "))
	}

	return nil
}
