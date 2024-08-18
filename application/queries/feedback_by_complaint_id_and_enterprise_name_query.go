package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_enterprise"
	"go-complaint/infrastructure/persistence/finders/find_feedback"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type FeedbackByComplaintIdAndEnterpriseName struct {
	ComplaintId    string `json:"complaintId"`
	EnterpriseName string `json:"enterpriseName"`
}

func NewFeedbackByComplaintIdAndEnterpriseName(complaintId, enterpriseName string) *FeedbackByComplaintIdAndEnterpriseName {
	return &FeedbackByComplaintIdAndEnterpriseName{
		ComplaintId:    complaintId,
		EnterpriseName: enterpriseName,
	}
}

func (q FeedbackByComplaintIdAndEnterpriseName) Execute(ctx context.Context) (*dto.Feedback, error) {
	complaintId, err := uuid.Parse(q.ComplaintId)
	if err != nil {
		return nil, ErrBadRequest
	}
	reg := repositories.MapperRegistryInstance()
	enterpriseRepository, ok := reg.Get("Enterprise").(repositories.EnterpriseRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	dbE, err := enterpriseRepository.Find(ctx, find_enterprise.ByName(q.EnterpriseName))
	if err != nil {
		return nil, err
	}
	r, ok := reg.Get("Feedback").(repositories.FeedbackRepository)
	f, err := r.Find(ctx, find_feedback.ByComplaintIdAndEnterpriseId(complaintId, dbE.Id()))
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	if err != nil {
		return nil, err
	}

	return dto.NewFeedbackDTO(*f), nil
}
