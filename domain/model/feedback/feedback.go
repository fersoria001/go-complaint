package feedback

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/erros"
	"slices"
	"time"

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
	complaintId     uuid.UUID
	enterpriseId    uuid.UUID
	replyReview     mapset.Set[*ReplyReview]
	feedbackAnswers mapset.Set[*Answer]
	reviewedAt      time.Time
	updatedAt       time.Time
	isDone          bool
}

func (f *Feedback) Answer(
	ctx context.Context,
	senderID uuid.UUID,
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
			f.complaintId,
			senderID,
			newAnswer.Id(),
		),
	)
	if err != nil {
		return nil, err
	}
	f.feedbackAnswers.Add(newAnswer)
	return newAnswer, nil
}

// ErrNilValue if the color key is empty
// ErrReplyReviewNotFound if the reply review assigned to the color key is not found
// ErrReplyNotFound if the reply is not found
func (f *Feedback) RemoveReply(
	colorKey string,
	reply complaint.Reply,
) error {
	replyReview, err := f.ReplyReview(colorKey)
	if err != nil {
		return err
	}
	err = replyReview.RemoveReply(reply)
	if err != nil {
		return err
	}
	return nil
}

// ErrReplyReviewNotFound if the reply review assigned to the color key is not found
func (f *Feedback) RemoveReplyReview(
	rr *ReplyReview,
) (uuid.UUID, error) {
	if !f.ReplyReviewExists(rr.Color()) {
		return uuid.Nil, ErrReplyReviewNotFound
	}
	var id uuid.UUID
	slice := f.replyReview.ToSlice()
	for i := range slice {
		if slice[i].Color() == rr.Color() {
			id = slice[i].ID()
		}
	}
	slice = slices.DeleteFunc(f.replyReview.ToSlice(), func(i *ReplyReview) bool {
		return i.Color() == rr.Color()
	})
	newSet := mapset.NewSet(slice...)
	f.replyReview = newSet
	return id, nil
}

func (f *Feedback) ReplyReviewExists(
	colorKey string,
) bool {
	exists := false
	slice := f.replyReview.ToSlice()
	for i := range slice {
		if slice[i].Color() == colorKey {
			exists = true
		}
	}
	return exists
}

