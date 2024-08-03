package commands

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type InviteToProjectCommand struct {
	EnterpriseId string `json:"enterpriseName"`
	Role         string `json:"role"`
	ProposeTo    string `json:"proposeTo"`
	ProposedBy   string `json:"proposedBy"`
}

func NewInviteToProjectCommand(enterpriseId, role, proposeTo, proposedBy string) *InviteToProjectCommand {
	return &InviteToProjectCommand{
		EnterpriseId: enterpriseId,
		Role:         role,
		ProposeTo:    proposeTo,
		ProposedBy:   proposedBy,
	}
}

func (c InviteToProjectCommand) Execute(ctx context.Context) error {
	role := enterprise.ParsePosition(c.Role)
	if role < 0 {
		return fmt.Errorf("role doesn't exists")
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Enterprise").(repositories.EnterpriseRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	enterpriseId, err := uuid.Parse(c.EnterpriseId)
	if err != nil {
		return err
	}
	dbE, err := r.Get(ctx, enterpriseId)
	if err != nil {
		return err
	}
	proposedToId, err := uuid.Parse(c.ProposeTo)
	if err != nil {
		return err
	}
	proposedById, err := uuid.Parse(c.ProposedBy)
	if err != nil {
		return err
	}
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	proposedBy, err := recipientRepository.Get(ctx, proposedById)
	if err != nil {
		return err
	}
	enterpriseRecipient, err := recipientRepository.Get(ctx, dbE.Id())
	if err != nil {
		return err
	}
	proposedTo, err := recipientRepository.Get(ctx, proposedToId)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if e, ok := event.(*enterprise.HiringInvitationSent); ok {
					n := NewSendNotificationCommand(
						e.ProposedTo().String(),
						e.EnterpriseId().String(),
						"You have been invited to join a project",
						fmt.Sprintf("%s has invited you to join %s project", proposedBy.SubjectName(), dbE.Name()),
						"/hiring-invitations")
					return n.Execute(ctx)
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&enterprise.HiringInvitationSent{})
			},
		},
	)
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*enterprise.HiringInvitationSent); ok {
					hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
					if !ok {
						return ErrWrongTypeAssertion
					}
					date := time.Now()
					newHp := enterprise.NewHiringProccess(uuid.New(), *enterpriseRecipient, *proposedTo,
						role, enterprise.PENDING, "", *proposedBy, date, date, *proposedBy)
					err := hiringProccessRepository.Save(ctx, *newHp)
					if err != nil {
						return err
					}
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&enterprise.HiringInvitationSent{})
			},
		},
	)
	err = dbE.InviteToProject(ctx, proposedById, proposedToId, role)
	if err != nil {
		return err
	}
	return nil
}
