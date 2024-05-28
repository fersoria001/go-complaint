package feedback

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/erros"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

// Package feedback
// <<Aggregate root>> Feedback
// <<Entity>> ReplyReview
// Related to the following <<Aggregate root>>:
// Feedback 1 .. 1 Complaint
// Feedback 1 .. 1 Reviewer
// Feedback 1 .. * Answer
type Feedback struct {
	complaintID uuid.UUID
	reviewerID  string
	reviewedID  string
	replyReview mapset.Set[*ReplyReview]
}

func (f *Feedback) Answer(
	ctx context.Context,
	feedbackID uuid.UUID,
	senderID string,
	senderIMG string,
	senderName string,
	body string,
	createdAt common.Date,
	read bool,
	readAt common.Date,
	updatedAt common.Date,
) (*Answer, error) {
	if feedbackID != f.ComplaintID() {
		return nil, &erros.InvalidTypeError{}
	}
	if senderID != f.ReviewerID() && senderID != f.ReviewedID() {
		return nil, &erros.UnauthorizedError{}
	}
	newAnswer, err := NewAnswer(
		uuid.New(),
		feedbackID,
		senderID,
		senderIMG,
		senderName,
		body,
		createdAt,
		read,
		readAt,
		updatedAt,
	)
	if err != nil {
		return nil, err
	}
	err = domain.DomainEventPublisherInstance().Publish(
		ctx,
		NewFeedbackReplied(
			feedbackID,
			f.ComplaintID(),
			f.ReviewerID(),
			f.ReviewedID(),
			newAnswer.ID(),
		),
	)
	if err != nil {
		return nil, err
	}
	return newAnswer, nil
}

func (f *Feedback) AddReplyReview(
	ctx context.Context,
	replyReview *ReplyReview,
) error {
	if replyReview == nil {
		return &erros.NullValueError{}
	}
	if f.ReviewerID() != replyReview.Review().ReviewerID() {
		return &erros.UnauthorizedError{}
	}
	ok := f.replyReview.Add(replyReview)
	if !ok {
		return &erros.AlreadyExistsError{}
	}
	err := domain.DomainEventPublisherInstance().Publish(
		ctx,
		NewAddedFeedback(
			replyReview.ID(),
			replyReview.FeedbackID(),
			replyReview.Reply().SenderID(),
			replyReview.Reply().SenderIMG(),
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func NewFeedback(
	complaintID uuid.UUID,
	reviewerID string,
	reviewedID string,
	replyReviews mapset.Set[*ReplyReview],
) (*Feedback, error) {
	var feedback *Feedback = new(Feedback)
	err := feedback.setComplaintID(complaintID)
	if err != nil {
		return nil, err
	}
	err = feedback.setReviewerID(reviewerID)
	if err != nil {
		return nil, err
	}
	err = feedback.setReviewedID(reviewedID)
	if err != nil {
		return nil, err
	}
	err = feedback.setReplyReviewSet(replyReviews)
	if err != nil {
		return nil, err
	}

	return feedback, nil
}

func (f *Feedback) setReviewedID(reviewedID string) error {
	if reviewedID == "" {
		return &erros.NullValueError{}
	}
	f.reviewedID = reviewedID
	return nil
}

func (f *Feedback) setComplaintID(complaintID uuid.UUID) error {
	if complaintID == uuid.Nil {
		return &erros.NullValueError{}
	}
	f.complaintID = complaintID
	return nil
}

func (f *Feedback) setReviewerID(reviewerID string) error {
	if reviewerID == "" {
		return &erros.NullValueError{}
	}
	f.reviewerID = reviewerID
	return nil
}

func (f *Feedback) setReplyReviewSet(replyReview mapset.Set[*ReplyReview]) error {
	if replyReview == nil {
		return &erros.NullValueError{}
	}
	f.replyReview = replyReview
	return nil
}

func (f *Feedback) ComplaintID() uuid.UUID {
	return f.complaintID
}

func (f *Feedback) ReviewerID() string {
	return f.reviewerID
}

func (f *Feedback) ReviewedID() string {
	return f.reviewedID
}

func (f *Feedback) ReplyReview() mapset.Set[*ReplyReview] {
	return f.replyReview
}
