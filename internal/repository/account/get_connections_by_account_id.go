package account_repository

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
)

func (self *accountRepositoryImplementation) GetConnectionsByAccountId(ctx context.Context, i *GetConnectionsByAccountIdInput) ([]*models.Connection, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	result, err := db.GetConnectionsByAccountId(ctx, int32(i.AccountId))
	if err != nil {
		return nil, err
	}

	connections := make([]*models.Connection, len(result))
	for k, v := range result {
		var externalHandle *string
		if v.ExternalHandle.Valid {
			externalHandle = &v.ExternalHandle.String
		}

		var accessToken *string
		if v.AccessToken.Valid {
			accessToken = &v.AccessToken.String
		}

		var refreshToken *string
		if v.RefreshToken.Valid {
			refreshToken = &v.RefreshToken.String
		}

		connections[k] = &models.Connection{
			AccountId:      int(v.AccountID),
			CreatedAt:      v.CreatedAt.Time,
			AccessToken:    accessToken,
			RefreshToken:   refreshToken,
			ExternalHandle: externalHandle,
			ExternalId:     v.ExternalID,
			Provider:       models.Provider(v.Provider),
		}
	}

	return connections, nil
}
