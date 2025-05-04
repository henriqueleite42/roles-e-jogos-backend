package account_usecase

import (
	"context"

	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *AccountUsecaseImplementation) CheckHandle(ctx context.Context, i *CheckHandleInput) (*CheckHandleOutput, error) {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return nil, err
	}

	_, err = self.AccountRepository.GetAccountByHandle(ctx, &account_repository.GetAccountByHandleInput{
		Handle: i.Handle,
	})
	if err != nil {
		if err.Error() == "not found" {
			return &CheckHandleOutput{
				Available: true,
			}, nil
		}

		self.Logger.Error().Err(err).Msg("fail to get account by handle")
		return &CheckHandleOutput{
			Available: false,
		}, nil
	}

	return &CheckHandleOutput{
		Available: false,
	}, nil
}
