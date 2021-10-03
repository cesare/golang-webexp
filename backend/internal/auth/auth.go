package auth

import "webexp/internal/configs"

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

func (*AuthStart) Execute() *AuthAttributes {
	return &AuthAttributes{
		State:       "dummy",
		CallbackUri: "http://dummy.localhost/auth/callback",
	}
}
