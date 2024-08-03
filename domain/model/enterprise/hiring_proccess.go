package enterprise

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/recipient"
	"time"

	"github.com/google/uuid"
)

type HiringProccessStatus int

const (
	PENDING HiringProccessStatus = iota
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

func (s HiringProccessStatus) String() string {
	switch s {
	case PENDING:
		return "PENDING"
	case REJECTED:
		return "REJECTED"
	case CANCELED:
		return "CANCELED"
	case ACCEPTED:
		return "ACCEPTED"
	case USER_ACCEPTED:
		return "USER_ACCEPTED"
	case HIRED:
		return "HIRED"
	case RATED:
		return "RATED"
	case WAITING:
		return "WAITING"
	case FIRED:
		return "FIRED"
	case LEAVED:
		return "LEAVED"
	default:
		return "UNKNOWN"
	}
}
func ParseHiringProccessStatus(s string) HiringProccessStatus {
	switch s {
	case "PENDING":
		return PENDING
	case "REJECTED":
		return REJECTED
	case "CANCELED":
		return CANCELED
	case "ACCEPTED":
		return ACCEPTED
	case "USER_ACCEPTED":
		return USER_ACCEPTED
	case "HIRED":
		return HIRED
	case "RATED":
		return RATED
	case "WAITING":
		return WAITING
	case "FIRED":
		return FIRED
	case "LEAVED":
		return LEAVED
	default:
		return -1
	}
}

type HiringProccess struct {
	id         uuid.UUID
	enterprise recipient.Recipient
	user       recipient.Recipient
	role       Position
	status     HiringProccessStatus
	reason     string
	emitedBy   recipient.Recipient
	occurredOn time.Time
	lastUpdate time.Time
	updatedBy  recipient.Recipient
}

func NewHiringProccess(
	id uuid.UUID,
	enterprise, user recipient.Recipient,
	role Position,
	status HiringProccessStatus,
	reason string,
	emitedBy recipient.Recipient,
	occurredOn, lastUpdate time.Time,
	updatedBy recipient.Recipient,
) *HiringProccess {
	return &HiringProccess{
		id:         id,
		enterprise: enterprise,
		user:       user,
		role:       role,
		status:     status,
		reason:     reason,
		emitedBy:   emitedBy,
		occurredOn: occurredOn,
		lastUpdate: lastUpdate,
		updatedBy:  updatedBy,
	}
}

func (h *HiringProccess) ChangeStatus(ctx context.Context, status HiringProccessStatus, updatedBy recipient.Recipient) error {
	h.status = status
	h.updatedBy = updatedBy
	h.lastUpdate = time.Now()
	pub := domain.DomainEventPublisherInstance()
	return pub.Publish(ctx, NewHiringProccessStatusChanged(
		h.id,
		h.status.String(),
		h.updatedBy.Id(),
	))
}

func (h *HiringProccess) WriteAReason(reason string, updatedBy recipient.Recipient) {
	h.reason = reason
	h.updatedBy = updatedBy
	h.lastUpdate = time.Now()
}

func (h HiringProccess) Id() uuid.UUID {
	return h.id
}
func (h HiringProccess) Enterprise() recipient.Recipient {
	return h.enterprise
}
func (h HiringProccess) User() recipient.Recipient {
	return h.user
}
func (h HiringProccess) Role() Position {
	return h.role
}
func (h HiringProccess) Status() HiringProccessStatus {
	return h.status
}
func (h HiringProccess) Reason() string {
	return h.reason
}
func (h HiringProccess) EmitedBy() recipient.Recipient {
	return h.emitedBy
}
func (h HiringProccess) OccurredOn() time.Time {
	return h.occurredOn
}
func (h HiringProccess) LastUpdate() time.Time {
	return h.lastUpdate
}
func (h HiringProccess) UpdatedBy() recipient.Recipient {
	return h.updatedBy
}
