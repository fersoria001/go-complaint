package commands

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type SendComplaintCommand struct {
	ComplaintId string `json:"complaintId"`
	Body        string `json:"body"`
}

func NewSendComplaintCommand(complaintId, body string) *SendComplaintCommand {
	return &SendComplaintCommand{
		ComplaintId: complaintId,
		Body:        body,
	}
}

func (scc SendComplaintCommand) Execute(ctx context.Context) error {
	id, err := uuid.Parse(scc.ComplaintId)
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
	err = c.SetBody(ctx, scc.Body)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*complaint.ComplaintSent); ok {
					recipientRepository, ok := reg.Get("ComplaintData").(repositories.ComplaintDataRepository)
					if !ok {
						return ErrWrongTypeAssertion
					}
					newComplaintData := complaint.NewComplaintData(uuid.New(), c.Receiver().Id(), c.Id(), time.Now(), complaint.RECEIVED)
					err := recipientRepository.Save(ctx, *newComplaintData)
					if err != nil {
						return err
					}
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&complaint.ComplaintSent{})
			},
		},
	)
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*complaint.ComplaintSent); ok {
					recipientRepository, ok := reg.Get("ComplaintData").(repositories.ComplaintDataRepository)
					if !ok {
						return ErrWrongTypeAssertion
					}
					newComplaintData := complaint.NewComplaintData(uuid.New(), c.Author().Id(), c.Id(), time.Now(), complaint.SENT)
					err := recipientRepository.Save(ctx, *newComplaintData)
					if err != nil {
						return err
					}
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&complaint.ComplaintSent{})
			},
		},
	)
	err = c.Send(ctx)
	if err != nil {
		return err
	}
	err = repository.Update(ctx, c)
	if err != nil {
		return err
	}
	return nil
}
