package dto

import "go-complaint/domain/model/enterprise"

type Industry struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
}

func NewIndustry(domainObj enterprise.Industry) Industry {
	return Industry{
		ID:   domainObj.ID(),
		Name: domainObj.Name(),
	}
}
