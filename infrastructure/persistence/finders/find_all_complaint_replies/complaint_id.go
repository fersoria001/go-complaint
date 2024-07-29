package find_all_complaint_replies

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
	author_id,
	body,
	is_read,
	read_at,
	created_at,
	updated_at
	FROM complaint_replies WHERE complaint_id = $1
	`),
		args: []interface{}{complaintID},
	}
}

func (e *ComplaintID) Query() string {
	return e.query
}

func (e *ComplaintID) Args() []interface{} {
	return e.args
}
