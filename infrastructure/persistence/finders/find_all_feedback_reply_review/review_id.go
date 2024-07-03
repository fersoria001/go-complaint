package find_all_feedback_reply_review

import "github.com/google/uuid"

type ReviewID struct {
	query string
	args  []interface{}
}

func ByReviewID(reviewID uuid.UUID) *ReviewID {
	return &ReviewID{
		query: string(`
		SELECT id, feedback_id, review_id, color,
		created_at
	FROM feedback_reply_review
	WHERE review_id = $1
		`),
		args: []interface{}{reviewID},
	}
}

func (e *ReviewID) Query() string {
	return e.query
}

func (e *ReviewID) Args() []interface{} {
	return e.args
}
