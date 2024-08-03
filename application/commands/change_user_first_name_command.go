package commands

import (
	"context"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type ChangeUserFirstNameCommand struct {
	UserId       string `json:"userId"`
	NewFirstName string `json:"newFirstName"`
}

func NewChangeUserFirstNameCommand(userId, newFirstName string) *ChangeUserFirstNameCommand {
	return &ChangeUserFirstNameCommand{
		UserId:       userId,
		NewFirstName: newFirstName,
	}
}

func (c ChangeUserFirstNameCommand) Execute(ctx context.Context) error {
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
	err = user.ChangeFirstName(ctx, c.NewFirstName)
	if err != nil {
		return err
	}
	err = r.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
