package google

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

type exchangeTokenApiOutput struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func (self *googleAdapter) ExchangeCode(i *adapters.ExchangeCodeInput) (*adapters.ExchangeCodeOutput, error) {
	self.logger.Trace().Msg("start ExchangeCode")

	self.logger.Trace().Msg("building exchange code body")
	// ALERT: The order of the properties is important, don't change it!
	body := url.Values{}
	body.Set("code", i.Code)
	body.Set("client_id", self.secrets.GoogleClientId)
	body.Set("client_secret", self.secrets.GoogleClientSecret)
	body.Set("redirect_uri", self.secrets.GoogleRedirectUri)
	body.Set("grant_type", "authorization_code")
	// ALERT: The order of the properties is important, don't change it!
	self.logger.Trace().Msg("exchange code built")
	encodedBody := body.Encode()
	self.logger.Debug().Str("body", encodedBody).Msg("Exchange code built")

	self.logger.Trace().Msg("building request to exchange code")
	req, err := http.NewRequest(
		http.MethodPost,
		"https://oauth2.googleapis.com/token",
		strings.NewReader(body.Encode()),
	)
	if err != nil {
		self.logger.Error().Err(err).Msg(
			"fail to build request to exchange code",
		)
		return nil, errors.New("fail to build request")
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	self.logger.Trace().Msg("request built")

	self.logger.Trace().Msg("do request to exchange code")
	codeRes, err := self.httpClient.Do(req)
	if err != nil {
		self.logger.Error().Err(err).Msg(
			"fail to make request to exchange code",
		)
		return nil, errors.New("fail to make request")
	}
	defer codeRes.Body.Close()
	self.logger.Trace().Msg("request to exchange code done")

	self.logger.Debug().Int("status", codeRes.StatusCode).Msg("response status")

	if codeRes.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(codeRes.Body)
		if err != nil {
			self.logger.Error().Err(err).Msg(
				"fail to read response body",
			)
			return nil, errors.New("fail to make request")
		}

		bodyString := string(bodyBytes)
		self.logger.Error().Str("body", bodyString).Msg(
			"fail to make request",
		)
		return nil, errors.New("fail to make request")
	}

	self.logger.Trace().Msg("try to decode response body")
	exchangeCode := exchangeTokenApiOutput{}
	err = json.NewDecoder(codeRes.Body).Decode(&exchangeCode)
	if err != nil {
		self.logger.Error().Err(err).Msg(
			"fail to decode response body",
		)
		return nil, errors.New("fail to decode request body")
	}
	self.logger.Debug().Interface("body", exchangeCode).Msg("response body decoded")

	expDate := time.
		Now().
		Add(
			time.Duration(
				exchangeCode.ExpiresIn,
			),
		)

	self.logger.Debug().
		Time("time", expDate).
		Msg("expDate")

	output := &adapters.ExchangeCodeOutput{
		AccessToken:  exchangeCode.AccessToken,
		RefreshToken: exchangeCode.RefreshToken,
		Scopes:       strings.Split(exchangeCode.Scope, " "),
		ExpiresAt:    expDate,
	}

	self.logger.Debug().
		Interface("output", output).
		Msg("successfully finish GoogleAdapter.ExchangeCode")

	return output, nil
}
