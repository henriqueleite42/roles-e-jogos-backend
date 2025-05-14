package account_repository

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5/pgtype"
)

func (self *accountRepositoryImplementation) LinkConnectionWithAccount(ctx context.Context, i *LinkConnectionWithAccountInput) error {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return err
	}

	var externalHandle pgtype.Text
	if i.ExternalHandle != nil {
		externalHandle = pgtype.Text{
			Valid:  true,
			String: *i.ExternalHandle,
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

	accountIdInt32 := int32(i.AccountId)

	err = db.CreateConnection(ctx, queries.CreateConnectionParams{
		AccountID:      accountIdInt32,
		ExternalHandle: externalHandle,
		ExternalID:     i.ExternalId,
		Provider:       queries.ProviderEnum(i.Provider),
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
	})
	if err != nil {
		return err
	}

	err = db.CreateValidatedEmailAddress(ctx, queries.CreateValidatedEmailAddressParams{
		AccountID:    accountIdInt32,
		EmailAddress: i.Email,
	})
	if err != nil {
		return err
	}

	return nil
}
