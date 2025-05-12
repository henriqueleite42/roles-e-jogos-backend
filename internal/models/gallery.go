package models

import (
	"time"
)

type Media struct {
	CreatedAt   time.Time `validate:"required" db:"created_at"`
	Date        time.Time `validate:"required" db:"date"`
	Description *string   `validate:"omitempty" db:"description"`
	GameId      *int      `validate:"omitempty" db:"game_id"`
	Id          int       `db:"id"`
	OwnerId     int       `db:"owner_id"`
	Path        string    `db:"path"`
}
