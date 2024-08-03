package commands

import (
	"context"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type ChangeUserAddressCommand struct {
	UserId       string `json:"userId"`
	NewCountryId int    `json:"countryId"`
	NewCountyId  int    `json:"countyId"`
	NewCityId    int    `json:"cityId"`
}

func NewChangeUserAddressCommand(userId string, newCountryId,
	newCountyId, newCityId int) *ChangeUserAddressCommand {
	return &ChangeUserAddressCommand{
		UserId:       userId,
		NewCountryId: newCountryId,
		NewCountyId:  newCountyId,
		NewCityId:    newCityId,
	}
}

func (c ChangeUserAddressCommand) Execute(ctx context.Context) error {
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
	err = user.ChangeCountry(ctx, c.NewCountryId)
	if err != nil {
		return err
	}
	err = user.ChangeCountryState(ctx, c.NewCountyId)
	if err != nil {
		return err
	}
	err = user.ChangeCity(ctx, c.NewCityId)
	if err != nil {
		return err
	}
	err = r.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
