package ludopedia

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

type getUserDataApiOutput struct {
	IdUsuario int    `json:"id_usuario"`
	Usuario   string `json:"usuario"`
	Thumb     string `json:"thumb"`
}

func (self *ludopediaAdapter) GetUserData(accessToken string) (*adapters.GetAuthenticatedUserDataOutput, error) {
	self.logger.Trace().Msg("start GetUserData")

	req, err := http.NewRequest(
		http.MethodGet,
		"https://ludopedia.com.br/api/v1/me",
		nil,
	)
	if err != nil {
		self.logger.Error().Err(err).Msg("fail to build request to get ludopedia user data")
		return nil, errors.New("fail to build request")
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)
	self.logger.Trace().Msg("request built")

	self.logger.Trace().Msg("do request to get ludopedia user data")
	userDataRes, err := self.httpClient.Do(req)
	if err != nil {
		self.logger.Error().Err(err).Msg("fail to make request to get ludopedia user data")
		return nil, errors.New("fail to make request")
	}
	defer userDataRes.Body.Close()
	self.logger.Trace().Msg("request to get ludopedia user data done")

	self.logger.Trace().Msg("decode req body")
	userData := getUserDataApiOutput{}
	err = json.NewDecoder(userDataRes.Body).Decode(&userData)
	if err != nil {
		self.logger.Error().Err(err).Msg("fail to decode req body")
		return nil, errors.New("fail to decode request body")
	}
	self.logger.Trace().Msg("req body decoded")

	output := &adapters.GetAuthenticatedUserDataOutput{
		Id:        strconv.Itoa(userData.IdUsuario),
		Handle:    &userData.Usuario,
		AvatarUrl: &userData.Thumb,
	}

	self.logger.Debug().
		Interface("output", output).
		Msg("successfully finish GetUserData")

	return output, nil
}
