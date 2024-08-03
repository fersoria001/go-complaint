package commands

import (
	"context"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
)

type ChangePasswordCommand struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func NewChangePasswordCommand(username, oldPassword, newPassword string) *ChangePasswordCommand {
	return &ChangePasswordCommand{
		Username:    username,
		OldPassword: oldPassword,
		NewPassword: newPassword,
	}
}

func (c ChangePasswordCommand) Execute(ctx context.Context) error {
	reg := repositories.MapperRegistryInstance()
	userRepository, ok := reg.Get("User").(repositories.UserRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	user, err := userRepository.Find(ctx, find_user.ByUsername(c.Username))
	if err != nil {
		return err
	}
	svc := infrastructure.EncryptionServiceInstance()
	err = svc.Compare(user.Password(), c.OldPassword)
	if err != nil {
		return err
	}
	newEncrypted, err := svc.Encrypt(c.NewPassword)
	if err != nil {
		return err
	}
	err = user.ChangePassword(ctx, c.OldPassword, c.NewPassword, string(newEncrypted))
	if err != nil {
		return err
	}
	err = userRepository.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
