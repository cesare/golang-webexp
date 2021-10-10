package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"webexp/internal/configs"
)

type CallbackAttributes struct {
	Code  string
	State string
}

type AuthResults struct {
	Token string
}

type AuthRejected struct {
}

func (*AuthRejected) Error() string {
	return "Authentication rejected"
}

type AuthFailed struct {
}

func (*AuthFailed) Error() string {
	return "Authentication Failed"
}

type Auth struct {
	config *configs.Config
	attrs  CallbackAttributes
}

func NewAuth(config *configs.Config, attrs CallbackAttributes) *Auth {
	return &Auth{config: config, attrs: attrs}
}

func (a *Auth) Execute() (*AuthResults, error) {
	_, err := a.requestToken()
	if err != nil {
		return nil, err
	}

	results := AuthResults{Token: "dummy"}
	return &results, nil
}

func (a *Auth) requestToken() (*tokenResponse, error) {
	request := tokenRequest{
		config: a.config,
		attrs:  a.attrs,
	}
	return request.execute()
}

type tokenRequest struct {
	config *configs.Config
	attrs  CallbackAttributes
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (r *tokenRequest) execute() (*tokenResponse, error) {
	params := url.Values{
		"client_id":     {r.config.Auth.ClientId},
		"client_secret": {r.config.Auth.ClientSecret},
		"code":          {r.attrs.Code},
		"state":         {r.attrs.State},
	}
	requestBody := strings.NewReader(params.Encode())

	uri := "https://github.com/login/oauth/access_token"
	request, err := http.NewRequest(http.MethodPost, uri, requestBody)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	bytes := []byte(body)
	var token tokenResponse
	e := json.Unmarshal(bytes, &token)
	if e != nil {
		return nil, e
	}

	return &token, nil
}
