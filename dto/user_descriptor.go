package dto

import (
	"go-complaint/domain/model/identity"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserDescriptor struct {
	Email              string             `json:"email"`
	FullName           string             `json:"full_name"`
	ProfileIMG         string             `json:"profile_img"`
	Gender             string             `json:"gender"`
	Pronoun            string             `json:"pronoun"`
	ClientData         ClientData         `json:"client_data"`
	RememberMe         bool               `json:"remember_me"`
	GrantedAuthorities []GrantedAuthority `json:"authorities"`
	jwt.StandardClaims
}

func NewUserDescriptor(
	clientData ClientData,
	user identity.User,
	rememberMe bool,
) UserDescriptor {
	thisDate := time.Now()
	ud := UserDescriptor{
		Email:              user.Email(),
		FullName:           user.FullName(),
		ProfileIMG:         user.ProfileIMG(),
		Gender:             user.Gender(),
		Pronoun:            user.Pronoun(),
		ClientData:         clientData,
		RememberMe:         rememberMe,
		GrantedAuthorities: NewGrantedAuthorities(user.Authorities()),
	}
	ud.IssuedAt = thisDate.Unix()
	ud.ExpiresAt = thisDate.Add(time.Hour * 24).Unix()
	return ud
}
