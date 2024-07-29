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

type Person struct {
	ProfileImg string  `json:"profileImg"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	Email      string  `json:"email"`
	Genre      string  `json:"genre"`
	Pronoun    string  `json:"pronoun"`
	Age        int     `json:"age"`
	Phone      string  `json:"phone"`
	Address    Address `json:"address"`
}
type User struct {
	Id       string  `json:"id"`
	Username string  `json:"username"`
	Person   *Person `json:"person"`
	Status   string  `json:"status"`
}

func NewUser(obj *identity.User) *User {
	return &User{
		Id:       obj.Id().String(),
		Username: obj.UserName(),
		Person: &Person{
			ProfileImg: obj.ProfileIMG(),
			Email:      obj.Email(),
			FirstName:  obj.FirstName(),
			LastName:   obj.LastName(),
			Genre:      obj.Genre(),
			Pronoun:    obj.Pronoun(),
			Age:        obj.Age(),
			Phone:      obj.Phone(),
			Address:    NewAddress(obj.Address()),
		},
	}
}

func (u *User) SetStatus(status UserStatus) {
	u.Status = status.String()
}
