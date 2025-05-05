package adapters

import (
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

type HasValidSessionInput struct {
	Req *http.Request
}

type Auth interface {
	HasValidSession(i *HasValidSessionInput) (*models.AccountDataDb, error)
}
