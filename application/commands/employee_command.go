package commands

type EmployeeCommand struct {
	EmployeeID   string
	EnterpriseID string
}

// func (command EmployeeCommand) LeaveEnterprise(
// 	ctx context.Context,
// ) error {
// 	parsedID, err := uuid.Parse(command.EmployeeID)
// 	if err != nil {
// 		return err
// 	}
// 	dbEnterprise, err := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Get(ctx, command.EnterpriseID)
// 	if err != nil {
// 		return err
// 	}
// 	r := repositories.MapperRegistryInstance().Get("Employee").(repositories.EmployeeRepository)
// 	emp, err := r.Get(ctx, parsedID)
// 	if err != nil {
// 		return err
// 	}
// 	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
// 		HandleEvent: func(event domain.DomainEvent) error {
// 			if e, ok := event.(*enterprise.EmployeeLeaved); ok {
// 				NotificationCommand{
// 					OwnerID:     e.EnterpriseID(),
// 					ThumbnailID: emp.Email(),
// 					Thumbnail:   emp.ProfileIMG(),
// 					Title:       fmt.Sprintf("%s has leaved %s", emp.FullName(), dbEnterprise.Name()),
// 					Content:     fmt.Sprintf("%s is no longer part of %s", emp.FullName(), dbEnterprise.Name()),
// 					Link:        fmt.Sprintf("/%s/employees", dbEnterprise.Name()),
// 				}.SaveNew(ctx)
// 				return nil
// 			}
// 			return &erros.ValueNotFoundError{}
// 		},
// 		SubscribedToEventType: func() reflect.Type {
// 			return reflect.TypeOf(&enterprise.EmployeeLeaved{})
// 		},
// 	})
// 	user, err := dbEnterprise.EmployeeLeave(ctx, parsedID)
// 	if err != nil {
// 		return err
// 	}
// 	err = repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Update(ctx, user)
// 	if err != nil {
// 		return err
// 	}
// 	err = repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Update(ctx, dbEnterprise)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
