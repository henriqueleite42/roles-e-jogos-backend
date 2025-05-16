package models

import (
	"time"
)

type PaginationInputId struct {
	After *int `validate:"omitempty"`
	Limit int  `validate:"min=1,max=100"`
}

type PaginationInputInt struct {
	After *int `validate:"omitempty"`
	Limit int  `validate:"min=1,max=100"`
}

type PaginationInputString struct {
	After *string `validate:"omitempty"`
	Limit int     `validate:"min=1,max=100"`
}

type PaginationInputTimestamp struct {
	After *time.Time `validate:"omitempty"`
	Limit int        `validate:"min=1,max=100"`
}

type PaginationOutputId struct {
	Current *int `validate:"omitempty"`
	Limit   int
	Next    *int `validate:"omitempty"`
}

type PaginationOutputInt struct {
	Current *int `validate:"omitempty"`
	Limit   int
	Next    *int `validate:"omitempty"`
}

type PaginationOutputString struct {
	Current *string `validate:"omitempty"`
	Limit   int
	Next    *string `validate:"omitempty"`
}

type PaginationOutputTimestamp struct {
	Current *time.Time `validate:"omitempty"`
	Limit   int
	Next    *time.Time `validate:"omitempty"`
}
