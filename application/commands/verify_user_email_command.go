package commands

type UserCommand struct {
}

// func (userCommand UserCommand) VerifyEmail(ctx context.Context) error {
// 	if userCommand.EmailVerificationToken == "" {
// 		return ErrBadRequest
// 	}
// 	token, ok := cache.InMemoryCacheInstance().Get(
// 		userCommand.EmailVerificationToken)
// 	if !ok {
// 		return ErrConfirmationNotFound
// 	}
// 	confirmation, ok := token.(bool)
// 	if !ok {
// 		return ErrWrongTypeAssertion
// 	}
// 	if confirmation {
// 		return ErrAlreadyVerified
// 	}
// 	emailVerification, err := application_services.JWTApplicationServiceInstance().ParseEmailVerification(userCommand.EmailVerificationToken)
// 	if err != nil {
// 		return err
// 	}
// 	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
// 		HandleEvent: func(event domain.DomainEvent) error {
// 			if _, ok := event.(*identity.UserEmailVerified); ok {
// 				SendEmailCommand{
// 					ToEmail: emailVerification.Email,
// 					ToName:  userCommand.FullName,
// 				}.EmailVerified(ctx)
// 			}
// 			return nil
// 		},
// 		SubscribedToEventType: func() reflect.Type {
// 			return reflect.TypeOf(&identity.UserEmailVerified{})
// 		},
// 	})

// 	mapper := repositories.MapperRegistryInstance().Get("User")
// 	if mapper == nil {
// 		return repositories.ErrMapperNotRegistered
// 	}
// 	userMapper, ok := mapper.(repositories.UserRepository)
// 	if !ok {
// 		return repositories.ErrWrongTypeAssertion
// 	}
// 	user, err := userMapper.Get(ctx, emailVerification.Email)
// 	if err != nil {
// 		return err
// 	}
// 	err = user.VerifyEmail(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	err = userMapper.Update(ctx, user)
// 	if err != nil {
// 		return err
// 	}
// 	cache.InMemoryCacheInstance().Delete(userCommand.EmailVerificationToken)
// 	return nil
// }
