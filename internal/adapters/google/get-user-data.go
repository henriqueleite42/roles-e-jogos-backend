package google

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

type getUserDataApiOutput struct {
	Sub           string `json:"sub"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Picture       string `json:"picture"`
}

func (self *googleAdapter) GetUserData(accessToken string) (*adapters.GetAuthenticatedUserDataOutput, error) {
	self.logger.Trace().Msg("start GetUserData")

	req, err := http.NewRequest(
		http.MethodGet,
		"https://openidconnect.googleapis.com/v1/userinfo",
		nil,
	)
	if err != nil {
		self.logger.Error().Err(err).Msg("fail to build request to get google user data")
		return nil, errors.New("fail to build request")
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)
	self.logger.Trace().Msg("request built")

	self.logger.Trace().Msg("do request to get google user data")
	userDataRes, err := self.httpClient.Do(req)
	if err != nil {
		self.logger.Error().Err(err).Msg("fail to make request to get google user data")
		return nil, errors.New("fail to make request")
	}
	defer userDataRes.Body.Close()
	self.logger.Trace().Msg("request to get google user data done")

	self.logger.Trace().Msg("decode req body")
	userData := getUserDataApiOutput{}
	err = json.NewDecoder(userDataRes.Body).Decode(&userData)
	if err != nil {
		self.logger.Error().Err(err).Msg("fail to decode req body")
		return nil, errors.New("fail to decode request body")
	}
	self.logger.Trace().Msg("req body decoded")

	handle := strings.Split(userData.Email, "@")[0]

	output := &adapters.GetAuthenticatedUserDataOutput{
		Id:              userData.Sub,
		Name:            userData.GivenName + " " + userData.FamilyName,
		Handle:          &handle,
		Email:           userData.Email,
		IsEmailVerified: userData.EmailVerified,
		AvatarUrl:       &userData.Picture,
	}

	self.logger.Debug().
		Interface("output", output).
		Msg("successfully finish GetUserData")

	return output, nil
}
