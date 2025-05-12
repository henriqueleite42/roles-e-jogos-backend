package collection_usecase

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

type CollectionUsecase interface {
	AddToPersonalCollection(ctx context.Context, i *AddToPersonalCollectionInput) error
	GetCollectiveCollection(ctx context.Context, i *GetCollectiveCollectionInput) (*GetCollectiveCollectionOutput, error)
}
