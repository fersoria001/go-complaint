package queries

import (
	"context"
	"encoding/json"
	"fmt"
	"go-complaint/domain/model/feedback"
	"go-complaint/dto"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/finders/find_all_events"
	"go-complaint/infrastructure/persistence/finders/find_all_feedback"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type FeedbackQuery struct {
	ComplaintID string
	FeedbackID  string
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
	v, ok := cache.InMemoryCacheInstance().Get(objectID)
	if !ok {
		return dto.FeedbackAnswer{}, fmt.Errorf("complaint not found in cache %v", objectID)
	}
	cachedReply, ok := v.(*feedback.Answer)
	if !ok {
		return dto.FeedbackAnswer{}, fmt.Errorf("incorrect type of cached object %v", objectID)
	}
	return dto.NewFeedbackAnswerDTO(*cachedReply), nil
}

func (query FeedbackQuery) FindByComplaintID(ctx context.Context) (dto.Feedback, error) {
	if query.ComplaintID == "" {
		return dto.Feedback{}, ErrBadRequest
	}
	parsedComplaintID, err := uuid.Parse(query.ComplaintID)
	if err != nil {
		return dto.Feedback{}, ErrBadRequest
	}
	feedbacks, err := repositories.MapperRegistryInstance().Get("Feedback").(repositories.FeedbackRepository).FindAll(
		ctx,
		find_all_feedback.ByComplaintID(parsedComplaintID),
	)
	if err != nil {
		return dto.Feedback{}, err
	}
	f, ok := feedbacks.Pop()
	if !ok {
		return dto.Feedback{}, ErrNotFound
	}
	result := dto.NewFeedbackDTO(*f)
	return result, nil
}

func (query FeedbackQuery) Feedback(ctx context.Context) (dto.Feedback, error) {
	if query.FeedbackID == "" {
		return dto.Feedback{}, ErrBadRequest
	}
	id, err := uuid.Parse(query.FeedbackID)
	if err != nil {
		return dto.Feedback{}, ErrBadRequest
	}
	f, err := repositories.MapperRegistryInstance().Get("Feedback").(repositories.FeedbackRepository).Get(
		ctx,
		id,
	)
	if err != nil {
		return dto.Feedback{}, err
	}
	result := dto.NewFeedbackDTO(*f)
	return result, nil
}

func (query FeedbackQuery) FindByRevieweeID(ctx context.Context) ([]dto.Feedback, error) {
	if query.RevieweeID == "" {
		return nil, ErrBadRequest
	}
	storedEvents, err := repositories.MapperRegistryInstance().Get("Event").(repositories.EventRepository).FindAll(
		ctx,
		find_all_events.By(),
	)
	if err != nil {
		return nil, err
	}
	s := storedEvents.ToSlice()
	s1 := make([]uuid.UUID, 0)
	for _, v := range s {
		if v.TypeName == "*feedback.FeedbackDone" {
			var e feedback.FeedbackDone
			err = json.Unmarshal(v.EventBody, &e)
			if err != nil {
				return nil, err
			}
			if e.ReviewedID() == query.RevieweeID {
				s1 = append(s1, e.FeedbackID())
			}
		}
	}
	if len(s1) <= 0 {
		return nil, ErrNotFound
	}
	r := repositories.MapperRegistryInstance().Get("Feedback").(repositories.FeedbackRepository)
	dbFeedbacks, err := r.FindAll(ctx, find_all_feedback.ByAnyIDs(s1))
	if err != nil {
		return nil, err
	}
	if dbFeedbacks.Cardinality() <= 0 {
		return nil, ErrNotFound
	}
	results := make([]dto.Feedback, 0)
	s2 := dbFeedbacks.ToSlice()
	for _, v := range s2 {
		feedbackDto := dto.NewFeedbackDTO(*v)
		results = append(results, feedbackDto)
	}
	return results, nil
}
