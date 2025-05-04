package game_usecase

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

type CreateGameInput struct {
	Description        string      `validate:"max=1000"`
	IconPath           *string     `db:"icon_path"`
	Kind               models.Kind `validate:"required" db:"kind"`
	LudopediaId        *int
	LudopediaUrl       *string     `validate:"max=500"`
	MaxAmountOfPlayers int
	MinAmountOfPlayers int
	Name               string
}

type GameUsecase interface {
	CreateGame(ctx context.Context, i *CreateGameInput) error
}
