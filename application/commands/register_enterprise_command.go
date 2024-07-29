package commands

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/domain/model/recipient"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"

	"github.com/google/uuid"
)

type RegisterEnterpriseCommand struct {
	OwnerId        string `json:"ownerId"`
	Name           string `json:"name"`
	LogoImg        string `json:"logoImg"`
	BannerImg      string `json:"bannerImg"`
	Website        string `json:"website"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	FoundationDate string `json:"foundationDate"`
	IndustryId     int    `json:"industryId"`
	CountryId      int    `json:"countryId"`
	CountryStateId int    `json:"countryStateId"`
	CityId         int    `json:"cityId"`
}

func NewRegisterEnterpriseCommand(
	ownerId,
	name,
	logoImg,
	bannerImg,
	website,
	email,
	phone,
	foundationDate string,
	industryId,
	countryId,
	countryStateId,
	cityId int,
) *RegisterEnterpriseCommand {
	return &RegisterEnterpriseCommand{
		OwnerId:        ownerId,
		Name:           name,
		LogoImg:        logoImg,
		BannerImg:      bannerImg,
		Website:        website,
		Email:          email,
		Phone:          phone,
		FoundationDate: foundationDate,
		IndustryId:     industryId,
		CountryId:      countryId,
		CountryStateId: countryStateId,
		CityId:         cityId,
	}
}

func (c RegisterEnterpriseCommand) Execute(ctx context.Context) error {
	repository, ok := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	newId := uuid.New()
	industry, err := enterprise.NewIndustry(c.IndustryId, "")
	if err != nil {
		return err
	}
	country := common.NewCountry(c.CountryId, "", "")
	countryState := common.NewCountryState(c.CountryStateId, "")
	city := common.NewCity(c.CityId, "", "", 0, 0)
	address := common.NewAddress(newId, country, countryState, city)
	foundationDate, err := common.NewDateFromString(c.FoundationDate)
	if err != nil {
		return err
	}
	ownerId, err := uuid.Parse(c.OwnerId)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*enterprise.EnterpriseCreated); ok {
					userRepository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
					if !ok {
						return ErrWrongTypeAssertion
					}
					owner, err := userRepository.Get(ctx, ownerId)
					if err != nil {
						return err
					}
					err = owner.AddRole(
						ctx,
						identity.OWNER,
						newId,
					)
					if err != nil {
						return err
					}
					err = userRepository.Update(ctx, owner)
					if err != nil {
						return err
					}
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&enterprise.EnterpriseCreated{})
			},
		},
	)
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*enterprise.EnterpriseCreated); ok {
					recipientRepository, ok := repositories.MapperRegistryInstance().Get("Recipient").(repositories.RecipientRepository)
					if !ok {
						return ErrWrongTypeAssertion
					}
					recipient := recipient.NewRecipient(newId, c.Name, c.LogoImg, c.Email, true)
					err := recipientRepository.Save(ctx, *recipient)
					if err != nil {
						return err
					}
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&enterprise.EnterpriseCreated{})
			},
		},
	)
	newEnterprise, err := enterprise.CreateEnterprise(
		ctx,
		newId,
		ownerId,
		c.Name,
		c.LogoImg,
		c.BannerImg,
		c.Website,
		c.Email,
		c.Phone,
		foundationDate,
		industry,
		address,
	)
	if err != nil {
		return err
	}
	err = repository.Save(ctx, newEnterprise)
	if err != nil {
		return err
	}
	return nil
}
