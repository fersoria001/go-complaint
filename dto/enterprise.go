package dto

import (
	"go-complaint/domain/model/employee"
	"go-complaint/domain/model/enterprise"

	mapset "github.com/deckarep/golang-set/v2"
)

type Enterprise struct {
	Name           string     `json:"name"`
	LogoIMG        string     `json:"logo_img"`
	BannerIMG      string     `json:"banner_img"`
	Website        string     `json:"website"`
	Email          string     `json:"email"`
	Phone          string     `json:"phone"`
	Address        Address    `json:"address"`
	Industry       string     `json:"industry"`
	FoundationDate string     `json:"foundation_date"`
	OwnerID        string     `json:"owner_id"`
	Employees      []Employee `json:"employees"`
}

func NewEnterprise(domainObj *enterprise.Enterprise) Enterprise {
	casted := mapset.NewSet[employee.Employee]()
	for emp := range domainObj.Employees().Iter() {
		cast, ok := emp.(*employee.Employee)
		if ok {
			casted.Add(*cast)
		}
	}
	list := NewEmployeeList(casted)
	dereferencedList := make([]Employee, 0)
	for _, emp := range list {
		dereferencedList = append(dereferencedList, *emp)
	}
	return Enterprise{
		Name:      domainObj.Name(),
		LogoIMG:   domainObj.LogoIMG(),
		BannerIMG: domainObj.BannerIMG(),
		Website:   domainObj.Website(),
		Email:     domainObj.Email(),
		Phone:     domainObj.Phone(),
		Address: Address{
			Country: domainObj.Address().Country().Name(),
			County:  domainObj.Address().CountryState().Name(),
			City:    domainObj.Address().City().Name(),
		},
		Industry:       domainObj.Industry().Name(),
		FoundationDate: domainObj.FoundationDate().StringRepresentation(),
		OwnerID:        domainObj.Owner(),
		Employees:      dereferencedList,
	}
}
