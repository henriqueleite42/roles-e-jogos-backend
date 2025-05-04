package models

import (
	"time"
)

type EventConfirmationStatus string

const (
	EventConfirmationStatus_Going    EventConfirmationStatus = "GOING"
	EventConfirmationStatus_Maybe    EventConfirmationStatus = "MAYBE"
	EventConfirmationStatus_NotGoing EventConfirmationStatus = "NOT_GOING"
)

type Media struct {
	CreatedAt   time.Time `validate:"required" db:"created_at"`
	Date        time.Time `validate:"required" db:"date"`
	Description *string   `db:"description"`
	GameId      *int      `db:"game_id"`
	Id          int       `db:"id"`
	OwnerId     int       `db:"owner_id"`
	Path        string    `db:"path"`
}
