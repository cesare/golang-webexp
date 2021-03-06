package auth

import (
	"time"
	"webexp/internal/configs"

	"github.com/golang-jwt/jwt/v4"
)

type TokenGenerator struct {
	config     *configs.Config
	identifier string
}

func NewTokenGenerator(config *configs.Config, identifier string) *TokenGenerator {
	return &TokenGenerator{
		config:     config,
		identifier: identifier,
	}
}

func (g *TokenGenerator) Generate() (string, error) {
	claims := g.createClaims()
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	return token.SignedString(g.config.Auth.TokenSigningKey.Bytes())
}

func (g *TokenGenerator) createClaims() jwt.RegisteredClaims {
	now := time.Now()
	issuedAt := jwt.NewNumericDate(now)

	ttl, _ := time.ParseDuration("1h")
	expiresAt := jwt.NewNumericDate(now.Add(ttl))

	claims := jwt.RegisteredClaims{
		Subject:   g.identifier,
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
	}
	return claims
}
