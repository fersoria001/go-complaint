package commands

import (
	"context"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type CreateNewComplaintCommand struct {
	AuthorId   string `json:"authorId"`
	ReceiverId string `json:"receiverId"`
}

func NewCreateNewComplaintCommand(authorId, receiverId string) *CreateNewComplaintCommand {
	return &CreateNewComplaintCommand{
		AuthorId:   authorId,
		ReceiverId: receiverId,
	}
}

func (c CreateNewComplaintCommand) Execute(ctx context.Context) error {
	id := uuid.New()
	authorId, err := uuid.Parse(c.AuthorId)
	if err != nil {
		return err
	}
	receiverId, err := uuid.Parse(c.ReceiverId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	author, err := recipientRepository.Get(ctx, authorId)
	if err != nil {
		return err
	}
	receiver, err := recipientRepository.Get(ctx, receiverId)
	if err != nil {
		return err
	}
	newComplaint, err := complaint.CreateNew(ctx, id, *author, *receiver)
	if err != nil {
		return err
	}
	err = complaintRepository.Save(ctx, newComplaint)
	if err != nil {
		return err
	}
	return nil
}
