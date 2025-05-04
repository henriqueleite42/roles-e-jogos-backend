package models

import (
	"time"
)

type GroupCollectionItem struct {
	Game   *Game                            `validate:"required"`
	Owners []*GroupCollectionItemOwnersItem `validate:"required"`
}

type GroupCollectionItemOwnersItem struct {
	AccountId int
	AvatarUrl string
	Handle    string
}

type PersonalCollection struct {
	AccountId  int        `db:"account_id"`
	AcquiredAt *time.Time `db:"acquired_at"`
	CreatedAt  time.Time  `validate:"required" db:"created_at"`
	GameId     int        `db:"game_id"`
	Id         int        `db:"id"`
	Paid       *int       `db:"paid"`
}
