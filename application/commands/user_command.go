package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"go-complaint/application/application_services"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/erros"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
)

type UserCommand struct {
	Email                  string `json:"email"`
	Password               string `json:"password"`
	ProfileIMG             string `json:"profileIMG"`
	FirstName              string `json:"firstName"`
	LastName               string `json:"lastName"`
	Gender                 string `json:"gender"`
	Pronoun                string `json:"pronoun"`
	BirthDate              string `json:"birthDate"`
	PhoneCode              string `json:"phoneCode"`
	Phone                  string `json:"phone"`
	CountryID              int    `json:"country"`
	CountryStateID         int    `json:"county"`
	CityID                 int    `json:"city"`
	OldPassword            string `json:"oldPassword"`
	FullName               string `json:"fullName"`
	UpdateType             string `json:"updateType"`
	EmailVerificationToken string `json:"emailVerificationToken"`
	EventID                string `json:"eventID"`
	RejectedReason         string `json:"rejected_reason"`
}

func (userCommand UserCommand) Register(ctx context.Context) error {
	if userCommand.Email == "" ||
		userCommand.Password == "" ||
		userCommand.FirstName == "" ||
		userCommand.LastName == "" ||
		userCommand.BirthDate == "" ||
		userCommand.PhoneCode == "" ||
		userCommand.Phone == "" {
		return ErrBadRequest
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if castedEvent, ok := event.(*identity.UserCreated); ok {
					SendEmailCommand{
						ToEmail:           userCommand.Email,
						ToName:            userCommand.FirstName + " " + userCommand.LastName,
						ConfirmationToken: castedEvent.ConfirmationToken(),
					}.Welcome(ctx)
					return nil
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&identity.UserCreated{})
			},
		},
	)
	mapper := repositories.MapperRegistryInstance().Get("User")
	if mapper == nil {
		return repositories.ErrMapperNotRegistered
	}
	userMapper, ok := mapper.(repositories.UserRepository)
	if !ok {
		return repositories.ErrWrongTypeAssertion
	}
	defaultProfileIMG := "/default.jpg"
	commonBirthDate, err := common.NewDateFromString(userCommand.BirthDate)
	if err != nil {
		return err
	}
	country := common.NewCountry(
		userCommand.CountryID,
		"",
		"",
	)
	countryState := common.NewCountryState(
		userCommand.CountryStateID,
		"",
	)
	city := common.NewCity(
		userCommand.CityID,
		"",
		"",
		0,
		0,
	)
	var phoneCode string
	if strings.HasPrefix(userCommand.PhoneCode, "+") {
		phoneCode = userCommand.PhoneCode
	} else {
		phoneCode = "+" + userCommand.PhoneCode
	}
	phone := phoneCode + userCommand.Phone
	address := common.NewAddress(
		uuid.New(),
		country,
		countryState,
		city,
	)
	newPerson, err := identity.NewPerson(
		userCommand.Email,
		defaultProfileIMG,
		userCommand.Gender,
		userCommand.Pronoun,
		userCommand.FirstName,
		userCommand.LastName,
		phone,
		commonBirthDate,
		address,
	)
	if err != nil {
		return err
	}
	userID := userCommand.Email
	registerDate := common.NewDate(time.Now())
	hash, err := infrastructure.EncryptionServiceInstance().Encrypt(userCommand.Password)
	if err != nil {
		return err
	}
	emailVerification := identity.NewEmailVerification(userCommand.Email)
	token, err := application_services.JWTApplicationServiceInstance().GenerateJWTToken(emailVerification)
	if err != nil {
		return err
	}
	newUser, err := identity.CreateUser(
		ctx,
		userID,
		string(hash),
		token.Token(),
		registerDate,
		newPerson,
	)
	if err != nil {
		return err
	}
	cache.InMemoryCacheInstance().Set(token.Token(), false)
	userCommand.FullName = newUser.FullName()
	err = userMapper.Save(ctx, newUser)
	if err != nil {
		return err
	}
	return nil
}

