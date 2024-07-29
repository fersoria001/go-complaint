package complaint

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type ComplaintTitleSet struct {
	complaintId uuid.UUID
	newTitle    string
	occurredOn  time.Time
}

func NewComplaintTitleSet(complaintId uuid.UUID, newTitle string) *ComplaintTitleSet {
	return &ComplaintTitleSet{
		complaintId: complaintId,
		newTitle:    newTitle,
		occurredOn:  time.Now(),
	}
}

func (cts *ComplaintTitleSet) OccurredOn() time.Time {
	return cts.occurredOn
}

func (cts *ComplaintTitleSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ComplaintId uuid.UUID
		NewTitle    string
		OccurredOn  time.Time
	}{
		ComplaintId: cts.complaintId,
		NewTitle:    cts.newTitle,
		OccurredOn:  cts.occurredOn,
	})
}

func (cts *ComplaintTitleSet) UnmarshalJSON(data []byte) error {
	aux := struct {
		ComplaintId uuid.UUID
		NewTitle    string
		OccurredOn  time.Time
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	cts.complaintId = aux.ComplaintId
	cts.newTitle = aux.NewTitle
	cts.occurredOn = aux.OccurredOn
	return nil
}
