package identity

type UserRole struct {
	userID       string
	enterpriseID string
	Role
}

func NewUserRole(
	role Role,
	userID string,
	enterpriseID string,
) *UserRole {
	return &UserRole{
		Role:         role,
		userID:       userID,
		enterpriseID: enterpriseID,
	}
}

func (ur *UserRole) SetRole(role RolesEnum) {
	ur.role = role
}

func (ur UserRole) UserID() string {
	return ur.userID
}

func (ur UserRole) EnterpriseID() string {
	return ur.enterpriseID
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
	if ur.UserID() != other.UserID() {
		return false
	}
	if ur.EnterpriseID() != other.EnterpriseID() {
		return false
	}
	return true
}
