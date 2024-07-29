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

type SendComplaintToReviewCommand struct {
	ReceiverId  string `json:"receiverId"`
	ComplaintId string `json:"complaintId"`
}

func NewSendComplaintToReviewCommand(receiverId, complaintId string) *SendComplaintToReviewCommand {
	return &SendComplaintToReviewCommand{
		ReceiverId:  receiverId,
		ComplaintId: complaintId,
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
	reg := repositories.MapperRegistryInstance()
	repository := reg.Get("Complaint").(repositories.ComplaintRepository)
	dbC, err := repository.Get(ctx, complaintId)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*complaint.ComplaintSentForReview); ok {
					recipientRepository, ok := reg.Get("ComplaintData").(repositories.ComplaintDataRepository)
					if !ok {
						return ErrWrongTypeAssertion
					}
					newComplaintData := complaint.NewComplaintData(uuid.New(), receiverId, complaintId, time.Now(), complaint.RESOLVED)
					err := recipientRepository.Save(ctx, *newComplaintData)
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
	err = dbC.MarkAsReviewable(ctx, receiverId)
	if err != nil {
		return err
	}
	err = repository.Update(ctx, dbC)
	if err != nil {
		return err
	}
	return nil
}
