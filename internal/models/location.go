package models

import (
	"time"
)

type LocationKind string

const (
	LocationKind_Business LocationKind = "BUSINESS"
	LocationKind_Personal LocationKind = "PERSONAL"
)

type Location struct {
	Address   string       `db:"address"`
	CreatedAt time.Time    `validate:"required" db:"created_at"`
	CreatedBy int          `db:"created_by"`
	IconPath  *string      `validate:"omitempty" db:"icon_path"`
	Id        int          `db:"id"`
	Kind      LocationKind `validate:"required" db:"kind"`
	Name      string       `db:"name"`
}
