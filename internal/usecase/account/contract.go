package account_usecase

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

type CheckHandleInput struct {
	Handle string `validate:"handle"`
}
type CheckHandleOutput struct {
	Available bool
}
type CreateWithGoogleProviderInput struct {
	Code string `validate:"min=1"`
}
type CreateWithLudopediaProviderInput struct {
	Code string `validate:"min=1"`
}
type EditHandleInput struct {
	AccountId int    `validate:"id"`
	NewHandle string `validate:"handle"`
}
type EditProfileInput struct {
	AccountId int     `validate:"id"`
	Name      *string `validate:"max=24"`
}
type ExchangeSignInOtpInput struct {
	AccountId int    `validate:"id"`
	Otp       string
}
type GetEmailListByIdInput struct {
	AccountsIds []int `validate:"gt=0,dive,min=1,required"`
}
type GetEmailListByIdOutput struct {
	Data []*models.EmailAddress `validate:"required"`
}
type GetListByIdInput struct {
	AccountsIds []int `validate:"gt=0,dive,min=1,required"`
}
type GetListByIdOutput struct {
	Data []*models.AccountDataDb `validate:"required"`
}
type GetProfileByHandleInput struct {
	Handle string `validate:"handle"`
}
type GetProfileByIdInput struct {
	AccountId int `validate:"id"`
}
type SendSignInOtpInput struct {
	Email string `validate:"email"`
}

type AccountUsecase interface {
	CheckHandle(ctx context.Context, i *CheckHandleInput) (*CheckHandleOutput, error)
	CreateWithGoogleProvider(ctx context.Context, i *CreateWithGoogleProviderInput) (*models.AccountData, error)
	CreateWithLudopediaProvider(ctx context.Context, i *CreateWithLudopediaProviderInput) (*models.AccountData, error)
	EditHandle(ctx context.Context, i *EditHandleInput) error
	EditProfile(ctx context.Context, i *EditProfileInput) error
	ExchangeSignInOtp(ctx context.Context, i *ExchangeSignInOtpInput) (*models.AccountData, error)
	GetEmailListById(ctx context.Context, i *GetEmailListByIdInput) (*GetEmailListByIdOutput, error)
	GetListById(ctx context.Context, i *GetListByIdInput) (*GetListByIdOutput, error)
	GetProfileByHandle(ctx context.Context, i *GetProfileByHandleInput) (*models.ProfileData, error)
	GetProfileById(ctx context.Context, i *GetProfileByIdInput) (*models.ProfileData, error)
	SendSignInOtp(ctx context.Context, i *SendSignInOtpInput) error
}
