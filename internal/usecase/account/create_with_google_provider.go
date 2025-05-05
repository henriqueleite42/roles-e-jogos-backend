package account_usecase

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *AccountUsecaseImplementation) CreateWithGoogleProvider(ctx context.Context, i *CreateWithGoogleProviderInput) (*models.AccountData, error) {
	exchangeResult, err := self.GoogleAdapter.ExchangeCode(&adapters.ExchangeCodeInput{
		Code: i.Code,
	})
	if err != nil {
		return nil, err
	}

	err = self.GoogleAdapter.CheckRequiredScopes(exchangeResult.Scopes)
	if err != nil {
		return nil, err
	}

	externalUserData, err := self.GoogleAdapter.GetUserData(exchangeResult.AccessToken)
	if err != nil {
		return nil, err
	}

	tx, ctx, err := utils.SetTxInCtx(self.Db, ctx)
	if err != nil {
		return nil, err
	}

	account, err := self.AccountRepository.GetAccountDataByEmailOrConnection(ctx, &account_repository.GetAccountDataByEmailOrConnectionInput{
		Email:      externalUserData.Email,
		ExternalId: externalUserData.Id,
		Provider:   models.Provider_Google,
	})
	if err != nil && err.Error() != "not found" {
		tx.Rollback(ctx)
		return nil, err
	}
	if account != nil {
		tx.Commit(ctx)
		return account, nil
	}

	newAccount, err := self.AccountRepository.CreateAccountWithConnection(ctx, &account_repository.CreateAccountWithConnectionInput{
		Email:          externalUserData.Email,
		ExternalHandle: externalUserData.Handle,
		ExternalId:     externalUserData.Id,
		Handle:         genHandle(),
		Name:           &externalUserData.Name,
		Provider:       models.Provider_Google,
		RefreshToken:   exchangeResult.RefreshToken,
	})
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	if externalUserData.AvatarUrl != nil {
		// Save image
	}

	tx.Commit(ctx)
	return newAccount, nil
}
