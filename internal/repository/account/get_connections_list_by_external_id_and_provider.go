package account_repository

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
)

func (self *accountRepositoryImplementation) GetConnectionsListByExternalIdAndProvider(ctx context.Context, i *GetConnectionsListByExternalIdAndProviderInput) (*GetConnectionsListByExternalIdAndProviderOutput, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	result, err := db.GetConnectionsByExternalIdsAndProvider(ctx, queries.GetConnectionsByExternalIdsAndProviderParams{
		Column1:  i.ExternalIds,
		Provider: queries.ProviderEnum(i.Provider),
	})
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

	return &GetConnectionsListByExternalIdAndProviderOutput{
		Data: connections,
	}, nil
}
