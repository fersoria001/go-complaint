package commands

import (
	"context"
	"go-complaint/application/application_services"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/identity"
	"go-complaint/domain/model/recipient"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type RegisterUserCommand struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Genre          string `json:"genre"`
	Pronoun        string `json:"pronoun"`
	BirthDate      string `json:"birthDate"`
	Phone          string `json:"phone"`
	ProfileImg     string `json:"profileImg"`
	CountryId      int    `json:"countryId"`
	CountryStateId int    `json:"countryStateId"`
	CityId         int    `json:"cityId"`
}

func NewRegisterUserCommand(
	username,
	password,
	firstName,
	lastName,
	genre,
	pronoun,
	birthDate,
	phone,
	profileImg string,
	countryId,
	countryStateId,
	cityId int,
) *RegisterUserCommand {
	return &RegisterUserCommand{
		Username:       username,
		Password:       password,
		FirstName:      firstName,
		LastName:       lastName,
		Genre:          genre,
		Pronoun:        pronoun,
		BirthDate:      birthDate,
		Phone:          phone,
		ProfileImg:     profileImg,
		CountryId:      countryId,
		CountryStateId: countryStateId,
		CityId:         cityId,
	}
}

func (c RegisterUserCommand) Execute(ctx context.Context) error {
	reg := repositories.MapperRegistryInstance()
	userRepository, ok := reg.Get("User").(repositories.UserRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	newId := uuid.New()
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*identity.UserCreated); ok {
					// SendEmailCommand{
					// 	ToEmail:           userCommand.Email,
					// 	ToName:            userCommand.FirstName + " " + userCommand.LastName,
					// 	ConfirmationToken: castedEvent.ConfirmationToken(),
					// }.Welcome(ctx)
					recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
					if !ok {
						return ErrWrongTypeAssertion
					}
					newRecipient := recipient.NewRecipient(newId, c.FirstName+" "+c.LastName, c.ProfileImg, c.Username, false)
					err := recipientRepository.Save(ctx, *newRecipient)
					if err != nil {
						return err
					}
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&identity.UserCreated{})
			},
		},
	)
	birthDate, err := common.NewDateFromString(c.BirthDate)
	if err != nil {
		return err
	}
	country := common.NewCountry(
		c.CountryId,
		"",
		"",
	)
	countryState := common.NewCountryState(
		c.CountryStateId,
		"",
	)
	city := common.NewCity(
		c.CityId,
		"",
		"",
		0,
		0,
	)
	address := common.NewAddress(
		newId,
		country,
		countryState,
		city,
	)
	newPerson, err := identity.NewPerson(
		newId,
		c.Username,
		c.ProfileImg,
		c.Genre,
		c.Pronoun,
		c.FirstName,
		c.LastName,
		c.Phone,
		birthDate,
		address,
	)
	if err != nil {
		return err
	}
	registerDate := common.NewDate(time.Now())
	hash, err := infrastructure.EncryptionServiceInstance().Encrypt(c.Password)
	if err != nil {
		return err
	}
	emailVerification := identity.NewEmailVerification(c.Username)
	token, err := application_services.JWTApplicationServiceInstance().GenerateJWTToken(emailVerification)
	if err != nil {
		return err
	}
	newUser, err := identity.CreateUser(
		ctx,
		newId,
		c.Username,
		string(hash),
		token.Token(),
		registerDate,
		newPerson,
	)
	if err != nil {
		return err
	}
	cache.InMemoryInstance().Set(token.Token(), false)
	err = userRepository.Save(ctx, newUser)
	if err != nil {
		return err
	}
	return nil
}
