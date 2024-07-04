package application

import (
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
)

type ConfirmationCode struct {
	Code int
	jwt.StandardClaims
}

func CreateConfirmationCode() ConfirmationCode {
	min := 1000000
	max := 9999999
	randomCode := rand.Intn(max-min+1) + min
	cc := ConfirmationCode{
		Code: randomCode,
	}
	cc.IssuedAt = time.Now().Unix()
	cc.ExpiresAt = time.Now().Add(time.Minute * 15).Unix()

	return cc
}
