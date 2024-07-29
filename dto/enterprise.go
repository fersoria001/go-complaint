package dto

import (
	"go-complaint/domain/model/enterprise"
)

type Enterprise struct {
	Id             string      `json:"id"`
	Name           string      `json:"name"`
	LogoIMG        string      `json:"logo_img"`
	BannerIMG      string      `json:"banner_img"`
	Website        string      `json:"website"`
	Email          string      `json:"email"`
	Phone          string      `json:"phone"`
	Address        *Address    `json:"address"`
	Industry       string      `json:"industry"`
	FoundationDate string      `json:"foundation_date"`
	OwnerID        string      `json:"owner_id"`
	Employees      []*Employee `json:"employees"`
}

func NewEnterprise(obj *enterprise.Enterprise) *Enterprise {
	list := NewEmployeeList(obj.Employees())
	return &Enterprise{
		Id:        obj.Id().String(),
		Name:      obj.Name(),
		LogoIMG:   obj.LogoIMG(),
		BannerIMG: obj.BannerIMG(),
		Website:   obj.Website(),
		Email:     obj.Email(),
		Phone:     obj.Phone(),
		Address: &Address{
			Country: obj.Address().Country().Name(),
			County:  obj.Address().CountryState().Name(),
			City:    obj.Address().City().Name(),
		},
		Industry:       obj.Industry().Name(),
		FoundationDate: obj.FoundationDate().StringRepresentation(),
		OwnerID:        obj.OwnerId().String(),
		Employees:      list,
	}
}
