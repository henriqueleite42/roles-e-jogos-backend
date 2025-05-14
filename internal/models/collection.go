package models

import (
	"time"
)

type CollectionImportStatus string

const (
	CollectionImportStatus_Started       CollectionImportStatus = "STARTED"
	CollectionImportStatus_Completed     CollectionImportStatus = "COMPLETED"
	CollectionImportStatus_Failed        CollectionImportStatus = "FAILED"
	CollectionImportStatus_NotYetStarted CollectionImportStatus = "NOT_YET_STARTED"
)

type CollectionImportTrigger string

const (
	CollectionImportTrigger_ManualByUser CollectionImportTrigger = "MANUAL_BY_USER"
	CollectionImportTrigger_AccountLink  CollectionImportTrigger = "ACCOUNT_LINK"
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

type ImportCollectionLog struct {
	AccountId  int                     `db:"account_id"`
	CreatedAt  time.Time               `validate:"required" db:"created_at"`
	EndedAt    *time.Time              `validate:"omitempty" db:"ended_at"`
	ExternalId string                  `db:"external_id"`
	Id         int                     `db:"id"`
	Provider   Provider                `validate:"required" db:"provider"`
	Status     CollectionImportStatus  `validate:"required" db:"status"`
	Trigger    CollectionImportTrigger `validate:"required" db:"trigger"`
}

type PersonalCollection struct {
	AccountId  int        `db:"account_id"`
	AcquiredAt *time.Time `validate:"omitempty" db:"acquired_at"`
	CreatedAt  time.Time  `validate:"required" db:"created_at"`
	GameId     int        `db:"game_id"`
	Id         int        `db:"id"`
	Paid       *int       `validate:"omitempty" db:"paid"`
}

type ImportCollectionEvent struct {
	AccountId  int                     `validate:"id"`
	ExternalId string
	Trigger    CollectionImportTrigger `validate:"required" db:"trigger"`
}
