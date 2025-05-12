package account_usecase

import (
	"context"

	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *AccountUsecaseImplementation) GetProfilesListByHandle(ctx context.Context, i *GetProfilesListByHandleInput) (*GetProfilesListByHandleOutput, error) {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return nil, err
	}

	accounts, err := self.AccountRepository.GetProfilesListByHandle(ctx, &account_repository.GetProfilesListByHandleInput{
		Handle: i.Handle,
	})
	if err != nil {
		return nil, err
	}

	return &GetProfilesListByHandleOutput{
		Data: accounts.Data,
	}, nil
}
