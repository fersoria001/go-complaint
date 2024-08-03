package commands

import (
	"context"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type ChangeUserLastNameCommand struct {
	UserId      string `json:"userId"`
	NewLastName string `json:"newLastName"`
}

func NewChangeUserLastNameCommand(userId, newLastName string) *ChangeUserLastNameCommand {
	return &ChangeUserLastNameCommand{
		UserId:      userId,
		NewLastName: newLastName,
	}
}

func (c ChangeUserLastNameCommand) Execute(ctx context.Context) error {
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
	err = user.ChangeLastName(ctx, c.NewLastName)
	if err != nil {
		return err
	}
	err = r.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
