package complaint

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type UserComplaintClosed struct {
	complaintID uuid.UUID
	authorID    string
	occurredOn  time.Time
}

func NewUserComplaintClosed(complaintID uuid.UUID, authorID string) *UserComplaintClosed {
	return &UserComplaintClosed{
		complaintID: complaintID,
		authorID:    authorID,
		occurredOn:  time.Now(),
	}
}

func (ucc *UserComplaintClosed) OccurredOn() time.Time {
	return ucc.occurredOn
}

func (ucc *UserComplaintClosed) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ComplaintID string `json:"complaint_id"`
		AuthorID    string `json:"author_id"`
		OccurredOn  string `json:"occurred_on"`
	}{
		ComplaintID: ucc.complaintID.String(),
		AuthorID:    ucc.authorID,
		OccurredOn:  common.StringDate(ucc.occurredOn),
	})
}

func (ucc *UserComplaintClosed) UnmarshalJSON(data []byte) error {
	aux := &struct {
		ComplaintID string `json:"complaint_id"`
		AuthorID    string `json:"author_id"`
		OccurredOn  string `json:"occurred_on"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	occurredOn, err := common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	complaintID, err := uuid.Parse(aux.ComplaintID)
	if err != nil {
		return err
	}
	ucc.complaintID = complaintID
	ucc.authorID = aux.AuthorID
	ucc.occurredOn = occurredOn
	return nil
}
