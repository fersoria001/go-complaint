package queries

import (
	"context"
	"encoding/json"
	"go-complaint/domain/model/feedback"
	"go-complaint/dto"
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

func (query FeedbackQuery) FeedbackById(ctx context.Context) (*dto.Feedback, error) {
	if query.FeedbackID == "" {
		return nil, ErrBadRequest
	}
	id, err := uuid.Parse(query.FeedbackID)
	if err != nil {
		return nil, ErrBadRequest
	}
	f, err := repositories.MapperRegistryInstance().Get("Feedback").(repositories.FeedbackRepository).Get(
		ctx,
		id,
	)
	if err != nil {
		return nil, err
	}
	result := dto.NewFeedbackDTO(*f)
	return result, nil
}

func (query FeedbackQuery) FindByRevieweeID(ctx context.Context) ([]*dto.Feedback, error) {
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
	s := storedEvents
	s1 := make([]uuid.UUID, 0)
	revieweeId, err := uuid.Parse(query.RevieweeID)
	if err != nil {
		return nil, err
	}
	for _, v := range s {
		if v.TypeName == "*feedback.FeedbackDone" {
			var e feedback.FeedbackDone
			err = json.Unmarshal(v.EventBody, &e)
			if err != nil {
				return nil, err
			}
			if e.ReviewedId() == revieweeId {
				s1 = append(s1, e.FeedbackId())
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
	results := make([]*dto.Feedback, 0)
	s2 := dbFeedbacks.ToSlice()
	for _, v := range s2 {
		feedbackDto := dto.NewFeedbackDTO(*v)
		results = append(results, feedbackDto)
	}
	return results, nil
}
