package dto

import (
	"go-complaint/domain/model/identity"
)

type GrantedAuthority struct {
	EnterpriseID string `json:"enterprise_id"`
	Authority    string `json:"authority"`
}

func NewGrantedAuthority(
	enterpriseID string,
	domainObj identity.GrantedAuthority,
) GrantedAuthority {
	return GrantedAuthority{
		EnterpriseID: enterpriseID,
		Authority:    domainObj.Authority(),
	}
}

func NewGrantedAuthorities(
	domainObjs map[string][]identity.GrantedAuthority,
) []GrantedAuthority {
	var dtos []GrantedAuthority
	for key, domainObj := range domainObjs {
		for _, value := range domainObj {
			dtos = append(dtos, NewGrantedAuthority(key, value))
		}
	}
	return dtos
}
