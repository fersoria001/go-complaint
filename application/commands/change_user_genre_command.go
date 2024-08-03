package commands

import (
	"context"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type ChangeUserGenreCommand struct {
	UserId   string `json:"userId"`
	NewGenre string `json:"newGenre"`
}

func NewChangeUserGenreCommand(userId, newGenre string) *ChangeUserGenreCommand {
	return &ChangeUserGenreCommand{
		UserId:   userId,
		NewGenre: newGenre,
	}
}

func (c ChangeUserGenreCommand) Execute(ctx context.Context) error {
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
	err = user.ChangeGenre(ctx, c.NewGenre)
	if err != nil {
		return err
	}
	err = r.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
