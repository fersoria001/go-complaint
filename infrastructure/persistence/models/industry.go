package models

import "go-complaint/domain/model/enterprise"

type Industry struct {
	ID   int
	Name string
}

func (i *Industry) Columns() Columns {
	return Columns{
		"id",
		"name",
	}
}

func (i *Industry) Values() Values {
	return Values{
		&i.ID,
		&i.Name,
	}
}

func (i *Industry) Args() string {
	return "$1, $2"
}

func (i *Industry) Table() string {
	return "industries"
}

func (i *Industry) FromDomain(industry *enterprise.Industry) {
	i.ID = industry.ID()
	i.Name = industry.Name()
}