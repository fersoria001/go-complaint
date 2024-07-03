package dto

type ComplaintInfo struct {
	ComplaintsReceived int     `json:"complaints_received"`
	ComplaintsResolved int     `json:"complaints_resolved"`
	ComplaintsReviewed int     `json:"complaints_reviewed"`
	ComplaintsPending  int     `json:"complaints_pending"`
	AverageRating      float64 `json:"average_rating"`
}
