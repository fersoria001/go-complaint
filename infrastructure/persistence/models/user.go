package models

import (
	"go-complaint/domain/model/identity"
)

type User struct {
	ProfileIMG   string
	RegisterDate string
	Email        string
	Password     string
}

func NewUser(domain *identity.User) *User {
	return &User{
		ProfileIMG:   domain.ProfileIMG(),
		RegisterDate: domain.RegisterDate().StringRepresentation(),
		Email:        domain.Email(),
		Password:     domain.Password(),
	}
}

func (u *User) Columns() Columns {
	return Columns{
		"profile_img",
		"register_date",
		"email",
		"password",
	}
}

func (u *User) Values() Values {
	return Values{
		&u.ProfileIMG,
		&u.RegisterDate,
		&u.Email,
		&u.Password,
	}
}

func (u *User) Args() string {
	return "$1, $2, $3, $4"
}

func (u *User) Table() string {
	return "users"
}
