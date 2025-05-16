package models

import (
	"time"
)

type EventAttendanceStatus string

const (
	EventAttendanceStatus_Going    EventAttendanceStatus = "GOING"
	EventAttendanceStatus_Maybe    EventAttendanceStatus = "MAYBE"
	EventAttendanceStatus_NotGoing EventAttendanceStatus = "NOT_GOING"
)

type EventConfidentiality string

const (
	EventConfidentiality_Public      EventConfidentiality = "PUBLIC"
	EventConfidentiality_InvitedOnly EventConfidentiality = "INVITED_ONLY"
)

type EventData struct {
	Attendances []*EventDataAttendancesItem `validate:"required"`
	Capacity    *int                        `validate:"omitempty"`
	Description string
	EndDate     *time.Time                  `validate:"omitempty"`
	Games       []*EventDataGamesItem       `validate:"required"`
	IconUrl     *string                     `validate:"omitempty"`
	Id          int
	Location    *EventDataLocation          `validate:"required"`
	Name        string
	OwnerId     int
	StartDate   time.Time                   `validate:"required"`
}

type EventDataAttendancesItem struct {
	AccountId int
	AvatarUrl *string               `validate:"omitempty"`
	Handle    string
	Status    EventAttendanceStatus `validate:"required" db:"status"`
}

type EventDataGamesItem struct {
	AverageDuration    int
	IconUrl            *string `validate:"omitempty"`
	Id                 int
	Kind               Kind    `validate:"required" db:"kind"`
	LudopediaUrl       *string `validate:"omitempty"`
	MaxAmountOfPlayers int
	MinAge             int
	MinAmountOfPlayers int
	Name               string
}

type EventDataLocation struct {
	Address string
	IconUrl *string `validate:"omitempty"`
	Id      int
	Name    string
}

type Event struct {
	Capacity    *int       `validate:"omitempty" db:"capacity"`
	CreatedAt   time.Time  `validate:"required" db:"created_at"`
	Description string     `db:"description"`
	EndDate     *time.Time `validate:"omitempty" db:"end_date"`
	IconPath    *string    `validate:"omitempty" db:"icon_path"`
	Id          int        `db:"id"`
	LocationId  int        `db:"location_id"`
	Name        string     `db:"name"`
	OwnerId     int        `db:"owner_id"`
	StartDate   time.Time  `validate:"required" db:"start_date"`
}

type EventAttendance struct {
	AccountId   int                   `db:"account_id"`
	ConfirmedAt time.Time             `validate:"required" db:"confirmed_at"`
	EventId     int                   `db:"event_id"`
	Id          int                   `db:"id"`
	Status      EventAttendanceStatus `validate:"required" db:"status"`
}

type EventGame struct {
	EventId int `db:"event_id"`
	GameId  int `db:"game_id"`
	Id      int `db:"id"`
	OwnerId int `db:"owner_id"`
}
