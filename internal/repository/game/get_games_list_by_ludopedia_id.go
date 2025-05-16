package game_repository

import (
	"context"
)

func (self *gameRepositoryImplementation) GetGamesListByLudopediaId(ctx context.Context, i *GetGamesListByLudopediaIdInput) (*GetGamesListByLudopediaIdOutput, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	ludopediaIds := make([]int32, len(i.LudopediaIds))
	for k, v := range i.LudopediaIds {
		ludopediaIds[k] = int32(v)
	}

	result, err := db.GetGamesListByLudopediaId(ctx, ludopediaIds)
	if err != nil {
		return nil, err
	}

	games := make([]*GetGamesListByLudopediaIdOutputDataItem, len(result))
	for k, v := range result {
		games[k] = &GetGamesListByLudopediaIdOutputDataItem{
			LudopediaId: int(v.LudopediaID.Int32),
			GameId:      int(v.ID),
		}
	}

	return &GetGamesListByLudopediaIdOutput{
		Data: games,
	}, nil
}
