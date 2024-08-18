package commands

import (
	"context"
	"go-complaint/application"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type RemoveFeedbackCommentCommand struct {
	Color      string `json:"color"`
	FeedbackId string `json:"feedbackId"`
}

func NewRemoveFeedbackCommentCommand(color, feedbackId string) *RemoveFeedbackCommentCommand {
	return &RemoveFeedbackCommentCommand{
		Color:      color,
		FeedbackId: feedbackId,
	}
}

func (c RemoveFeedbackCommentCommand) Execute(ctx context.Context) error {
	feedbackId, err := uuid.Parse(c.FeedbackId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Feedback").(repositories.FeedbackRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	f, err := r.Get(ctx, feedbackId)
	if err != nil {
		return err
	}
	err = f.DeleteComment(c.Color)
	if err != nil {
		return err
	}
	err = r.Update(ctx, f)
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
