package account_repository

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/jackc/pgx/v5/pgtype"
)

func (self *accountRepositoryImplementation) GetProfilesListByHandle(ctx context.Context, i *GetProfilesListByHandleInput) (*GetProfilesListByHandleOutput, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := db.GetProfilesListByHandle(ctx, pgtype.Text{
		Valid:  true,
		String: i.Handle,
	})
	if err != nil {
		return nil, err
	}

	profiles := make([]*models.MinimumProfileData, len(rows))
	for k, v := range rows {
		var avatarUrl *string
		if v.AvatarPath.Valid {
			url := self.secretsAdapter.MediasCloudfrontUrl + v.AvatarPath.String
			avatarUrl = &url
		}

		profiles[k] = &models.MinimumProfileData{
			AccountId: int(v.ID),
			AvatarUrl: avatarUrl,
			Handle:    v.Handle,
		}
	}

	return &GetProfilesListByHandleOutput{
		Data: profiles,
	}, nil
}
