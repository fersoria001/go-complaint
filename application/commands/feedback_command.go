package commands

import (
	"context"
	"fmt"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/feedback"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/persistence/finders/find_all_feedback"
	"go-complaint/infrastructure/persistence/repositories"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type FeedbackCommand struct {
	FeedbackID       string
	ComplaintID      string
	ReviewedID       string
	ReviewerID       string
	ReviewComment    string
	ReplyReviewColor string
	Replies          []string
	AnswerBody       string
}

func (command FeedbackCommand) AnswerFeedback(
	ctx context.Context,
) error {
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
	infrastructure.PushNotificationInMemoryQueueInstance().QueueNotification(
		infrastructure.Operation{
			ID:        fmt.Sprintf("feedbackLastReply:%s", dbFeedback.ID().String()),
			Operation: lastReply,
		})
	return nil
}

func (command FeedbackCommand) EndFeedback(
	ctx context.Context,
) error {
	if command.ComplaintID == "" ||
		command.ReviewerID == "" {
		return ErrBadRequest
	}
	parsedID, err := uuid.Parse(command.ComplaintID)
	if err != nil {
		return ErrBadRequest
	}
	dbComplaint, err := repositories.MapperRegistryInstance().Get(
		"Complaint",
	).(repositories.ComplaintRepository).Get(
		ctx,
		parsedID,
	)
	if err != nil {
		return err
	}
	err = dbComplaint.SendToHistory(ctx, command.ReviewerID)
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
	return nil
}

func (command FeedbackCommand) AddReplyReview(
	ctx context.Context,
) error {
	if command.ComplaintID == "" || command.ReviewedID == "" ||
		command.ReviewerID == "" || command.ReviewComment == "" ||
		command.ReplyReviewColor == "" || len(command.Replies) < 1 {
		return ErrBadRequest
	}
	thisDate := common.NewDate(time.Now())
	parsedID, err := uuid.Parse(command.ReviewerID)
	if err != nil {
		return ErrBadRequest
	}
	findFeedback, err := repositories.MapperRegistryInstance().Get(
		"Feedback",
	).(repositories.FeedbackRepository).FindAll(
		ctx,
		find_all_feedback.ByComplaintIDAndReviewedID(
			parsedID,
			command.ReviewedID,
		),
	)
	if err != nil {
		return err
	}
	var dbFeedback *feedback.Feedback
	var ok bool
	if dbFeedback, ok = findFeedback.Pop(); !ok {
		replyReviews := mapset.NewSet[*feedback.ReplyReview]()
		feedbackAnswers := mapset.NewSet[*feedback.Answer]()
		dbFeedback, err = feedback.NewFeedback(
			uuid.New(),
			parsedID,
			command.ReviewedID,
			replyReviews,
			feedbackAnswers,
		)
		if err != nil {
			return err
		}
	}
	feedbackID := dbFeedback.ID()
	replyReviewNewID := uuid.New()
	reviewedReplies := mapset.NewSet[complaint.Reply]()
	for _, replyID := range command.Replies {
		parseReplyID, err := uuid.Parse(replyID)
		if err != nil {
			return ErrBadRequest
		}
		reviewedReply, err := repositories.MapperRegistryInstance().Get(
			"Reply",
		).(repositories.ComplaintRepliesRepository).Get(
			ctx,
			parseReplyID,
		)
		if err != nil {
			return err
		}
		reviewedReplies.Add(*reviewedReply)
	}
	replyReviewID := uuid.New()
	review, err := feedback.NewReview(
		replyReviewID,
		command.ReviewerID,
		thisDate,
		command.ReviewComment,
	)
	if err != nil {
		return err
	}
	replyReview, err := feedback.NewReplyReview(
		replyReviewNewID,
		feedbackID,
		reviewedReplies,
		review,
		command.ReplyReviewColor,
	)
	if err != nil {
		return err
	}
	err = dbFeedback.AddReplyReview(
		ctx,
		replyReview,
	)
	if err != nil {
		return err
	}
	if !ok {
		err = repositories.MapperRegistryInstance().Get(
			"Feedback",
		).(repositories.FeedbackRepository).Save(
			ctx,
			dbFeedback,
		)
		if err != nil {
			return err
		}
	} else {
		err = repositories.MapperRegistryInstance().Get(
			"Feedback",
		).(repositories.FeedbackRepository).Update(
			ctx,
			dbFeedback,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
