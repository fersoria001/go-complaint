package dto

import "go-complaint/domain/model/common"

type CountryState struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewCountryState(countryState common.CountryState) CountryState {
	return CountryState{
		ID:   countryState.ID(),
		Name: countryState.Name(),
	}
}
