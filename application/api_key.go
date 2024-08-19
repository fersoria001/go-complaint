package application

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type ApiKey struct {
	User string `json:"user"`
	jwt.StandardClaims
}

func NewApiKey(user string) ApiKey {
	apiKey := new(ApiKey)
	year := time.Minute * 60 * 60 * 365
	apiKey.User = user
	apiKey.IssuedAt = time.Now().Unix()
	apiKey.ExpiresAt = time.Now().Add(year * 2).Unix()
	return *apiKey
}
