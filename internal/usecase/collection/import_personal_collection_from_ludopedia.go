package collection_usecase

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	collection_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/collection"
	game_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/game"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

// Return the game and if it was imported, or an error
func (self *CollectionUsecaseImplementation) getOrImportGame(
	ctx context.Context,
	connection *models.Connection,
	ludopediaGame *adapters.LudopediaCollectionItem,
) (*models.Game, error) {
	game, err := self.GameRepository.GetGameByLudopediaId(ctx, &game_repository.GetGameByLudopediaIdInput{
		LudopediaId: ludopediaGame.Id,
	})
	if err != nil && err.Error() != "not found" {
		self.Logger.Error().Err(err).Int("gameId", ludopediaGame.Id).Msg("fail to get game from database")
		return nil, err
	}

	// If game already exists, we don't need to import it
	if game != nil {
		return game, nil
	}

	// If the game doesn't exists in our database, then we need to import it

	ludopediaGameToImport, err := self.LudopediaAdapter.GetGame(&adapters.GetGameInput{
		AccessToken: *connection.AccessToken,
		LudopediaId: ludopediaGame.Id,
	})
	if err != nil {
		self.Logger.Error().Err(err).Int("gameId", ludopediaGame.Id).Msg("fail to get ludopedia game")
		return nil, err
	}

	path := fmt.Sprintf("games/%s.{{ext}}", self.IdAdapter.GenId())
	iconPath, err := self.StorageAdapter.DownloadFromUrl(&adapters.DownloadFromUrlInput{
		Url:       ludopediaGame.ImageUrl,
		StorageId: self.SecretsAdapter.MediasS3BucketName,
		FileName:  path,
	})
	if err != nil {
		self.Logger.Error().Err(err).Msg("fail to download game img")
		return nil, err
	}

	newGame, err := self.GameRepository.CreateGame(ctx, &game_repository.CreateGameInput{
		Name:               ludopediaGame.Name,
		Description:        "",
		IconPath:           &iconPath,
		Kind:               models.Kind_Game,
		LudopediaId:        &ludopediaGame.Id,
		LudopediaUrl:       &ludopediaGame.LudopediaUrl,
		AverageDuration:    ludopediaGameToImport.AverageDuration,
		MinAmountOfPlayers: ludopediaGameToImport.MinAmountOfPlayers,
		MaxAmountOfPlayers: ludopediaGameToImport.MaxAmountOfPlayers,
		MinAge:             ludopediaGameToImport.MinAge,
	})
	if err != nil {
		self.Logger.Error().Err(err).Int("gameId", ludopediaGame.Id).Msg("fail to import ludopedia game")
		return nil, err
	}
	return newGame, nil
}

func (self *CollectionUsecaseImplementation) ImportPersonalCollectionFromLudopedia(ctx context.Context, i *ImportPersonalCollectionFromLudopediaInput) error {
	ctx, err := utils.SetDbInCtx(self.Db, ctx)
	if err != nil {
		return err
	}

	connections, err := self.AccountRepository.GetConnectionsByAccountIdAndProvider(ctx, &account_repository.GetConnectionsByAccountIdAndProviderInput{})
	if err != nil {
		self.Logger.Warn().Err(err).Msg("fail to get collective collection")
		return err
	}

	for _, connection := range connections {
		page := 1

		for {
			collectionPage, err := self.LudopediaAdapter.GetCollection(&adapters.GetCollectionInput{
				AccessToken: *connection.AccessToken,
				Page:        strconv.Itoa(page),
			})
			if err != nil {
				self.Logger.Error().Err(err).Int("page", page).Msg("fail to get ludopedia collection")
				break
			}
			// Empty page
			if len(collectionPage.Collection) == 0 {
				break
			}

			for _, ludopediaGame := range collectionPage.Collection {
				game, err := self.getOrImportGame(ctx, connection, ludopediaGame)
				if err != nil {
					break
				}

				err = self.CollectionRepository.AddToPersonalCollection(ctx, &collection_repository.AddToPersonalCollectionInput{
					AccountId: i.AccountId,
					GameId:    game.Id,
					Paid:      &ludopediaGame.Paid,
				})
				if err != nil {
					self.Logger.Error().Err(err).Int("gameId", ludopediaGame.Id).Msg("fail to add game to personal collection")
					break
				}

				// Sleeps for half a second to avoid ludopedia rate limit
				time.Sleep((1 * time.Second) / 2)
			}

			// Has less items than the maximum amount
			if len(collectionPage.Collection) < 100 {
				break
			}

			page++
		}
	}

	return nil
}
