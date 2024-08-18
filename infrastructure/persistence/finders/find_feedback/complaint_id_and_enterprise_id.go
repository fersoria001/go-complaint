package find_feedback

import "github.com/google/uuid"

type ComplaintIdAndEnterpriseId struct {
	query string
	args  []interface{}
}

func ByComplaintIdAndEnterpriseId(complaintId, enterpriseId uuid.UUID) *ComplaintIdAndEnterpriseId {
	return &ComplaintIdAndEnterpriseId{
		query: string(`
		SELECT 
			id,
			complaint_id,
			enterprise_id,
			reviewed_at,
			updated_at,
			is_done
		FROM feedback
		WHERE complaint_id = $1 AND enterprise_id = $2`),
		args: []interface{}{complaintId, enterpriseId},
	}
}

func (e *ComplaintIdAndEnterpriseId) Query() string {
	return e.query
}

func (e *ComplaintIdAndEnterpriseId) Args() []interface{} {
	return e.args
}
