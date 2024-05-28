package feedback

import (
	"go-complaint/domain/model/common"
	"go-complaint/erros"
)

type Review struct {
	reviewerID   string
	reviewerIMG  string
	reviewerName string
	reviewedAt   common.Date
	comment      string
}

func NewReview(
	reviewerID string,
	reviewerIMG string,
	reviewerName string,
	reviewedAt common.Date,
	comment string,
) (*Review, error) {
	var review *Review = new(Review)
	err := review.setReviewerID(reviewerID)
	if err != nil {
		return nil, err
	}
	err = review.setReviewerIMG(reviewerIMG)
	if err != nil {
		return nil, err
	}
	err = review.setReviewerName(reviewerName)
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
	return review, nil
}

func (r *Review) setReviewerID(reviewerID string) error {
	if reviewerID == "" {
		return &erros.NullValueError{}
	}
	r.reviewerID = reviewerID
	return nil
}

func (r *Review) setReviewerIMG(reviewerIMG string) error {
	if reviewerIMG == "" {
		return &erros.NullValueError{}
	}
	r.reviewerIMG = reviewerIMG
	return nil
}

func (r *Review) setReviewerName(reviewerName string) error {
	if reviewerName == "" {
		return &erros.NullValueError{}
	}
	r.reviewerName = reviewerName
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

func (r *Review) ReviewerID() string {
	return r.reviewerID
}

func (r *Review) ReviewerIMG() string {
	return r.reviewerIMG
}

func (r *Review) ReviewerName() string {
	return r.reviewerName
}

func (r *Review) ReviewedAt() common.Date {
	return r.reviewedAt
}

func (r *Review) Comment() string {
	return r.comment
}
