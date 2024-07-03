package find_all_feedback

import "github.com/google/uuid"

type AnyIDs struct {
	query string
	args  []interface{}
}

func ByAnyIDs(feedbackIDs []uuid.UUID) *AnyIDs {
	return &AnyIDs{
		query: string(`
		SELECT 
			id,
			complaint_id,
			enterprise_id,
			reviewed_at,
			updated_at,
			is_done
		FROM feedback
		WHERE id = any ($1)`),
		args: []interface{}{feedbackIDs},
	}
}

func (e *AnyIDs) Query() string {
	return e.query
}

func (e *AnyIDs) Args() []interface{} {
	return e.args
}
