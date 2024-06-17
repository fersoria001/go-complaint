package find_all_feedback_review

import "github.com/google/uuid"

type ReplyReviewID struct {
	query string
	args  []interface{}
}

func ByReplyReviewID(replyReviewID uuid.UUID) *ReplyReviewID {
	return &ReplyReviewID{
		query: string(`
		SELECT 
		id,
		reviewer_id,
		reviewed_at,
		review_comment
		FROM feedback_reviews
		WHERE ID = $1
		`),
		args: []interface{}{replyReviewID},
	}
}

func (e *ReplyReviewID) Query() string {
	return e.query
}

func (e *ReplyReviewID) Args() []interface{} {
	return e.args
}
