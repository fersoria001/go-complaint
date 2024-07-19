package dto

import "go-complaint/domain/model/identity"

type UserTypeList struct {
	Users         []User `json:"users"`
	Count         int    `json:"count"`
	CurrentLimit  int    `json:"current_limit"`
	CurrentOffset int    `json:"current_offset"`
	NextCursor    int    `json:"nextCursor"`
}

func NewUserTypeList(domainObjects []identity.User, limit, offset int) *UserTypeList {
	users := make([]User, 0)
	for _, user := range domainObjects {
		users = append(users, NewUser(user))
	}
	count := len(users)
	return &UserTypeList{
		Users:         users,
		Count:         count,
		CurrentLimit:  limit,
		CurrentOffset: offset,
	}
}
