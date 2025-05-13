package game_repository

import (
	"context"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (self *gameRepositoryImplementation) GetGameByLudopediaId(ctx context.Context, i *GetGameByLudopediaIdInput) (*models.Game, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	row, err := db.GetGameByLudopediaId(ctx, pgtype.Int4{
		Valid: true,
		Int32: int32(i.LudopediaId),
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("not found")
		}

		return nil, err
	}

	var iconPath *string
	if row.IconPath.Valid {
		iconPath = &row.IconPath.String
	}
	var ludopediaId *int
	if row.LudopediaID.Valid {
		ludopediaIdInt := int(row.LudopediaID.Int32)
		ludopediaId = &ludopediaIdInt
	}
	var ludopediaUrl *string
	if row.LudopediaUrl.Valid {
		ludopediaUrl = &row.LudopediaUrl.String
	}

	return &models.Game{
		Id:                 int(row.ID),
		Name:               row.Name,
		Description:        row.Description,
		IconPath:           iconPath,
		Kind:               models.Kind(row.Kind),
		LudopediaId:        ludopediaId,
		LudopediaUrl:       ludopediaUrl,
		MaxAmountOfPlayers: int(row.MaxAmountOfPlayers),
		MinAmountOfPlayers: int(row.MinAmountOfPlayers),
		MinAge:             int(row.MinAge),
		AverageDuration:    int(row.AverageDuration),
	}, nil
}
