package account_usecase

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *AccountUsecaseImplementation) GetProfileByHandle(ctx context.Context, i *GetProfileByHandleInput) (*models.ProfileData, error) {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return nil, err
	}

	account, err := self.AccountRepository.GetAccountByHandle(ctx, &account_repository.GetAccountByHandleInput{
		Handle: i.Handle,
	})
	if err != nil {
		return nil, err
	}

	var avatarUrl *string
	if account.AvatarPath != nil {
		str := self.SecretsAdapter.MediasCloudfrontUrl + *account.AvatarPath
		avatarUrl = &str
	}

	return &models.ProfileData{
		AvatarUrl: avatarUrl,
		Handle:    account.Handle,
		Name:      account.Name,
	}, nil
}
