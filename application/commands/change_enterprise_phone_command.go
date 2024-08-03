package commands

import (
	"context"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type ChangeEnterprisePhoneCommand struct {
	EnterpriseId string `json:"enterpriseId"`
	NewPhone     string `json:"newPhone"`
}

func NewChangeEnterprisePhoneCommand(enterpriseId, newPhone string) *ChangeEnterprisePhoneCommand {
	return &ChangeEnterprisePhoneCommand{
		EnterpriseId: enterpriseId,
		NewPhone:     newPhone,
	}
}

func (c ChangeEnterprisePhoneCommand) Execute(ctx context.Context) error {
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
	err = dbEnterprise.ChangePhone(ctx, c.NewPhone)
	if err != nil {
		return err
	}
	err = r.Update(ctx, dbEnterprise)
	if err != nil {
		return err
	}
	return nil
}
