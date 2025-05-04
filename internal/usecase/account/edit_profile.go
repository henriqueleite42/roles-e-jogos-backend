package account_usecase

import (
	"context"

	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *AccountUsecaseImplementation) EditProfile(ctx context.Context, i *EditProfileInput) error {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return err
	}

	err = self.AccountRepository.EditProfile(ctx, &account_repository.EditProfileInput{
		AccountId: i.AccountId,
		Name:      i.Name,
	})
	if err != nil {
		return err
	}

	return nil
}
