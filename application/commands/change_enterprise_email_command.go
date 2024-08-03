package commands

import (
	"context"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type ChangeEnterpriseEmailCommand struct {
	EnterpriseId string `json:"enterpriseId"`
	NewEmail     string `json:"newEmail"`
}

func NewChangeEnterpriseEmailCommand(enterpriseId, newEmail string) *ChangeEnterpriseEmailCommand {
	return &ChangeEnterpriseEmailCommand{
		EnterpriseId: enterpriseId,
		NewEmail:     newEmail,
	}
}

func (c ChangeEnterpriseEmailCommand) Execute(ctx context.Context) error {
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
	err = dbEnterprise.ChangeEmail(ctx, c.NewEmail)
	if err != nil {
		return err
	}
	err = r.Update(ctx, dbEnterprise)
	if err != nil {
		return err
	}
	return nil
}
