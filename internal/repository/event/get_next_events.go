package event_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5/pgtype"
)

type gamesJsonData struct {
	Id                 int     `json:"id"`
	Name               string  `json:"name"`
	IconPath           *string `json:"icon_path"`
	Kind               string  `json:"kind"`
	LudopediaUrl       *string `json:"ludopedia_url"`
	MinAmountOfPlayers int     `json:"min_amount_of_players"`
	MaxAmountOfPlayers int     `json:"max_amount_of_players"`
	AverageDuration    int     `json:"average_duration"`
	MinAge             int     `json:"min_age"`
}

type attendancesJsonData struct {
	AccountId  int     `json:"account_id"`
	Handle     string  `json:"handle"`
	AvatarPath *string `json:"avatar_path"`
	Status     string  `json:"status"`
}

func (self *eventRepositoryImplementation) GetNextEvents(ctx context.Context, i *GetNextEventsInput) (*GetNextEventsOutput, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	after := pgtype.Timestamptz{
		Valid: true,
	}
	if i.Pagination.After == nil {
		after.Time = time.Now()
	} else {
		after.Time = *i.Pagination.After
	}

	rows, err := db.GetNextEvents(ctx, queries.GetNextEventsParams{
		StartDate: after,
		Limit:     int32(i.Pagination.Limit),
	})
	if err != nil {
		return nil, err
	}

	events := make([]*models.EventData, len(rows))
	for k, v := range rows {
		var iconUrl *string
		if v.IconPath.Valid {
			url := self.secretsAdapter.MediasCloudfrontUrl + v.IconPath.String
			iconUrl = &url
		}
		var endDate *time.Time
		if v.EndDate.Valid {
			endDate = &v.EndDate.Time
		}
		var capacity *int
		if v.Capacity.Valid {
			capacityInt := int(v.Capacity.Int32)
			capacity = &capacityInt
		}

		var locationIconUrl *string
		if v.LocationIconPath.Valid {
			url := self.secretsAdapter.MediasCloudfrontUrl + v.LocationIconPath.String
			locationIconUrl = &url
		}

		gamesJsonParsed := []*gamesJsonData{}
		err := json.Unmarshal(v.Games, &gamesJsonParsed)
		if err != nil {
			return nil, fmt.Errorf("fail to decode games json")
		}
		games := make([]*models.EventDataGamesItem, 0, len(gamesJsonParsed))
		for _, vv := range gamesJsonParsed {
			if vv.Id == 0 {
				continue
			}

			var iconUrl *string
			if vv.IconPath != nil {
				url := self.secretsAdapter.MediasCloudfrontUrl + *vv.IconPath
				iconUrl = &url
			}

			games = append(games, &models.EventDataGamesItem{
				Id:                 vv.Id,
				Name:               vv.Name,
				AverageDuration:    vv.AverageDuration,
				IconUrl:            iconUrl,
				Kind:               models.Kind(vv.Kind),
				LudopediaUrl:       vv.LudopediaUrl,
				MaxAmountOfPlayers: vv.MaxAmountOfPlayers,
				MinAge:             vv.MinAge,
				MinAmountOfPlayers: vv.MinAmountOfPlayers,
			})
		}

		attendancesJsonParsed := []*attendancesJsonData{}
		err = json.Unmarshal(v.Attendances, &attendancesJsonParsed)
		if err != nil {
			return nil, fmt.Errorf("fail to decode attendances json")
		}
		attendances := make([]*models.EventDataAttendancesItem, 0, len(attendancesJsonParsed))
		for _, vv := range attendancesJsonParsed {
			if vv.AccountId == 0 {
				continue
			}

			var avatarUrl *string
			if vv.AvatarPath != nil {
				url := self.secretsAdapter.MediasCloudfrontUrl + *vv.AvatarPath
				avatarUrl = &url
			}

			attendances = append(attendances, &models.EventDataAttendancesItem{
				AccountId: vv.AccountId,
				AvatarUrl: avatarUrl,
				Handle:    vv.Handle,
				Status:    models.EventAttendanceStatus(vv.Status),
			})
		}

		events[k] = &models.EventData{
			Id:          int(v.ID),
			Name:        v.Name,
			Description: v.Description,
			StartDate:   v.StartDate.Time,
			EndDate:     endDate,
			IconUrl:     iconUrl,
			OwnerId:     int(v.OwnerID),
			Capacity:    capacity,
			Location: &models.EventDataLocation{
				Address: v.LocationAddress,
				IconUrl: locationIconUrl,
				Id:      int(v.LocationID),
				Name:    v.LocationName,
			},
			Games:       games,
			Attendances: attendances,
		}
	}

	var next *time.Time
	if len(events) == i.Pagination.Limit {
		next = &events[len(events)-1].StartDate
	}

	return &GetNextEventsOutput{
		Data: events,
		Pagination: &models.PaginationOutputTimestamp{
			Limit:   i.Pagination.Limit,
			Next:    next,
			Current: i.Pagination.After,
		},
	}, nil
}
