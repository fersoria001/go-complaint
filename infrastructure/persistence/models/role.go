package models

import "go-complaint/domain/model/identity"

type Role struct {
	ID           string
	EnterpriseID string
}

func NewRole(domain identity.Role) Role {
	return Role{
		ID: domain.ID(),
	}
}

func (r Role) Columns() Columns {
	return Columns{
		"id",
	}
}

func (r Role) Values() Values {
	return Values{
		&r.ID,
	}
}

func (r Role) Args() string {
	return "$1"
}

func (r Role) Table() string {
	return "roles"
}
