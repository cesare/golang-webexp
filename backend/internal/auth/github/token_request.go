package github

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"webexp/internal/configs"
)

type TokenRequest struct {
	config *configs.Config
	state  string
	code   string
}

func NewTokenRequest(config *configs.Config, state string, code string) *TokenRequest {
	return &TokenRequest{
		config: config,
		state:  state,
		code:   code,
	}
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (r *TokenRequest) Execute() (*TokenResponse, error) {
	response, err := r.request()
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	bytes := []byte(body)
	var token TokenResponse
	e := json.Unmarshal(bytes, &token)
	if e != nil {
		return nil, e
	}

	return &token, nil
}

func (r *TokenRequest) request() (*http.Response, error) {
	body := r.createRequestBody()
	request, err := r.createRequest(body)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(request)
}

func (r *TokenRequest) createRequest(body io.Reader) (*http.Request, error) {
	uri := "https://github.com/login/oauth/access_token"
	request, err := http.NewRequest(http.MethodPost, uri, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/json")

	return request, nil
}

func (r *TokenRequest) createRequestBody() io.Reader {
	params := url.Values{
		"client_id":     {r.config.Auth.ClientId},
		"client_secret": {r.config.Auth.ClientSecret},
		"code":          {r.code},
		"state":         {r.state},
	}
	return strings.NewReader(params.Encode())
}
