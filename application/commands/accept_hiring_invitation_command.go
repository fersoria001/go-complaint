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

type AcceptHiringInvitationCommand struct {
	UserId           string `json:"userId"`
	HiringProccessId string `json:"hiringProccessId"`
}

func NewAcceptHiringInvitationCommand(userId, hiringProccessId string) *AcceptHiringInvitationCommand {
	return &AcceptHiringInvitationCommand{
		UserId:           userId,
		HiringProccessId: hiringProccessId,
	}
}

func (c AcceptHiringInvitationCommand) Execute(ctx context.Context) error {
	hiringProccessId, err := uuid.Parse(c.HiringProccessId)
	if err != nil {
		return err
	}
	userId, err := uuid.Parse(c.UserId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	hiringProccess, err := hiringProccessRepository.Get(ctx, hiringProccessId)
	if err != nil {
		return err
	}
	user, err := recipientRepository.Get(ctx, userId)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if event, ok := event.(*enterprise.HiringProccessStatusChanged); ok {
				if enterprise.ParseHiringProccessStatus(event.NewStatus()) == enterprise.USER_ACCEPTED {
					c := NewSendNotificationCommand(
						hiringProccess.Enterprise().Id().String(),
						user.Id().String(),
						fmt.Sprintf(`%s accepted your invitation"`, user.SubjectName()),
						fmt.Sprintf("User %s accepted your invitation to join %s", user.SubjectName(), hiringProccess.Enterprise().SubjectName()),
						fmt.Sprintf("/%s/hiring-procceses?id=%s", hiringProccess.Enterprise().Id(), hiringProccess.Id()),
					)
					return c.Execute(ctx)
				}
				return fmt.Errorf("unexpected event in accept hiring invitation command")
			}
			return ErrWrongTypeAssertion
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.HiringProccessStatusChanged{})
		},
	})
	err = hiringProccess.ChangeStatus(ctx, enterprise.USER_ACCEPTED, *user)
	if err != nil {
		return err
	}
	err = hiringProccessRepository.Update(ctx, *hiringProccess)
	if err != nil {
		return err
	}
	return nil
}
