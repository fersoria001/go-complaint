package models

import (
	"go-complaint/domain/model/identity"

	mapset "github.com/deckarep/golang-set/v2"
)

type UserRole struct {
	UserID       string
	RoleID       string
	EnterpriseID string
}

func NewUserRoles(domain mapset.Set[identity.UserRole]) []UserRole {
	var userRoles []UserRole = make([]UserRole, 0)
	for userRole := range domain.Iter() {
		userRoles = append(userRoles, NewUserRole(userRole))
	}
	return userRoles
}

func NewUserRole(domain identity.UserRole) UserRole {
	return UserRole{
		UserID:       domain.User().Email(),
		RoleID:       domain.Role().ID(),
		EnterpriseID: domain.Enterprise(),
	}
}

func (ur *UserRole) Columns() Columns {
	return Columns{
		"user_id",
		"role_id",
		"enterprise_id",
	}
}

func (ur *UserRole) Values() Values {
	return Values{
		&ur.UserID,
		&ur.RoleID,
		&ur.EnterpriseID,
	}
}

func (ur *UserRole) Args() string {
	return "$1, $2, $3"
}

func (ur *UserRole) Table() string {
	return "user_roles"
}
