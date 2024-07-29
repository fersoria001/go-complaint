package dto

import "go-complaint/domain/model/common"

type City struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	CountryCode string  `json:"country_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

func NewCity(domainObj common.City) *City {
	return &City{
		Id:          domainObj.ID(),
		Name:        domainObj.Name(),
		CountryCode: domainObj.CountryCode(),
		Latitude:    domainObj.Latitude(),
		Longitude:   domainObj.Longitude(),
	}
}
