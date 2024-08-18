package commands

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"

	"github.com/google/uuid"
)

type SendComplaintToReviewCommand struct {
	ReceiverId    string `json:"receiverId"`
	ComplaintId   string `json:"complaintId"`
	CurrentUserId string `json:"currentUserId"`
}

func NewSendComplaintToReviewCommand(receiverId, complaintId, currentUserId string) *SendComplaintToReviewCommand {
	return &SendComplaintToReviewCommand{
		ReceiverId:    receiverId,
		ComplaintId:   complaintId,
		CurrentUserId: currentUserId,
	}
}

func (c SendComplaintToReviewCommand) Execute(ctx context.Context) error {
	complaintId, err := uuid.Parse(c.ComplaintId)
	if err != nil {
		return err
	}
	receiverId, err := uuid.Parse(c.ReceiverId)
	if err != nil {
		return err
	}
	currentUserId, err := uuid.Parse(c.CurrentUserId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	triggeredBy, err := recipientRepository.Get(ctx, receiverId)
	if err != nil {
		return err
	}
	dbC, err := repository.Get(ctx, complaintId)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*complaint.ComplaintSentForReview); ok {
					if currentUserId != receiverId {
						err := NewLogEnterpriseActivityCommand(
							currentUserId.String(),
							dbC.Id().String(),
							dbC.Receiver().Id().String(),
							dbC.Receiver().SubjectName(),
							enterprise.ComplaintResolved.String(),
						).Execute(ctx)
						if err != nil {
							return err
						}
					}
					err = NewSendNotificationCommand(
						dbC.Author().Id().String(),
						dbC.Receiver().Id().String(),
						"A complaint attention needs your review!",
						fmt.Sprintf("%s ask for you to review his/her attention on your complaint", dbC.Author().SubjectName()),
						"/reviews",
					).Execute(ctx)
					if err != nil {
						return err
					}
					err := NewLogComplaintDataCommand(
						receiverId.String(),
						dbC.Author().Id().String(),
						dbC.Receiver().Id().String(),
						complaintId.String(), complaint.RESOLVED.String(),
					).Execute(ctx)
					if err != nil {
						return err
					}
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&complaint.ComplaintSentForReview{})
			},
		},
	)
	err = dbC.MarkAsReviewable(ctx, *triggeredBy)
	if err != nil {
		return err
	}
	err = repository.Update(ctx, dbC)
	if err != nil {
		return err
	}
	return nil
}
