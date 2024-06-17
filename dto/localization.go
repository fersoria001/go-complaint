package dto

import "go-complaint/domain/model/common"

type Localization struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func NewLocalization(domainObj common.Localization) Localization {
	return Localization{
		Latitude:  domainObj.Latitude(),
		Longitude: domainObj.Longitude(),
	}
}
