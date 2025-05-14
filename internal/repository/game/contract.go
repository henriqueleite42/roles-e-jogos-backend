package game_repository

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

type CreateGameInput struct {
	AverageDuration    int
	Description        string      `validate:"max=1000"`
	IconPath           *string     `validate:"omitempty" db:"icon_path"`
	Kind               models.Kind `validate:"required" db:"kind"`
	LudopediaId        *int        `validate:"omitempty"`
	LudopediaUrl       *string     `validate:"omitempty,max=500"`
	MaxAmountOfPlayers int
	MinAge             int
	MinAmountOfPlayers int
	Name               string
}
type GetGameByLudopediaIdInput struct {
	LudopediaId int
}
type GetGamesListByLudopediaIdInput struct {
	LudopediaIds []int `validate:"required"`
}
type GetGamesListByLudopediaIdOutput struct {
	Data []*GetGamesListByLudopediaIdOutputDataItem `validate:"required"`
}
type GetGamesListByLudopediaIdOutputDataItem struct {
	ExternalId string
	GameId     int
}

type GameRepository interface {
	CreateGame(ctx context.Context, i *CreateGameInput) (*models.Game, error)
	GetGameByLudopediaId(ctx context.Context, i *GetGameByLudopediaIdInput) (*models.Game, error)
	GetGamesListByLudopediaId(ctx context.Context, i *GetGamesListByLudopediaIdInput) (*GetGamesListByLudopediaIdOutput, error)
}
