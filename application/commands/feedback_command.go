package commands

import (
	"context"
	"errors"
	"fmt"
	"go-complaint/application/application_services"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/feedback"
	"go-complaint/dto"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/finders/find_all_complaint_replies"
	"go-complaint/infrastructure/persistence/removers/remove_all_feedback_replies"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type FeedbackCommand struct {
	FeedbackID   string
	EnterpriseID string
	ComplaintID  string
	ReviewedID   string
	ReviewerID   string
	Comment      string
	Color        string
	Replies      []string
	AnswerBody   string
}

func (command FeedbackCommand) AnswerFeedback(ctx context.Context) error {
	if command.FeedbackID == "" || command.ReviewerID == "" ||
		command.AnswerBody == "" {
		return ErrBadRequest
	}
	parsedID, err := uuid.Parse(command.FeedbackID)
	if err != nil {
		return ErrBadRequest
	}
	dbFeedback, err := repositories.MapperRegistryInstance().Get(
		"Feedback",
	).(repositories.FeedbackRepository).Get(
		ctx,
		parsedID,
	)
	if err != nil {
		return err
	}
	sender, err := repositories.MapperRegistryInstance().Get(
		"User",
	).(repositories.UserRepository).Get(
		ctx,
		command.ReviewerID,
	)
	if err != nil {
		return err
	}
	thisDate := common.NewDate(time.Now())
	lastReply, err := dbFeedback.Answer(
		ctx,
		sender.Email(),
		sender.ProfileIMG(),
		sender.FullName(),
		command.AnswerBody,
		thisDate,
		false,
		thisDate,
		thisDate,
		false,
		"",
	)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get(
		"Feedback",
	).(repositories.FeedbackRepository).Update(
		ctx,
		dbFeedback,
	)
	if err != nil {
		return err
	}
	cache.InMemoryCacheInstance().SetPublish(
		fmt.Sprintf("feedbackLastReply:%s", dbFeedback.ID().String()),
		dto.NewFeedbackAnswerDTO(*lastReply),
	)
	return nil
}

func (command FeedbackCommand) EndFeedback(
	ctx context.Context,
) error {
	if command.FeedbackID == "" ||
		command.ReviewerID == "" {
		return ErrBadRequest
	}
	parsedID, err := uuid.Parse(command.FeedbackID)
	if err != nil {
		return ErrBadRequest
	}
	dbFeedback, err := repositories.MapperRegistryInstance().Get(
		"Feedback",
	).(repositories.FeedbackRepository).Get(
		ctx,
		parsedID,
	)
	if err != nil {
		return err
	}
	dbComplaint, err := repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).Get(
		ctx,
		dbFeedback.ComplaintID(),
	)
	if err != nil {
		return err
	}
	reviewer, err := repositories.MapperRegistryInstance().Get(
		"User",
	).(repositories.UserRepository).Get(
		ctx,
		command.ReviewerID,
	)
	if err != nil {
		return err
	}
	param := fmt.Sprintf("id=%s", dbFeedback.ID().String())
	replaced := application_services.EncodingApplicationServiceInstance().SafeUtf16Encode(param)
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if ev, ok := event.(*feedback.FeedbackDone); ok {
					NotificationCommand{
						OwnerID:     ev.ReviewedID(),
						ThumbnailID: dbComplaint.ReceiverID(),
						Thumbnail:   dbComplaint.ReceiverProfileIMG(),
						Title:       fmt.Sprintf("%s has made a feedback on your attention", reviewer.FullName()),
						Content:     fmt.Sprintf("You have received a feedback from %s on your attention", reviewer.FullName()),
						Link:        fmt.Sprintf("/%s/feedbacks?%s", dbComplaint.ReceiverID(), replaced),
					}.SaveNew(ctx)

				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&feedback.FeedbackDone{})
			},
		})
	err = dbFeedback.EndFeedback(ctx)
	if err != nil {
		return err
	}
	err = dbComplaint.SendToHistory(ctx, command.ReviewerID)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get(
		"Feedback",
	).(repositories.FeedbackRepository).Update(
		ctx,
		dbFeedback,
	)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get(
		"Complaint",
	).(repositories.ComplaintRepository).Update(
		ctx,
		dbComplaint,
	)
	if err != nil {
		return err
	}
	cache.RequestChannel <- cache.Request{
		Key:     fmt.Sprintf("feedback:%s", dbFeedback.ID().String()),
		Payload: dto.NewFeedbackDTO(*dbFeedback),
	}
	return nil
}

