package application

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/feedback"
	"go-complaint/dto"
	"go-complaint/erros"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type FeedbackService struct {
	feedbackRepository *repositories.FeedbackRepository
	answerRepository   *repositories.AnswerRepository
}

func NewFeedbackService(
	feedbackRepository *repositories.FeedbackRepository,
	answerRepository *repositories.AnswerRepository,
) *FeedbackService {
	return &FeedbackService{
		feedbackRepository: feedbackRepository,
		answerRepository:   answerRepository,
	}
}

func (feedbackService *FeedbackService) MarkAnswerAsRead(
	ctx context.Context,
	answerID string,
) error {
	dbObj, err := feedbackService.answerRepository.Get(ctx, answerID)
	if err != nil {
		return err
	}
	dbAnswer, ok := dbObj.(*feedback.Answer)
	if !ok {
		return &erros.InvalidTypeError{}
	}
	err = dbAnswer.MarkAsRead()
	if err != nil {
		return err
	}
	err = feedbackService.answerRepository.Update(ctx, dbAnswer)
	if err != nil {
		return err
	}
	return nil
}

func (feedbackService *FeedbackService) AnswerAFeedback(
	ctx context.Context,
	complaintID string,
	senderID string,
	senderIMG string,
	senderName string,
	body string,
	createdAt string,
	read bool,
	readAt string,
	updatedAt string,
) error {
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				fmt.Println(`handle event at AnswerAFeedback`)
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&feedback.FeedbackReplied{})
			},
		},
	)

	dbObj, err := feedbackService.feedbackRepository.Get(ctx, complaintID)
	if err != nil {
		return err
	}
	dbFeedback, ok := dbObj.(*feedback.Feedback)
	if !ok {
		return &erros.InvalidTypeError{}
	}
	parsedComplaintID, err := uuid.Parse(complaintID)
	if err != nil {
		return err
	}
	commonCreatedAt, err := common.NewDateFromString(createdAt)
	if err != nil {
		return err
	}
	commonReadAt, err := common.NewDateFromString(readAt)
	if err != nil {
		return err
	}
	commonUpdatedAt, err := common.NewDateFromString(updatedAt)
	if err != nil {
		return err
	}
	newAnswer, err := dbFeedback.Answer(
		ctx,
		parsedComplaintID,
		senderID,
		senderIMG,
		senderName,
		body,
		commonCreatedAt,
		read,
		commonReadAt,
		commonUpdatedAt,
	)
	if err != nil {
		return err
	}
	err = feedbackService.answerRepository.Save(ctx, newAnswer)
	if err != nil {
		return err
	}
	return nil
}

func (feedbackService *FeedbackService) Feedback(
	ctx context.Context,
	complaintID string,
) (*dto.Feedback, error) {
	dbObj, err := feedbackService.feedbackRepository.Get(ctx, complaintID)
	if err != nil {
		return nil, err
	}
	dbFeedback, ok := dbObj.(*feedback.Feedback)
	if !ok {
		return nil, &erros.NoElementError{}
	}
	dbAnswers, err := feedbackService.answerRepository.FindByFeedbackID(ctx, complaintID)
	if err != nil {
		return nil, err
	}
	return dto.NewFeedback(dbFeedback, dbAnswers), nil
}

func (feedbackService *FeedbackService) CreateAFeedback(
	ctx context.Context,
	complaintID string,
	reviewerID string,
	reviewedID string,
	reviewerIMG string,
	reviewerName string,
	senderID string,
	senderIMG string,
	senderName string,
	body string,
	createdAt string,
	read bool,
	readAt string,
	updatedAt string,
	comment string,
) error {
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				fmt.Println(`handle event at CreateAFeedback`)
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&feedback.AddedFeedback{})
			},
		},
	)
	//maybe senderID == reviewedID in all cases
	feedbackID, err := uuid.Parse(complaintID)
	if err != nil {
		return err
	}
	commonCreatedAt, err := common.NewDateFromString(createdAt)
	if err != nil {
		return err
	}
	commonReadAt, err := common.NewDateFromString(readAt)
	if err != nil {
		return err
	}
	commonUpdatedAt, err := common.NewDateFromString(updatedAt)
	if err != nil {
		return err
	}
	reply, err := feedback.NewReply(
		senderID,
		senderIMG,
		senderName,
		body,
		commonCreatedAt,
		read,
		commonReadAt,
		commonUpdatedAt,
	)
	if err != nil {
		return err
	}
	//create the review for that reply
	reviewedAt, err := common.NewDateFromString(createdAt)
	if err != nil {
		return err
	}
	review, err := feedback.NewReview(
		reviewerID,
		reviewerIMG,
		reviewerName,
		reviewedAt,
		comment,
	)
	if err != nil {
		return err
	}
	replyReview, err := feedback.NewReplyReview(
		uuid.New(),
		feedbackID,
		reply,
		review,
	)
	if err != nil {
		return err
	}
	//add the reply review to the feedback
	emptyReplyReviews := mapset.NewSet[*feedback.ReplyReview]()

	newFeedback, err := feedback.NewFeedback(
		feedbackID,
		reviewerID,
		reviewedID,
		emptyReplyReviews,
	)
	if err != nil {
		return err
	}
	err = newFeedback.AddReplyReview(ctx, replyReview)
	if err != nil {
		return err
	}
	err = feedbackService.feedbackRepository.Save(ctx, newFeedback)
	if err != nil {
		return err
	}

	return nil
}
