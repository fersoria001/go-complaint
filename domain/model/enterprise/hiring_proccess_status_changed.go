package enterprise

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type HiringProccessStatusChanged struct {
	hiringProccessId uuid.UUID
	newStatus        string
	updatedById      uuid.UUID
	occurredOn       time.Time
}

func NewHiringProccessStatusChanged(
	hiringProccessId uuid.UUID,
	newStatus string,
	updatedById uuid.UUID,
) *HiringProccessStatusChanged {
	return &HiringProccessStatusChanged{
		hiringProccessId: hiringProccessId,
		newStatus:        newStatus,
		updatedById:      updatedById,
		occurredOn:       time.Now(),
	}
}

func (hpsc HiringProccessStatusChanged) HiringProccessId() uuid.UUID {
	return hpsc.hiringProccessId
}

func (hpsc HiringProccessStatusChanged) NewStatus() string {
	return hpsc.newStatus
}

func (hpsc HiringProccessStatusChanged) UpdatedById() uuid.UUID {
	return hpsc.updatedById
}

func (hpsc HiringProccessStatusChanged) OccurredOn() time.Time {
	return hpsc.occurredOn
}

func (hpsc *HiringProccessStatusChanged) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		HiringProccessId uuid.UUID `json:"hiringProccessId"`
		NewStatus        string    `json:"newStatus"`
		UpdatedById      uuid.UUID `json:"updatedById"`
		OccurredOn       time.Time `json:"occurredOn"`
	}{
		HiringProccessId: hpsc.hiringProccessId,
		NewStatus:        hpsc.newStatus,
		UpdatedById:      hpsc.updatedById,
		OccurredOn:       hpsc.occurredOn,
	})
}

func (hpsc *HiringProccessStatusChanged) UnmarshalJSON(data []byte) error {
	aux := struct {
		HiringProccessId uuid.UUID `json:"hiringProccessId"`
		NewStatus        string    `json:"newStatus"`
		UpdatedById      uuid.UUID `json:"updatedById"`
		OccurredOn       time.Time `json:"occurredOn"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	hpsc.hiringProccessId = aux.HiringProccessId
	hpsc.newStatus = aux.NewStatus
	hpsc.updatedById = aux.UpdatedById
	hpsc.occurredOn = aux.OccurredOn
	return nil
}
