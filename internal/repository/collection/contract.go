package collection_repository

import (
	"context"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

type AddToPersonalCollectionInput struct {
	AccountId  int        `validate:"id" db:"id"`
	AcquiredAt *time.Time `validate:"omitempty"`
	Paid       *int       `validate:"omitempty"`
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

type CollectionRepository interface {
	AddToPersonalCollection(ctx context.Context, i *AddToPersonalCollectionInput) error
	GetCollectiveCollection(ctx context.Context, i *GetCollectiveCollectionInput) (*GetCollectiveCollectionOutput, error)
}
