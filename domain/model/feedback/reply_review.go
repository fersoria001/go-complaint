package feedback

import (
	"go-complaint/erros"

	"github.com/google/uuid"
)

// Package feedback
// <<Entity>> ReplyReview
// <<Value Object>> Reply
// <<Value Object>> Review
// ReplyReview is a struct that represents the  attached reply with it's single review of a feedback.
// a feedback can have multiple reply reviews.
type ReplyReview struct {
	id         uuid.UUID
	feedbackID uuid.UUID
	reply      *Reply
	review     *Review
}

func NewReplyReview(
	id,
	feedbackID uuid.UUID,
	reply *Reply,
	review *Review,
) (*ReplyReview, error) {
	var rr *ReplyReview = new(ReplyReview)
	err := rr.setID(id)
	if err != nil {
		return nil, err
	}
	err = rr.setFeedbackID(feedbackID)
	if err != nil {
		return nil, err
	}
	err = rr.setReply(reply)
	if err != nil {
		return nil, err
	}
	err = rr.setReview(review)
	if err != nil {
		return nil, err
	}
	return rr, nil
}

func (rr *ReplyReview) setID(id uuid.UUID) error {
	if id == uuid.Nil {
		return &erros.NullValueError{}
	}
	rr.id = id
	return nil
}

func (rr *ReplyReview) setFeedbackID(feedbackID uuid.UUID) error {
	if feedbackID == uuid.Nil {
		return &erros.NullValueError{}
	}
	rr.feedbackID = feedbackID
	return nil
}

func (rr *ReplyReview) setReply(reply *Reply) error {
	if reply == nil {
		return &erros.NullValueError{}
	}
	rr.reply = reply
	return nil
}

func (rr *ReplyReview) setReview(review *Review) error {
	if review == nil {
		return &erros.NullValueError{}
	}
	rr.review = review
	return nil
}

func (rr *ReplyReview) ID() uuid.UUID {
	return rr.id
}

func (rr *ReplyReview) FeedbackID() uuid.UUID {
	return rr.feedbackID
}

func (rr *ReplyReview) Reply() *Reply {
	return rr.reply
}

func (rr *ReplyReview) Review() *Review {
	return rr.review
}