/*
return ErrColorKeyNotFound if the colorKey is not found in replyReviews
or ErrReplyAlreadyExists if the reply already exist in the replyReview
*/
func (f *Feedback) AddReply(
	ctx context.Context,
	colorKey string,
	reply complaint.Reply,
) error {
	if !f.ReplyReviewExists(colorKey) {
		return ErrColorKeyNotFound
	}
	slice := f.replyReview.ToSlice()
	for i := range slice {
		if slice[i].Color() == colorKey {
			err := slice[i].AddReply(reply)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

/*
return err if replyReview is nil
return err if failed to publish event
*/
func (f *Feedback) AddReplyReview(
	ctx context.Context,
	replyReview *ReplyReview,
) error {
	if replyReview == nil {
		return ErrNilValue
	}
	if f.replyReview == nil {
		f.replyReview = mapset.NewSet[*ReplyReview]()
	}
	f.replyReview.Add(replyReview)
	reviewedIds := mapset.NewSet[uuid.UUID]()
	for i := range replyReview.replies.Iter() {
		reviewedIds.Add(i.Sender().Id())
	}
	for i := range reviewedIds.Iter() {
		err := domain.DomainEventPublisherInstance().Publish(
			ctx,
			NewAddedFeedback(
				replyReview.id,
				replyReview.feedbackId,
				replyReview.reviewer.Id(),
				i,
			),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateFeedback(
	ctx context.Context,
	id uuid.UUID,
	complaintId uuid.UUID,
	enterpriseId uuid.UUID,
) (*Feedback, error) {
	newFeedback := NewFeedbackEntity(
		id,
		complaintId,
		enterpriseId,
	)
	err := domain.DomainEventPublisherInstance().Publish(
		ctx,
		NewFeedbackCreated(
			newFeedback.id,
			newFeedback.complaintId,
			newFeedback.enterpriseId,
		),
	)
	if err != nil {
		return nil, err
	}
	return newFeedback, nil
}

func NewFeedbackEntity(
	id uuid.UUID,
	complaintId uuid.UUID,
	enterpriseId uuid.UUID,
) *Feedback {
	return &Feedback{
		id:              id,
		complaintId:     complaintId,
		enterpriseId:    enterpriseId,
		replyReview:     mapset.NewSet[*ReplyReview](),
		feedbackAnswers: mapset.NewSet[*Answer](),
		reviewedAt:      time.Now(),
		updatedAt:       time.Now(),
	}
}

func NewFeedback(
	feedbackId,
	complaintId uuid.UUID,
	enterpriseId uuid.UUID,
	replyReviews mapset.Set[*ReplyReview],
	feedbackAnswers mapset.Set[*Answer],
	reviewedAt time.Time,
	updatedAt time.Time,
	isDone bool,
) (*Feedback, error) {
	var feedback *Feedback = new(Feedback)
	feedback.reviewedAt = reviewedAt
	feedback.updatedAt = updatedAt
	feedback.isDone = isDone
	err := feedback.SetEnterpriseId(enterpriseId)
	if err != nil {
		return nil, err
	}
	err = feedback.setComplaintId(complaintId)
	if err != nil {
		return nil, err
	}
	err = feedback.setID(feedbackId)
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

func (f *Feedback) EndFeedback(
	ctx context.Context,
) error {
	if f.ReplyReviews().Cardinality() < 3 {
		return ErrFeedbackIsNotDone
	}
	f.updatedAt = time.Now()
	f.isDone = true
	ids := mapset.NewSet[uuid.UUID]()
	for i := range f.replyReview.Iter() {
		for j := range i.replies.Iter() {
			ids.Add(j.Sender().Id())
		}
	}
	for i := range ids.Iter() {
		err := domain.DomainEventPublisherInstance().Publish(
			ctx,
			NewFeedbackDone(
				f.id,
				f.complaintId,
				f.enterpriseId,
				i,
				time.Now(),
			),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f Feedback) IsDone() bool {
	return f.isDone
}

func (f Feedback) Id() uuid.UUID {
	return f.id
}
func (f *Feedback) SetEnterpriseId(enterpriseId uuid.UUID) error {
	if enterpriseId == uuid.Nil {
		return &erros.NullValueError{}
	}
	f.enterpriseId = enterpriseId
	return nil
}
func (f Feedback) EnterpriseId() uuid.UUID {
	return f.enterpriseId
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

func (f *Feedback) setComplaintId(complaintId uuid.UUID) error {
	if complaintId == uuid.Nil {
		return &erros.NullValueError{}
	}
	f.complaintId = complaintId
	return nil
}

func (f *Feedback) setReplyReviewSet(replyReview mapset.Set[*ReplyReview]) error {
	if replyReview == nil {
		return &erros.NullValueError{}
	}
	f.replyReview = replyReview
	return nil
}

func (f Feedback) ComplaintId() uuid.UUID {
	return f.complaintId
}

func (f *Feedback) DeleteComment(colorKey string) (uuid.UUID, error) {
	if colorKey == "" {
		return uuid.Nil, ErrNilValue
	}
	_, err := f.ReplyReview(colorKey)
	if err != nil {
		return uuid.Nil, err
	}
	var id uuid.UUID
	slice := f.replyReview.ToSlice()
	for i := range slice {
		if slice[i].Color() == colorKey {
			id = slice[i].ID()
		}
	}
	slice = slices.DeleteFunc(f.replyReview.ToSlice(), func(i *ReplyReview) bool {
		return i.Color() == colorKey
	})
	newSet := mapset.NewSet[*ReplyReview](slice...)
	f.replyReview = newSet
	return id, nil
}

func (f *Feedback) AddComment(colorKey string, comment string) error {
	if colorKey == "" {
		return ErrNilValue
	}
	rr, err := f.ReplyReview(colorKey)
	if err != nil {
		return err
	}
	err = rr.AddComment(comment)
	if err != nil {
		return err
	}
	return nil
}

// return ErrReplyReviewNotFound if the colorKey is not found in replyReviews
// return ErrNilValue if the colorKey is empty
func (f Feedback) ReplyReview(colorKey string) (*ReplyReview, error) {
	if colorKey == "" {
		return nil, ErrNilValue
	}
	for i := range f.replyReview.Iter() {
		if i.color == colorKey {
			return i, nil
		}
	}
	return nil, ErrReplyReviewNotFound
}

func (f Feedback) ReplyReviews() mapset.Set[ReplyReview] {
	if f.replyReview == nil {
		return nil
	}
	valueCopy := mapset.NewSet[ReplyReview]()
	for replyReview := range f.replyReview.Iter() {
		valueCopy.Add(*replyReview)
	}
	return valueCopy
}

func (f Feedback) FeedbackAnswers() mapset.Set[Answer] {
	if f.feedbackAnswers == nil {
		return nil
	}
	valueCopy := mapset.NewSet[Answer]()
	for answer := range f.feedbackAnswers.Iter() {
		valueCopy.Add(*answer)
	}
	return valueCopy
}

func (f Feedback) ReviewedAt() time.Time {
	return f.reviewedAt
}

func (f Feedback) UpdatedAt() time.Time {
	return f.updatedAt
}
