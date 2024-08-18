package commands

import (
	"context"
	"errors"
	"go-complaint/domain/model/feedback"
	"go-complaint/infrastructure/persistence/finders/find_enterprise"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type CreateFeedbackCommand struct {
	ComplaintId    string `json:"complaintId"`
	EnterpriseName string `json:"enterpriseName"`
}

func NewCreateFeedbackCommand(complaintId, enterpriseName string) *CreateFeedbackCommand {
	return &CreateFeedbackCommand{
		ComplaintId:    complaintId,
		EnterpriseName: enterpriseName,
	}
}

func (c CreateFeedbackCommand) Execute(ctx context.Context) error {
	complaintId, err := uuid.Parse(c.ComplaintId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	enterpriseRepository, ok := reg.Get("Enterprise").(repositories.EnterpriseRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	dbE, err := enterpriseRepository.Find(ctx, find_enterprise.ByName(c.EnterpriseName))
	if err != nil {
		return err
	}
	newId := uuid.New()
	f, err := feedback.CreateFeedback(ctx, newId, complaintId, dbE.Id())
	if err != nil {
		return err
	}
	r, ok := reg.Get("Feedback").(repositories.FeedbackRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = r.Save(ctx, f)
	if err != nil {
		if errors.Is(err, repositories.ErrFeedbackAlreadyExists) {
			return ErrFeedbackAlreadyExists
		}
		return err
	}
	return nil
}
