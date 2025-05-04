package account_usecase

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *AccountUsecaseImplementation) GetProfileById(ctx context.Context, i *GetProfileByIdInput) (*models.ProfileData, error) {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return nil, err
	}

	account, err := self.AccountRepository.GetAccountById(ctx, &account_repository.GetAccountByIdInput{
		AccountId: i.AccountId,
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