func (userCommand UserCommand) VerifyEmail(ctx context.Context) error {
	if userCommand.EmailVerificationToken == "" {
		return ErrBadRequest
	}
	token, ok := cache.InMemoryCacheInstance().Get(
		userCommand.EmailVerificationToken)
	if !ok {
		return ErrConfirmationNotFound
	}
	confirmation, ok := token.(bool)
	if !ok {
		return ErrWrongTypeAssertion
	}
	if confirmation {
		return ErrAlreadyVerified
	}
	emailVerification, err := application_services.JWTApplicationServiceInstance().ParseEmailVerification(userCommand.EmailVerificationToken)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*identity.UserEmailVerified); ok {
				SendEmailCommand{
					ToEmail: emailVerification.Email,
					ToName:  userCommand.FullName,
				}.EmailVerified(ctx)
			}
			return nil
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&identity.UserEmailVerified{})
		},
	})

	mapper := repositories.MapperRegistryInstance().Get("User")
	if mapper == nil {
		return repositories.ErrMapperNotRegistered
	}
	userMapper, ok := mapper.(repositories.UserRepository)
	if !ok {
		return repositories.ErrWrongTypeAssertion
	}
	user, err := userMapper.Get(ctx, emailVerification.Email)
	if err != nil {
		return err
	}
	err = user.VerifyEmail(ctx)
	if err != nil {
		return err
	}
	err = userMapper.Update(ctx, user)
	if err != nil {
		return err
	}
	cache.InMemoryCacheInstance().Delete(userCommand.EmailVerificationToken)
	return nil
}

func (userCommand UserCommand) RecoverPassword(ctx context.Context) error {
	if userCommand.Email == "" {
		return ErrBadRequest
	}
	newRandomGeneratedPassword := strings.Join(strings.Split(uuid.New().String(), "-")[0:2], "-")
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*identity.PasswordReset); ok {
				SendEmailCommand{
					ToEmail:        userCommand.Email,
					RandomPassword: newRandomGeneratedPassword,
				}.PasswordRecovery(ctx)
			}
			return nil
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&identity.PasswordReset{})
		},
	})
	mapper := repositories.MapperRegistryInstance().Get("User")
	if mapper == nil {
		return repositories.ErrMapperNotRegistered
	}
	userMapper, ok := mapper.(repositories.UserRepository)
	if !ok {
		return repositories.ErrWrongTypeAssertion
	}
	user, err := userMapper.Get(ctx, userCommand.Email)
	if err != nil {
		return err
	}
	encryptedPassword, err := infrastructure.EncryptionServiceInstance().Encrypt(newRandomGeneratedPassword)
	if err != nil {
		return err
	}
	err = user.ResetPassword(ctx, newRandomGeneratedPassword, string(encryptedPassword))
	if err != nil {
		return err
	}
	userCommand.FullName = user.FullName()
	err = userMapper.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (userCommand UserCommand) ChangePassword(
	ctx context.Context,
) error {
	if userCommand.Email == "" ||
		userCommand.OldPassword == "" ||
		userCommand.Password == "" {
		return ErrBadRequest
	}
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*identity.PasswordChanged); ok {
				SendEmailCommand{
					ToEmail: userCommand.Email,
					ToName:  userCommand.FullName,
				}.PasswordChanged(ctx)

				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&identity.PasswordChanged{})
		},
	})
	mapper := repositories.MapperRegistryInstance().Get("User")
	if mapper == nil {
		return repositories.ErrMapperNotRegistered
	}
	userMapper, ok := mapper.(repositories.UserRepository)
	if !ok {
		return repositories.ErrWrongTypeAssertion
	}
	user, err := userMapper.Get(ctx, userCommand.Email)
	if err != nil {
		return err
	}
	err = infrastructure.EncryptionServiceInstance().Compare(
		user.Password(), userCommand.OldPassword)
	if err != nil {
		return err
	}
	encryptedPassword, err := infrastructure.EncryptionServiceInstance().Encrypt(
		userCommand.Password)
	if err != nil {
		return err
	}
	err = user.ChangePassword(ctx, userCommand.OldPassword,
		userCommand.Password, string(encryptedPassword))
	if err != nil {
		return err
	}
	userCommand.FullName = user.FullName()
	err = userMapper.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (userCommand UserCommand) UpdatePersonalData(
	ctx context.Context,
) error {
	if userCommand.Email == "" || userCommand.UpdateType == "" {
		return ErrBadRequest
	}
	mapper := repositories.MapperRegistryInstance().Get("User")
	if mapper == nil {
		return repositories.ErrMapperNotRegistered
	}
	userMapper, ok := mapper.(repositories.UserRepository)
	if !ok {
		return repositories.ErrWrongTypeAssertion
	}
	user, err := userMapper.Get(ctx, userCommand.Email)
	if err != nil {
		return err
	}
	switch userCommand.UpdateType {
	case "pronoun":
		if userCommand.Pronoun == "" {
			return ErrBadRequest
		}
		err = user.ChangePronoun(ctx, userCommand.Pronoun)
	case "gender":
		if userCommand.Gender == "" {
			return ErrBadRequest
		}
		err = user.ChangeGender(ctx, userCommand.Gender)
	case "firstName":
		if userCommand.FirstName == "" {
			return ErrBadRequest
		}
		err = user.ChangeFirstName(ctx, userCommand.FirstName)
	case "lastName":
		if userCommand.LastName == "" {
			return ErrBadRequest
		}
		err = user.ChangeLastName(ctx, userCommand.LastName)
	case "profileIMG":
		if userCommand.ProfileIMG == "" {
			return ErrBadRequest
		}
		err = user.ChangeProfileIMG(ctx, userCommand.ProfileIMG)
	case "phone":
		if userCommand.Phone == "" {
			return ErrBadRequest
		}
		err = user.ChangePhone(ctx, userCommand.Phone)
	case "country":
		err = user.ChangeCountry(ctx, userCommand.CountryID)
	case "countryState":
		err = user.ChangeCountryState(ctx, userCommand.CountryStateID)
	case "city":
		err = user.ChangeCity(ctx, userCommand.CityID)
	default:
		return ErrBadRequest
	}
	if err != nil {
		return err
	}
	err = userMapper.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (userCommand UserCommand) AcceptHiringInvitation(
	ctx context.Context,
) error {
	if userCommand.Email == "" || userCommand.EventID == "" {
		return ErrBadRequest
	}
	user, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, userCommand.Email)
	if err != nil {
		return err
	}
	parsedID, err := uuid.Parse(userCommand.EventID)
	if err != nil {
		return err
	}
	storedEvent, err := repositories.MapperRegistryInstance().Get("Event").(repositories.EventRepository).Get(
		ctx,
		parsedID,
	)
	if err != nil {
		return err
	}
	var invitation enterprise.HiringInvitationSent
	err = json.Unmarshal(storedEvent.EventBody, &invitation)
	if err != nil {
		return err
	}
	role, err := identity.ParseRole(invitation.ProposalPosition().String())
	if err != nil {
		return err
	}
	ep, err := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Get(ctx, invitation.EnterpriseID())
	if err != nil {
		return err
	}
	param := fmt.Sprintf("id=%s", storedEvent.EventId.String())
	replaced := application_services.EncodingApplicationServiceInstance().SafeUtf16Encode(param)
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if event, ok := event.(*identity.HiringInvitationAccepted); ok {
				NotificationCommand{
					OwnerID:     invitation.EnterpriseID(),
					ThumbnailID: user.Email(),
					Thumbnail:   ep.LogoIMG(),
					Title:       fmt.Sprintf(`%s accepted your invitation"`, user.FullName()),
					Content:     fmt.Sprintf("User %s accepted your invitation to join %s", user.FullName(), ep.Name()),
					Link:        fmt.Sprintf("/%s/hiring-procceses?%s", event.EnterpriseID(), replaced),
				}.SaveNew(ctx)
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&identity.HiringInvitationAccepted{})
		},
	})
	err = user.AcceptHiringInvitation(
		ctx,
		invitation.EnterpriseID(),
		role,
	)
	if err != nil {
		return err
	}

	return nil
}

