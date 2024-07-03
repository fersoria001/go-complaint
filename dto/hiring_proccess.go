package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"time"
)

type HiringProccessState int

const (
	PENDING HiringProccessState = iota
	ACCEPTED
	REJECTED
	USER_ACCEPTED
	CANCELED
	HIRED
	RATED
	WAITING
	FIRED
	LEAVED
)

func (s HiringProccessState) String() string {
	switch s {
	case PENDING:
		return "pending"
	case REJECTED:
		return "rejected"
	case CANCELED:
		return "canceled"
	case ACCEPTED:
		return "accepted"
	case USER_ACCEPTED:
		return "user_accepted"
	case HIRED:
		return "hired"
	case RATED:
		return "rated"
	case WAITING:
		return "waiting"
	case FIRED:
		return "fired"
	case LEAVED:
		return "leaved"
	default:
		return "unknown"
	}
}
func ParseHiringProccessState(s string) HiringProccessState {
	switch s {
	case "pending":
		return PENDING
	case "rejected":
		return REJECTED
	case "canceled":
		return CANCELED
	case "accepted":
		return ACCEPTED
	case "user_accepted":
		return USER_ACCEPTED
	case "hired":
		return HIRED
	case "rated":
		return RATED
	case "waiting":
		return WAITING
	case "fired":
		return FIRED
	case "leaved":
		return LEAVED
	default:
		return -1
	}
}

type HiringProccessList struct {
	HiringProccesses []HiringProccess `json:"hiring_proccesses"`
	Count            int              `json:"count"`
	CurrentOffset    int              `json:"current_offset"`
	CurrentLimit     int              `json:"current_limit"`
}

type HiringProccess struct {
	EventID    string `json:"event_id"`
	User       User   `json:"user"`
	Position   string `json:"position"`
	Status     string `json:"status"`
	Reason     string `json:"reason"`
	EmitedByID string `json:"emited_by_id"`
	EmitedBy   User   `json:"emited_by"`
	OccurredOn string `json:"occurred_on"`
	LastUpdate string `json:"last_update"`
}

func NewHiringProccess(
	occurredOn time.Time,
	user identity.User,
	position enterprise.Position,
	emitedByID string,
) *HiringProccess {
	stringDate := common.StringDate(occurredOn)
	return &HiringProccess{
		OccurredOn: stringDate,
		User:       NewUser(user),
		Position:   position.String(),
		EmitedByID: emitedByID,
	}
}

func (hiringProccess *HiringProccess) SetOcurredOn(occurredOn time.Time) {
	hiringProccess.OccurredOn = common.StringDate(occurredOn)
}

func (hiringProccess HiringProccess) GetOcurredOn() time.Time {
	d, _ := common.ParseDate(hiringProccess.OccurredOn)
	return d
}

func (hiringProccess *HiringProccess) SetEventID(eventID string) {
	hiringProccess.EventID = eventID
}

func (hiringProccess *HiringProccess) SetEmitedByID(emitedByID string) {
	hiringProccess.EmitedByID = emitedByID
}

func (hiringProccess *HiringProccess) SetEmitedBy(user identity.User) {
	hiringProccess.EmitedBy = NewUser(user)
}

func (hiringProccess *HiringProccess) SetStatus(status HiringProccessState) {
	parsed := ParseHiringProccessState(hiringProccess.Status)
	if parsed == -1 {
		hiringProccess.Status = status.String()
		return
	}
	if parsed < status {
		hiringProccess.Status = status.String()
		return
	}
}

func (hiringProccess *HiringProccess) SetReason(reason string) {
	hiringProccess.Reason = reason
}

func (hiringProccess *HiringProccess) SetLastUpdate(occurredOn time.Time) {
	hiringProccess.LastUpdate = common.StringDate(occurredOn)
}

func (hiringProccess HiringProccess) GetLastUpdate() time.Time {
	d, _ := common.ParseDate(hiringProccess.LastUpdate)
	return d
}
