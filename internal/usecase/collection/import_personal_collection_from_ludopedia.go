package collection_usecase

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/models"
	account_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/account"
	collection_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/collection"
	game_repository "github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/game"
	collection_utils "github.com/henriqueleite42/roles-e-jogos-backend/internal/usecase/collection/utils"
	"github.com/henriqueleite42/roles-e-jogos-backend/internal/utils"
)

type filteredEvent struct {
	// Step 1
	AccountId             int
	ExternalId            string
	ImportCollectionLogId int
	// Step 2
	AccessToken string
}

func convertLudopediaPaidToOurPaid(paid *float64) (*int, error) {
	if paid == nil {
		return nil, nil
	}

	// Convert float64 to string with precision
	str := strconv.FormatFloat(*paid, 'f', -1, 64)

	var paidStr string

	parts := strings.Split(str, ".")
	if len(parts) >= 1 {
		paidStr += parts[0]
	}
	if len(parts) >= 2 {
		frac := parts[1]
		if len(frac) > 2 {
			frac = frac[:2]
		}
		paidStr += frac
	}

	paidInt, err := strconv.Atoi(paidStr)
	if err != nil {
		return nil, err
	}

	return &paidInt, nil
}

func (self *CollectionUsecaseImplementation) importGame(
	ctx context.Context,
	accessToken string,
	ludopediaGameId int,
) (*models.Game, error) {
	ludopediaGameToImport, err := self.LudopediaAdapter.GetGame(&adapters.GetGameInput{
		AccessToken: accessToken,
		LudopediaId: ludopediaGameId,
	})
	if err != nil {
		self.Logger.Error().Err(err).Int("gameId", ludopediaGameId).Msg("fail to get ludopedia game")
		return nil, err
	}

	var iconPath *string
	if ludopediaGameToImport.ImageUrl != nil {
		path := fmt.Sprintf("games/%s.{{ext}}", self.IdAdapter.GenId())
		iconPathStr, err := self.StorageAdapter.DownloadFromUrl(&adapters.DownloadFromUrlInput{
			Url:       *ludopediaGameToImport.ImageUrl,
			StorageId: self.SecretsAdapter.MediasS3BucketName,
			FileName:  path,
		})
		if err != nil {
			self.Logger.Error().Err(err).Msg("fail to download game img")
		} else {
			iconPath = &iconPathStr
		}
	}

	newGame, err := self.GameRepository.CreateGame(ctx, &game_repository.CreateGameInput{
		Name:               ludopediaGameToImport.Name,
		Description:        "",
		IconPath:           iconPath,
		Kind:               models.Kind_Game,
		LudopediaId:        &ludopediaGameToImport.Id,
		LudopediaUrl:       &ludopediaGameToImport.LudopediaUrl,
		AverageDuration:    ludopediaGameToImport.AverageDuration,
		MinAmountOfPlayers: ludopediaGameToImport.MinAmountOfPlayers,
		MaxAmountOfPlayers: ludopediaGameToImport.MaxAmountOfPlayers,
		MinAge:             ludopediaGameToImport.MinAge,
	})
	if err != nil {
		self.Logger.Error().Err(err).Int("gameId", ludopediaGameToImport.Id).Msg("fail to import ludopedia game")
		return nil, err
	}
	return newGame, nil
}

func (self *CollectionUsecaseImplementation) getFullLudopediaCollection(
	ctx context.Context,
	i *filteredEvent,
	collectionManager *collection_utils.CollectionManager,
) error {
	page := 1

	for {
		collectionPage, err := self.LudopediaAdapter.GetCollection(&adapters.GetCollectionInput{
			AccessToken: i.AccessToken,
			Page:        strconv.Itoa(page),
		})
		if err != nil {
			self.Logger.Error().Err(err).Int("page", page).Msg("fail to get ludopedia collection")
			return err
		}
		// Empty page
		if len(collectionPage.Collection) == 0 {
			break
		}

		for _, v := range collectionPage.Collection {
			paid, err := convertLudopediaPaidToOurPaid(v.Paid)
			if err != nil {
				self.Logger.Error().Err(err).Msg("fail to convert ludopedia paid")
			}

			collectionManager.AddAccountLudopediaGame(&collection_utils.AddAccountLudopediaGameInput{
				AccountId:       i.AccountId,
				AccessToken:     i.AccessToken,
				LudopediaGameId: v.Id,
				Paid:            paid,
			})
		}

		// Has less items than the maximum amount
		if len(collectionPage.Collection) < 100 {
			break
		}

		page++
	}

	return nil
}

