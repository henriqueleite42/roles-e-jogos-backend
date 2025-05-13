package game_repository

import (
	"context"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5/pgtype"
)

func (self *gameRepositoryImplementation) CreateGame(ctx context.Context, i *CreateGameInput) (*models.Game, error) {
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
	var ludopediaId pgtype.Int4
	if i.LudopediaId != nil {
		ludopediaId = pgtype.Int4{
			Valid: true,
			Int32: int32(*i.LudopediaId),
		}
	}
	var ludopediaUrl pgtype.Text
	if i.LudopediaUrl != nil {
		ludopediaUrl = pgtype.Text{
			Valid:  true,
			String: *i.LudopediaUrl,
		}
	}

	gameId, err := db.CreateGame(ctx, queries.CreateGameParams{
		Name:               i.Name,
		Description:        i.Description,
		IconPath:           iconPath,
		Kind:               queries.KindEnum(i.Kind),
		LudopediaID:        ludopediaId,
		LudopediaUrl:       ludopediaUrl,
		MaxAmountOfPlayers: int32(i.MaxAmountOfPlayers),
		MinAmountOfPlayers: int32(i.MinAmountOfPlayers),
		MinAge:             int32(i.MinAge),
		AverageDuration:    int32(i.AverageDuration),
	})
	if err != nil {
		return nil, err
	}

	return &models.Game{
		Id:                 int(gameId),
		Name:               i.Name,
		Description:        i.Description,
		IconPath:           i.IconPath,
		Kind:               i.Kind,
		LudopediaId:        i.LudopediaId,
		LudopediaUrl:       i.LudopediaUrl,
		MaxAmountOfPlayers: i.MaxAmountOfPlayers,
		MinAmountOfPlayers: i.MinAmountOfPlayers,
		MinAge:             i.MinAge,
		AverageDuration:    i.AverageDuration,
		CreatedAt:          time.Now(),
	}, nil
}
