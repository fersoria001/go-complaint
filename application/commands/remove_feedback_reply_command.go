package commands

import (
	"context"
	"go-complaint/application"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_all_complaint_replies"
	"go-complaint/infrastructure/persistence/removers/remove_all_feedback_replies"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type RemoveFeedbackReplyCommand struct {
	FeedbackId string   `json:"feedbackId"`
	Color      string   `json:"color"`
	RepliesIds []string `json:"repliesIds"`
}

func NewRemoveFeedbackReplyCommand(feedbackId, color string, repliesIds []string) *RemoveFeedbackReplyCommand {
	return &RemoveFeedbackReplyCommand{
		FeedbackId: feedbackId,
		Color:      color,
		RepliesIds: repliesIds,
	}
}

func (c RemoveFeedbackReplyCommand) Execute(ctx context.Context) error {
	feedbackId, err := uuid.Parse(c.FeedbackId)
	if err != nil {
		return err
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
	feedbackReplyRepository, ok := reg.Get("FeedbackReply").(repositories.FeedbackRepliesRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	f, err := feedbackRepository.Get(ctx, feedbackId)
	if err != nil {
		return err
	}
	s := make([]uuid.UUID, len(c.RepliesIds))
	for _, v := range c.RepliesIds {
		replyId, err := uuid.Parse(v)
		if err != nil {
			return err
		}
		s = append(s, replyId)
	}
	replies, err := repliesRepository.FindAll(ctx, find_all_complaint_replies.ByAnyIDs(s))
	if err != nil {
		return err
	}
	for _, v := range replies.ToSlice() {
		err = f.RemoveReply(c.Color, *v)
		if err != nil {
			return err
		}
	}
	replyReview, err := f.ReplyReview(c.Color)
	if err != nil {
		return err
	}
	if replyReview.Replies().IsEmpty() {
		id, err := f.RemoveReplyReview(replyReview)
		if err != nil {
			return err
		}
		err = feedbackReplyRepository.DeleteAll(ctx, remove_all_feedback_replies.WhereReplyReviewID(id))
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
