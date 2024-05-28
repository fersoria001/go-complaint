package models

import (
	"go-complaint/domain/model/enterprise"
)

type Enterprise struct {
	Name           string
	OwnerID        string
	LogoIMG        string
	BannerIMG      string
	Website        string
	Email          string
	Phone          string
	Country        string
	County         string
	City           string
	Industry       string
	RegisterAt     string
	FoundationDate string
}

func NewEnterprise(domain *enterprise.Enterprise) *Enterprise {
	return &Enterprise{
		Name:           domain.Name(),
		OwnerID:        domain.OwnerID(),
		LogoIMG:        domain.LogoIMG(),
		BannerIMG:      domain.BannerIMG(),
		Website:        domain.Website(),
		Email:          domain.Email(),
		Phone:          domain.Phone(),
		Country:        domain.Address().Country(),
		County:         domain.Address().County(),
		City:           domain.Address().City(),
		Industry:       domain.Industry().Name(),
		RegisterAt:     domain.RegisterAt().StringRepresentation(),
		FoundationDate: domain.FoundationDate().StringRepresentation(),
	}

}

func (e *Enterprise) Columns() Columns {
	return Columns{
		"id",
		"owner_id",
		"logo_img",
		"banner_img",
		"website",
		"email",
		"phone",
		"country",
		"county",
		"city",
		"industry",
		"register_at",
		"foundation_date",
	}
}

func (e *Enterprise) Values() Values {
	return Values{
		&e.Name,
		&e.OwnerID,
		&e.LogoIMG,
		&e.BannerIMG,
		&e.Website,
		&e.Email,
		&e.Phone,
		&e.Country,
		&e.County,
		&e.City,
		&e.Industry,
		&e.RegisterAt,
		&e.FoundationDate,
	}
}

func (e *Enterprise) Args() string {
	return "$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13"
}

func (e *Enterprise) Table() string {
	return "enterprises"
}