func (command FeedbackCommand) CreateFeedback(
	ctx context.Context,
) error {
	if command.Color == "" ||
		command.ComplaintID == "" ||
		command.EnterpriseID == "" {
		return ErrBadRequest
	}
	complaintID, err := uuid.Parse(command.ComplaintID)
	if err != nil {
		return ErrBadRequest
	}

	f, err := feedback.CreateFeedback(
		ctx,
		uuid.New(),
		complaintID,
		command.EnterpriseID,
	)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get(
		"Feedback",
	).(repositories.FeedbackRepository).Save(
		ctx,
		f,
	)
	if err != nil {
		return err
	}
	cache.RequestChannel <- cache.Request{
		Key:     fmt.Sprintf("feedback:%s", f.ID().String()),
		Payload: dto.NewFeedbackDTO(*f),
	}
	return nil
}

func (command FeedbackCommand) AddReply(
	ctx context.Context,
) error {
	if command.FeedbackID == "" ||
		command.ReviewerID == "" ||
		command.Color == "" ||
		command.Replies == nil {
		return ErrBadRequest
	}
	parsedID, err := uuid.Parse(command.FeedbackID)
	if err != nil {
		return ErrBadRequest
	}
	f, err := repositories.MapperRegistryInstance().Get(
		"Feedback",
	).(repositories.FeedbackRepository).Get(
		ctx,
		parsedID,
	)
	if err != nil {
		return err
	}
	reviewer, err := repositories.MapperRegistryInstance().Get(
		"User",
	).(repositories.UserRepository).Get(
		ctx,
		command.ReviewerID,
	)
	if err != nil {
		return err
	}
	slice := make([]uuid.UUID, len(command.Replies))
	for _, replyID := range command.Replies {
		parsedReplyID, err := uuid.Parse(replyID)
		if err != nil {
			return ErrBadRequest
		}
		slice = append(slice, parsedReplyID)
	}
	replies, err := repositories.MapperRegistryInstance().Get(
		"Reply").(repositories.ComplaintRepliesRepository).FindAll(
		ctx,
		find_all_complaint_replies.ByAnyIDs(slice),
	)
	if err != nil {
		return err
	}
	_, err = f.ReplyReview(command.Color)
	if err != nil {
		if errors.Is(err, feedback.ErrReplyReviewNotFound) {
			replyReview := feedback.NewReplyReviewEntity(
				uuid.New(),
				parsedID,
				*reviewer,
				command.Color,
			)
			err = f.AddReplyReview(ctx, replyReview)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	for reply := range replies.Iter() {
		err = f.AddReply(ctx, command.Color, *reply)
		if err != nil {
			return err
		}
	}
	err = repositories.MapperRegistryInstance().Get(
		"Feedback",
	).(repositories.FeedbackRepository).Update(
		ctx,
		f,
	)
	if err != nil {
		return err
	}
	cache.RequestChannel <- cache.Request{
		Key:     fmt.Sprintf("feedback:%s", f.ID().String()),
		Payload: dto.NewFeedbackDTO(*f),
	}
	return nil
}

func (command FeedbackCommand) RemoveReply(
	ctx context.Context,
) error {
	if command.FeedbackID == "" ||
		command.Color == "" ||
		command.Replies == nil {
		return ErrBadRequest
	}
	parsedID, err := uuid.Parse(command.FeedbackID)
	if err != nil {
		return ErrBadRequest
	}
	f, err := repositories.MapperRegistryInstance().Get(
		"Feedback",
	).(repositories.FeedbackRepository).Get(
		ctx,
		parsedID,
	)
	if err != nil {
		return err
	}
	slice := make([]uuid.UUID, len(command.Replies))
	for _, replyID := range command.Replies {
		parsedReplyID, err := uuid.Parse(replyID)
		if err != nil {
			return ErrBadRequest
		}
		slice = append(slice, parsedReplyID)
	}
	replies, err := repositories.MapperRegistryInstance().Get(
		"Reply").(repositories.ComplaintRepliesRepository).FindAll(
		ctx,
		find_all_complaint_replies.ByAnyIDs(slice),
	)
	if err != nil {
		return err
	}
	replySlice := replies.ToSlice()
	for _, reply := range replySlice {
		err = f.RemoveReply(command.Color, *reply)
		if err != nil {
			return err
		}
	}
	replyReview, err := f.ReplyReview(command.Color)
	if err != nil {
		return err
	}
	if replyReview.Replies().Cardinality() == 0 {
		id, err := f.RemoveReplyReview(replyReview)
		if err != nil {
			return err
		}
		err = repositories.MapperRegistryInstance().Get(
			"FeedbackReply").(repositories.FeedbackRepliesRepository).DeleteAll(
			ctx,
			remove_all_feedback_replies.WhereReplyReviewID(id),
		)
		if err != nil {
			return err
		}
	}
	err = repositories.MapperRegistryInstance().Get(
		"Feedback",
	).(repositories.FeedbackRepository).Update(
		ctx,
		f,
	)
	if err != nil {
		return err
	}
	cache.RequestChannel <- cache.Request{
		Key:     fmt.Sprintf("feedback:%s", f.ID().String()),
		Payload: dto.NewFeedbackDTO(*f),
	}
	return nil
}

func (command FeedbackCommand) AddComment(
	ctx context.Context,
) error {
	if command.FeedbackID == "" ||
		command.Color == "" {
		return ErrBadRequest
	}
	parsedID, err := uuid.Parse(command.FeedbackID)
	if err != nil {
		return ErrBadRequest
	}
	f, err := repositories.MapperRegistryInstance().Get(
		"Feedback",
	).(repositories.FeedbackRepository).Get(
		ctx,
		parsedID,
	)
	if err != nil {
		return err
	}
	err = f.AddComment(
		command.Color,
		command.Comment,
	)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get(
		"Feedback",
	).(repositories.FeedbackRepository).Update(
		ctx,
		f,
	)
	if err != nil {
		return err
	}
	cache.RequestChannel <- cache.Request{
		Key:     fmt.Sprintf("feedback:%s", f.ID().String()),
		Payload: dto.NewFeedbackDTO(*f),
	}
	return nil
}

func (command FeedbackCommand) DeleteComment(
	ctx context.Context,
) error {
	if command.Color == "" ||
		command.FeedbackID == "" {
		return ErrBadRequest
	}
	parsedID, err := uuid.Parse(command.FeedbackID)
	if err != nil {
		return ErrBadRequest
	}
	f, err := repositories.MapperRegistryInstance().Get(
		"Feedback",
	).(repositories.FeedbackRepository).Get(
		ctx,
		parsedID,
	)
	if err != nil {
		return err
	}
	id, err := f.DeleteComment(command.Color)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get(
		"Feedback",
	).(repositories.FeedbackRepository).Update(
		ctx,
		f,
	)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get(
		"FeedbackReply").(repositories.FeedbackRepliesRepository).DeleteAll(
		ctx,
		remove_all_feedback_replies.WhereReplyReviewID(id),
	)
	if err != nil {
		return err
	}
	cache.RequestChannel <- cache.Request{
		Key:     fmt.Sprintf("feedback:%s", f.ID().String()),
		Payload: dto.NewFeedbackDTO(*f),
	}
	return nil
}
