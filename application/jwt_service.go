package application

import (
	"go-complaint/dto"

	"github.com/golang-jwt/jwt"
)

type JWTService struct {
}

func NewJWTService() *JWTService {
	return &JWTService{}
}
func (jwts *JWTService) GenerateJWTToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("JWT-supersecret-sign-password"))
}

func (jwts *JWTService) ParseUserDescriptor(jwtToken string) (*dto.UserDescriptor, error) {
	var claims = new(dto.UserDescriptor)
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("JWT-supersecret-sign-password"), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}
