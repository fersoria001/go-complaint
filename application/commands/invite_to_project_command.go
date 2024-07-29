package commands

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/finders/find_enterprise"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"

	"github.com/google/uuid"
)

type InviteToProjectCommand struct {
	EnterpriseName string `json:"enterpriseName"`
	Role           string `json:"role"`
	ProposeTo      string `json:"proposeTo"`
	ProposedBy     string `json:"proposedBy"`
}

func NewInviteToProjectCommand(enterpriseName, role, proposeTo, proposedBy string) *InviteToProjectCommand {
	return &InviteToProjectCommand{
		EnterpriseName: enterpriseName,
		Role:           role,
		ProposeTo:      proposeTo,
		ProposedBy:     proposedBy,
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
	dbE, err := r.Find(ctx, find_enterprise.ByName(c.EnterpriseName))
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
	for _, v := range dbE.Employees().ToSlice() {
		if v.User.Id() == proposedToId {
			return fmt.Errorf("user's already hired")
		}
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if e, ok := event.(*enterprise.HiringInvitationSent); ok {
					recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
					if !ok {
						return ErrWrongTypeAssertion
					}
					proposedBy, err := recipientRepository.Get(ctx, proposedById)
					if err != nil {
						return err
					}
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
	err = dbE.InviteToProject(ctx, proposedById, proposedToId, role)
	if err != nil {
		return err
	}
	return nil
}
