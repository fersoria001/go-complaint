package dto

import "go-complaint/domain/model/enterprise"

type Enterprise struct {
	Name           string  `json:"name"`
	LogoIMG        string  `json:"logo_img"`
	BannerIMG      string  `json:"banner_img"`
	Website        string  `json:"website"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	Address        Address `json:"address"`
	Industry       string  `json:"industry"`
	FoundationDate string  `json:"foundation_date"`
	OwnerID        string  `json:"owner_id"`
}

func NewEnterprise(domainObj *enterprise.Enterprise) *Enterprise {
	return &Enterprise{
		Name:      domainObj.Name(),
		LogoIMG:   domainObj.LogoIMG(),
		BannerIMG: domainObj.BannerIMG(),
		Website:   domainObj.Website(),
		Email:     domainObj.Email(),
		Phone:     domainObj.Phone(),
		Address: Address{
			Country: domainObj.Address().Country(),
			County:  domainObj.Address().County(),
			City:    domainObj.Address().City(),
		},
		Industry:       domainObj.Industry().Name(),
		FoundationDate: domainObj.FoundationDate().StringRepresentation(),
		OwnerID:        domainObj.OwnerID(),
	}
}
