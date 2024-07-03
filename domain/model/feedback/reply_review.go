package feedback

import (
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/identity"
	"go-complaint/erros"
	"slices"
	"time"

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
	reviewer   identity.User
	color      string
	createdAt  time.Time
}

func NewReplyReviewEntity(
	id,
	feedbackID uuid.UUID,
	reviewer identity.User,
	color string,
) *ReplyReview {
	return &ReplyReview{
		id:         id,
		feedbackID: feedbackID,
		reviewer:   reviewer,
		replies:    mapset.NewSet[complaint.Reply](),
		review:     NewReview(feedbackID, id, ""),
		color:      color,
		createdAt:  time.Now(),
	}
}

func NewReplyReview(
	id,
	feedbackID uuid.UUID,
	replies mapset.Set[complaint.Reply],
	reviewer identity.User,
	review *Review,
	color string,
	createdAt time.Time,
) (*ReplyReview, error) {
	return &ReplyReview{
		id:         id,
		feedbackID: feedbackID,
		replies:    replies,
		reviewer:   reviewer,
		review:     review,
		color:      color,
		createdAt:  createdAt,
	}, nil
}

// ErrReplyNotFound if reply doesn't exist in the set.
func (rr *ReplyReview) RemoveReply(reply complaint.Reply) error {
	if !rr.ReplyExists(reply) {
		return ErrReplyNotFound
	}
	slice := rr.replies.ToSlice()
	slice = slices.DeleteFunc(slice, func(i complaint.Reply) bool {
		return i.ID() == reply.ID()
	})
	newSet := mapset.NewSet(slice...)
	rr.replies = newSet
	return nil
}

func (rr *ReplyReview) AddComment(comment string) error {
	return rr.review.setComment(comment)
}

func (rr *ReplyReview) Reviewer() identity.User {
	return rr.reviewer
}

func (rr *ReplyReview) Color() string {
	return rr.color
}

func (rr *ReplyReview) CreatedAt() time.Time {
	return rr.createdAt
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

func (rr ReplyReview) ReplyExists(reply complaint.Reply) bool {
	slice := rr.replies.ToSlice()
	for i := range slice {
		if slice[i].ID() == reply.ID() {
			return true
		}
	}
	return false
}

/*
return ErrReplyAlreadyExists if the reply already exists in the set.
*/
func (rr *ReplyReview) AddReply(reply complaint.Reply) error {
	if rr.ReplyExists(reply) {
		return ErrReplyAlreadyExists
	}
	rr.replies.Add(reply)
	return nil
}

func (rr *ReplyReview) Review() Review {
	if rr.review != nil {
		return *rr.review
	}
	return Review{}
}
