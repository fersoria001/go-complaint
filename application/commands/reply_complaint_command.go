package commands

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"

	"github.com/google/uuid"
)

type ReplyComplaintCommand struct {
	SenderId    string `json:"senderId"`
	ComplaintId string `json:"complaintId"`
	Body        string `json:"body"`
}

func NewReplyComplaintCommand(senderId, complaintId, body string) *ReplyComplaintCommand {
	return &ReplyComplaintCommand{
		SenderId:    senderId,
		ComplaintId: complaintId,
		Body:        body,
	}
}

func (c ReplyComplaintCommand) Execute(ctx context.Context) error {
	newReplyId := uuid.New()
	complaintId, err := uuid.Parse(c.ComplaintId)
	if err != nil {
		return err
	}
	senderId, err := uuid.Parse(c.SenderId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	sender, err := recipientRepository.Get(ctx, senderId)
	if err != nil {
		return err
	}
	dbC, err := complaintRepository.Get(ctx, complaintId)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				// if _, ok := event.(*complaint.ComplaintReplied); ok {
				// }
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&complaint.ComplaintReplied{})
			},
		},
	)
	err = dbC.Reply(
		ctx,
		newReplyId,
		*sender,
		c.Body,
	)
	if err != nil {
		return err
	}
	cache.InMemoryInstance().Set(c.ComplaintId, newReplyId.String())
	err = complaintRepository.Update(ctx, dbC)
	if err != nil {
		return err
	}
	return nil
}
