package event_usecase

import (
	"context"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

type ConfirmAttendanceInput struct {
	AccountId    int                          `validate:"id" db:"id"`
	Confirmation models.EventAttendanceStatus `validate:"required" db:"confirmation"`
}
type CreateEventInput struct {
	AccountId          int                   `validate:"id" db:"id"`
	Date               time.Time             `validate:"required"`
	Description        string                `validate:"min=1,max=1000"`
	GamesList          []int                 `validate:"required"`
	Icon               *CreateEventInputIcon `validate:"required"`
	LocationAddress    string                `validate:"min=1,max=500"`
	LocationName       string                `validate:"min=1,max=100"`
	MaxAmountOfPlayers *int                  `validate:"omitempty,min=1,max=9999"`
	Name               string                `validate:"min=1,max=50"`
}
type CreateEventInputIcon struct {
	CustomIcon *string `validate:"omitempty,path"`
	GameIcon   *int    `validate:"omitempty"`
}
type GetEventsInput struct {
	Pagination *models.PaginationInputString `validate:"omitempty"`
}
type GetEventsOutput struct {
	Data       []*models.EventData            `validate:"required"`
	Pagination *models.PaginationOutputString `validate:"required"`
}

type EventUsecase interface {
	ConfirmAttendance(ctx context.Context, i *ConfirmAttendanceInput) error
	CreateEvent(ctx context.Context, i *CreateEventInput) error
	GetEvents(ctx context.Context, i *GetEventsInput) (*GetEventsOutput, error)
}