func (command UserCommand) RejectHiringInvitation(
	ctx context.Context,
) error {
	if command.Email == "" || command.EventID == "" {
		return ErrBadRequest
	}
	user, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, command.Email)
	if err != nil {
		return err
	}
	parsedID, err := uuid.Parse(command.EventID)
	if err != nil {
		return err
	}
	storedEvent, err := repositories.MapperRegistryInstance().Get("Event").(repositories.EventRepository).Get(
		ctx,
		parsedID,
	)
	if err != nil {
		return err
	}
	var invitation enterprise.HiringInvitationSent
	err = json.Unmarshal(storedEvent.EventBody, &invitation)
	if err != nil {
		return err
	}
	role, err := identity.ParseRole(invitation.ProposalPosition().String())
	if err != nil {
		return err
	}
	ep, err := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Get(
		ctx, invitation.EnterpriseID())
	if err != nil {
		return err
	}
	param := fmt.Sprintf("id=%s", storedEvent.EventId.String())
	replaced := application_services.EncodingApplicationServiceInstance().SafeUtf16Encode(param)
	if command.RejectedReason == "" {
		command.RejectedReason = "No reason provided"
	}
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if event, ok := event.(*identity.HiringInvitationRejected); ok {
				NotificationCommand{
					OwnerID:   invitation.EnterpriseID(),
					Thumbnail: user.Email(),
					Title:     fmt.Sprintf(`%s rejected your invitation"`, user.FullName()),
					Content:   fmt.Sprintf("User %s rejected your invitation to join %s : %s", user.FullName(), ep.Name(), command.RejectedReason),
					Link:      fmt.Sprintf("/%s/hiring-procceses?%s", event.EnterpriseID(), replaced),
				}.SaveNew(ctx)
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&identity.HiringInvitationRejected{})
		},
	})
	err = user.RejectHiringInvitation(
		ctx,
		invitation.EnterpriseID(),
		command.RejectedReason,
		role,
	)
	if err != nil {
		return err
	}

	return nil
}
