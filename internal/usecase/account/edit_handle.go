package account_usecase

import (
	"context"

	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *AccountUsecaseImplementation) EditHandle(ctx context.Context, i *EditHandleInput) error {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return err
	}

	err = self.AccountRepository.EditAccountHandle(ctx, &account_repository.EditAccountHandleInput{
		AccountId: i.AccountId,
		Handle:    i.NewHandle,
	})
	if err != nil {
		return err
	}

	return nil
}
