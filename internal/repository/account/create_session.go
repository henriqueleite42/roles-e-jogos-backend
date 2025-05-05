package account_repository

import (
	"context"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
)

func (self *accountRepositoryImplementation) CreateSession(ctx context.Context, i *CreateSessionInput) (*models.Session, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	sessionId := self.idAdapter.GenSessionId()

	err = db.CreateSession(ctx, queries.CreateSessionParams{
		SessionID: sessionId,
		AccountID: int32(i.AccountId),
	})
	if err != nil {
		return nil, err
	}

	return &models.Session{
		SessionId: sessionId,
		AccountId: i.AccountId,
		CreatedAt: time.Now(),
	}, nil
}
