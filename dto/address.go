package dto

import (
	"go-complaint/domain/model/common"
)

type Address struct {
	Country string `json:"country"`
	County  string `json:"county"`
	City    string `json:"city"`
}

func NewAddress(domainObj common.Address) *Address {
	return &Address{
		Country: domainObj.Country(),
		County:  domainObj.County(),
		City:    domainObj.City(),
	}
}
