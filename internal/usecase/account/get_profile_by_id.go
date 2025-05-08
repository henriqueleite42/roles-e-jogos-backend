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

	connections, err := self.AccountRepository.GetConnectionsByAccountId(ctx, &account_repository.GetConnectionsByAccountIdInput{
		AccountId: account.Id,
	})
	if err != nil {
		return nil, err
	}

	secureConnectionsData := make([]*models.ProfileDataConnectionsItem, len(connections))
	for k, v := range connections {
		secureConnectionsData[k] = &models.ProfileDataConnectionsItem{
			ExternalHandle: v.ExternalHandle,
			ExternalId:     v.ExternalId,
			Provider:       v.Provider,
		}
	}

	return &models.ProfileData{
		AccountId:   account.Id,
		Connections: secureConnectionsData,
		AvatarUrl:   avatarUrl,
		Handle:      account.Handle,
		Name:        account.Name,
	}, nil
}
