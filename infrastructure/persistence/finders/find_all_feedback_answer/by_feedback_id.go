package find_all_feedback_answer

import "github.com/google/uuid"

type FeedbackID struct {
	query string
	args  []interface{}
}

func ByFeedbackID(feedbackID uuid.UUID) *FeedbackID {
	return &FeedbackID{
		query: string(`
		SELECT 
			id,
			feedback_id,
			sender_id,
			answer_body,
			created_at,
			read_status,
			read_at, 
			updated_at,
			is_enterprise,
			enterprise_id
		FROM feedback_answers
		WHERE FEEDBACK_ID = $1
		`),
		args: []interface{}{feedbackID},
	}
}

func (e *FeedbackID) Query() string {
	return e.query
}

func (e *FeedbackID) Args() []interface{} {
	return e.args
}
