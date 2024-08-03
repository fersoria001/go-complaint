package find_feedback

import "github.com/google/uuid"

type ComplaintId struct {
	query string
	args  []interface{}
}

func ByComplaintId(complaintId uuid.UUID) *ComplaintId {
	return &ComplaintId{
		query: string(`
		SELECT 
			id,
			complaint_id,
			enterprise_id,
			reviewed_at,
			updated_at,
			is_done
		FROM feedback
		WHERE complaint_id = $1`),
		args: []interface{}{complaintId},
	}
}

func (e *ComplaintId) Query() string {
	return e.query
}

func (e *ComplaintId) Args() []interface{} {
	return e.args
}
