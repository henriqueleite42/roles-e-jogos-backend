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
	EventConfidentiality_OnlyInvited EventConfidentiality = "ONLY_INVITED"
)

type EventData struct {
	Attendances []*EventDataAttendancesItem `validate:"required"`
	Event       *Event                      `validate:"required"`
	Games       []*Game                     `validate:"required"`
}

type EventDataAttendancesItem struct {
	AccountId int
	AvatarUrl string
	Handle    string
}

type Event struct {
	CreatedAt          time.Time `validate:"required" db:"created_at"`
	Date               time.Time `validate:"required" db:"date"`
	Description        string    `db:"description"`
	IconPath           *string   `db:"icon_path"`
	Id                 int       `db:"id"`
	LocationAddress    string    `db:"location_address"`
	LocationName       string    `db:"location_name"`
	MaxAmountOfPlayers *int      `db:"max_amount_of_players"`
	Name               string    `db:"name"`
	OwnerId            int       `db:"owner_id"`
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
}
