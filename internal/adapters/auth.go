package adapters

import (
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

type HasValidSessionInput struct {
	Req *http.Request
}

type SetSessionOnResInput struct {
	Res       http.ResponseWriter
	SessionId string
}

type Auth interface {
	SetSessionOnRes(i *SetSessionOnResInput)
	HasValidSession(i *HasValidSessionInput) (*models.AccountDataDb, error)
}
