package find_all_feedback

import "github.com/google/uuid"

type ID struct {
	query string
	args  []interface{}
}

func ByID(feedbackID uuid.UUID) *ID {
	return &ID{
		query: string(`
		SELECT 
			id,
			complaint_id,
			reviewed_id
		FROM feedback
		WHERE id = $1`),
		args: []interface{}{feedbackID},
	}
}

func (e *ID) Query() string {
	return e.query
}

func (e *ID) Args() []interface{} {
	return e.args
}
