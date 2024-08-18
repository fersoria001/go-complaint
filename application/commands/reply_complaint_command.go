package commands

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/finders/find_all_user_roles"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"
	"slices"

	"github.com/google/uuid"
)

type ReplyComplaintCommand struct {
	SenderId    string `json:"senderId"`
	AliasId     string `json:"aliasId"`
	ComplaintId string `json:"complaintId"`
	Body        string `json:"body"`
}

func NewReplyComplaintCommand(senderId, aliasId, complaintId, body string) *ReplyComplaintCommand {
	return &ReplyComplaintCommand{
		SenderId:    senderId,
		AliasId:     aliasId,
		ComplaintId: complaintId,
		Body:        body,
	}
}

func (c ReplyComplaintCommand) Execute(ctx context.Context) error {
	newReplyId := uuid.New()
	complaintId, err := uuid.Parse(c.ComplaintId)

	if err != nil {
		return err
	}

	senderId, err := uuid.Parse(c.SenderId)

	if err != nil {
		return err
	}

	reg := repositories.MapperRegistryInstance()
	complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)

	if !ok {
		return ErrWrongTypeAssertion
	}

	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)

	if !ok {
		return ErrWrongTypeAssertion
	}

	sender, err := recipientRepository.Get(ctx, senderId)

	if err != nil {
		return err
	}
	dbC, err := complaintRepository.Get(ctx, complaintId)

	if err != nil {
		return err
	}

	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*complaint.ComplaintReplied); ok {
					userRolesRepository, ok := reg.Get("UserRole").(repositories.UserRoleRepository)
					if !ok {
						return ErrWrongTypeAssertion
					}
					dbRoles, err := userRolesRepository.FindAll(ctx, find_all_user_roles.ByUserId(senderId))
					if err != nil {
						return err
					}
					roles := dbRoles.ToSlice()

					if dbC.Author().IsEnterprise() && senderId != dbC.Receiver().Id() {
						if slices.ContainsFunc(roles, func(e *identity.UserRole) bool { return e.EnterpriseId() == dbC.Author().Id() }) {
							logEa := NewLogEnterpriseActivityCommand(
								senderId.String(),
								dbC.Id().String(),
								dbC.Author().Id().String(),
								dbC.Author().SubjectName(),
								enterprise.ComplaintReplied.String(),
							)
							err := logEa.Execute(ctx)
							if err != nil {
								return err
							}
						}
					}

					if dbC.Receiver().IsEnterprise() && senderId != dbC.Author().Id() {
						contains := slices.ContainsFunc(roles, func(e *identity.UserRole) bool {
							return e.EnterpriseId() == dbC.Receiver().Id()
						})
						if contains {
							logEa := NewLogEnterpriseActivityCommand(
								senderId.String(),
								dbC.Id().String(),
								dbC.Receiver().Id().String(),
								dbC.Receiver().SubjectName(),
								enterprise.ComplaintReplied.String(),
							)
							err := logEa.Execute(ctx)
							if err != nil {
								return err
							}
						}
					}

				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&complaint.ComplaintReplied{})
			},
		},
	)
	alias, err := uuid.Parse(c.AliasId)
	if err != nil {
		alias = uuid.Nil
	}
	err = dbC.Reply(
		ctx,
		newReplyId,
		*sender,
		c.Body,
		alias,
	)

	if err != nil {
		return err
	}

	cache.InMemoryInstance().Set(c.ComplaintId, newReplyId.String())
	err = complaintRepository.Update(ctx, dbC)

	if err != nil {
		return err
	}

	return nil
}
