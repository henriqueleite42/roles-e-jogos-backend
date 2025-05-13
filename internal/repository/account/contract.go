package account_repository

import (
	"context"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

type CreateAccountWithConnectionInput struct {
	AccessToken    *string         `validate:"omitempty"`
	AvatarPath     *string         `validate:"omitempty"`
	Email          string
	ExternalHandle *string         `validate:"omitempty"`
	ExternalId     string
	Handle         string
	Name           *string         `validate:"omitempty"`
	Provider       models.Provider `validate:"required" db:"provider"`
	RefreshToken   *string         `validate:"omitempty"`
}
type CreateAccountWithEmailInput struct {
	Email  string
	Handle string
}
type CreateOtpInput struct {
	AccountId int
	Code      string
	Purpose   models.OtpPurpose `validate:"required" db:"purpose"`
}
type CreateSessionInput struct {
	AccountId int
}
type EditAccountHandleInput struct {
	AccountId int
	Handle    string
}
type EditProfileInput struct {
	AccountId int
	Name      *string `validate:"omitempty"`
}
type GetAccountByHandleInput struct {
	Handle string
}
type GetAccountByIdInput struct {
	AccountId int
}
type GetAccountDataByConnectionInput struct {
	ExternalId string
	Provider   models.Provider `validate:"required" db:"provider"`
}
type GetAccountDataByEmailInput struct {
	Email string
}
type GetAccountDataByEmailOrConnectionInput struct {
	Email      string
	ExternalId string
	Provider   models.Provider `validate:"required" db:"provider"`
}
type GetAccountDataByHandleInput struct {
	Handle string
}
type GetAccountDataByIdInput struct {
	AccountId int
}
type GetAccountDataBySessionIdInput struct {
	SessionId string
}
type GetConnectionInput struct {
	ExternalId string
	Provider   models.Provider `validate:"required" db:"provider"`
}
type GetConnectionsByAccountIdAndProviderInput struct {
	AccountId int
	Provider  models.Provider `validate:"required" db:"provider"`
}
type GetConnectionsByAccountIdInput struct {
	AccountId int
}
type GetEmailListByIdsInput struct {
	AccountsIds   []int `validate:"required"`
	ValidatedOnly bool
}
type GetEmailListByIdsOutput struct {
	Data []*models.EmailAddress `validate:"required"`
}
type GetListByIdsInput struct {
	AccountsIds []int `validate:"required"`
}
type GetListByIdsOutput struct {
	Data []*models.AccountDataDb `validate:"required"`
}
type GetOtpInput struct {
	AccountId int
	Code      string
	Purpose   models.OtpPurpose `validate:"required" db:"purpose"`
}
type GetOtpOutput struct {
	CreatedAt time.Time `validate:"required"`
}
type GetProfilesListByHandleInput struct {
	Handle string
}
type GetProfilesListByHandleOutput struct {
	Data []*models.MinimumProfileData `validate:"required"`
}
type LinkConnectionWithAccountInput struct {
	AccountId      int
	Email          string
	ExternalHandle *string         `validate:"omitempty"`
	ExternalId     string
	Provider       models.Provider `validate:"required" db:"provider"`
	RefreshToken   *string         `validate:"omitempty"`
}

type AccountRepository interface {
	CreateAccountWithConnection(ctx context.Context, i *CreateAccountWithConnectionInput) (*models.AccountData, error)
	CreateAccountWithEmail(ctx context.Context, i *CreateAccountWithEmailInput) (*models.AccountData, error)
	CreateOtp(ctx context.Context, i *CreateOtpInput) error
	CreateSession(ctx context.Context, i *CreateSessionInput) (*models.Session, error)
	EditAccountHandle(ctx context.Context, i *EditAccountHandleInput) error
	EditProfile(ctx context.Context, i *EditProfileInput) error
	GetAccountByHandle(ctx context.Context, i *GetAccountByHandleInput) (*models.Account, error)
	GetAccountById(ctx context.Context, i *GetAccountByIdInput) (*models.Account, error)
	GetAccountDataByConnection(ctx context.Context, i *GetAccountDataByConnectionInput) (*models.AccountData, error)
	GetAccountDataByEmail(ctx context.Context, i *GetAccountDataByEmailInput) (*models.AccountData, error)
	GetAccountDataByEmailOrConnection(ctx context.Context, i *GetAccountDataByEmailOrConnectionInput) (*models.AccountData, error)
	GetAccountDataByHandle(ctx context.Context, i *GetAccountDataByHandleInput) (*models.AccountData, error)
	GetAccountDataById(ctx context.Context, i *GetAccountDataByIdInput) (*models.AccountData, error)
	GetAccountDataBySessionId(ctx context.Context, i *GetAccountDataBySessionIdInput) (*models.AccountDataDb, error)
	GetConnection(ctx context.Context, i *GetConnectionInput) (*models.Connection, error)
	GetConnectionsByAccountId(ctx context.Context, i *GetConnectionsByAccountIdInput) ([]*models.Connection, error)
	GetConnectionsByAccountIdAndProvider(ctx context.Context, i *GetConnectionsByAccountIdAndProviderInput) ([]*models.Connection, error)
	GetEmailListByIds(ctx context.Context, i *GetEmailListByIdsInput) (*GetEmailListByIdsOutput, error)
	GetListByIds(ctx context.Context, i *GetListByIdsInput) (*GetListByIdsOutput, error)
	GetOtp(ctx context.Context, i *GetOtpInput) (*GetOtpOutput, error)
	GetProfilesListByHandle(ctx context.Context, i *GetProfilesListByHandleInput) (*GetProfilesListByHandleOutput, error)
	LinkConnectionWithAccount(ctx context.Context, i *LinkConnectionWithAccountInput) error
}
