package event_usecase

import (
	"context"
	"errors"
)

func (self *EventUsecaseImplementation) GetEvents(ctx context.Context, i *GetEventsInput) (*GetEventsOutput, error) {
	return nil, errors.New("\"GetEvents\": unimplemented")
}
