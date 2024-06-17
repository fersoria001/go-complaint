package find_all_feedback_reply_review

import "github.com/google/uuid"

type FeedbackID struct {
	query string
	args  []interface{}
}

func ByFeedbackID(feedbackID uuid.UUID) *FeedbackID {
	return &FeedbackID{
		query: string(`
		SELECT id, feedback_id, review_id, color
	FROM public.feedback_reply_review
	WHERE feedback_id = $1
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
