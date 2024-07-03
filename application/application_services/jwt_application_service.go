package application_services

import (
	"go-complaint/application"
	"go-complaint/domain/model/identity"
	"go-complaint/dto"
	"os"
	"sync"

	"github.com/golang-jwt/jwt"
)

var jwtServiceInstance *JWTApplicationService
var jwtServiceOnce sync.Once

func JWTApplicationServiceInstance() *JWTApplicationService {
	jwtServiceOnce.Do(func() {
		jwtServiceInstance = NewJWTApplicationService()
	})
	return jwtServiceInstance
}

type JWTApplicationService struct {
}

func NewJWTApplicationService() *JWTApplicationService {
	return &JWTApplicationService{}
}
func (jwts *JWTApplicationService) GenerateJWTToken(claims jwt.Claims) (
	application.JWTToken,
	error,
) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return application.JWTToken{}, err
	}
	return application.NewJWTToken(tokenString), nil
}

func (jwts *JWTApplicationService) ParseUserDescriptor(jwtToken string) (dto.UserDescriptor, error) {
	var claims dto.UserDescriptor
	token, err := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return claims, err
	}
	if !token.Valid {
		return claims, err
	}
	if claims.Email == "" {
		return claims, ErrInvalidToken
	}
	return claims, nil
}

func (jwts *JWTApplicationService) ParseConfirmationCode(jwtToken string) (
	application.ConfirmationCode,
	error,
) {
	var claims application.ConfirmationCode
	token, err := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return claims, err
	}
	if !token.Valid {
		if claims.ExpiresAt > 0 && claims.ExpiresAt < jwt.TimeFunc().Unix() {
			return claims, ErrTokenExpired
		}
		return claims, err
	}
	return claims, nil
}

func (jwts *JWTApplicationService) ParseEmailVerification(jwtToken string) (
	identity.EmailVerification,
	error,
) {
	var claims identity.EmailVerification
	token, err := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return claims, err
	}
	if !token.Valid {
		if claims.ExpiresAt > 0 && claims.ExpiresAt < jwt.TimeFunc().Unix() {
			return claims, ErrTokenExpired
		}
		return claims, err
	}
	if claims.Email == "" || claims.Code == 0 {
		return claims, ErrInvalidToken
	}
	return claims, nil
}
