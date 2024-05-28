package dto

type UsersForHiring struct {
	Users         []*User `json:"users"`
	Count         int     `json:"count"`
	CurrentLimit  int     `json:"current_limit"`
	CurrentOffset int     `json:"current_offset"`
}
