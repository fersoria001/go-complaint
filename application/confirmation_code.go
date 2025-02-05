package application

import (
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type ConfirmationCode struct {
	Code int
	jwt.StandardClaims
}

func CreateConfirmationCode() ConfirmationCode {
	env := os.Getenv("ENVIRONMENT")
	code := 9999999
	if env == "PROD" {
		min := 1000000
		max := 9999999
		code = rand.Intn(max-min+1) + min
	}
	cc := ConfirmationCode{
		Code: code,
	}
	cc.IssuedAt = time.Now().Unix()
	cc.ExpiresAt = time.Now().Add(time.Minute * 15).Unix()

	return cc
}
