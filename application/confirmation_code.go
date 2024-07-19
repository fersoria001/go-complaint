package application

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type ConfirmationCode struct {
	Code int
	jwt.StandardClaims
}

func CreateConfirmationCode() ConfirmationCode {
	// min := 1000000
	// max := 9999999
	// randomCode := rand.Intn(max-min+1) + min
	cc := ConfirmationCode{
		Code: 9999999,
	}
	cc.IssuedAt = time.Now().Unix()
	cc.ExpiresAt = time.Now().Add(time.Minute * 15).Unix()

	return cc
}
