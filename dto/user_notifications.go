package dto

type UserNotifications struct {
	Count            int                `json:"count"`
	HiringInvitation []HiringInvitation `json:"hiring_invitation"`
	WaitingForReview []WaitingForReview `json:"waiting_for_review"`
}

type WaitingForReview struct {
	EventID     string `json:"event_id"`
	ComplaintID string `json:"complaint_id"`
	ReceiverID  string `json:"receiver_id"`
	TriggeredBy string `json:"triggered_by"`
	AuthorID    string `json:"author_id"`
	OccurredOn  string `json:"occurred_on"`
	Seen        bool   `json:"seen"`
}

// func NewWaitingForReview(eventID string, seen bool) WaitingForReview {
// 	stringDate := common.NewDate(domainEvent.OccurredOn()).StringRepresentation()
// 	return WaitingForReview{
// 		EventID:     eventID,
// 		ComplaintID: domainEvent.ComplaintID().String(),
// 		ReceiverID:  domainEvent.ReceiverID(),
// 		TriggeredBy: domainEvent.TriggeredBy(),
// 		AuthorID:    domainEvent.AuthorID(),
// 		OccurredOn:  stringDate,
// 		Seen:        seen,
// 	}
// }
