package enterprise

import "strings"

type Position int

//There's more roles to be added
const (
	ASSISTANT Position = iota
	MANAGER
)

func (p Position) String() string {
	switch p {
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
		return -1
	}
}
