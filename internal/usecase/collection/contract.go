package collection_usecase

import (
	"context"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

type AddToPersonalCollectionInput struct {
	AccountId  int        `validate:"id" db:"id"`
	AcquiredAt *time.Time
	Paid       *int
}
type GetCollectiveCollectionInput struct {
	Pagination *models.PaginationInputString
}
type GetCollectiveCollectionOutput struct {
	Data       []*models.GroupCollectionItem  `validate:"required"`
	Pagination *models.PaginationOutputString `validate:"required"`
}

type CollectionUsecase interface {
	AddToPersonalCollection(ctx context.Context, i *AddToPersonalCollectionInput) error
	GetCollectiveCollection(ctx context.Context, i *GetCollectiveCollectionInput) (*GetCollectiveCollectionOutput, error)
}
