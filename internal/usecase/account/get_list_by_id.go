package account_usecase

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *AccountUsecaseImplementation) GetListById(ctx context.Context, i *GetListByIdInput) (*GetListByIdOutput, error) {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return nil, err
	}

	accounts, err := self.AccountRepository.GetListByIds(ctx, &account_repository.GetListByIdsInput{
		AccountsIds: i.AccountsIds,
	})
	if err != nil {
		return nil, err
	}

	result := &GetListByIdOutput{
		Data: make([]*models.AccountDataDb, len(accounts.Data)),
	}

	for k, v := range accounts.Data {
		result.Data[k] = &models.AccountDataDb{
			AvatarPath: v.AvatarPath,
			Handle:     v.Handle,
			AccountId:  v.AccountId,
		}
	}

	return result, nil
}
