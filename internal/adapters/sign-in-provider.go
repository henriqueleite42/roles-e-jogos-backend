package adapters

import "time"

type ExchangeCodeInput struct {
	Code string
}

type ExchangeCodeOutput struct {
	Scopes       []string
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

type GetAuthenticatedUserDataOutput struct {
	Id              string
	Handle          *string
	Name            string
	Email           string
	AvatarUrl       *string
	BannerUrl       *string
	IsEmailVerified bool
}

type SignInProvider interface {
	ExchangeCode(i *ExchangeCodeInput) (*ExchangeCodeOutput, error)

	GetUserData(accessToken string) (*GetAuthenticatedUserDataOutput, error)

	CheckRequiredScopes(scopes []string) error
}
