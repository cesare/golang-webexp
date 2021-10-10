package auth

import "webexp/internal/configs"

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
	results := AuthResults{Token: "dummy"}
	return &results, nil
}
