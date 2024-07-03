package feedback

import (
	"go-complaint/erros"

	"github.com/google/uuid"
)

type Review struct {
	feedbackID    uuid.UUID
	replyReviewID uuid.UUID
	comment       string
}

func NewReview(
	feedbackID uuid.UUID,
	replyReviewID uuid.UUID,
	comment string,
) *Review {
	return &Review{
		feedbackID:    feedbackID,
		replyReviewID: replyReviewID,
		comment:       comment,
	}
}

func (r Review) FeedbackID() uuid.UUID {
	return r.feedbackID
}

func (r *Review) setReplyReviewID(replyReviewID uuid.UUID) error {
	if replyReviewID == uuid.Nil {
		return &erros.NullValueError{}
	}
	r.replyReviewID = replyReviewID
	return nil
}

func (r *Review) setComment(comment string) error {
	if len(comment) < 5 {
		return ValidationError{Message: "comment is too short"}
	}
	if len(comment) > 500 {
		return ValidationError{Message: "comment is too long"}
	}

	r.comment = comment
	return nil
}

func (r Review) Comment() string {
	return r.comment
}

func (r Review) ReplyReviewID() uuid.UUID {
	return r.replyReviewID
}
