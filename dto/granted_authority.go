package dto

import (
	"go-complaint/domain/model/identity"

	"github.com/google/uuid"
)

type GrantedAuthority struct {
	EnterpriseID string `json:"enterpriseId"`
	Principal    string `json:"principal"`
	Authority    string `json:"authority"`
}

func NewGrantedAuthority(
	enterpriseID string,
	domainObj identity.GrantedAuthority,
) GrantedAuthority {
	return GrantedAuthority{
		EnterpriseID: enterpriseID,
		Principal:    domainObj.Principal(),
		Authority:    domainObj.Authority(),
	}
}

func NewGrantedAuthorities(
	domainObjs map[uuid.UUID][]identity.GrantedAuthority,
) []GrantedAuthority {
	var dtos []GrantedAuthority
	for key, domainObj := range domainObjs {
		for _, value := range domainObj {
			dtos = append(dtos, NewGrantedAuthority(key.String(), value))
		}
	}
	return dtos
}
