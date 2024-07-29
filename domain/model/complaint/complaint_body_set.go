package complaint

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type ComplaintBodySet struct {
	complaintId uuid.UUID
	newBody     string
	occurredOn  time.Time
}

func NewComplaintBodySet(complaintId uuid.UUID, newBody string) *ComplaintBodySet {
	return &ComplaintBodySet{
		complaintId: complaintId,
		newBody:     newBody,
		occurredOn:  time.Now(),
	}
}

func (cbs *ComplaintBodySet) OccurredOn() time.Time {
	return cbs.occurredOn
}

func (cbs *ComplaintBodySet) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ComplaintId uuid.UUID
		NewBody     string
		OccurredOn  time.Time
	}{
		ComplaintId: cbs.complaintId,
		NewBody:     cbs.newBody,
		OccurredOn:  cbs.occurredOn,
	})
}

func (cbs *ComplaintBodySet) UnmarshalJSON(data []byte) error {
	aux := struct {
		ComplaintId uuid.UUID
		NewBody     string
		OccurredOn  time.Time
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	cbs.complaintId = aux.ComplaintId
	cbs.newBody = aux.NewBody
	cbs.occurredOn = aux.OccurredOn
	return nil
}
