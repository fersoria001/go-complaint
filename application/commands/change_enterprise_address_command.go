package commands

import (
	"context"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type ChangeEnterpriseAddressCommand struct {
	EnterpriseId string `json:"enterpriseId"`
	NewCountryId int    `json:"newCountryId"`
	NewCountyId  int    `json:"newCountyId"`
	NewCityId    int    `json:"newCityId"`
}

func NewChangeEnterpriseAddressCommand(enterpriseId string, newCountryId, newCountyId, newCityId int) *ChangeEnterpriseAddressCommand {
	return &ChangeEnterpriseAddressCommand{
		EnterpriseId: enterpriseId,
		NewCountryId: newCountryId,
		NewCountyId:  newCountyId,
		NewCityId:    newCityId,
	}
}

func (c ChangeEnterpriseAddressCommand) Execute(ctx context.Context) error {
	enterpriseId, err := uuid.Parse(c.EnterpriseId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Enterprise").(repositories.EnterpriseRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	dbEnterprise, err := r.Get(ctx, enterpriseId)
	if err != nil {
		return err
	}
	err = dbEnterprise.ChangeCountry(ctx, c.NewCountryId)
	if err != nil {
		return err
	}
	err = dbEnterprise.ChangeCountryState(ctx, c.NewCountyId)
	if err != nil {
		return err
	}
	err = dbEnterprise.ChangeCity(ctx, c.NewCityId)
	if err != nil {
		return err
	}
	err = r.Update(ctx, dbEnterprise)
	if err != nil {
		return err
	}
	return nil
}
