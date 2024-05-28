package complaint

import "go-complaint/erros"

type Status int

const (
	OPEN Status = iota
	STARTED
	IN_DISCUSSION
	IN_REVIEW
	CLOSED
	IN_HISTORY
)

func (s Status) String() string {
	switch s {
	case OPEN:
		return "OPEN"
	case STARTED:
		return "STARTED"
	case IN_DISCUSSION:
		return "IN_DISCUSSION"
	case IN_REVIEW:
		return "IN_REVIEW"
	case CLOSED:
		return "CLOSED"
	case IN_HISTORY:
		return "IN_HISTORY"
	default:
		return "UNKNOWN"
	}
}

func ParseStatus(s string) (Status, error) {
	switch s {
	case "OPEN":
		return OPEN, nil
	case "STARTED":
		return STARTED, nil
	case "IN_DISCUSSION":
		return IN_DISCUSSION, nil
	case "IN_REVIEW":
		return IN_REVIEW, nil
	case "CLOSED":
		return CLOSED, nil
	case "IN_HISTORY":
		return IN_HISTORY, nil
	default:
		return -1, &erros.OutOfRangeError{}
	}
}
