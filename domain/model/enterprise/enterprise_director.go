package enterprise

import (
	"go-complaint/domain/model/common"
)

type EnterpriseDirector struct {
	enterprise *Enterprise
}

func (ed *EnterpriseDirector) Show() interface{} {
	return EnterpriseInterest{
		Name:      ed.enterprise.Name(),
		LogoIMG:   ed.enterprise.LogoIMG(),
		BannerIMG: ed.enterprise.BannerIMG(),
		Website:   ed.enterprise.Website(),
		Email:     ed.enterprise.Email(),
		Phone:     ed.enterprise.Phone(),
		Address: common.AddressInterest{
			Country:      ed.enterprise.Address().Country().Name(),
			CountryState: ed.enterprise.Address().CountryState().Name(),
			City:         ed.enterprise.Address().City().Name(),
		},
		Industry:       ed.enterprise.Industry().Name(),
		FoundationDate: ed.enterprise.FoundationDate().StringRepresentation(),
		OwnerID:        ed.enterprise.Owner(),
	}
}

// func (ed *EnterpriseDirector) Create() {
// 	ed.enterprise = NewEnterpriseWithDirector(ed)
// }

func (ed *EnterpriseDirector) Changed(obj interface{}) error {
	switch obj := obj.(type) {
	case *Enterprise:
		ed.enterprise = obj
		return nil
	default:
		return ErrUnknownColleagueType
	}
}
