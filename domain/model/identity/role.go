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
	case "Not assigned":
		return -1, &erros.OutOfRangeError{}
	case "Assistant":
		return ASSISTANT, nil
	case "Manager":
		return MANAGER, nil
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
	id string
}

func NewRole(id RolesEnum) *Role {
	return &Role{id: id.String()}
}

func (r Role) ID() string {
	return r.id
}
