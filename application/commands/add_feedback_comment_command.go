package commands

import (
	"context"

	"go-complaint/application"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type AddFeedbackCommentCommand struct {
	FeedbackId string `json:"feedbackId"`
	Color      string `json:"color"`
	Comment    string `json:"comment"`
}

func NewAddFeedbackCommentCommand(feedbackId, color, comment string) *AddFeedbackCommentCommand {
	return &AddFeedbackCommentCommand{
		FeedbackId: feedbackId,
		Color:      color,
		Comment:    comment,
	}
}

func (c AddFeedbackCommentCommand) Execute(ctx context.Context) error {
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
	err = f.AddComment(c.Color, c.Comment)
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
