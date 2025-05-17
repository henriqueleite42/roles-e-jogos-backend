package adapters

import (
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

type GetSessionIdInput struct {
	Req *http.Request
}

type HasValidSessionInput struct {
	Req *http.Request
}

type SetSessionOnResInput struct {
	Res       http.ResponseWriter
	SessionId string
}

type DeleteSessionFromResInput struct {
	Res http.ResponseWriter
}

type Auth interface {
	GetSessionId(i *GetSessionIdInput) (string, error)
	SetSessionOnRes(i *SetSessionOnResInput)
	DeleteSessionFromRes(i *DeleteSessionFromResInput)
	HasValidSession(i *HasValidSessionInput) (*models.AccountDataDb, error)
}
