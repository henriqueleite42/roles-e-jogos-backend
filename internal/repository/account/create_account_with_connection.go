package account_repository

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5/pgtype"
)

func (self *accountRepositoryImplementation) CreateAccountWithConnection(ctx context.Context, i *CreateAccountWithConnectionInput) (*models.AccountData, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	var name pgtype.Text
	if i.Name != nil {
		name = pgtype.Text{
			Valid:  true,
			String: *i.Name,
		}
	}
	var avatarPath pgtype.Text
	if i.AvatarPath != nil {
		avatarPath = pgtype.Text{
			Valid:  true,
			String: *i.AvatarPath,
		}
	}
	var accessToken pgtype.Text
	if i.AccessToken != nil {
		accessToken = pgtype.Text{
			Valid:  true,
			String: *i.AccessToken,
		}
	}
	var refreshToken pgtype.Text
	if i.RefreshToken != nil {
		refreshToken = pgtype.Text{
			Valid:  true,
			String: *i.RefreshToken,
		}
	}
	var externalHandle pgtype.Text
	if i.ExternalHandle != nil {
		externalHandle = pgtype.Text{
			Valid:  true,
			String: *i.ExternalHandle,
		}
	}
	if i.ExternalHandle == nil && i.Name != nil {
		externalHandle = pgtype.Text{
			Valid:  true,
			String: *i.Name,
		}
	}

	accountIdInt32, err := db.CreateAccount(ctx, queries.CreateAccountParams{
		Handle:     i.Handle,
		Name:       name,
		AvatarPath: avatarPath,
	})
	if err != nil {
		return nil, err
	}
	accountId := int(accountIdInt32)

	err = db.CreateConnection(ctx, queries.CreateConnectionParams{
		AccountID:      accountIdInt32,
		ExternalHandle: externalHandle,
		ExternalID:     i.ExternalId,
		Provider:       queries.ProviderEnum(i.Provider),
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
	})
	if err != nil {
		return nil, err
	}

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
