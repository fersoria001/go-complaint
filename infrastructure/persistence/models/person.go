package models

import (
	"go-complaint/domain/model/identity"
)

type Person struct {
	Email     string
	FirstName string
	LastName  string
	BirthDate string
	Phone     string
	Country   string
	County    string
	City      string
}

func NewPerson(domain *identity.Person) *Person {
	return &Person{
		Email:     domain.Email(),
		FirstName: domain.FirstName(),
		LastName:  domain.LastName(),
		BirthDate: domain.BirthDate().StringRepresentation(),
		Phone:     domain.Phone(),
		Country:   domain.Address().Country(),
		County:    domain.Address().County(),
		City:      domain.Address().City(),
	}
}

func (p *Person) Columns() Columns {
	return Columns{
		"email",
		"first_name",
		"last_name",
		"birth_date",
		"phone",
		"country",
		"county",
		"city",
	}
}

func (p *Person) Values() Values {
	return Values{
		&p.Email,
		&p.FirstName,
		&p.LastName,
		&p.BirthDate,
		&p.Phone,
		&p.Country,
		&p.County,
		&p.City,
	}
}

func (p *Person) Args() string {
	return "$1, $2, $3, $4, $5, $6, $7, $8"
}

func (p *Person) Table() string {
	return "persons"
}
