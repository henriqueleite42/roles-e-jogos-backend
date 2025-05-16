package event_usecase

import (
	"context"

	event_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/event"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *EventUsecaseImplementation) GetNextEvents(ctx context.Context, i *GetNextEventsInput) (*GetNextEventsOutput, error) {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return nil, err
	}

	events, err := self.EventRepository.GetNextEvents(ctx, &event_repository.GetNextEventsInput{
		Pagination: i.Pagination,
	})
	if err != nil {
		self.Logger.Warn().Err(err).Msg("fail to confirm attendance")
		return nil, err
	}

	return &GetNextEventsOutput{
		Data:       events.Data,
		Pagination: events.Pagination,
	}, nil
}
