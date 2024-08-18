package commands

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"

	"github.com/google/uuid"
)

type RejectHiringInvitationCommand struct {
	UserId           string `json:"userId"`
	HiringProccessId string `json:"hiringProccessId"`
	RejectionReason  string `json:"rejectionReason"`
}

func NewRejectHiringInvitationCommand(userId, hiringProccessId, rejectionReason string) *RejectHiringInvitationCommand {
	return &RejectHiringInvitationCommand{
		UserId:           userId,
		HiringProccessId: hiringProccessId,
		RejectionReason:  rejectionReason,
	}
}

func (c RejectHiringInvitationCommand) Execute(ctx context.Context) error {
	userId, err := uuid.Parse(c.UserId)
	if err != nil {
		return err
	}
	hiringProccessId, err := uuid.Parse(c.HiringProccessId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	user, err := recipientRepository.Get(ctx, userId)
	if err != nil {
		return err
	}
	hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	hiringProccess, err := hiringProccessRepository.Get(ctx, hiringProccessId)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if e, ok := event.(*enterprise.HiringProccessStatusChanged); ok {
				if enterprise.ParseHiringProccessStatus(e.NewStatus()) == enterprise.REJECTED {
					c := NewSendNotificationCommand(
						hiringProccess.Enterprise().Id().String(),
						userId.String(),
						fmt.Sprintf(`%s rejected your invitation"`, user.SubjectName()),
						fmt.Sprintf("User %s rejected your invitation to join %s : %s", user.SubjectName(), hiringProccess.Enterprise().SubjectName(), c.RejectionReason),
						"/hiring",
					)
					return c.Execute(ctx)
				}
				return fmt.Errorf("unexpected event at reject hiring invitation handler")
			}
			return ErrWrongTypeAssertion
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.HiringProccessStatusChanged{})
		},
	})
	err = hiringProccess.ChangeStatus(ctx, enterprise.REJECTED, *user)
	if err != nil {
		return err
	}
	hiringProccess.WriteAReason(c.RejectionReason, *user)
	err = hiringProccessRepository.Update(ctx, *hiringProccess)
	if err != nil {
		return err
	}
	return nil
}
