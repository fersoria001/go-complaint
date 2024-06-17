package feedback

import (
	"go-complaint/domain/model/common"
	"go-complaint/erros"

	"github.com/google/uuid"
)

type Review struct {
	replyReviewID uuid.UUID
	reviewerID    string
	reviewedAt    common.Date
	comment       string
}

func NewReview(
	replyReviewID uuid.UUID,
	reviewerID string,
	reviewedAt common.Date,
	comment string,
) (*Review, error) {
	var review *Review = new(Review)
	err := review.setReviewerID(reviewerID)
	if err != nil {
		return nil, err
	}
	err = review.setReviewedAt(reviewedAt)
	if err != nil {
		return nil, err
	}
	err = review.setComment(comment)
	if err != nil {
		return nil, err
	}
	err = review.setReplyReviewID(replyReviewID)
	if err != nil {
		return nil, err
	}

	return review, nil
}

func (r Review) ReviewerID() string {
	return r.reviewerID
}

func (r *Review) setReviewerID(reviewer string) error {
	if reviewer == "" {
		return &erros.NullValueError{}
	}
	r.reviewerID = reviewer
	return nil
}

func (r *Review) setReplyReviewID(replyReviewID uuid.UUID) error {
	if replyReviewID == uuid.Nil {
		return &erros.NullValueError{}
	}
	r.replyReviewID = replyReviewID
	return nil
}

func (r *Review) setReviewedAt(reviewedAt common.Date) error {

	if reviewedAt == (common.Date{}) {
		return &erros.EmptyStructError{}
	}
	r.reviewedAt = reviewedAt
	return nil
}

func (r *Review) setComment(comment string) error {
	if comment == "" {
		return &erros.NullValueError{}
	}
	r.comment = comment
	return nil
}

func (r Review) ReviewedAt() common.Date {
	return r.reviewedAt
}

func (r Review) Comment() string {
	return r.comment
}

func (r Review) ReplyReviewID() uuid.UUID {
	return r.replyReviewID
}
