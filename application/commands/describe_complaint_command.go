package commands

import (
	"context"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type DescribeComplaintCommand struct {
	ComplaintId string `json:"complaintId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewDescribeComplaintCommand(complaintId, title, description string) *DescribeComplaintCommand {
	return &DescribeComplaintCommand{
		ComplaintId: complaintId,
		Title:       title,
		Description: description,
	}
}

func (dcc DescribeComplaintCommand) Execute(ctx context.Context) error {
	id, err := uuid.Parse(dcc.ComplaintId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	c, err := repository.Get(ctx, id)
	if err != nil {
		return err
	}
	err = c.SetTitle(ctx, dcc.Title)
	if err != nil {
		return err
	}
	err = c.SetDescription(ctx, dcc.Description)
	if err != nil {
		return err
	}
	err = repository.Update(ctx, c)
	if err != nil {
		return err
	}
	return nil
}
