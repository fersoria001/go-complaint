package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/employee"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/erros"
	employeefindall "go-complaint/infrastructure/persistence/finders/employee_findall"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type EnterpriseCommand struct {
	OwnerID        string `json:"owner_id"`
	Name           string `json:"name"`
	LogoIMG        string `json:"logo_img"`
	BannerIMG      string `json:"banner_img"`
	Website        string `json:"website"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	CountryID      int    `json:"country_id"`
	CountryStateID int    `json:"country_state_id"`
	CityID         int    `json:"city_id"`
	IndustryID     int    `json:"industry_id"`
	FoundationDate string `json:"foundation_date"`
	UpdateType     string `json:"update_type"`
	ProposeTo      string `json:"propose_to"`
	Position       string `json:"position"`
	EventID        string `json:"event_id"`
	EmployeeID     string `json:"employee_id"`
	TriggeredByID  string `json:"triggered_by_id"`
}

func (enterpriseCommand EnterpriseCommand) Register(
	ctx context.Context,
) error {
	if enterpriseCommand.OwnerID == "" ||
		enterpriseCommand.Name == "" ||
		enterpriseCommand.Website == "" ||
		enterpriseCommand.Email == "" ||
		enterpriseCommand.Phone == "" ||
		enterpriseCommand.FoundationDate == "" {
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
	user, err := userMapper.Get(ctx, enterpriseCommand.OwnerID)
	if err != nil {
		return err
	}
	industry, err := enterprise.NewIndustry(
		enterpriseCommand.IndustryID,
		"",
	)
	if err != nil {
		return err
	}
	country := common.NewCountry(
		enterpriseCommand.CountryID,
		"",
		"",
	)
	countryState := common.NewCountryState(
		enterpriseCommand.CountryStateID,
		"",
	)
	city := common.NewCity(
		enterpriseCommand.CityID,
		"",
		"",
		0,
		0,
	)
	address := common.NewAddress(
		uuid.New(),
		country,
		countryState,
		city,
	)
	foundationDate, err := common.NewDateFromString(
		enterpriseCommand.FoundationDate,
	)
	if err != nil {
		return err
	}
	newEnterprise, err := enterprise.CreateEnterprise(
		ctx,
		user,
		enterpriseCommand.Name,
		enterpriseCommand.Website,
		enterpriseCommand.Email,
		enterpriseCommand.Phone,
		foundationDate,
		industry,
		address,
	)
	if err != nil {
		return err
	}
	mapper = repositories.MapperRegistryInstance().Get("Enterprise")
	if mapper == nil {
		return repositories.ErrMapperNotRegistered
	}
	enterpriseMapper, ok := mapper.(repositories.EnterpriseRepository)
	if !ok {
		return repositories.ErrWrongTypeAssertion
	}
	err = enterpriseMapper.Save(ctx, newEnterprise)
	if err != nil {
		return err
	}
	err = userMapper.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (enterpriseCommand EnterpriseCommand) UpdateEnterprise(
	ctx context.Context,
) error {
	if enterpriseCommand.UpdateType == "" ||
		enterpriseCommand.OwnerID == "" ||
		enterpriseCommand.Name == "" {
		return ErrBadRequest
	}
	mapper := repositories.MapperRegistryInstance().Get("Enterprise")
	if mapper == nil {
		return repositories.ErrMapperNotRegistered
	}
	enterpriseMapper, ok := mapper.(repositories.EnterpriseRepository)
	if !ok {
		return repositories.ErrWrongTypeAssertion
	}
	if enterpriseCommand.LogoIMG == "" {
		return ErrBadRequest
	}
	enterprise, err := enterpriseMapper.Get(ctx, enterpriseCommand.Name)
	if err != nil {
		return err
	}
	if enterprise.Owner() != enterpriseCommand.OwnerID {
		return ErrUnauthorized
	}
	switch enterpriseCommand.UpdateType {
	case "logoIMG":
		if enterpriseCommand.LogoIMG == "" {
			return ErrBadRequest
		}
		err = enterprise.ChangeLogoIMG(ctx, enterpriseCommand.LogoIMG)
	case "bannerIMG":
		if enterpriseCommand.BannerIMG == "" {
			return ErrBadRequest
		}
		err = enterprise.ChangeBannerIMG(ctx, enterpriseCommand.BannerIMG)
	case "website":
		if enterpriseCommand.Website == "" {
			return ErrBadRequest
		}
		err = enterprise.ChangeWebsite(ctx, enterpriseCommand.Website)
	case "email":
		if enterpriseCommand.Email == "" {
			return ErrBadRequest
		}
		err = enterprise.ChangeEmail(ctx, enterpriseCommand.Email)
	case "phone":
		if enterpriseCommand.Phone == "" {
			return ErrBadRequest
		}
		err = enterprise.ChangePhone(ctx, enterpriseCommand.Phone)
	case "country":
		err = enterprise.ChangeCountry(ctx, enterpriseCommand.CountryID)
	case "countryState":
		err = enterprise.ChangeCountryState(ctx, enterpriseCommand.CountryStateID)
	case "city":
		err = enterprise.ChangeCity(ctx, enterpriseCommand.CityID)
	default:
		return ErrBadRequest
	}
	if err != nil {
		return err
	}
	err = enterpriseMapper.Update(ctx, enterprise)
	if err != nil {
		return err
	}
	return nil
}

/**
 * InviteToProject
 * @param ctx context.Context
 * @return error
 Check if the enterpriseCommand.Name,
  enterpriseCommand.Position, and enterpriseCommand.ProposeTo are not empty.
 Check if the enterpriseCommand.Position is a valid position.
 Check if the user is authorized to invite someone to the project.
 Check if the user is already hired.
 Get the target user.
 Invite the target user to the project.
 Subscribe to the HiringInvitationSent event.
 Send an email to the target user.
**/

func (enterpriseCommand EnterpriseCommand) InviteToProject(
	ctx context.Context,
) error {
	if enterpriseCommand.Name == "" ||
		enterpriseCommand.Position == "" ||
		enterpriseCommand.OwnerID == "" ||
		enterpriseCommand.ProposeTo == "" {
		return ErrBadRequest
	}
	if enterprise.ParsePosition(enterpriseCommand.Position) == enterprise.NOT_EXISTS {
		return ErrBadRequest
	}
	enterpriseOwner, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, enterpriseCommand.OwnerID)
	if err != nil {
		return ErrNotFound
	}
	dbemployee, err := repositories.MapperRegistryInstance().Get("Employee").(repositories.EmployeeRepository).FindAll(
		ctx,
		employeefindall.NewByUserIDAndEnterpriseID(
			enterpriseCommand.ProposeTo,
			enterpriseCommand.Name,
		),
	)
	if err != nil {
		return ErrNotFound
	}
	if dbemployee.Cardinality() != 0 {
		return ErrAlreadyHired
	}
	targetUser, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, enterpriseCommand.ProposeTo)
	if err != nil {
		return ErrNotFound
	}
	dbEnterprise, err := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Get(
		ctx, enterpriseCommand.Name)
	if err != nil {
		return err
	}

	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if castedEvent, ok := event.(*enterprise.HiringInvitationSent); ok {
					SendEmailCommand{
						ToEmail: castedEvent.ProposedTo(),
						ToName:  targetUser.FullName(),
					}.HiringInvitationSent(ctx)
					NotificationCommand{
						OwnerID:   castedEvent.ProposedTo(),
						Thumbnail: "",
						Title:     "You have been invited to join a project",
						Content:   enterpriseOwner.FullName() + " has invited you to join the project " + dbEnterprise.Name(),
						Link:      "/hiring-invitations",
					}.SaveNew(ctx)
					return nil
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&enterprise.HiringInvitationSent{})
			},
		},
	)
	err = dbEnterprise.InviteToProject(
		ctx,
		enterpriseCommand.OwnerID,
		enterpriseCommand.ProposeTo,
		enterprise.ParsePosition(enterpriseCommand.Position),
	)
	if err != nil {
		return err
	}
	return nil
}

func (enterpriseCommand EnterpriseCommand) HireEmployee(
	ctx context.Context,
) error {
	if enterpriseCommand.Name == "" ||
		enterpriseCommand.OwnerID == "" ||
		enterpriseCommand.EventID == "" {
		return ErrBadRequest
	}
	dbEnterprise, err := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Get(
		ctx,
		enterpriseCommand.Name,
	)
	if err != nil {
		return ErrNotFound
	}
	if dbEnterprise.Owner() != enterpriseCommand.OwnerID {
		return ErrForbidden
	}
	parseID, err := uuid.Parse(enterpriseCommand.EventID)
	if err != nil {
		return ErrBadRequest
	}
	storedEvent, err := repositories.MapperRegistryInstance().Get(
		"StoredEvent",
	).(repositories.EventRepository).Get(
		ctx,
		parseID,
	)
	if err != nil {
		return ErrNotFound
	}
	if storedEvent.TypeName != "*identity.HiringInvitationAccepted" {
		return ErrBadRequest
	}
	var acceptedInvitation identity.HiringInvitationAccepted
	err = json.Unmarshal(storedEvent.EventBody, &acceptedInvitation)
	if err != nil {
		return err
	}
	invitedUser, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(
		ctx,
		acceptedInvitation.InvitedUserID(),
	)
	if err != nil {
		return err
	}
	position := enterprise.ParsePosition(acceptedInvitation.ProposedPosition().String())
	hiringDate := common.NewDate(time.Now())
	newEmployee, err := employee.NewEmployee(
		uuid.New(),
		acceptedInvitation.EnterpriseID(),
		invitedUser,
		position,
		hiringDate,
		false,
		hiringDate,
	)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*enterprise.EmployeeHired); ok {
				NotificationCommand{
					OwnerID:   invitedUser.Email(),
					Thumbnail: dbEnterprise.LogoIMG(),
					Title:     fmt.Sprintf("You have been hired by %s", dbEnterprise.Name()),
					Content:   fmt.Sprintf("You have been hired by %s as %s", dbEnterprise.Name(), position.String()),
					Link:      fmt.Sprintf("/offices/%s", dbEnterprise.Name()),
				}.SaveNew(ctx)
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&identity.HiringInvitationAccepted{})
		},
	})
	err = dbEnterprise.HireEmployee(ctx, invitedUser, newEmployee)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Update(ctx, dbEnterprise)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Update(ctx, invitedUser)
	if err != nil {
		return err
	}
	return nil
}

func (command EnterpriseCommand) CancelHiringProccess(
	ctx context.Context,
) error {
	if command.Name == "" ||
		command.OwnerID == "" ||
		command.EventID == "" {
		return ErrBadRequest
	}
	parsedID, err := uuid.Parse(command.EventID)
	if err != nil {
		return ErrBadRequest
	}
	storedEvent, err := repositories.MapperRegistryInstance().Get("StoredEvent").(repositories.EventRepository).Get(
		ctx,
		parsedID,
	)
	if err != nil {
		return ErrNotFound
	}
	if storedEvent.TypeName != "*identity.HiringInvitationAccepted" {
		return ErrBadRequest
	}
	var acceptedInvitation identity.HiringInvitationAccepted
	err = json.Unmarshal(storedEvent.EventBody, &acceptedInvitation)
	if err != nil {
		return err
	}
	dbEnterprise, err := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Get(
		ctx,
		command.Name,
	)
	if err != nil {
		return ErrNotFound
	}
	if dbEnterprise.Owner() != command.OwnerID {
		return ErrForbidden
	}
	position := enterprise.ParsePosition(acceptedInvitation.ProposedPosition().String())
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*enterprise.HiringProccessCanceled); ok {
				NotificationCommand{
					OwnerID:   acceptedInvitation.InvitedUserID(),
					Thumbnail: dbEnterprise.LogoIMG(),
					Title:     fmt.Sprintf("Your hiring process at %s has been canceled", dbEnterprise.Name()),
					Content:   fmt.Sprintf("Your hiring process at %s as %s has been canceled", dbEnterprise.Name(), position.String()),
					Link:      "/hiring-invitations",
				}.SaveNew(ctx)
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.HiringProccessCanceled{})
		},
	})
	err = dbEnterprise.CancelHiringProccess(
		ctx,
		acceptedInvitation.InvitedUserID(),
		position,
	)
	if err != nil {
		return err
	}
	return nil
}

func (command EnterpriseCommand) FireEmployee(
	ctx context.Context,
) error {
	if command.Name == "" ||
		command.EmployeeID == "" {
		return ErrBadRequest
	}
	dbEnterprise, err := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Get(
		ctx,
		command.Name,
	)
	if err != nil {
		return ErrNotFound
	}
	employeeID, err := uuid.Parse(command.EmployeeID)
	if err != nil {
		return ErrBadRequest
	}
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if e, ok := event.(*enterprise.EmployeeFired); ok {
				NotificationCommand{
					OwnerID:   e.UserID(),
					Thumbnail: dbEnterprise.LogoIMG(),
					Title:     fmt.Sprintf("Your work has end at %s", dbEnterprise.Name()),
					Content:   fmt.Sprintf("You no longer hold the position %s at %s ", e.Position().String(), dbEnterprise.Name()),
					Link:      "/",
				}.SaveNew(ctx)
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.EmployeeFired{})
		},
	})
	user, err := dbEnterprise.FireEmployee(ctx, employeeID)
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

func (command EnterpriseCommand) PromoteEmployee(
	ctx context.Context,
) error {
	if command.Name == "" ||
		command.EmployeeID == "" ||
		command.Position == "" {
		return ErrBadRequest
	}
	dbEnterprise, err := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Get(
		ctx,
		command.Name,
	)
	if err != nil {
		return ErrNotFound
	}
	employeeID, err := uuid.Parse(command.EmployeeID)
	if err != nil {
		return ErrBadRequest
	}
	position := enterprise.ParsePosition(command.Position)
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if e, ok := event.(*enterprise.EmployeeFired); ok {
				NotificationCommand{
					OwnerID:   e.UserID(),
					Thumbnail: dbEnterprise.LogoIMG(),
					Title:     fmt.Sprintf("Your work has end at %s", dbEnterprise.Name()),
					Content:   fmt.Sprintf("You no longer hold the position %s at %s ", e.Position().String(), dbEnterprise.Name()),
					Link:      "/",
				}.SaveNew(ctx)
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.EmployeeFired{})
		},
	})
	updatedUser, err := dbEnterprise.PromoteEmployee(ctx, command.TriggeredByID, employeeID, position)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Update(ctx, updatedUser)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Update(ctx, dbEnterprise)
	if err != nil {
		return err
	}
	return nil

}
