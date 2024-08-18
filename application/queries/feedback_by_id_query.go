package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type FeedbackByIdQuery struct {
	FeedbackId string `json:"feedbackId"`
}

func NewFeedbackByIdQuery(feedbackId string) *FeedbackByIdQuery {
	return &FeedbackByIdQuery{
		FeedbackId: feedbackId,
	}
}

func (q FeedbackByIdQuery) Execute(ctx context.Context) (*dto.Feedback, error) {
	feedbackId, err := uuid.Parse(q.FeedbackId)
	if err != nil {
		return nil, ErrBadRequest
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Feedback").(repositories.FeedbackRepository)
	f, err := r.Get(ctx, feedbackId)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	if err != nil {
		return nil, err
	}
	return dto.NewFeedbackDTO(*f), nil
}
