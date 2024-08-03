package commands

import (
	"context"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
	"strings"

	"github.com/google/uuid"
)

type RecoverPasswordCommand struct {
	Username string `json:"username"`
}

func NewRecoverPasswordCommand(username string) *RecoverPasswordCommand {
	return &RecoverPasswordCommand{
		Username: username,
	}
}

func (c RecoverPasswordCommand) Execute(ctx context.Context) error {
	reg := repositories.MapperRegistryInstance()
	userRepository, ok := reg.Get("User").(repositories.UserRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	user, err := userRepository.Find(ctx, find_user.ByUsername(c.Username))
	if err != nil {
		return err
	}
	newRandomGeneratedPassword := strings.Join(strings.Split(uuid.New().String(), "-")[0:2], "-")
	newRandomGeneratedPassword = strings.Replace(newRandomGeneratedPassword, "-", "", 1)
	encryptedPassword, err := infrastructure.EncryptionServiceInstance().Encrypt(newRandomGeneratedPassword)
	if err != nil {
		return err
	}
	// 	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
	// 		HandleEvent: func(event domain.DomainEvent) error {
	// 			if _, ok := event.(*identity.PasswordReset); ok {
	// 				SendEmailCommand{
	// 					ToEmail:        userCommand.Email,
	// 					RandomPassword: newRandomGeneratedPassword,
	// 				}.PasswordRecovery(ctx)
	// 			}
	// 			return nil
	// 		},
	// 		SubscribedToEventType: func() reflect.Type {
	// 			return reflect.TypeOf(&identity.PasswordReset{})
	// 		},
	// 	})
	err = user.ResetPassword(ctx, newRandomGeneratedPassword, string(encryptedPassword))
	if err != nil {
		return err
	}
	err = userRepository.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
