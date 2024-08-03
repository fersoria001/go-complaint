package commands

import (
	"context"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type ChangeUserPhoneCommand struct {
	UserId   string `json:"userId"`
	NewPhone string `json:"newPhone"`
}

func NewChangeUserPhoneCommand(userId, newPhone string) *ChangeUserPhoneCommand {
	return &ChangeUserPhoneCommand{
		UserId:   userId,
		NewPhone: newPhone,
	}
}

func (c ChangeUserPhoneCommand) Execute(ctx context.Context) error {
	userId, err := uuid.Parse(c.UserId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("User").(repositories.UserRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	user, err := r.Get(ctx, userId)
	if err != nil {
		return err
	}
	err = user.ChangePhone(ctx, c.NewPhone)
	if err != nil {
		return err
	}
	err = r.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
