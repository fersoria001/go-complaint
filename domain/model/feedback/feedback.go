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
	id              uuid.UUID
	complaintID     uuid.UUID
	reviewedID      string
	replyReview     mapset.Set[*ReplyReview]
	feedbackAnswers mapset.Set[*Answer]
}

func (f *Feedback) Answer(
	ctx context.Context,
	senderID string,
	senderIMG string,
	senderName string,
	body string,
	createdAt common.Date,
	read bool,
	readAt common.Date,
	updatedAt common.Date,
	isEnterprise bool,
	enterpriseID string,
) (*Answer, error) {
	newAnswer, err := NewAnswer(
		uuid.New(),
		f.id,
		senderID,
		senderIMG,
		senderName,
		body,
		createdAt,
		read,
		readAt,
		updatedAt,
		isEnterprise,
		enterpriseID,
	)
	if err != nil {
		return nil, err
	}
	err = domain.DomainEventPublisherInstance().Publish(
		ctx,
		NewFeedbackReplied(
			f.id,
			f.complaintID,
			f.ReviewedID(),
			newAnswer.ID(),
		),
	)
	if err != nil {
		return nil, err
	}
	f.feedbackAnswers.Add(newAnswer)
	return newAnswer, nil
}

func (f *Feedback) AddReplyReview(
	ctx context.Context,
	replyReview *ReplyReview,
) error {
	if replyReview == nil {
		return &erros.NullValueError{}
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
			replyReview.review.reviewerID,
			f.reviewedID,
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func NewFeedback(
	feedbackID,
	complaintID uuid.UUID,
	reviewedID string,
	replyReviews mapset.Set[*ReplyReview],
	feedbackAnswers mapset.Set[*Answer],
) (*Feedback, error) {
	var feedback *Feedback = new(Feedback)

	err := feedback.setComplaintID(complaintID)
	if err != nil {
		return nil, err
	}
	err = feedback.setID(feedbackID)
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
	err = feedback.setFeedbackAnswers(feedbackAnswers)
	if err != nil {
		return nil, err
	}
	return feedback, nil
}

func (f Feedback) ID() uuid.UUID {
	return f.id
}

func (f *Feedback) setID(id uuid.UUID) error {
	if id == uuid.Nil {
		return &erros.NullValueError{}
	}
	f.id = id
	return nil
}

func (f *Feedback) setFeedbackAnswers(fba mapset.Set[*Answer]) error {
	if fba == nil {
		return &erros.NullValueError{}
	}
	f.feedbackAnswers = fba
	return nil
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

func (f *Feedback) setReplyReviewSet(replyReview mapset.Set[*ReplyReview]) error {
	if replyReview == nil {
		return &erros.NullValueError{}
	}
	f.replyReview = replyReview
	return nil
}

func (f Feedback) ComplaintID() uuid.UUID {
	return f.complaintID
}

func (f Feedback) ReviewedID() string {
	return f.reviewedID
}

func (f Feedback) ReplyReview() mapset.Set[ReplyReview] {
	valueCopy := mapset.NewSet[ReplyReview]()
	for replyReview := range f.replyReview.Iter() {
		valueCopy.Add(*replyReview)
	}
	return valueCopy
}

func (f Feedback) FeedbackAnswers() mapset.Set[Answer] {
	valueCopy := mapset.NewSet[Answer]()
	for answer := range f.feedbackAnswers.Iter() {
		valueCopy.Add(*answer)
	}
	return valueCopy
}
