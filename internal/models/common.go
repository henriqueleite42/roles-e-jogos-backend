package models

import (
	"time"
)

type PaginationInputId struct {
	After *int
	Limit *int `validate:"min=1,max=100"`
}

type PaginationInputInt struct {
	After *int
	Limit *int `validate:"min=1,max=100"`
}

type PaginationInputString struct {
	After *string
	Limit *int    `validate:"min=1,max=100"`
}

type PaginationInputTimestamp struct {
	After *time.Time
	Limit *int       `validate:"min=1,max=100"`
}

type PaginationOutputId struct {
	Limit    int
	Next     *int
	Previous *int
}

type PaginationOutputInt struct {
	Limit    int
	Next     *int
	Previous *int
}

type PaginationOutputString struct {
	Limit    int
	Next     *string
	Previous *string
}

type PaginationOutputTimestamp struct {
	Limit    int
	Next     *time.Time
	Previous *time.Time
}
