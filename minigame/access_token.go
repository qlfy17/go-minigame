package minigame

import (
	"context"
	"fmt"
	"net/url"
)

type AccessTokenService struct {
	client *Client
}

type GetStableAccessTokenRequest struct {
	GrantType    string `json:"grant_type"`
	AppID        string `json:"appid"`
	Secret       string `json:"secret"`
	ForceRefresh bool   `json:"force_refresh"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (s *AccessTokenService) GetStableAccessToken(ctx context.Context) (*AccessToken, error) {
	url_ := fmt.Sprintf("%s/cgi-bin/stable_token", s.client.BaseURL)

	params := url.Values{
		"grant_type":    []string{""},
		"appid":         []string{""},
		"secret":        []string{""},
		"force_refresh": []string{"false"},
	}
	url_ += "?" + params.Encode()

	req, err := s.client.Get(url_, nil)
	if err != nil {
		return nil, err
	}

	at := new(AccessToken)
	_, err = s.client.Do(req, at)
	if err != nil {
		return nil, err
	}

	return at, nil
}
