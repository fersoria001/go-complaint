package identity

type UserRole struct {
	role       *Role
	user       *User
	enterprise string
}

func NewUserRole(role *Role, user *User,
	enterprise string) *UserRole {
	return &UserRole{role: role, user: user, enterprise: enterprise}
}

func (ur *UserRole) Role() Role {
	return *ur.role
}

func (ur *UserRole) User() User {
	return *ur.user
}

func (ur *UserRole) Enterprise() string {
	return ur.enterprise
}

func (ur *UserRole) Equals(obj any) bool {
	var other *UserRole
	if obj == nil {
		return false
	}
	if obj != obj {
		return false
	}
	obj, ok := obj.(*UserRole)
	if !ok {
		return false
	}
	other = obj.(*UserRole)
	if ur.role.ID() != other.role.ID() {
		return false
	}
	if ur.user.Email() != other.user.Email() {
		return false
	}
	if ur.Enterprise() != other.Enterprise() {
		return false
	}
	return true
}
