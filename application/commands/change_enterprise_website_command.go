package commands

import (
	"context"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type ChangeEnterpriseWebsiteCommand struct {
	EnterpriseId string `json:"enterpriseId"`
	NewWebsite   string `json:"newWebsite"`
}

func NewChangeEnterpriseWebsiteCommand(enterpriseId, newWebsite string) *ChangeEnterpriseWebsiteCommand {
	return &ChangeEnterpriseWebsiteCommand{
		EnterpriseId: enterpriseId,
		NewWebsite:   newWebsite,
	}
}

func (c ChangeEnterpriseWebsiteCommand) Execute(ctx context.Context) error {
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
	err = dbEnterprise.ChangeWebsite(ctx, c.NewWebsite)
	if err != nil {
		return err
	}
	err = r.Update(ctx, dbEnterprise)
	if err != nil {
		return err
	}
	return nil
}
