package models

import (
	"time"
)

type Kind string

const (
	Kind_Rpg       Kind = "RPG"
	Kind_Game      Kind = "GAME"
	Kind_Expansion Kind = "EXPANSION"
)

type Game struct {
	CreatedAt          time.Time `validate:"required" db:"created_at"`
	Description        string    `db:"description"`
	IconPath           *string   `db:"icon_path"`
	Id                 int       `db:"id"`
	Kind               Kind      `validate:"required" db:"kind"`
	LudopediaId        *int      `db:"ludopedia_id"`
	LudopediaUrl       *string   `db:"ludopedia_url"`
	MaxAmountOfPlayers int       `db:"max_amount_of_players"`
	MinAmountOfPlayers int       `db:"min_amount_of_players"`
	Name               string    `db:"name"`
}
