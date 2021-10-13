package auth

import (
	"webexp/internal/auth/github"
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
	token, err := a.requestToken()
	if err != nil {
		return nil, err
	}

	user, err := a.requestUser(token.AccessToken)
	if err != nil {
		return nil, err
	}

	t, err := NewTokenGenerator(a.config, user.Id).Generate()
	if err != nil {
		return nil, err
	}

	results := AuthResults{Token: t}
	return &results, nil
}

func (a *Auth) requestToken() (*github.TokenResponse, error) {
	request := github.NewTokenRequest(a.config, a.attrs.State, a.attrs.Code)
	return request.Execute()
}

func (a *Auth) requestUser(accessToken string) (*github.UserResponse, error) {
	request := github.NewUserRequest(accessToken)
	return request.Execute()
}
