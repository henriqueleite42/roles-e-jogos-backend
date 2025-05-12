package models

import (
	"time"
)

type GroupCollectionItem struct {
	Game   *GroupCollectionItemGame `validate:"required"`
	Owners []*MinimumProfileData    `validate:"required"`
}

type GroupCollectionItemGame struct {
	IconUrl            *string `validate:"omitempty"`
	Id                 int
	LudopediaUrl       string
	MaxAmountOfPlayers int
	MinAmountOfPlayers int
	Name               string
}

type PersonalCollection struct {
	AccountId  int        `db:"account_id"`
	AcquiredAt *time.Time `validate:"omitempty" db:"acquired_at"`
	CreatedAt  time.Time  `validate:"required" db:"created_at"`
	GameId     int        `db:"game_id"`
	Id         int        `db:"id"`
	Paid       *int       `validate:"omitempty" db:"paid"`
}
