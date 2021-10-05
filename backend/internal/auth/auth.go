package auth

import (
	"crypto/rand"
	"encoding/base64"
	"webexp/internal/configs"
)

type AuthAttributes struct {
	State       string
	CallbackUri string
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
	attrs := AuthAttributes{
		State:       state,
		CallbackUri: "http://dummy.localhost/auth/callback",
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
