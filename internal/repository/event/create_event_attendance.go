package event_repository

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
)

func (self *eventRepositoryImplementation) CreateEventAttendance(ctx context.Context, i *CreateEventAttendanceInput) error {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return err
	}

	err = db.CreateEventAttendance(ctx, queries.CreateEventAttendanceParams{
		AccountID: int32(i.AccountId),
		EventID:   int32(i.EventId),
		Status:    queries.EventAttendanceStatusEnum(i.Status),
	})
	if err != nil {
		return err
	}

	return nil
}
