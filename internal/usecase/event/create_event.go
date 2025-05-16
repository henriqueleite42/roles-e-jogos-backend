package event_usecase

import (
	"context"

	event_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/event"
	game_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/game"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

func (self *EventUsecaseImplementation) CreateEvent(ctx context.Context, i *CreateEventInput) error {
	tx, ctx, err := utils.SetTxInCtx(self.Db, ctx)
	if err != nil {
		return err
	}

	var iconPath *string
	if i.Icon != nil {
		if i.Icon.CustomIconPath != nil {
			iconPath = i.Icon.CustomIconPath
		}
		if i.Icon.UseGameIconGameId != nil {
			game, err := self.GameRepository.GetGameById(ctx, &game_repository.GetGameByIdInput{
				Id: *i.Icon.UseGameIconGameId,
			})
			if err == nil {
				iconPath = game.IconPath
			} else {
				self.Logger.Warn().Int("GameId", *i.Icon.UseGameIconGameId).Err(err).Msg("fail to get game by id")
			}
		}
	}

	event, err := self.EventRepository.CreateEvent(ctx, &event_repository.CreateEventInput{
		Capacity:    i.Capacity,
		Description: i.Description,
		EndDate:     i.EndDate,
		IconPath:    iconPath,
		LocationId:  i.LocationId,
		Name:        i.Name,
		OwnerId:     i.AccountId,
		StartDate:   i.StartDate,
	})
	if err != nil {
		tx.Rollback(ctx)
		self.Logger.Warn().Err(err).Msg("fail to create event")
		return err
	}

	for _, v := range i.GamesList {
		err := self.EventRepository.CreateEventGame(ctx, &event_repository.CreateEventGameInput{
			EventId: event.Id,
			GameId:  v,
			OwnerId: i.AccountId,
		})
		if err != nil {
			tx.Rollback(ctx)
			self.Logger.Warn().Err(err).Msg("fail to create event game")
			return err
		}
	}

	return nil
}
