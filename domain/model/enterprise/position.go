package enterprise

import "strings"

type Position int

//There's more roles to be added
const (
	NOT_EXISTS Position = iota
	ASSISTANT
	MANAGER
)

func (p Position) String() string {
	switch p {
	case NOT_EXISTS:
		return "NOT_EXISTS"
	case ASSISTANT:
		return "ASSISTANT"
	case MANAGER:
		return "MANAGER"
	default:
		return "NOT_EXISTS"
	}
}
func ParsePosition(s string) Position {
	upper := strings.ToUpper(s)
	switch upper {
	case "ASSISTANT":
		return ASSISTANT
	case "MANAGER":
		return MANAGER
	default:
		return NOT_EXISTS
	}
}
