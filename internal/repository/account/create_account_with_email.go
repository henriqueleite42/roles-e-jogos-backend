package account_repository

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
)

func (self *accountRepositoryImplementation) CreateAccountWithEmail(ctx context.Context, i *CreateAccountWithEmailInput) (*models.AccountData, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	accountIdInt32, err := db.CreateAccount(ctx, queries.CreateAccountParams{
		Handle: i.Handle,
	})
	if err != nil {
		return nil, err
	}
	accountId := int(accountIdInt32)

	err = db.CreateValidatedEmailAddress(ctx, queries.CreateValidatedEmailAddressParams{
		AccountID:    accountIdInt32,
		EmailAddress: i.Email,
	})
	if err != nil {
		return nil, err
	}

	return &models.AccountData{
		AccountId: accountId,
		IsAdmin:   false,
	}, nil
}
