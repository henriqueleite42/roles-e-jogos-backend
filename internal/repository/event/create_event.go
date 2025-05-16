package event_repository

import (
	"context"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5/pgtype"
)

func (self *eventRepositoryImplementation) CreateEvent(ctx context.Context, i *CreateEventInput) (*models.Event, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	var iconPath pgtype.Text
	if i.IconPath != nil {
		iconPath = pgtype.Text{
			Valid:  true,
			String: *i.IconPath,
		}
	}
	var endDate pgtype.Timestamptz
	if i.EndDate != nil {
		endDate = pgtype.Timestamptz{
			Valid: true,
			Time:  *i.EndDate,
		}
	}
	var capacity pgtype.Int4
	if i.Capacity != nil {
		capacity = pgtype.Int4{
			Valid: true,
			Int32: int32(*i.Capacity),
		}
	}

	eventId, err := db.CreateEvent(ctx, queries.CreateEventParams{
		Name:        i.Name,
		Description: i.Description,
		IconPath:    iconPath,
		StartDate: pgtype.Timestamptz{
			Valid: true,
			Time:  i.StartDate,
		},
		EndDate:  endDate,
		Capacity: capacity,
		OwnerID:  int32(i.OwnerId),
	})
	if err != nil {
		return nil, err
	}

	return &models.Event{
		Capacity:    i.Capacity,
		CreatedAt:   time.Now(),
		Description: i.Description,
		EndDate:     i.EndDate,
		IconPath:    i.IconPath,
		Id:          int(eventId),
		LocationId:  i.LocationId,
		Name:        i.Name,
		OwnerId:     i.OwnerId,
		StartDate:   i.StartDate,
	}, nil
}
