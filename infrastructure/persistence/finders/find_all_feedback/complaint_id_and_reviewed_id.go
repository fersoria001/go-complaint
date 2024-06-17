package find_all_feedback

import "github.com/google/uuid"

type ComplaintIDAndReviewedID struct {
	query string
	args  []interface{}
}

func ByComplaintIDAndReviewedID(complaintID uuid.UUID, reviewedID string) *ComplaintIDAndReviewedID {
	return &ComplaintIDAndReviewedID{
		query: string(`
		SELECT 
			id,
			complaint_id,
			reviewed_id
		FROM feedback
		WHERE complaint_id = $1 AND REVIEWED_ID = $2`),
		args: []interface{}{complaintID, reviewedID},
	}
}

func (e *ComplaintIDAndReviewedID) Query() string {
	return e.query
}

func (e *ComplaintIDAndReviewedID) Args() []interface{} {
	return e.args
}
