package commands

import (
	"context"
	"go-complaint/application/application_services"
	"go-complaint/domain"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
	"os"
	"reflect"
)

type VerifyUserEmailCommand struct {
	ValidationToken string `json:"validationToken"`
}

func NewVerifyUserEmailCommand(validationToken string) *VerifyUserEmailCommand {
	return &VerifyUserEmailCommand{
		ValidationToken: validationToken,
	}
}

func (c VerifyUserEmailCommand) Execute(ctx context.Context) error {
	token, ok := cache.InMemoryInstance().Get(c.ValidationToken)
	if !ok {
		return ErrConfirmationNotFound
	}
	confirmation, ok := token.(bool)
	if !ok {
		return ErrWrongTypeAssertion
	}
	if confirmation {
		return ErrAlreadyVerified
	}
	emailVerification, err := application_services.JWTApplicationServiceInstance().ParseEmailVerification(c.ValidationToken)
	if err != nil {
		return err
	}
	environment := os.Getenv("ENVIRONMENT")
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("User").(repositories.UserRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	user, err := r.Find(ctx, find_user.ByUsername(emailVerification.Email))
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*identity.UserEmailVerified); ok {
				if environment == "PROD" {
					return NewSendEmailVerifiedEmailCommand(user.Email(), user.FullName()).Execute(ctx)
				}
			}
			return nil
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&identity.UserEmailVerified{})
		},
	})
	err = user.VerifyEmail(ctx)
	if err != nil {
		return err
	}
	err = r.Update(ctx, user)
	if err != nil {
		return err
	}
	cache.InMemoryInstance().Delete(c.ValidationToken)
	return nil
}
