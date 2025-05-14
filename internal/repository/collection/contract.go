package collection_repository

import (
	"context"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

type AddToPersonalCollectionInput struct {
	AccountId  int
	AcquiredAt *time.Time `validate:"omitempty"`
	GameId     int
	Paid       *int       `validate:"omitempty"`
}
type CreateImportCollectionLogInput struct {
	AccountId  int
	ExternalId string
	Provider   models.Provider                `validate:"required" db:"provider"`
	Status     models.CollectionImportStatus  `validate:"required" db:"status"`
	Trigger    models.CollectionImportTrigger `validate:"required" db:"trigger"`
}
type CreateImportCollectionLogOutput struct {
	Id int
}
type GetCollectiveCollectionInput struct {
	AccountId          *int                          `validate:"omitempty"`
	GameName           *string                       `validate:"omitempty"`
	Kind               models.Kind                   `validate:"required" db:"kind"`
	MaxAmountOfPlayers *int                          `validate:"omitempty"`
	Pagination         *models.PaginationInputString `validate:"omitempty"`
}
type GetCollectiveCollectionOutput struct {
	Data       []*models.GroupCollectionItem  `validate:"required"`
	Pagination *models.PaginationOutputString `validate:"required"`
}
type GetOngoingImportCollectionLogInput struct {
	ExternalIds []string        `validate:"required"`
	Provider    models.Provider `validate:"required" db:"provider"`
}
type GetOngoingImportCollectionLogOutput struct {
	Data []*models.ImportCollectionLog `validate:"required"`
}
type UpdateManyImportCollectionsLogsInput struct {
	Ids    []int                         `validate:"required"`
	Status models.CollectionImportStatus `validate:"required" db:"status"`
}

type CollectionRepository interface {
	AddToPersonalCollection(ctx context.Context, i *AddToPersonalCollectionInput) error
	CreateImportCollectionLog(ctx context.Context, i *CreateImportCollectionLogInput) (*CreateImportCollectionLogOutput, error)
	GetCollectiveCollection(ctx context.Context, i *GetCollectiveCollectionInput) (*GetCollectiveCollectionOutput, error)
	GetOngoingImportCollectionLog(ctx context.Context, i *GetOngoingImportCollectionLogInput) (*GetOngoingImportCollectionLogOutput, error)
	UpdateManyImportCollectionsLogs(ctx context.Context, i *UpdateManyImportCollectionsLogsInput) error
}
