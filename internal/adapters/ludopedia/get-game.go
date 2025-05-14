package ludopedia

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

func (self *ludopediaAdapter) GetGame(i *adapters.GetGameInput) (*adapters.GetGameOutput, error) {
	self.logger.Trace().Msg("start GetGame")

	req, err := http.NewRequest(
		http.MethodGet,
		"https://ludopedia.com.br/api/v1/jogos/"+strconv.Itoa(i.LudopediaId),
		nil,
	)
	if err != nil {
		self.logger.Error().Err(err).Msg("fail to build request to get ludopedia game")
		return nil, errors.New("fail to build request")
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+i.AccessToken)
	self.logger.Trace().Msg("request built")

	self.logger.Trace().Msg("do request to get ludopedia game")
	gameDataRes, err := self.httpClient.Do(req)
	if err != nil {
		self.logger.Error().Err(err).Msg("fail to make request to get ludopedia game")
		return nil, errors.New("fail to make request")
	}
	defer gameDataRes.Body.Close()
	self.logger.Trace().Msg("request to get ludopedia game done")

	if gameDataRes.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(gameDataRes.Body)
		if err != nil {
			self.logger.Error().
				Int("status", gameDataRes.StatusCode).
				Err(err).
				Msg("fail to read response body")
			return nil, errors.New("fail to make request")
		}

		bodyString := string(bodyBytes)
		self.logger.Error().
			Int("status", gameDataRes.StatusCode).
			Str("body", bodyString).
			Msg("fail to make request")
		return nil, errors.New("fail to make request")
	}

	self.logger.Trace().Msg("decode req body")
	game := &adapters.GetGameOutput{}
	err = json.NewDecoder(gameDataRes.Body).Decode(game)
	if err != nil {
		self.logger.Error().Err(err).Msg("fail to decode req body")
		return nil, errors.New("fail to decode request body")
	}
	self.logger.Trace().Msg("req body decoded")

	self.logger.Debug().
		Interface("output", *game).
		Msg("successfully finish GetGame")

	return game, nil

}
