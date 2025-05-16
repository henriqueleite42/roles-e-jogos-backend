package event_usecase

import (
	"context"

	event_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/event"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *EventUsecaseImplementation) ConfirmAttendance(ctx context.Context, i *ConfirmAttendanceInput) error {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return err
	}

	err = self.EventRepository.CreateEventAttendance(ctx, &event_repository.CreateEventAttendanceInput{
		AccountId: i.AccountId,
		EventId:   i.EventId,
		Status:    i.Confirmation,
	})
	if err != nil {
		self.Logger.Warn().Err(err).Msg("fail to confirm attendance")
		return err
	}

	return nil
}
