package enterprise

import "go-complaint/erros"

type Position int

//There's more roles to be added
const (
	NOT_ASSIGNED Position = iota
	ASSISTANT
	MANAGER
)

func (p Position) String() string {
	switch p {
	case NOT_ASSIGNED:
		return "Not assigned"
	case ASSISTANT:
		return "Assistant"
	case MANAGER:
		return "Manager"
	default:
		return "Unknown"
	}
}

// this should return an error
func ParsePosition(s string) (Position, error) {
	switch s {
	case "Not assigned":
		return NOT_ASSIGNED, nil
	case "NOT_ASSIGNED":
		return NOT_ASSIGNED, nil
	case "ASSISTANT":
		return ASSISTANT, nil
	case "Assistant":
		return ASSISTANT, nil
	case "MANAGER":
		return MANAGER, nil
	case "Manager":
		return MANAGER, nil
	default:
		return -1, &erros.OutOfRangeError{}
	}
}
