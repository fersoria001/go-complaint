package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

// Package enterprise
// <<Domain event>> implements domain.DomainEvent
type HiringProccessCanceled struct {
	enterpriseId uuid.UUID
	candidateId  uuid.UUID
	emitedBy     uuid.UUID
	reason       string
	position     Position
	occurredOn   time.Time
}

func NewHiringProccessCanceled(enterpriseId,
	candidateId,
	emitedBy uuid.UUID,
	reason string,
	position Position,
) *HiringProccessCanceled {
	return &HiringProccessCanceled{
		enterpriseId: enterpriseId,
		candidateId:  candidateId,
		emitedBy:     emitedBy,
		reason:       reason,
		position:     position,
		occurredOn:   time.Now(),
	}
}

func (h HiringProccessCanceled) Reason() string {
	return h.reason
}

func (h HiringProccessCanceled) EmitedBy() uuid.UUID {
	return h.emitedBy
}

func (h HiringProccessCanceled) EnterpriseId() uuid.UUID {
	return h.enterpriseId
}

func (h HiringProccessCanceled) CandidateId() uuid.UUID {
	return h.candidateId
}

func (h HiringProccessCanceled) Position() Position {
	return h.position
}

func (h HiringProccessCanceled) OccurredOn() time.Time {
	return h.occurredOn
}

func (h *HiringProccessCanceled) MarshalJSON() ([]byte, error) {
	ms := common.StringDate(h.occurredOn)
	return json.Marshal(struct {
		EnterpriseId uuid.UUID
		CandidateId  uuid.UUID
		EmitedBy     uuid.UUID
		Reason       string
		Position     string
		OccurredOn   string
	}{
		EnterpriseId: h.enterpriseId,
		CandidateId:  h.candidateId,
		EmitedBy:     h.emitedBy,
		Reason:       h.reason,
		Position:     h.position.String(),
		OccurredOn:   ms,
	})
}

func (h *HiringProccessCanceled) UnmarshalJSON(data []byte) error {
	var v struct {
		EnterpriseId uuid.UUID
		CandidateId  uuid.UUID
		EmitedBy     uuid.UUID
		Reason       string
		Position     string
		OccurredOn   string
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	occurredOn, err := common.ParseDate(v.OccurredOn)
	if err != nil {
		return err
	}
	h.reason = v.Reason
	h.enterpriseId = v.EnterpriseId
	h.candidateId = v.CandidateId
	position := v.Position
	h.position = ParsePosition(position)
	h.emitedBy = v.EmitedBy
	h.occurredOn = occurredOn
	return nil
}
