package collection_repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5/pgtype"
)

type DbOwner struct {
	AccountId  int     `json:"account_id"`
	Handle     string  `json:"handle"`
	AvatarPath *string `json:"avatar_path"`
}

func (self *collectionRepositoryImplementation) GetCollectiveCollection(ctx context.Context, i *GetCollectiveCollectionInput) (*GetCollectiveCollectionOutput, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	gameName := pgtype.Text{Valid: true}
	if i.GameName != nil {
		gameName.String = *i.GameName
	}
	accountId := pgtype.Int4{Valid: true}
	if i.AccountId != nil {
		accountId.Int32 = int32(*i.AccountId)
	}
	maxAmountOfPlayers := pgtype.Int4{Valid: true}
	if i.MaxAmountOfPlayers != nil {
		maxAmountOfPlayers.Int32 = int32(*i.MaxAmountOfPlayers)
	}
	after := pgtype.Text{Valid: true}
	if i.Pagination.After != nil {
		after.String = *i.Pagination.After
	}

	rows, err := db.GetCollectiveCollection(ctx, queries.GetCollectiveCollectionParams{
		Kind:    queries.KindEnum(i.Kind),
		Column2: gameName,
		Column3: accountId,
		Column4: maxAmountOfPlayers,
		Column5: after,
		Limit:   int32(*i.Pagination.Limit),
	})
	if err != nil {
		return nil, err
	}

	games := make([]*models.GroupCollectionItem, len(rows))
	for k, v := range rows {
		ownerJsonParsed := []*DbOwner{}
		err := json.Unmarshal(v.Owners, &ownerJsonParsed)
		if err != nil {
			return nil, fmt.Errorf("fail to decode owners json")
		}

		owners := make([]*models.MinimumProfileData, len(ownerJsonParsed))
		for kk, vv := range ownerJsonParsed {
			var avatarUrl *string
			if vv.AvatarPath != nil {
				url := self.secretsAdapter.MediasCloudfrontUrl + *vv.AvatarPath
				avatarUrl = &url
			}

			owners[kk] = &models.MinimumProfileData{
				AccountId: vv.AccountId,
				AvatarUrl: avatarUrl,
				Handle:    vv.Handle,
			}
		}

		var iconUrl *string
		if v.IconPath.Valid {
			str := self.secretsAdapter.MediasCloudfrontUrl + v.IconPath.String
			iconUrl = &str
		}

		games[k] = &models.GroupCollectionItem{
			Game: &models.GroupCollectionItemGame{
				IconUrl:            iconUrl,
				Id:                 int(v.GameID),
				LudopediaUrl:       v.LudopediaUrl.String,
				MaxAmountOfPlayers: int(v.MaxAmountOfPlayers),
				MinAmountOfPlayers: int(v.MinAmountOfPlayers),
				Name:               v.Name,
			},
			Owners: owners,
		}
	}

	var next *string
	if len(games) > 0 {
		next = &games[len(games)-1].Game.Name
	}

	return &GetCollectiveCollectionOutput{
		Data: games,
		Pagination: &models.PaginationOutputString{
			Limit:   *i.Pagination.Limit,
			Next:    next,
			Current: i.Pagination.After,
		},
	}, nil
}
