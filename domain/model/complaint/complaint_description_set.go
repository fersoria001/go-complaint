package complaint

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type ComplaintDescriptionSet struct {
	complaintId    uuid.UUID
	newDescription string
	occurredOn     time.Time
}

func NewComplaintDescriptionSet(complaintId uuid.UUID, newDescription string) *ComplaintDescriptionSet {
	return &ComplaintDescriptionSet{
		complaintId:    complaintId,
		newDescription: newDescription,
		occurredOn:     time.Now(),
	}
}

func (cds *ComplaintDescriptionSet) OccurredOn() time.Time {
	return cds.occurredOn
}

func (cds *ComplaintDescriptionSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ComplaintId    uuid.UUID
		NewDescription string
		OccurredOn     time.Time
	}{
		ComplaintId:    cds.complaintId,
		NewDescription: cds.newDescription,
		OccurredOn:     cds.occurredOn,
	})
}

func (cds *ComplaintDescriptionSet) UnmarshalJSON(data []byte) error {
	aux := struct {
		ComplaintId    uuid.UUID
		NewDescription string
		OccurredOn     time.Time
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	cds.complaintId = aux.ComplaintId
	cds.newDescription = aux.NewDescription
	cds.occurredOn = aux.OccurredOn
	return nil
}
