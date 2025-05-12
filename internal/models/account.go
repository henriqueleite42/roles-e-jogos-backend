package models

import (
	"time"
)

type OtpPurpose string

const (
	OtpPurpose_SignIn OtpPurpose = "SIGN_IN"
)

type Provider string

const (
	Provider_Google    Provider = "GOOGLE"
	Provider_Ludopedia Provider = "LUDOPEDIA"
)

type AccountData struct {
	AccountId int
	IsAdmin   bool
}

type AccountDataDb struct {
	AccountId  int
	AvatarPath *string `validate:"omitempty"`
	Handle     string
	IsAdmin    bool
}

type ProfileData struct {
	AccountId   int
	AvatarUrl   *string                       `validate:"omitempty"`
	Connections []*ProfileDataConnectionsItem `validate:"required"`
	Handle      string
	Name        *string                       `validate:"omitempty"`
}

type ProfileDataConnectionsItem struct {
	ExternalHandle *string  `validate:"omitempty"`
	ExternalId     string
	Provider       Provider `validate:"required" db:"provider"`
}

type SessionData struct {
	SessionId string
}

type Account struct {
	AvatarPath *string   `validate:"omitempty" db:"avatar_path"`
	CreatedAt  time.Time `validate:"required" db:"created_at"`
	Handle     string    `db:"handle"`
	Id         int       `validate:"id" db:"id"`
	IsAdmin    bool      `db:"is_admin"`
	Name       *string   `validate:"omitempty" db:"name"`
}

type Connection struct {
	AccessToken    *string   `validate:"omitempty" db:"access_token"`
	AccountId      int       `db:"account_id"`
	CreatedAt      time.Time `validate:"required" db:"created_at"`
	ExternalHandle *string   `validate:"omitempty" db:"external_handle"`
	ExternalId     string    `db:"external_id"`
	Provider       Provider  `validate:"required" db:"provider"`
	RefreshToken   *string   `validate:"omitempty" db:"refresh_token"`
}

type EmailAddress struct {
	AccountId    int        `db:"account_id"`
	CreatedAt    time.Time  `validate:"required" db:"created_at"`
	EmailAddress string     `validate:"email" db:"email_address"`
	ValidatedAt  *time.Time `validate:"omitempty" db:"validated_at"`
}

type OneTimePassword struct {
	AccountId int        `db:"account_id"`
	Code      string     `db:"code"`
	CreatedAt time.Time  `validate:"required" db:"created_at"`
	Purpose   OtpPurpose `validate:"required" db:"purpose"`
}

type Session struct {
	AccountId int       `db:"account_id"`
	CreatedAt time.Time `validate:"required" db:"created_at"`
	SessionId string    `db:"session_id"`
}
