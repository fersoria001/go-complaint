package complaint

type ComplaintDataType int

const (
	UNKNOWN ComplaintDataType = iota
	SENT
	RESOLVED
	REVIEWED
	RECEIVED
)

func (cdt ComplaintDataType) String() string {
	switch cdt {
	case SENT:
		return "SENT"
	case RESOLVED:
		return "RESOLVED"
	case REVIEWED:
		return "REVIEWED"
	case RECEIVED:
		return "RECEIVED"
	default:
		return ""
	}
}

func ParseComplaintDataType(status string) ComplaintDataType {
	switch status {
	case "SENT":
		return SENT
	case "RESOLVED":
		return RESOLVED
	case "REVIEWED":
		return REVIEWED
	case "RECEIVED":
		return RECEIVED
	default:
		return UNKNOWN
	}
}
