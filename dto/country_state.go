package dto

import "go-complaint/domain/model/common"

type CountryState struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewCountryState(countryState common.CountryState) *CountryState {
	return &CountryState{
		Id:   countryState.ID(),
		Name: countryState.Name(),
	}
}
