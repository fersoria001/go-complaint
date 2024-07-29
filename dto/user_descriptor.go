package dto

import (
	"go-complaint/domain/model/identity"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserDescriptor struct {
	Id                 string             `json:"id"`
	Email              string             `json:"email"`
	FullName           string             `json:"full_name"`
	ProfileImg         string             `json:"profile_img"`
	Genre              string             `json:"gender"`
	Pronoun            string             `json:"pronoun"`
	ClientData         ClientData         `json:"client_data"`
	GrantedAuthorities []GrantedAuthority `json:"authorities"`
	jwt.StandardClaims
}

func NewUserDescriptor(clientData ClientData, user identity.User) *UserDescriptor {
	thisDate := time.Now()
	ud := &UserDescriptor{
		Id:                 user.Id().String(),
		Email:              user.Email(),
		FullName:           user.FullName(),
		ProfileImg:         user.ProfileIMG(),
		Genre:              user.Genre(),
		Pronoun:            user.Pronoun(),
		ClientData:         clientData,
		GrantedAuthorities: NewGrantedAuthorities(user.Authorities()),
	}
	ud.IssuedAt = thisDate.Unix()
	ud.ExpiresAt = thisDate.Add(time.Hour * 24).Unix()
	return ud
}
