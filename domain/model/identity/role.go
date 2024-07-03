package identity

import "go-complaint/erros"

type RolesEnum int

const (
	ASSISTANT RolesEnum = iota
	MANAGER
	OWNER
)

func (r RolesEnum) String() string {
	switch r {
	case ASSISTANT:
		return "ASSISTANT"
	case MANAGER:
		return "MANAGER"
	case OWNER:
		return "OWNER"
	default:
		return "UNKNOWN"
	}
}

func ParseRole(role string) (RolesEnum, error) {
	switch role {
	case "ASSISTANT":
		return ASSISTANT, nil
	case "MANAGER":
		return MANAGER, nil
	case "OWNER":
		return OWNER, nil
	default:
		return -1, &erros.OutOfRangeError{}
	}
}

type Role struct {
	role RolesEnum
}

func NewRole(id string) (Role, error) {
	role, err := ParseRole(id)
	if err != nil {
		return Role{}, err
	}
	return Role{
		role: role,
	}, nil
}

func (r Role) GetRole() RolesEnum {
	return r.role
}
