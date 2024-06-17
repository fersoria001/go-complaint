package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
)

type UserLog struct {
	Count          int              `json:"count"`
	ComplaintRated []ComplaintRated `json:"complaint_rated"`
}

type ComplaintRated struct {
	EventID         string `json:"event_id"`
	ComplaintID     string `json:"complaint_id"`
	RatedBy         string `json:"rated_by"`
	AssistantUserID string `json:"assistant_user_id"`
	OccurredOn      string `json:"occurred_on"`
}

func NewComplaintRated(eventID string, domainEvent complaint.ComplaintRated) ComplaintRated {
	stringDate := common.NewDate(domainEvent.OccurredOn()).StringRepresentation()
	return ComplaintRated{
		EventID:         eventID,
		ComplaintID:     domainEvent.ComplaintID().String(),
		RatedBy:         domainEvent.RatedBy(),
		AssistantUserID: domainEvent.AssistantUserID(),
		OccurredOn:      stringDate,
	}
}
