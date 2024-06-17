package queries

import (
	"context"
	"fmt"
	"go-complaint/domain/model/feedback"
	"go-complaint/dto"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/persistence/finders/find_all_feedback"
	"go-complaint/infrastructure/persistence/finders/find_all_feedback_reply_review"
	"go-complaint/infrastructure/persistence/finders/find_all_feedback_review"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type FeedbackQuery struct {
	ComplaintID string
	ReviewerID  string
	RevieweeID  string
	ReplyID     string
}

func (query FeedbackQuery) FeedbackLastReply(
	ctx context.Context,
) (dto.FeedbackAnswer, error) {
	if query.ReplyID == "" {
		return dto.FeedbackAnswer{}, ErrBadRequest
	}
	objectID := "feedbackLastReply:" + query.ReplyID
	v, ok := infrastructure.InMemoryCacheInstance().Get(objectID)
	if !ok {
		return dto.FeedbackAnswer{}, fmt.Errorf("complaint not found in cache %v", objectID)
	}
	cachedReply, ok := v.(*feedback.Answer)
	if !ok {
		return dto.FeedbackAnswer{}, fmt.Errorf("incorrect type of cached object %v", objectID)
	}
	return dto.NewFeedbackAnswerDTO(*cachedReply), nil
}

func (query FeedbackQuery) FindByComplaintID(ctx context.Context) (
	[]dto.Feedback,
	error) {
	if query.ComplaintID == "" {
		return nil, ErrBadRequest
	}
	parsedComplaintID, err := uuid.Parse(query.ComplaintID)
	if err != nil {
		return nil, ErrBadRequest
	}
	feedbacks, err := repositories.MapperRegistryInstance().Get("Feedback").(repositories.FeedbackRepository).FindAll(
		ctx,
		find_all_feedback.ByComplaintID(parsedComplaintID),
	)
	if err != nil {
		return nil, err
	}
	result := make([]dto.Feedback, 0, feedbacks.Cardinality())
	for _, f := range feedbacks.ToSlice() {
		result = append(result, dto.NewFeedbackDTO(*f))
	}
	return result, nil
}
func (query FeedbackQuery) FindByReviewerID(
	ctx context.Context,
) ([]dto.Feedback, error) {
	if query.ReviewerID == "" {
		return nil, ErrBadRequest
	}
	feedbacks, err := repositories.MapperRegistryInstance().Get("Feedback").(repositories.FeedbackRepository).FindAll(
		ctx,
		find_all_feedback.ByReviewedID(query.ReviewerID),
	)
	if err != nil {
		return nil, err
	}
	result := make([]dto.Feedback, 0, feedbacks.Cardinality())
	for _, f := range feedbacks.ToSlice() {
		result = append(result, dto.NewFeedbackDTO(*f))
	}
	return result, nil
}
func (query FeedbackQuery) FindByRevieweeID(
	ctx context.Context,
) ([]dto.Feedback, error) {
	if query.RevieweeID == "" {
		return nil, ErrBadRequest
	}
	reviews, err := repositories.MapperRegistryInstance().Get("Review").(repositories.FeedbackReviewRepository).FindAll(
		ctx,
		find_all_feedback_review.ByReviewerID(query.RevieweeID),
	)
	if err != nil {
		return nil, err
	}
	if reviews.Cardinality() == 0 {
		return nil, ErrNotFound
	}
	for review := range reviews.Iter() {
		replyReviews, err := repositories.MapperRegistryInstance().Get("ReplyReview").(repositories.FeedbackReplyReviewRepository).FindAll(
			ctx,
			find_all_feedback_reply_review.ByReviewID(review.ReplyReviewID()),
		)
		if err != nil {
			return nil, err
		}
		for replyReview := range replyReviews.Iter() {
			feedbacks, err := repositories.MapperRegistryInstance().Get("Feedback").(repositories.FeedbackRepository).FindAll(
				ctx,
				find_all_feedback.ByID(replyReview.FeedbackID()),
			)
			if err != nil {
				return nil, err
			}
			result := make([]dto.Feedback, 0, feedbacks.Cardinality())
			for _, f := range feedbacks.ToSlice() {
				result = append(result, dto.NewFeedbackDTO(*f))
			}
			return result, nil
		}
	}
	return nil, ErrNotFound
}
