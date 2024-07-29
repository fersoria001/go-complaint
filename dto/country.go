package dto

import "go-complaint/domain/model/common"

type Country struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	PhoneCode string `json:"phone_code"`
}

func NewCountry(country common.Country) *Country {
	return &Country{
		Id:        country.ID(),
		Name:      country.Name(),
		PhoneCode: country.PhoneCode(),
	}
}
