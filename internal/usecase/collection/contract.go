package collection_usecase

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
type GetCollectiveCollectionInput struct {
	AccountId          *int                          `validate:"omitempty,id"`
	GameName           *string                       `validate:"omitempty,min=1,max=128"`
	Kind               models.Kind                   `validate:"required" db:"kind"`
	MaxAmountOfPlayers *int                          `validate:"omitempty,min=1,max=99"`
	Pagination         *models.PaginationInputString `validate:"omitempty"`
}
type GetCollectiveCollectionOutput struct {
	Data       []*models.GroupCollectionItem  `validate:"required"`
	Pagination *models.PaginationOutputString `validate:"required"`
}
type GetLatestImportCollectionLogStatusInput struct {
	AccountId  int             `validate:"id"`
	ExternalId string
	Provider   models.Provider `validate:"required" db:"provider"`
}
type GetLatestImportCollectionLogStatusOutput struct {
	Status models.CollectionImportStatus `validate:"required" db:"status"`
}
type RequestImportPersonalCollectionFromLudopediaInput struct {
	AccountId   int `validate:"id"`
	LudopediaId int
}

type CollectionUsecase interface {
	AddToPersonalCollection(ctx context.Context, i *AddToPersonalCollectionInput) error
	GetCollectiveCollection(ctx context.Context, i *GetCollectiveCollectionInput) (*GetCollectiveCollectionOutput, error)
	GetLatestImportCollectionLogStatus(ctx context.Context, i *GetLatestImportCollectionLogStatusInput) (*GetLatestImportCollectionLogStatusOutput, error)
	ImportPersonalCollectionFromLudopedia(ctx context.Context, i []*models.ImportCollectionEvent) error
	RequestImportPersonalCollectionFromLudopedia(ctx context.Context, i *RequestImportPersonalCollectionFromLudopediaInput) error
}
