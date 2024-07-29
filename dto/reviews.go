package dto

type ComplaintReviewState int

type PendingComplaintReview struct {
	EventId     string    `json:"event_id"`
	Complaint   Complaint `json:"complaint"`
	TriggeredBy User      `json:"triggered_by"`
	RatedBy     User      `json:"rated_by"`
	Status      string    `json:"status"`
	OccurredOn  string    `json:"occurred_on"`
}

func (p *PendingComplaintReview) SetStatus(status string) {
	p.Status = status
}

func (p *PendingComplaintReview) SetRatedBy(user User) {
	p.RatedBy = user
}

func (p *PendingComplaintReview) SetOccurredOn(occurredOn string) {
	p.OccurredOn = occurredOn
}
