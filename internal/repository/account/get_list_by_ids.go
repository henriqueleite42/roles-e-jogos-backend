package account_repository

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

func (self *accountRepositoryImplementation) GetListByIds(ctx context.Context, i *GetListByIdsInput) (*GetListByIdsOutput, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	ids := make([]int32, len(i.AccountsIds))
	for k, v := range i.AccountsIds {
		ids[k] = int32(v)
	}

	rows, err := db.GetAccountsListByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.AccountDataDb, len(rows))
	for k, v := range rows {
		accounts[k] = &models.AccountDataDb{
			AvatarPath: &v.AvatarPath.String,
		}
	}

	return &GetListByIdsOutput{
		Data: accounts,
	}, nil
}
