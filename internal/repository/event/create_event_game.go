package event_repository

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
)

func (self *eventRepositoryImplementation) CreateEventGame(ctx context.Context, i *CreateEventGameInput) error {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return err
	}

	err = db.CreateEventGame(ctx, queries.CreateEventGameParams{
		EventID: int32(i.EventId),
		GameID:  int32(i.GameId),
		OwnerID: int32(i.OwnerId),
	})
	if err != nil {
		return err
	}

	return nil
}
