package find_all_feedback

import "github.com/google/uuid"

type ComplaintID struct {
	query string
	args  []interface{}
}

func ByComplaintID(complaintID uuid.UUID) *ComplaintID {
	return &ComplaintID{
		query: string(`
		SELECT 
			id,
			complaint_id,
			reviewed_id
		FROM feedback
		WHERE complaint_id = $1`),
		args: []interface{}{complaintID},
	}
}

func (e *ComplaintID) Query() string {
	return e.query
}

func (e *ComplaintID) Args() []interface{} {
	return e.args
}
