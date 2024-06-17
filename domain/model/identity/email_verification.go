package identity

import (
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
)

type EmailVerification struct {
	Email string
	Code  int
	jwt.StandardClaims
}

func NewEmailVerification(email string) EmailVerification {
	min := 1000000
	max := 9999999
	code := rand.Intn(max-min+1) + min
	cc := EmailVerification{
		Email: email,
		Code:  code,
	}
	cc.IssuedAt = time.Now().Unix()
	cc.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	return cc
}
