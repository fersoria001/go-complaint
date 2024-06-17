package dto

type PendingComplaintReview struct {
	EventID     string       `json:"event_id"`
	Complaint   ComplaintDTO `json:"complaint"`
	TriggeredBy User         `json:"triggered_by"`
}
