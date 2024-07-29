package identity

import "github.com/google/uuid"

type UserRole struct {
	userId       uuid.UUID
	enterpriseId uuid.UUID
	Role
}

func NewUserRole(
	role Role,
	userId,
	enterpriseId uuid.UUID,
) *UserRole {
	return &UserRole{
		Role:         role,
		userId:       userId,
		enterpriseId: enterpriseId,
	}
}

func (ur *UserRole) SetRole(role RolesEnum) {
	ur.role = role
}

func (ur UserRole) UserId() uuid.UUID {
	return ur.userId
}

func (ur UserRole) EnterpriseId() uuid.UUID {
	return ur.enterpriseId
}

func (ur UserRole) Equals(obj any) bool {
	var other UserRole
	if obj == nil {
		return false
	}
	if obj != obj {
		return false
	}
	other, ok := obj.(UserRole)
	if !ok {
		return false
	}
	if ur.GetRole() != other.GetRole() {
		return false
	}
	if ur.UserId() != other.UserId() {
		return false
	}
	if ur.EnterpriseId() != other.EnterpriseId() {
		return false
	}
	return true
}
