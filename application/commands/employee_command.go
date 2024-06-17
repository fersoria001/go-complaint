package commands

import (
	"context"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type EmployeeCommand struct {
	EmployeeID   string
	EnterpriseID string
}

func (command EmployeeCommand) LeaveEnterprise(
	ctx context.Context,
) error {
	parsedID, err := uuid.Parse(command.EmployeeID)
	if err != nil {
		return err
	}
	dbEnterprise, err := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Get(ctx, command.EnterpriseID)
	if err != nil {
		return err
	}
	user, err := dbEnterprise.EmployeeLeave(ctx, parsedID)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Update(ctx, user)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Update(ctx, dbEnterprise)
	if err != nil {
		return err
	}
	return nil
}
