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

type RateComplaintCommand struct {
	UserId      string `json:"userId"`
	ComplaintId string `json:"complaintId"`
	Rate        int    `json:"rate"`
	Comment     string `json:"comment"`
}

func NewRateComplaintCommand(userId, complaintId, comment string, rate int) *RateComplaintCommand {
	return &RateComplaintCommand{
		UserId:      userId,
		ComplaintId: complaintId,
		Comment:     comment,
		Rate:        rate,
	}
}

func (rcc RateComplaintCommand) Execute(ctx context.Context) error {
	userId, err := uuid.Parse(rcc.UserId)
	if err != nil {
		return err
	}
	complaintId, err := uuid.Parse(rcc.ComplaintId)
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
	ratedBy, err := recipientRepository.Get(ctx, userId)
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
				if _, ok := event.(*complaint.ComplaintClosed); ok {
					if dbC.Author().IsEnterprise() {
						err := NewLogEnterpriseActivityCommand(
							userId.String(),
							dbC.Id().String(),
							dbC.Author().Id().String(),
							dbC.Author().SubjectName(),
							enterprise.ComplaintReviewed.String(),
						).Execute(ctx)
						if err != nil {
							return err
						}
					}
					err := NewSendNotificationCommand(
						dbC.Receiver().Id().String(),
						userId.String(),
						"Your assistance has been rated",
						fmt.Sprintf("%s has rated your complaint assistance", ratedBy.SubjectName()),
						"/reviews",
					).Execute(ctx)
					if err != nil {
						return err
					}
					err = NewLogComplaintDataCommand(
						dbC.Author().Id().String(),
						dbC.Author().Id().String(),
						dbC.Receiver().Id().String(),
						complaintId.String(),
						complaint.REVIEWED.String(),
					).Execute(ctx)
					if err != nil {
						return err
					}
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&complaint.ComplaintClosed{})
			},
		},
	)
	err = dbC.Rate(ctx, *ratedBy, rcc.Rate, rcc.Comment)
	if err != nil {
		return err
	}
	err = repository.Update(ctx, dbC)
	if err != nil {
		return err
	}
	return nil
}
