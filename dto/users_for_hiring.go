package dto

import "go-complaint/domain/model/identity"

type UsersForHiring struct {
	Users         []User `json:"users"`
	Count         int    `json:"count"`
	CurrentLimit  int    `json:"current_limit"`
	CurrentOffset int    `json:"current_offset"`
}

func NewUsersForHiring(domainObjects []identity.User, limit, offset int) *UsersForHiring {
	users := make([]User, 0)
	for _, user := range domainObjects {
		users = append(users, NewUser(user))
	}
	return &UsersForHiring{
		Users:         users[offset : limit+offset],
		Count:         len(users),
		CurrentLimit:  limit,
		CurrentOffset: offset,
	}
}