func (self *CollectionUsecaseImplementation) updateOngoingImports(
	ctx context.Context,
	events []*filteredEvent,
	newStatus models.CollectionImportStatus,
) {
	ids := make([]int, len(events))
	for _, v := range events {
		ids = append(ids, v.ImportCollectionLogId)
	}

	err := self.CollectionRepository.UpdateManyImportCollectionsLogs(ctx, &collection_repository.UpdateManyImportCollectionsLogsInput{
		Ids:    ids,
		Status: newStatus,
	})
	if err != nil {
		self.Logger.
			Error().
			Err(err).
			Any("ImportsIds", ids).
			Any("NewStatus", newStatus).
			Msg("fail to update ongoing imports")
	}
}

// ===============================================
//
// THIS IS NOT UPDATED, DO NOT TRUST IT 100%!!!
//
// ==============================================
//
// Step 1:
// - Get the import status of every user
// - Ignores the ones that are already im progress
// - Create a new import status for every one that don't have one
//
// Step 2:
// - Get the user's ludopedia credentials
// - Ignore the ones that don't have Ludopedia linked
//   - Update import_log
//
// Step 3:
// - Get's the user collection from ludopedia
//   - Because probably they will not have sooooo many games, maybe we can get all the games for all the users at once?
//   - We will get only the ludopediaGameId and the accountId, put it on 2 maps and save it for later
//   - Map1: map[ludopediaGameId]: true
//   - Map2: map[accountId]: []ludopediaGameId
//
// Step 4:
// - With the map of ludopediaGameId, transform it into a slice
// - Get from the database all the games with these IDs
// - Check which games doesn't exist, and import then
// - Create a map of map[ludopediaGameId]: gameId
//
// Step 5:
// - With the map map[accountId]: []ludopediaGameId, save all the games on the database
// - Update import_log
func (self *CollectionUsecaseImplementation) ImportPersonalCollectionFromLudopedia(ctx context.Context, i []*models.ImportCollectionEvent) error {
	logger := utils.GetLoggerFromCtx(ctx, self.Logger)

	// =====================================
	//
	// Step 1
	//
	// =====================================

	// Group all external IDs into an slice
	externalIds := make([]string, len(i))
	for k, v := range i {
		externalIds[k] = v.ExternalId
	}
	logger.Trace().
		Any("externalIds", externalIds).
		Msg("externalIds")

	// Get all the ongoing imports
	ongoingImports, err := self.CollectionRepository.GetOngoingImportCollectionLog(ctx, &collection_repository.GetOngoingImportCollectionLogInput{
		ExternalIds: externalIds,
		Provider:    models.Provider_Ludopedia,
	})
	if err != nil {
		logger.Error().Err(err).Msg("fail to get ongoing imports")
		return err
	}

	// Transform ongoing imports into a map, to have a better performance on the next step
	ongoingImportsMap := make(map[string]bool, len(i))
	for _, v := range ongoingImports.Data {
		ongoingImportsMap[v.ExternalId] = true
	}
	logger.Trace().
		Any("ongoingImportsMap", ongoingImportsMap).
		Msg("got ongoing imports")

	// Filter all events: If the connection already has an import in progress, ignore the event
	filteredLength := len(i) - len(ongoingImportsMap)
	step1Events := make([]*filteredEvent, 0, filteredLength)
	filteredExternalIds := make([]string, 0, filteredLength)
	for _, v := range i {
		if ongoingImportsMap[v.ExternalId] {
			continue
		}

		// Also creates the import_log on the database
		importCollectionLog, err := self.CollectionRepository.CreateImportCollectionLog(ctx, &collection_repository.CreateImportCollectionLogInput{
			AccountId:  v.AccountId,
			ExternalId: v.ExternalId,
			Provider:   models.Provider_Ludopedia,
			Status:     models.CollectionImportStatus_Started,
			Trigger:    v.Trigger,
		})
		if err != nil {
			logger.Error().Err(err).Msg("fail to create import collection log")
			continue
		}

		step1Events = append(step1Events, &filteredEvent{
			AccountId:             v.AccountId,
			ExternalId:            v.ExternalId,
			ImportCollectionLogId: importCollectionLog.Id,
		})
		filteredExternalIds = append(filteredExternalIds, v.ExternalId)
	}
	logger.Trace().
		Int("len(step1Events)", len(step1Events)).
		Any("filteredExternalIds", filteredExternalIds).
		Msg("got step1Events")

	// =====================================
	//
	// Step 2
	//
	// =====================================

	// Get all connections to get their credentials
	connections, err := self.AccountRepository.GetConnectionsListByExternalIdAndProvider(ctx, &account_repository.GetConnectionsListByExternalIdAndProviderInput{
		ExternalIds: filteredExternalIds,
		Provider:    models.Provider_Ludopedia,
	})
	if err != nil {
		logger.Error().Err(err).Msg("fail to get connections")
		return err
	}
	connectionsExternalId := make(map[string]*models.Connection, filteredLength)
	for _, v := range connections.Data {
		logger.Trace().
			Str("ExternalId", v.ExternalId).
			Any("Connection", *v).
			Msg("connection")
		connectionsExternalId[v.ExternalId] = v
	}
	logger.Trace().
		Int("len(connections.Data)", len(connections.Data)).
		Any("connectionsExternalId", connectionsExternalId).
		Msg("got connections")

	step2Events := make([]*filteredEvent, 0, len(step1Events))
	step2EventsErrors := make([]*filteredEvent, 0, len(step1Events))
	for _, event := range step1Events {
		if connectionsExternalId[event.ExternalId] != nil &&
			connectionsExternalId[event.ExternalId].AccessToken != nil {
			connection := connectionsExternalId[event.ExternalId]

			step2Events = append(step2Events, &filteredEvent{
				AccountId:             event.AccountId,
				ExternalId:            event.ExternalId,
				ImportCollectionLogId: event.ImportCollectionLogId,
				AccessToken:           *connection.AccessToken,
			})
		} else {
			step2EventsErrors = append(step2EventsErrors, event)
		}
	}
	if len(step2EventsErrors) > 0 {
		self.updateOngoingImports(ctx, step2EventsErrors, models.CollectionImportStatus_Failed)
	}
	if len(step2Events) == 0 {
		logger.Warn().Err(err).Msg("no events with connection")
		return fmt.Errorf("no events with connection")
	}

	// =====================================
	//
	// Step 3
	//
	// =====================================

	logger.Trace().Msg("try to get ludopedia collection")
	collectionManager := collection_utils.NewCollectionManager()
	for _, v := range step2Events {
		err := self.getFullLudopediaCollection(ctx, v, collectionManager)
		if err != nil {
			self.updateOngoingImports(ctx, step2Events, models.CollectionImportStatus_Failed)
			return err
		}
	}
	logger.Trace().Msg("successfully got ludopedia collection")

	// =====================================
	//
	// Step 4
	//
	// =====================================

	logger.Trace().Msg("try to get games")
	ludopediaGamesIds := collectionManager.GetLudopediaGamesIds()
	gamesRelations, err := self.GameRepository.GetGamesListByLudopediaId(ctx, &game_repository.GetGamesListByLudopediaIdInput{
		LudopediaIds: ludopediaGamesIds,
	})
	if err != nil {
		logger.Error().Err(err).Msg("fail to get games relation on ludopedia collection import")
		return err
	}
	logger.Trace().Msg("successfully got games")
	gamesIdsMap := make(map[int]int, len(gamesRelations.Data))
	for _, v := range gamesRelations.Data {
		externalId, err := strconv.Atoi(v.ExternalId)
		if err != nil {
			gamesIdsMap[externalId] = v.GameId
		} else {
			logger.Error().Err(err).Str("ExternalId", v.ExternalId).Msg("fail to convert ludopedia game ID to int")
		}
	}
	logger.Trace().
		Any("ludopediaGamesIds", ludopediaGamesIds).
		Any("gamesIdsMap", gamesIdsMap).
		Msg("successfully created gamesIdsMap")

	logger.Trace().Msg("try to import missing games")
	gamesImportCounterForLogs := 1
	totalGamesToImport := len(ludopediaGamesIds) - len(gamesIdsMap)
	gamesIdsWithError := make(map[int]bool, totalGamesToImport)
	for _, ludopediaGameId := range ludopediaGamesIds {
		// If game already exists, theres no need to import it
		if gamesIdsMap[ludopediaGameId] != 0 {
			logger.Trace().Int("foo", gamesIdsMap[ludopediaGameId]).Msg("gamesIdsMap[ludopediaGameId]")
			gamesImportCounterForLogs++
			continue
		}

		newGame, err := self.importGame(
			ctx,
			collectionManager.AccessTokenByLudopediaGameId[ludopediaGameId],
			ludopediaGameId,
		)
		if err == nil {
			gamesIdsMap[ludopediaGameId] = newGame.Id
			logger.Trace().Msgf("successfully imported game %d / %d", gamesImportCounterForLogs, totalGamesToImport)
			gamesImportCounterForLogs++
		} else {
			gamesIdsWithError[ludopediaGameId] = true
			logger.Trace().Msgf("fail to import game %d / %d", gamesImportCounterForLogs, totalGamesToImport)
			gamesImportCounterForLogs++
		}

		// Sleeps to be sure that Ludopedia doesn't rate limit us
		time.Sleep(1 * time.Second)
	}

	if len(gamesIdsWithError) == len(ludopediaGamesIds) {
		logger.Trace().
			Any("gamesIdsWithError", gamesIdsWithError).
			Msg("fail to import all games")
		self.updateOngoingImports(ctx, step2Events, models.CollectionImportStatus_Failed)
		return fmt.Errorf("fail to import all games")
	}
	logger.Trace().
		Any("gamesIdsMap", gamesIdsMap).
		Any("gamesIdsWithError", gamesIdsWithError).
		Msg("successfully imported games")

	// =====================================
	//
	// Step 5
	//
	// =====================================

	errorEvents := make([]*filteredEvent, 0, len(step2Events))
	successEvents := make([]*filteredEvent, 0, len(step2Events))
	for _, event := range step2Events {
		accountGames := collectionManager.AccountLudopediaGamesMap[event.AccountId]

		logger.Trace().
			Int("AccountId", event.AccountId).
			Str("ExternalId", event.ExternalId).
			Msg("try to add games to personal collection")

		// If the account has no games, so we don't need to import anything and it succeed
		if len(accountGames) == 0 {
			logger.Trace().
				Int("AccountId", event.AccountId).
				Str("ExternalId", event.ExternalId).
				Msg("account has no games to add")
			continue
		}

		// Remove games that we could not import
		var filteredGames []*collection_utils.GameToImport
		for _, v := range accountGames {
			if !gamesIdsWithError[v.LudopediaGameId] {
				filteredGames = append(filteredGames, v)
			}
		}

		// If we failed to import all the games from ludopedia
		if len(filteredGames) == 0 {
			logger.Trace().
				Int("AccountId", event.AccountId).
				Str("ExternalId", event.ExternalId).
				Msg("all account games failed to be imported")
			errorEvents = append(errorEvents, event)
			continue
		}

		errorCounter := 0
		for _, v := range filteredGames {
			err := self.CollectionRepository.AddToPersonalCollection(ctx, &collection_repository.AddToPersonalCollectionInput{
				AccountId: event.AccountId,
				GameId:    gamesIdsMap[v.LudopediaGameId],
				Paid:      v.Paid,
			})
			if err != nil {
				logger.Error().Err(err).
					Int("AccountId", event.AccountId).
					Int("LudopediaGameId", v.LudopediaGameId).
					Int("GameId", gamesIdsMap[v.LudopediaGameId]).
					Any("Paid", v.Paid).
					Msg("fail to add ludopedia game to personal collection")
				errorCounter++
			}
		}

		// If all games failed be added to personal collection, so the processing failed
		if errorCounter == len(filteredGames) {
			logger.Trace().
				Int("AccountId", event.AccountId).
				Str("ExternalId", event.ExternalId).
				Msg("all account games failed to be added to personal collection")
			errorEvents = append(errorEvents, event)
			continue
		}

		logger.Trace().
			Int("AccountId", event.AccountId).
			Str("ExternalId", event.ExternalId).
			Msg("successfully add account games to personal collection")
		successEvents = append(successEvents, event)
	}

	if len(errorEvents) > 0 {
		logger.Trace().Msg("register failed imports")
		self.updateOngoingImports(ctx, errorEvents, models.CollectionImportStatus_Failed)

		// If all events failed, we don't need to do anything else
		if len(errorEvents) == len(step2Events) {
			logger.Trace().Msg("fail to add all games to personal collection")
			return fmt.Errorf("fail to process all events")
		}
	}

	if len(successEvents) == 0 {
		logger.Error().
			Any("errorEvents", errorEvents).
			Any("successEvents", successEvents).
			Msg("for some reason, the successEvents are empty")
		return fmt.Errorf("for some reason, the successEvents are empty")
	}

	self.updateOngoingImports(ctx, successEvents, models.CollectionImportStatus_Completed)
	logger.Trace().Msg("successfully added games to personal collection")

	return nil
}
