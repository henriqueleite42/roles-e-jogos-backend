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

func getOwners(ownersSource []byte) ([]*models.GroupCollectionItemOwnersItem, error) {
	ownerJsonParsed := []*DbOwner{}
	err := json.Unmarshal(ownersSource, &ownerJsonParsed)
	if err != nil {
		return nil, fmt.Errorf("fail to decode owners json")
	}

	owners := make([]*models.GroupCollectionItemOwnersItem, len(ownerJsonParsed))
	for kk, vv := range ownerJsonParsed {
		owners[kk] = &models.GroupCollectionItemOwnersItem{
			AccountId: vv.AccountId,
			AvatarUrl: *vv.AvatarPath,
			Handle:    vv.Handle,
		}
	}

	return owners, nil
}

func (self *collectionRepositoryImplementation) getIconUrl(iconPath pgtype.Text) *string {
	var iconUrl *string
	if iconPath.Valid {
		str := self.secretsAdapter.MediasCloudfrontUrl + iconPath.String
		iconUrl = &str
	}
	return iconUrl
}

func getNext(games []*models.GroupCollectionItem) *string {
	if len(games) > 0 {
		return &games[len(games)-1].Game.Name
	}

	return nil
}

func (self *collectionRepositoryImplementation) GetCollectiveCollection(ctx context.Context, i *GetCollectiveCollectionInput) (*GetCollectiveCollectionOutput, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	games := []*models.GroupCollectionItem{}

	after := ""
	if i.Pagination.After != nil {
		after = *i.Pagination.After
	}

	if i.AccountId != nil && i.GameName != nil && i.MaxAmountOfPlayers != nil {
		rows, err := db.GetCollectiveCollectionAllFilters(ctx, queries.GetCollectiveCollectionAllFiltersParams{
			Kind:               queries.KindEnum(i.Kind),
			AccountID:          int32(*i.AccountId),
			Name:               *i.GameName,
			MaxAmountOfPlayers: int32(*i.MaxAmountOfPlayers),
			Name_2:             *i.Pagination.After,
			Limit:              int32(*i.Pagination.Limit),
		})
		if err != nil {
			return nil, err
		}

		games = make([]*models.GroupCollectionItem, len(rows))
		for k, v := range rows {
			owners, err := getOwners(v.Owners)
			if err != nil {
				return nil, err
			}

			games[k] = &models.GroupCollectionItem{
				Game: &models.GroupCollectionItemGame{
					IconUrl:            self.getIconUrl(v.IconPath),
					Id:                 int(v.GameID),
					LudopediaUrl:       v.LudopediaUrl.String,
					MaxAmountOfPlayers: int(v.MaxAmountOfPlayers),
					MinAmountOfPlayers: int(v.MinAmountOfPlayers),
					Name:               v.Name,
				},
				Owners: owners,
			}
		}
	}
	if i.AccountId != nil {
		rows, err := db.GetCollectiveCollectionByOwner(ctx, queries.GetCollectiveCollectionByOwnerParams{
			Kind:      queries.KindEnum(i.Kind),
			AccountID: int32(*i.AccountId),
			Name:      *i.Pagination.After,
			Limit:     int32(*i.Pagination.Limit),
		})
		if err != nil {
			return nil, err
		}

		games = make([]*models.GroupCollectionItem, len(rows))
		for k, v := range rows {
			owners, err := getOwners(v.Owners)
			if err != nil {
				return nil, err
			}

			games[k] = &models.GroupCollectionItem{
				Game: &models.GroupCollectionItemGame{
					IconUrl:            self.getIconUrl(v.IconPath),
					Id:                 int(v.GameID),
					LudopediaUrl:       v.LudopediaUrl.String,
					MaxAmountOfPlayers: int(v.MaxAmountOfPlayers),
					MinAmountOfPlayers: int(v.MinAmountOfPlayers),
					Name:               v.Name,
				},
				Owners: owners,
			}
		}
	}
	if i.GameName != nil {
		rows, err := db.GetCollectiveCollectionByGameName(ctx, queries.GetCollectiveCollectionByGameNameParams{
			Kind:   queries.KindEnum(i.Kind),
			Name:   *i.GameName,
			Name_2: *i.Pagination.After,
			Limit:  int32(*i.Pagination.Limit),
		})
		if err != nil {
			return nil, err
		}

		games = make([]*models.GroupCollectionItem, len(rows))
		for k, v := range rows {
			owners, err := getOwners(v.Owners)
			if err != nil {
				return nil, err
			}

			games[k] = &models.GroupCollectionItem{
				Game: &models.GroupCollectionItemGame{
					IconUrl:            self.getIconUrl(v.IconPath),
					Id:                 int(v.GameID),
					LudopediaUrl:       v.LudopediaUrl.String,
					MaxAmountOfPlayers: int(v.MaxAmountOfPlayers),
					MinAmountOfPlayers: int(v.MinAmountOfPlayers),
					Name:               v.Name,
				},
				Owners: owners,
			}
		}
	}
	if i.MaxAmountOfPlayers != nil {
		rows, err := db.GetCollectiveCollectionByMaxAmountOfPlayers(ctx, queries.GetCollectiveCollectionByMaxAmountOfPlayersParams{
			Kind:               queries.KindEnum(i.Kind),
			MaxAmountOfPlayers: int32(*i.MaxAmountOfPlayers),
			Name:               *i.Pagination.After,
			Limit:              int32(*i.Pagination.Limit),
		})
		if err != nil {
			return nil, err
		}

		games = make([]*models.GroupCollectionItem, len(rows))
		for k, v := range rows {
			owners, err := getOwners(v.Owners)
			if err != nil {
				return nil, err
			}

			games[k] = &models.GroupCollectionItem{
				Game: &models.GroupCollectionItemGame{
					IconUrl:            self.getIconUrl(v.IconPath),
					Id:                 int(v.GameID),
					LudopediaUrl:       v.LudopediaUrl.String,
					MaxAmountOfPlayers: int(v.MaxAmountOfPlayers),
					MinAmountOfPlayers: int(v.MinAmountOfPlayers),
					Name:               v.Name,
				},
				Owners: owners,
			}
		}
	}
	if i.AccountId == nil && i.GameName == nil && i.MaxAmountOfPlayers == nil {
		rows, err := db.GetCollectiveCollection(ctx, queries.GetCollectiveCollectionParams{
			Kind:  queries.KindEnum(i.Kind),
			Name:  after,
			Limit: int32(*i.Pagination.Limit),
		})
		if err != nil {
			return nil, err
		}

		games = make([]*models.GroupCollectionItem, len(rows))
		for k, v := range rows {
			owners, err := getOwners(v.Owners)
			if err != nil {
				return nil, err
			}

			games[k] = &models.GroupCollectionItem{
				Game: &models.GroupCollectionItemGame{
					IconUrl:            self.getIconUrl(v.IconPath),
					Id:                 int(v.GameID),
					LudopediaUrl:       v.LudopediaUrl.String,
					MaxAmountOfPlayers: int(v.MaxAmountOfPlayers),
					MinAmountOfPlayers: int(v.MinAmountOfPlayers),
					Name:               v.Name,
				},
				Owners: owners,
			}
		}
	}

	return &GetCollectiveCollectionOutput{
		Data: games,
		Pagination: &models.PaginationOutputString{
			Limit:   *i.Pagination.Limit,
			Next:    getNext(games),
			Current: i.Pagination.After,
		},
	}, nil
}
