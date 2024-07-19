package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-complaint/application/application_services"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/employee"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/dto"
	"go-complaint/erros"
	"go-complaint/infrastructure/cache"
	employeefindall "go-complaint/infrastructure/persistence/finders/employee_findall"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"
	"slices"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type EnterpriseCommand struct {
	ID             string   `json:"id"`
	OwnerID        string   `json:"owner_id"`
	Name           string   `json:"name"`
	LogoIMG        string   `json:"logo_img"`
	BannerIMG      string   `json:"banner_img"`
	Website        string   `json:"website"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	CountryID      int      `json:"country_id"`
	CountryStateID int      `json:"country_state_id"`
	CityID         int      `json:"city_id"`
	IndustryID     int      `json:"industry_id"`
	FoundationDate string   `json:"foundation_date"`
	UpdateType     string   `json:"update_type"`
	ProposeTo      string   `json:"propose_to"`
	Position       string   `json:"position"`
	EventID        string   `json:"event_id"`
	EmployeeID     string   `json:"employee_id"`
	TriggeredByID  string   `json:"triggered_by_id"`
	CancelReason   string   `json:"cancel_reason"`
	SenderID       string   `json:"sender_id"`
	Content        string   `json:"content"`
	RepliesID      []string `json:"replies_id"`
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
	phone := enterpriseCommand.Phone
	newEnterprise, err := enterprise.CreateEnterprise(
		ctx,
		user,
		enterpriseCommand.Name,
		enterpriseCommand.Website,
		enterpriseCommand.Email,
		phone,
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
	enterprise, err := enterpriseMapper.Get(ctx, enterpriseCommand.Name)
	if err != nil {
		return err
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
	enterpriseOwner, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(
		ctx, enterpriseCommand.OwnerID)
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
					NotificationCommand{
						OwnerID:     castedEvent.ProposedTo(),
						ThumbnailID: dbEnterprise.Name(),
						Thumbnail:   dbEnterprise.LogoIMG(),
						Title:       "You have been invited to join a project",
						Content:     enterpriseOwner.FullName() + " has invited you to join the project " + dbEnterprise.Name(),
						Link:        "/hiring-invitations",
					}.SaveNew(ctx)
					SendEmailCommand{
						ToEmail: castedEvent.ProposedTo(),
						ToName:  targetUser.FullName(),
					}.HiringInvitationSent(ctx)
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
	parseID, err := uuid.Parse(enterpriseCommand.EventID)
	if err != nil {
		return ErrBadRequest
	}
	storedEvent, err := repositories.MapperRegistryInstance().Get(
		"Event",
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
					OwnerID:     invitedUser.Email(),
					ThumbnailID: dbEnterprise.Name(),
					Thumbnail:   dbEnterprise.LogoIMG(),
					Title:       fmt.Sprintf("You have been hired by %s", dbEnterprise.Name()),
					Content:     fmt.Sprintf("You have been hired by %s as %s", dbEnterprise.Name(), position.String()),
					Link:        fmt.Sprintf("/%s", dbEnterprise.Name()),
				}.SaveNew(ctx)

				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.EmployeeHired{})
		},
	})
	err = dbEnterprise.HireEmployee(ctx, enterpriseCommand.OwnerID, newEmployee)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Update(ctx, dbEnterprise)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Update(ctx, newEmployee.GetUser())
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
	storedEvent, err := repositories.MapperRegistryInstance().Get("Event").(repositories.EventRepository).Get(
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
	employee, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(
		ctx,
		command.OwnerID,
	)
	if err != nil {
		return ErrNotFound
	}
	param := fmt.Sprintf("id=%s", command.EventID)
	replaced := application_services.EncodingApplicationServiceInstance().SafeUtf16Encode(param)
	position := enterprise.ParsePosition(acceptedInvitation.ProposedPosition().String())
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*enterprise.HiringProccessCanceled); ok {
				NotificationCommand{
					OwnerID:   acceptedInvitation.InvitedUserID(),
					Thumbnail: dbEnterprise.LogoIMG(),
					Title:     fmt.Sprintf("Your hiring process at %s has been canceled", dbEnterprise.Name()),
					Content: fmt.Sprintf("Your hiring process at %s as %s has been canceled by %s",
						dbEnterprise.Name(), position.String(), employee.FullName()),
					Link: fmt.Sprintf("/hiring-invitations?%s", replaced),
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
		employee.Email(),
		command.CancelReason,
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
		command.EmployeeID == "" || command.OwnerID == "" {
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
					OwnerID:     e.UserID(),
					ThumbnailID: dbEnterprise.Name(),
					Thumbnail:   dbEnterprise.LogoIMG(),
					Title:       fmt.Sprintf("Your work has end at %s", dbEnterprise.Name()),
					Content:     fmt.Sprintf("You no longer hold the position %s at %s ", e.Position().String(), dbEnterprise.Name()),
					Link:        "/",
				}.SaveNew(ctx)
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.EmployeeFired{})
		},
	})
	user, err := dbEnterprise.FireEmployee(ctx, command.OwnerID, employeeID)
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
	slice := dbEnterprise.Employees().ToSlice()
	index, ok := slices.BinarySearchFunc(slice, employeeID, func(i enterprise.Employee, j uuid.UUID) int {
		if i.ID() == j {
			return 0
		}
		return -1
	})
	if !ok {
		return ErrNotFound
	}
	user := slice[index].GetUser()
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if e, ok := event.(*enterprise.EmployeePromoted); ok {
				NotificationCommand{
					OwnerID:     user.Email(),
					Thumbnail:   dbEnterprise.LogoIMG(),
					ThumbnailID: dbEnterprise.Name(),
					Title:       fmt.Sprintf("You have been promoted in %s", dbEnterprise.Name()),
					Content:     fmt.Sprintf("You have been promoted to %s at %s", e.Position().String(), dbEnterprise.Name()),
					Link:        "/",
				}.SaveNew(ctx)
				return nil
			}
			return nil
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.EmployeePromoted{})
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

func (command EnterpriseCommand) ReplyChat(
	ctx context.Context,
) error {
	if command.ID == "" ||
		command.SenderID == "" ||
		command.Content == "" {
		return ErrBadRequest
	}
	sender, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(
		ctx,
		command.SenderID,
	)
	if err != nil {
		return ErrNotFound
	}
	chatID, err := enterprise.NewChatID(command.ID)
	if err != nil {
		return err
	}
	dbChat, err := repositories.MapperRegistryInstance().Get("Chat").(repositories.ChatRepository).Get(
		ctx,
		*chatID,
	)
	var reply enterprise.Reply
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			dbChat = enterprise.NewChatEntity(
				*chatID,
			)
			reply = dbChat.Reply(uuid.New(), *sender, command.Content)
			err = repositories.MapperRegistryInstance().Get("Chat").(repositories.ChatRepository).Save(ctx, dbChat)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		reply = dbChat.Reply(uuid.New(), *sender, command.Content)
		err = repositories.MapperRegistryInstance().Get("Chat").(repositories.ChatRepository).Update(ctx, dbChat)
		if err != nil {
			return err
		}
	}
	cache.RequestChannel <- cache.Request{
		Type:    cache.WRITE,
		Key:     command.ID,
		Payload: dto.NewChatReply(reply),
	}
	return nil
}
func (command EnterpriseCommand) MarkAsSeen(ctx context.Context) error {
	if command.RepliesID == nil || len(command.RepliesID) == 0 || command.ID == "" {
		return ErrBadRequest
	}
	parseIds := make([]uuid.UUID, len(command.RepliesID))
	for _, replyID := range command.RepliesID {
		parsedReplyID, err := uuid.Parse(replyID)
		if err != nil {
			return ErrBadRequest
		}
		parseIds = append(parseIds, parsedReplyID)
	}
	chatID, err := enterprise.NewChatID(command.ID)
	if err != nil {
		return err
	}
	parsedIds := mapset.NewSet(parseIds...)
	r := repositories.MapperRegistryInstance().Get("Chat").(repositories.ChatRepository)
	dbChat, err := r.Get(ctx, *chatID)
	if err != nil {
		return err
	}
	dbChat.MarkAsSeen(parsedIds)
	err = r.Update(ctx, dbChat)
	if err != nil {
		return err
	}
	replies := dbChat.Replies()
	dtos := make([]*dto.ChatReply, 0, len(replies))
	for _, reply := range replies {
		dtos = append(dtos, dto.NewChatReply(*reply))
	}
	cache.RequestChannel <- cache.Request{
		Key:     fmt.Sprintf("complaintLastReply:%s", chatID.String()),
		Payload: dtos,
	}
	return nil
}
