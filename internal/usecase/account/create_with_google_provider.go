package account_usecase

import (
	"context"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *AccountUsecaseImplementation) CreateWithGoogleProvider(ctx context.Context, i *CreateWithGoogleProviderInput) (*models.SessionData, error) {
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
		session, err := self.AccountRepository.CreateSession(ctx, &account_repository.CreateSessionInput{
			AccountId: account.AccountId,
		})
		if err != nil {
			tx.Rollback(ctx)
			return nil, err
		}

		tx.Commit(ctx)
		return &models.SessionData{
			SessionId: session.SessionId,
		}, nil
	}

	var avatarPath *string
	if externalUserData.AvatarUrl != nil {
		path := fmt.Sprintf("avatars/%s.{{ext}}", self.IdAdapter.GenId())
		formattedPath, err := self.StorageAdapter.DownloadFromUrl(&adapters.DownloadFromUrlInput{
			Url:       *externalUserData.AvatarUrl,
			StorageId: self.SecretsAdapter.MediasS3BucketName,
			FileName:  path,
		})
		if err == nil {
			avatarPath = &formattedPath
		} else {
			self.Logger.Error().Err(err).Msg("fail to download avatar img")
		}
	}

	newAccount, err := self.AccountRepository.CreateAccountWithConnection(ctx, &account_repository.CreateAccountWithConnectionInput{
		Email:          externalUserData.Email,
		ExternalHandle: externalUserData.Handle,
		ExternalId:     externalUserData.Id,
		Handle:         genHandle(),
		Name:           &externalUserData.Name,
		AccessToken:    &exchangeResult.AccessToken,
		AvatarPath:     avatarPath,
		Provider:       models.Provider_Google,
		RefreshToken:   exchangeResult.RefreshToken,
	})
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	session, err := self.AccountRepository.CreateSession(ctx, &account_repository.CreateSessionInput{
		AccountId: newAccount.AccountId,
	})
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	tx.Commit(ctx)
	return &models.SessionData{
		SessionId: session.SessionId,
	}, nil
}
