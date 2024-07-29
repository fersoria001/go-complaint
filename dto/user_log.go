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
	EventId         string `json:"event_id"`
	ComplaintId     string `json:"complaint_id"`
	RatedBy         string `json:"rated_by"`
	AssistantUserId string `json:"assistant_user_id"`
	OccurredOn      string `json:"occurred_on"`
}

func NewComplaintRated(eventID string, domainEvent complaint.ComplaintRated) ComplaintRated {
	stringDate := common.NewDate(domainEvent.OccurredOn()).StringRepresentation()
	return ComplaintRated{
		EventId:         eventID,
		ComplaintId:     domainEvent.ComplaintId().String(),
		RatedBy:         domainEvent.RatedBy().String(),
		AssistantUserId: domainEvent.AssistantUserId().String(),
		OccurredOn:      stringDate,
	}
}
