package dto

import "go-complaint/domain/model/identity"

type User struct {
	ProfileIMG string  `json:"profile_img"`
	Email      string  `json:"email"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Age        int     `json:"age"`
	Phone      string  `json:"phone"`
	Address    Address `json:"address"`
}

func NewUser(domainObj *identity.User) *User {
	return &User{
		ProfileIMG: domainObj.ProfileIMG(),
		Email:      domainObj.Email(),
		FirstName:  domainObj.Person().FirstName(),
		LastName:   domainObj.Person().LastName(),
		Age:        domainObj.Person().Age(),
		Phone:      domainObj.Person().Phone(),
		Address:    *NewAddress(domainObj.Person().Address()),
	}
}
