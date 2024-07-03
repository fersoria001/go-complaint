package remove_all_feedback_replies

import "github.com/google/uuid"

type ReplyReviewID struct {
	query string
	args  []interface{}
}

func WhereReplyReviewID(replyReviewID uuid.UUID) *ReplyReviewID {
	return &ReplyReviewID{
		query: string(`
		DELETE FROM feedback_replies
		where reply_review_id = $1
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
