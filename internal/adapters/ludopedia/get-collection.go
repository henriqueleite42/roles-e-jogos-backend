package ludopedia

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

func (self *ludopediaAdapter) GetCollection(i *adapters.GetCollectionInput) (*adapters.GetCollectionOutput, error) {
	self.logger.Trace().Msg("start GetCollection")

	page := "1"
	if i.Page != "" {
		page = i.Page
	}

	req, err := http.NewRequest(
		http.MethodGet,
		"https://ludopedia.com.br/api/v1/colecao?lista=colecao&tp_jogo=b&rows=100&page="+page,
		nil,
	)
	if err != nil {
		self.logger.Error().Err(err).Msg("fail to build request to get ludopedia collection")
		return nil, errors.New("fail to build request")
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+i.AccessToken)
	self.logger.Trace().Msg("request built")

	self.logger.Trace().Msg("do request to get ludopedia collection")
	collectionDataRes, err := self.httpClient.Do(req)
	if err != nil {
		self.logger.Error().Err(err).Msg("fail to make request to get ludopedia collection")
		return nil, errors.New("fail to make request")
	}
	defer collectionDataRes.Body.Close()
	self.logger.Trace().Msg("request to get ludopedia collection done")

	if collectionDataRes.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(collectionDataRes.Body)
		if err != nil {
			self.logger.Error().
				Int("status", collectionDataRes.StatusCode).
				Err(err).
				Msg("fail to read response body")
			return nil, errors.New("fail to make request")
		}

		bodyString := string(bodyBytes)
		self.logger.Error().
			Int("status", collectionDataRes.StatusCode).
			Str("body", bodyString).
			Msg("fail to make request")
		return nil, errors.New("fail to make request")
	}

	self.logger.Trace().Msg("decode req body")
	collection := &adapters.GetCollectionOutput{}
	err = json.NewDecoder(collectionDataRes.Body).Decode(collection)
	if err != nil {
		self.logger.Error().Err(err).Msg("fail to decode req body")
		return nil, errors.New("fail to decode request body")
	}
	self.logger.Trace().Msg("req body decoded")

	self.logger.Debug().
		Interface("output", *collection).
		Msg("successfully finish GetCollection")

	return collection, nil

}
