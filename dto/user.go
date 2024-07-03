package dto

import "go-complaint/domain/model/identity"

type UserStatus int

const (
	OFFLINE UserStatus = iota
	ONLINE
)

func (s UserStatus) String() string {
	switch s {
	case OFFLINE:
		return "OFFLINE"
	case ONLINE:
		return "ONLINE"
	default:
		return "UNKNOWN"
	}
}

type User struct {
	ProfileIMG string  `json:"profileIMG"`
	Email      string  `json:"email"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	Gender     string  `json:"gender"`
	Pronoun    string  `json:"pronoun"`
	Age        int     `json:"age"`
	Phone      string  `json:"phone"`
	Address    Address `json:"address"`
	Status     string  `json:"status"`
}

func NewUser(domainObj identity.User) User {
	return User{
		ProfileIMG: domainObj.ProfileIMG(),
		Email:      domainObj.Email(),
		FirstName:  domainObj.FirstName(),
		LastName:   domainObj.LastName(),
		Gender:     domainObj.Gender(),
		Pronoun:    domainObj.Pronoun(),
		Age:        domainObj.Age(),
		Phone:      domainObj.Phone(),
		Address:    NewAddress(domainObj.Address()),
	}
}

func NewUserPtr(domainObj identity.User) *User {
	return &User{
		ProfileIMG: domainObj.ProfileIMG(),
		Email:      domainObj.Email(),
		FirstName:  domainObj.FirstName(),
		LastName:   domainObj.LastName(),
		Gender:     domainObj.Gender(),
		Pronoun:    domainObj.Pronoun(),
		Age:        domainObj.Age(),
		Phone:      domainObj.Phone(),
		Address:    NewAddress(domainObj.Address()),
	}
}

func (u *User) SetStatus(status UserStatus) {
	u.Status = status.String()
}
