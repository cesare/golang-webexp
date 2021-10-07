package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/url"
	"webexp/internal/configs"
)

type AuthAttributes struct {
	State            string
	AuthorizationUri string
}

type AuthStart struct {
	config *configs.Config
}

func NewAuthStart(config *configs.Config) *AuthStart {
	return &AuthStart{config: config}
}

func (as *AuthStart) Execute() (*AuthAttributes, error) {
	rawState, err := as.generateRawState()
	if err != nil {
		return nil, err
	}

	state := base64.RawURLEncoding.EncodeToString(rawState)
	authorizationUri := as.authUri(state)

	attrs := AuthAttributes{
		State:            state,
		AuthorizationUri: authorizationUri,
	}
	return &attrs, nil
}

func (as *AuthStart) generateRawState() ([]byte, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (as *AuthStart) authUri(state string) string {
	clientId := as.config.Auth.ClientId
	callbackUri := as.config.Frontend.CallbackUri()

	params := url.Values{}
	params.Set("state", state)
	params.Set("client_id", clientId)
	params.Set("redirect_uri", callbackUri)

	uri := url.URL{
		Scheme:   "https",
		Host:     "github.com",
		Path:     "/login/oauth/authorize",
		RawQuery: params.Encode(),
	}
	return uri.String()
}
