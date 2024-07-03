package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

// Package enterprise
// <<Domain event>> implements domain.DomainEvent
type HiringProccessCanceled struct {
	enterpriseID string
	candidateID  string
	emitedBy     string
	reason       string
	position     Position
	occurredOn   time.Time
}

func NewHiringProccessCanceled(enterpriseID string,
	candidateID string,
	emitedBy string,
	reason string,
	position Position,
) *HiringProccessCanceled {
	return &HiringProccessCanceled{
		enterpriseID: enterpriseID,
		candidateID:  candidateID,
		emitedBy:     emitedBy,
		reason:       reason,
		position:     position,
		occurredOn:   time.Now(),
	}
}

func (h HiringProccessCanceled) Reason() string {
	return h.reason
}

func (h HiringProccessCanceled) EmitedBy() string {
	return h.emitedBy
}

func (h HiringProccessCanceled) EnterpriseID() string {
	return h.enterpriseID
}

func (h HiringProccessCanceled) CandidateID() string {
	return h.candidateID
}

func (h HiringProccessCanceled) Position() Position {
	return h.position
}

func (h HiringProccessCanceled) OccurredOn() time.Time {
	return h.occurredOn
}

func (h *HiringProccessCanceled) MarshalJSON() ([]byte, error) {
	ms := common.StringDate(h.occurredOn)
	return json.Marshal(map[string]interface{}{
		"enterprise_id": h.enterpriseID,
		"candidate_id":  h.candidateID,
		"emited_by":     h.emitedBy,
		"reason":        h.reason,
		"position":      h.position.String(),
		"occurred_on":   ms,
	})
}

func (h *HiringProccessCanceled) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	occurredOn, err := common.ParseDate(v["occurred_on"].(string))
	if err != nil {
		return err
	}
	h.reason = v["reason"].(string)
	h.enterpriseID = v["enterprise_id"].(string)
	h.candidateID = v["candidate_id"].(string)
	position := v["position"].(string)
	h.position = ParsePosition(position)
	h.emitedBy = v["emited_by"].(string)
	h.occurredOn = occurredOn
	return nil
}
