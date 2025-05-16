package event_usecase

import (
	"context"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

type ConfirmAttendanceInput struct {
	AccountId    int
	Confirmation models.EventAttendanceStatus `validate:"required" db:"confirmation"`
	EventId      int
}
type CreateEventInput struct {
	AccountId   int                   `validate:"id"`
	Capacity    *int                  `validate:"omitempty,min=1,max=9999"`
	Description string                `validate:"min=1,max=1000"`
	EndDate     *time.Time            `validate:"omitempty"`
	GamesList   []int                 `validate:"required"`
	Icon        *CreateEventInputIcon `validate:"required"`
	LocationId  int                   `validate:"id"`
	Name        string                `validate:"min=1,max=50"`
	StartDate   time.Time             `validate:"required"`
}
type CreateEventInputIcon struct {
	CustomIconPath    *string `validate:"omitempty,path"`
	UseGameIconGameId *int    `validate:"omitempty"`
}
type GetNextEventsInput struct {
	Pagination *models.PaginationInputTimestamp `validate:"required"`
}
type GetNextEventsOutput struct {
	Data       []*models.EventData               `validate:"required"`
	Pagination *models.PaginationOutputTimestamp `validate:"required"`
}

type EventUsecase interface {
	ConfirmAttendance(ctx context.Context, i *ConfirmAttendanceInput) error
	CreateEvent(ctx context.Context, i *CreateEventInput) error
	GetNextEvents(ctx context.Context, i *GetNextEventsInput) (*GetNextEventsOutput, error)
}
