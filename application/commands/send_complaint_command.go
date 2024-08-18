package commands

import (
	"context"
	"fmt"
	"go-complaint/application"
	"go-complaint/domain"
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/enterprise"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"

	"github.com/google/uuid"
)

type SendComplaintCommand struct {
	ComplaintId   string `json:"complaintId"`
	CurrentUserId string `json:"currentUserId"`
	Body          string `json:"body"`
}

func NewSendComplaintCommand(complaintId, currentUserId, body string) *SendComplaintCommand {
	return &SendComplaintCommand{
		ComplaintId:   complaintId,
		CurrentUserId: currentUserId,
		Body:          body,
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
	currentUserId, err := uuid.Parse(scc.CurrentUserId)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*complaint.ComplaintSent); ok {
					err := NewLogComplaintDataCommand(
						c.Receiver().Id().String(), c.Author().Id().String(),
						c.Receiver().Id().String(), c.Id().String(), complaint.RECEIVED.String(),
					).Execute(ctx)
					if err != nil {
						return err
					}
					err = NewSendNotificationCommand(
						c.Receiver().Id().String(),
						c.Author().Id().String(),
						"You received a new complaint!",
						fmt.Sprintf("%s sent you a complaint complaint", c.Author().SubjectName()),
						"/complaints/",
					).Execute(ctx)
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
					if currentUserId != c.Author().Id() {
						err := NewLogEnterpriseActivityCommand(
							currentUserId.String(),
							c.Id().String(),
							c.Author().Id().String(),
							c.Author().SubjectName(),
							enterprise.ComplaintSent.String(),
						).Execute(ctx)
						if err != nil {
							return err
						}
					}
					err := NewLogComplaintDataCommand(
						c.Author().Id().String(), c.Author().Id().String(),
						c.Receiver().Id().String(), c.Id().String(), complaint.SENT.String(),
					).Execute(ctx)
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
	svc := application.ApplicationMessagePublisherInstance()
	svc.Publish(application.NewApplicationMessage(c.Author().Id().String(), "complaint", *dto.NewComplaint(*c)))
	svc.Publish(application.NewApplicationMessage(c.Receiver().Id().String(), "complaint", *dto.NewComplaint(*c)))

	return nil
}
