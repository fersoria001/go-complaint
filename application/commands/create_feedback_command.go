package commands

import (
	"context"
	"go-complaint/domain/model/feedback"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type CreateFeedbackCommand struct {
	ComplaintId  string `json:"complaintId"`
	EnterpriseId string `json:"enterpriseId"`
}

func NewCreateFeedbackCommand(complaintId, enterpriseId string) *CreateFeedbackCommand {
	return &CreateFeedbackCommand{
		ComplaintId:  complaintId,
		EnterpriseId: enterpriseId,
	}
}

func (c CreateFeedbackCommand) Execute(ctx context.Context) error {
	complaintId, err := uuid.Parse(c.ComplaintId)
	if err != nil {
		return err
	}
	enterpriseId, err := uuid.Parse(c.EnterpriseId)
	if err != nil {
		return err
	}
	newId := uuid.New()
	f, err := feedback.CreateFeedback(ctx, newId, complaintId, enterpriseId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Feedback").(repositories.FeedbackRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = r.Save(ctx, f)
	if err != nil {
		return err
	}
	return nil
}
