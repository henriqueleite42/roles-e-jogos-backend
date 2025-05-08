package ludopedia

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

const ONE_YEAR_IN_HOURS = 8766 * time.Hour

type exchangeTokenApiOutput struct {
	AccessToken string `json:"access_token"`
}

func (self *ludopediaAdapter) ExchangeCode(i *adapters.ExchangeCodeInput) (*adapters.ExchangeCodeOutput, error) {
	self.logger.Trace().Msg("start ExchangeCode")

	self.logger.Trace().Msg("building exchange code body")
	// ALERT: The order of the properties is important, don't change it!
	body := url.Values{}
	body.Set("code", i.Code)
	body.Set("client_id", self.secrets.LudopediaClientId)
	body.Set("client_secret", self.secrets.LudopediaClientSecret)
	body.Set("redirect_uri", self.secrets.LudopediaRedirectUri)
	body.Set("grant_type", "authorization_code")
	// ALERT: The order of the properties is important, don't change it!
	self.logger.Trace().Msg("exchange code built")
	encodedBody := body.Encode()
	self.logger.Debug().Str("body", encodedBody).Msg("Exchange code built")

	self.logger.Trace().Msg("building request to exchange code")
	req, err := http.NewRequest(
		http.MethodPost,
		"https://ludopedia.com.br/tokenrequest",
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

	// As the token doesn't expire, we set the date to a loooong time
	expDate := time.
		Now().
		Add(10 * ONE_YEAR_IN_HOURS)

	self.logger.Debug().
		Time("time", expDate).
		Msg("expDate")

	output := &adapters.ExchangeCodeOutput{
		AccessToken: exchangeCode.AccessToken,
		Scopes:      []string{},
		ExpiresAt:   expDate,
	}

	self.logger.Debug().
		Interface("output", output).
		Msg("successfully finish LudopediaAdapter.ExchangeCode")

	return output, nil
}
