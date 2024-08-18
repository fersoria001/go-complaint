package commands

import (
	"context"
	"errors"
	"go-complaint/application"
	"go-complaint/domain/model/feedback"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_all_complaint_replies"
	"go-complaint/infrastructure/persistence/repositories"
	"log"

	"github.com/google/uuid"
)

type AddFeedbackReplyCommand struct {
	FeedbackId string   `json:"feedbackId"`
	ReviewerId string   `json:"reviewerId"`
	Color      string   `json:"color"`
	RepliesIds []string `json:"repliesIds"`
}

func NewAddFeedbackReplyCommand(feedbackId, reviewerId, color string, repliesIds []string) *AddFeedbackReplyCommand {
	return &AddFeedbackReplyCommand{
		FeedbackId: feedbackId,
		ReviewerId: reviewerId,
		Color:      color,
		RepliesIds: repliesIds,
	}
}

func (c AddFeedbackReplyCommand) Execute(ctx context.Context) error {
	feedbackId, err := uuid.Parse(c.FeedbackId)
	if err != nil {
		return err
	}
	reviewerId, err := uuid.Parse(c.ReviewerId)
	if err != nil {
		return err
	}
	repliesIds := make([]uuid.UUID, len(c.RepliesIds))
	for i := range c.RepliesIds {
		replyId, err := uuid.Parse(c.RepliesIds[i])
		if err != nil {
			return ErrBadRequest
		}
		repliesIds = append(repliesIds, replyId)
	}
	reg := repositories.MapperRegistryInstance()
	feedbackRepository, ok := reg.Get("Feedback").(repositories.FeedbackRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	repliesRepository, ok := reg.Get("Reply").(repositories.ComplaintRepliesRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	userRepository, ok := reg.Get("User").(repositories.UserRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	reviewer, err := userRepository.Get(ctx, reviewerId)
	if err != nil {
		log.Println("reviewer not found")
		return err
	}
	replies, err := repliesRepository.FindAll(ctx, find_all_complaint_replies.ByAnyIDs(repliesIds))
	if err != nil {
		log.Println("replies not found")
		return err
	}
	f, err := feedbackRepository.Get(ctx, feedbackId)
	if err != nil {
		log.Println("feedback not found")
		return err
	}
	_, err = f.ReplyReview(c.Color)
	if err != nil {
		if errors.Is(err, feedback.ErrReplyReviewNotFound) {
			replyReview := feedback.NewReplyReviewEntity(
				uuid.New(),
				feedbackId,
				*reviewer,
				c.Color,
			)
			err = f.AddReplyReview(ctx, replyReview)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	for _, v := range replies.ToSlice() {
		err = f.AddReply(ctx, c.Color, *v)
		if err != nil {
			return err
		}
	}
	err = feedbackRepository.Update(ctx, f)
	if err != nil {
		return err
	}
	svc := application.ApplicationMessagePublisherInstance()
	svc.Publish(application.NewApplicationMessage(
		f.Id().String(),
		"feedback",
		*dto.NewFeedbackDTO(*f),
	))
	return nil
}
