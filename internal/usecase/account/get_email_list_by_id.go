package account_usecase

import (
	"context"

	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *AccountUsecaseImplementation) GetEmailListById(ctx context.Context, i *GetEmailListByIdInput) (*GetEmailListByIdOutput, error) {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return nil, err
	}

	emails, err := self.AccountRepository.GetEmailListByIds(ctx, &account_repository.GetEmailListByIdsInput{
		AccountsIds:   i.AccountsIds,
		ValidatedOnly: true,
	})
	if err != nil {
		return nil, err
	}

	return &GetEmailListByIdOutput{
		Data: emails.Data,
	}, nil
}
