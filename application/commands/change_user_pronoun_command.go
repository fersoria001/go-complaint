package commands

import (
	"context"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type ChangeUserPronounCommand struct {
	UserId     string `json:"userId"`
	NewPronoun string `json:"newPronoun"`
}

func NewChangeUserPronounCommand(userId, newPronoun string) *ChangeUserPronounCommand {
	return &ChangeUserPronounCommand{
		UserId:     userId,
		NewPronoun: newPronoun,
	}
}

func (c ChangeUserPronounCommand) Execute(ctx context.Context) error {
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
	err = user.ChangePronoun(ctx, c.NewPronoun)
	if err != nil {
		return err
	}
	err = r.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
