package event_repository

import (
	"context"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

type CreateEventAttendanceInput struct {
	AccountId int                          `db:"account_id"`
	EventId   int                          `db:"event_id"`
	Status    models.EventAttendanceStatus `validate:"required" db:"status"`
}
type CreateEventGameInput struct {
	EventId int
	GameId  int
	OwnerId int
}
type CreateEventInput struct {
	Capacity    *int       `validate:"omitempty"`
	Description string
	EndDate     *time.Time `validate:"omitempty"`
	IconPath    *string    `validate:"omitempty"`
	LocationId  int
	Name        string
	OwnerId     int
	StartDate   time.Time  `validate:"required"`
}
type GetNextEventsInput struct {
	Pagination *models.PaginationInputTimestamp `validate:"required"`
}
type GetNextEventsOutput struct {
	Data       []*models.EventData               `validate:"required"`
	Pagination *models.PaginationOutputTimestamp `validate:"required"`
}

type EventRepository interface {
	CreateEvent(ctx context.Context, i *CreateEventInput) (*models.Event, error)
	CreateEventAttendance(ctx context.Context, i *CreateEventAttendanceInput) error
	CreateEventGame(ctx context.Context, i *CreateEventGameInput) error
	GetNextEvents(ctx context.Context, i *GetNextEventsInput) (*GetNextEventsOutput, error)
}
