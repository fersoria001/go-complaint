package application

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/identity"
	"go-complaint/erros"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Package application
type IdentityService struct {
	repository *repositories.UserRepository
}

func NewIdentityService(
	repository *repositories.UserRepository,
) *IdentityService {
	return &IdentityService{repository: repository}
}

func (is *IdentityService) ChangeRole(
	ctx context.Context,
	email,
	enterprise,
	oldRole,
	newRole string) error {
	user, err := is.repository.Get(ctx, email)
	if err != nil {
		return err
	}
	parsedOldRole, err := identity.ParseRole(oldRole)
	if err != nil {
		return err
	}
	parsedNewRole, err := identity.ParseRole(newRole)
	if err != nil {
		return err
	}
	err = user.AddRole(ctx, identity.NewRole(parsedNewRole), enterprise)
	if err != nil {
		return err
	}
	oldUserRole := identity.NewUserRole(identity.NewRole(parsedOldRole), user, enterprise)
	err = user.RemoveUserRole(ctx, oldUserRole)
	if err != nil {
		return err
	}
	return is.repository.Update(ctx, user)
}

func (is *IdentityService) RemoveRole(
	ctx context.Context,
	email,
	role,
	enterprise string) error {
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*identity.RoleRemoved); ok {
				fmt.Println("Role Removed Handler")
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&identity.RoleRemoved{})
		},
	})

	user, err := is.repository.Get(ctx, email)
	if err != nil {
		return err
	}
	parsedRole, err := identity.ParseRole(role)
	if err != nil {
		return err
	}
	userRole := identity.NewUserRole(identity.NewRole(parsedRole), user, enterprise)
	err = user.RemoveUserRole(ctx, userRole)
	if err != nil {
		return err
	}
	return is.repository.Update(ctx, user)
}

func (is *IdentityService) AddNewRole(
	ctx context.Context,
	email,
	role,
	enterprise string) error {
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*identity.RoleAdded); ok {
				fmt.Println("Role Added Handler")
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&identity.RoleAdded{})
		},
	})
	user, err := is.repository.Get(ctx, email)
	if err != nil {
		return err
	}
	parsedRole, err := identity.ParseRole(role)
	if err != nil {
		return err
	}
	err = user.AddRole(ctx, identity.NewRole(parsedRole), enterprise)
	if err != nil {
		return err
	}
	return is.repository.Update(ctx, user)
}

func (is *IdentityService) UpdateUserProfile(
	ctx context.Context,
	id string,
	profileIMG,
	firstName,
	lastName,
	phone,
	country,
	county,
	city string) error {
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*identity.UserUpdated); ok {
				fmt.Println("Personal Data Changed Handler")
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&identity.UserUpdated{})
		},
	})

	user, err := is.repository.Get(ctx, id)
	if err != nil {
		return err
	}
	err = user.ChangePersonalData(
		ctx,
		profileIMG,
		firstName,
		lastName,
		phone,
		country,
		county,
		city)
	if err != nil {
		return err
	}
	return is.repository.Update(ctx, user)
}

func (is *IdentityService) ChangePassword(
	ctx context.Context,
	email,
	oldPassword,
	newPassword string) error {
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*identity.PasswordChanged); ok {
				fmt.Println("Password Changed Handler")
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&identity.PasswordChanged{})
		},
	})
	user, err := is.repository.Get(ctx, email)
	if err != nil {
		return err
	}
	encryptionService := infrastructure.NewEncryptionService()
	err = encryptionService.Compare(user.Password(), oldPassword)
	if err != nil {
		return err
	}
	encryptedPassword, err := encryptionService.Encrypt(newPassword)
	if err != nil {
		return err
	}
	err = user.ChangePassword(ctx, oldPassword, newPassword, string(encryptedPassword))
	if err != nil {
		return err
	}
	return is.repository.Update(ctx, user)
}

func (is *IdentityService) RecoverPassword(
	ctx context.Context,
	email string) error {
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{

		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*identity.PasswordReset); ok {
				fmt.Println("Password Reset Handler")
				return nil
			}
			return &erros.ValueNotFoundError{}
		},

		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&identity.PasswordReset{})
		},
	})

	user, err := is.repository.Get(ctx, email)
	if err != nil {
		return err
	}
	newRandomGeneratedPassword := strings.Join(strings.Split(uuid.New().String(), "-")[0:2], "-")
	encryptedPassword, err := infrastructure.NewEncryptionService().Encrypt(newRandomGeneratedPassword)
	if err != nil {
		return err
	}
	user.ResetPassword(ctx, newRandomGeneratedPassword, string(encryptedPassword))
	return is.repository.Update(ctx, user)
}

func (is *IdentityService) Login(
	ctx context.Context,
	email string,
	password string,
	rememberMe bool) (string, error) {
	authenticationService := infrastructure.NewAuthenticationService(is.repository)
	jwtService := NewJWTService()
	userDescriptor, err := authenticationService.AuthenticateUser(ctx, email, password, rememberMe)
	if err != nil {
		return "", err
	}
	return jwtService.GenerateJWTToken(userDescriptor)
}

func (is *IdentityService) RegisterUser(
	ctx context.Context,
	profileIMG string,
	email string,
	password string,
	firstName string,
	lastName string,
	birthDate string,
	phone string,
	country string,
	county string,
	city string,
) error {
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*identity.UserCreated); ok {
					fmt.Println("User Registered Handler")
					return nil
				}
				return &erros.ValueNotFoundError{}
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&identity.UserCreated{})
			},
		},
	)
	bdate, err := common.NewDateFromString(birthDate)
	if err != nil {
		return err
	}
	address, err := common.NewAddress(country, county, city)
	if err != nil {
		return err
	}
	person, err := identity.NewPerson(email, firstName, lastName, phone, bdate, address)
	if err != nil {
		return err
	}
	newDate := common.NewDate(time.Now())
	encryptedPassword, err := infrastructure.NewEncryptionService().Encrypt(password)
	if err != nil {
		return err
	}
	encryptedPasswordString := string(encryptedPassword)
	user, err := identity.CreateUser(ctx, profileIMG, newDate, email, encryptedPasswordString, person)
	if err != nil {
		return err
	}
	err = is.repository.Save(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
func (is *IdentityService) Users(ctx context.Context) ([]*identity.User, error) {
	return is.repository.GetAll(ctx)
}

func (is *IdentityService) PossibleEmployees(ctx context.Context, name string, emails []string, limit, offset int) ([]*identity.User, int, error) {
	return is.repository.FindByNotEmailsAndName(ctx, name, emails, limit, offset)
}

func (is *IdentityService) User(ctx context.Context, email string) (*identity.User, error) {
	return is.repository.Get(ctx, email)
}
