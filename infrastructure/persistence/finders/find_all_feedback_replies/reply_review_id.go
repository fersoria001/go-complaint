package find_all_feedback_replies

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
			reply_review_id,
			reply_id
		FROM feedback_replies
		WHERE REPLY_REVIEW_ID = $1
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
