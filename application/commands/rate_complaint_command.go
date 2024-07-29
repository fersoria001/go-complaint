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
	dbC, err := repository.Get(ctx, complaintId)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*complaint.ComplaintClosed); ok {
					recipientRepository, ok := reg.Get("ComplaintData").(repositories.ComplaintDataRepository)
					if !ok {
						return ErrWrongTypeAssertion
					}
					newComplaintData := complaint.NewComplaintData(uuid.New(), userId, complaintId, time.Now(), complaint.REVIEWED)
					err := recipientRepository.Save(ctx, *newComplaintData)
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
	err = dbC.Rate(ctx, userId, rcc.Rate, rcc.Comment)
	if err != nil {
		return err
	}
	err = repository.Update(ctx, dbC)
	if err != nil {
		return err
	}
	return nil
}
