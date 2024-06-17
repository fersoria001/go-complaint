package feedback

import (
	"go-complaint/domain/model/complaint"
	"go-complaint/erros"

	mapset "github.com/deckarep/golang-set/v2"
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
	replies    mapset.Set[complaint.Reply]
	review     *Review
	color      string
}

func NewReplyReview(
	id,
	feedbackID uuid.UUID,
	replies mapset.Set[complaint.Reply],
	review *Review,
	color string,
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
	err = rr.setReplies(replies)
	if err != nil {
		return nil, err
	}
	err = rr.setReview(review)
	if err != nil {
		return nil, err
	}
	err = rr.setColor(color)
	if err != nil {
		return nil, err
	}
	return rr, nil
}

func (rr *ReplyReview) Color() string {
	return rr.color
}

func (rr *ReplyReview) setColor(color string) error {
	if color == "" {
		return &erros.NullValueError{}
	}
	rr.color = color
	return nil
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

func (rr *ReplyReview) setReplies(replies mapset.Set[complaint.Reply]) error {
	if replies == nil {
		return &erros.NullValueError{}
	}
	rr.replies = replies
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

func (rr *ReplyReview) Replies() mapset.Set[complaint.Reply] {
	return rr.replies
}

func (rr *ReplyReview) AddReply(reply complaint.Reply) {
	rr.replies.Add(reply)
}

func (rr *ReplyReview) RemoveReply(reply complaint.Reply) {
	rr.replies.Remove(reply)
}

func (rr *ReplyReview) Review() Review {
	return *rr.review
}
