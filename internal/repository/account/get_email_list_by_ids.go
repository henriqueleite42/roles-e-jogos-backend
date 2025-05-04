package account_repository

import (
	"context"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

func (self *accountRepositoryImplementation) GetEmailListByIds(ctx context.Context, i *GetEmailListByIdsInput) (*GetEmailListByIdsOutput, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	accountsIds := make([]int32, len(i.AccountsIds))
	for k, v := range i.AccountsIds {
		accountsIds[k] = int32(v)
	}

	if i.ValidatedOnly {
		result, err := db.GetValidatedEmailsByAccountsIds(ctx, accountsIds)
		if err != nil {
			return nil, err
		}

		emails := make([]*models.EmailAddress, len(result))

		for k, v := range result {
			var validatedAt *time.Time = nil
			if v.ValidatedAt.Valid {
				validatedAt = &v.ValidatedAt.Time
			}

			emails[k] = &models.EmailAddress{
				AccountId:    int(v.AccountID),
				CreatedAt:    v.CreatedAt.Time,
				EmailAddress: v.EmailAddress,
				ValidatedAt:  validatedAt,
			}
		}

		return &GetEmailListByIdsOutput{
			Data: emails,
		}, nil
	}

	result, err := db.GetEmailsByAccountsIds(ctx, accountsIds)
	if err != nil {
		return nil, err
	}

	emails := make([]*models.EmailAddress, len(result))

	for k, v := range result {
		var validatedAt *time.Time = nil
		if v.ValidatedAt.Valid {
			validatedAt = &v.ValidatedAt.Time
		}

		emails[k] = &models.EmailAddress{
			AccountId:    int(v.AccountID),
			CreatedAt:    v.CreatedAt.Time,
			EmailAddress: v.EmailAddress,
			ValidatedAt:  validatedAt,
		}
	}

	return &GetEmailListByIdsOutput{
		Data: emails,
	}, nil
}
