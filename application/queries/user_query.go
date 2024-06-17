package queries

import (
	"context"
	"encoding/json"
	"go-complaint/application"
	"go-complaint/application/application_services"
	"go-complaint/application/commands"
	"go-complaint/domain"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/dto"
	"go-complaint/erros"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/persistence/finders/find_all_events"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"
	"time"
)

type UserQuery struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	RememberMe       bool   `json:"rememberMe"`
	Token            string `json:"token"`
	ConfirmationCode int    `json:"confirmation_code"`
	EventID          string `json:"event_id"`
}

func (userQuery UserQuery) SignIn(
	ctx context.Context,
) (dto.JWTToken, error) {
	if userQuery.Email == "" {
		return dto.JWTToken{}, &erros.ValueNotFoundError{}
	}
	if userQuery.Password == "" {
		return dto.JWTToken{}, &erros.ValueNotFoundError{}
	}
	ok, err := infrastructure.AuthenticationServiceInstance().AuthenticateUser(
		ctx,
		userQuery.Email,
		userQuery.Password,
		userQuery.RememberMe,
	)
	if err != nil {
		return dto.JWTToken{}, err
	}
	if !ok {
		return dto.JWTToken{}, ErrAuthenticationFailed
	}
	mapper := repositories.MapperRegistryInstance().Get("User")
	userRepository, ok := mapper.(repositories.UserRepository)
	if !ok {
		return dto.JWTToken{}, repositories.ErrWrongTypeAssertion
	}
	user, err := userRepository.Get(ctx, userQuery.Email)
	if err != nil {
		return dto.JWTToken{}, err
	}
	code := application.CreateConfirmationCode()

	clientData := application_services.AuthorizationApplicationServiceInstance().ClientData(ctx)
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if userSignedIn, ok := event.(*identity.UserSignedIn); ok {
				commands.SendEmailCommand{
					ToEmail:          user.Email(),
					ToName:           user.FullName(),
					ConfirmationCode: userSignedIn.ConfirmationCode(),
				}.VerifySignIn(ctx)
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&identity.UserSignedIn{})
		},
	})
	user.SignIn(
		ctx,
		code.Code,
		clientData.IP,
		clientData.Geolocalization.Latitude,
		clientData.Geolocalization.Longitude,
		clientData.Device,
	)
	token, err := application_services.JWTApplicationServiceInstance().GenerateJWTToken(code)
	if err != nil {
		return dto.JWTToken{}, err
	}
	confirmation := application.NewLoginConfirmation(userQuery.Email, token, false)
	infrastructure.InMemoryCacheInstance().Set(token.Token(), confirmation)
	return dto.JWTToken{
		Token: token.Token(),
	}, nil
}

func (userQuery UserQuery) Login(
	ctx context.Context,
) (dto.JWTToken, error) {
	if userQuery.Token == "" {
		return dto.JWTToken{}, &erros.ValueNotFoundError{}
	}
	confirmation, ok := infrastructure.InMemoryCacheInstance().Get(userQuery.Token)
	if !ok {
		return dto.JWTToken{}, ErrConfirmationNotFound
	}
	loginConfirmation, ok := confirmation.(*application.LoginConfirmation)
	if !ok {
		return dto.JWTToken{}, ErrWrongTypeAssertion
	}
	if loginConfirmation.IsConfirmed() {
		infrastructure.InMemoryCacheInstance().Delete(userQuery.Token)
		return dto.JWTToken{}, ErrConfirmationAlreadyDone
	}
	code, err := application_services.JWTApplicationServiceInstance().ParseConfirmationCode(
		userQuery.Token,
	)
	if err != nil {
		return dto.JWTToken{}, err
	}
	if code.Code != userQuery.ConfirmationCode {
		err := loginConfirmation.RetryConfirmation()
		if err != nil {
			return dto.JWTToken{}, err
		}
		return dto.JWTToken{}, ErrConfirmationCodeNotMatch
	}
	loginConfirmation.Confirm()
	_, swap := infrastructure.InMemoryCacheInstance().Swap(userQuery.Token, loginConfirmation)
	if !swap {
		return dto.JWTToken{}, &erros.ValueNotFoundError{}
	}
	mapper := repositories.MapperRegistryInstance().Get("User")
	userRepository, ok := mapper.(repositories.UserRepository)
	if !ok {
		return dto.JWTToken{}, repositories.ErrWrongTypeAssertion
	}
	user, err := userRepository.Get(ctx, loginConfirmation.Email())
	if err != nil {
		return dto.JWTToken{}, err
	}
	clientData := application_services.AuthorizationApplicationServiceInstance().ClientData(ctx)
	userDescriptor := dto.NewUserDescriptor(
		clientData,
		*user,
		userQuery.RememberMe,
	)
	token, err := application_services.JWTApplicationServiceInstance().GenerateJWTToken(
		userDescriptor,
	)
	if err != nil {
		return dto.JWTToken{}, err
	}
	return dto.JWTToken{
		Token: token.Token(),
	}, nil
}

func (userQuery UserQuery) User(
	ctx context.Context,
) (dto.User, error) {
	if userQuery.Email == "" {
		return dto.User{}, ErrNilValue
	}
	mapper := repositories.MapperRegistryInstance().Get("User")
	userRepository, ok := mapper.(repositories.UserRepository)
	if !ok {
		return dto.User{}, repositories.ErrWrongTypeAssertion
	}
	user, err := userRepository.Get(ctx, userQuery.Email)
	if err != nil {
		return dto.User{}, err
	}
	if !user.IsConfirmed() {
		return dto.User{}, ErrUserNotConfirmed
	}
	return dto.NewUser(*user), nil
}

func (userQuery UserQuery) HiringInvitations(
	ctx context.Context,
) ([]dto.HiringInvitation, error) {
	if userQuery.Email == "" {
		return []dto.HiringInvitation{}, ErrBadRequest
	}
	user, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, userQuery.Email)
	if err != nil {
		return []dto.HiringInvitation{}, err
	}
	hiringInvitations, err := repositories.MapperRegistryInstance().Get(
		"Event",
	).(repositories.EventRepository).FindAll(
		ctx,
		find_all_events.ByTypeName("*enterprise.HiringInvitationSent"),
	)
	if err != nil {
		return []dto.HiringInvitation{}, err
	}
	invitationsAccepted, err := repositories.MapperRegistryInstance().Get(
		"Event",
	).(repositories.EventRepository).FindAll(
		ctx,
		find_all_events.ByTypeName("*identity.HiringInvitationAccepted"),
	)
	if err != nil {
		return []dto.HiringInvitation{}, err
	}
	myInvitations := make(map[string]enterprise.HiringInvitationSent, 0)
	alreadyAcceptedInvitations := make(map[string]identity.HiringInvitationAccepted, 0)
	thisDate := time.Now()
	for invitation := range hiringInvitations.Iter() {
		var hiringInvitation enterprise.HiringInvitationSent
		err := json.Unmarshal(invitation.EventBody, &hiringInvitation)
		if err != nil {
			return []dto.HiringInvitation{}, err
		}
		expired := thisDate.Sub(hiringInvitation.OccurredOn()) > 24*5*time.Hour
		if hiringInvitation.ProposedTo() == userQuery.Email && !expired {
			myInvitations[invitation.EventId.String()] = hiringInvitation
		}
	}
	for invitation := range invitationsAccepted.Iter() {
		var hiringInvitation identity.HiringInvitationAccepted
		err := json.Unmarshal(invitation.EventBody, &hiringInvitation)
		if err != nil {
			return []dto.HiringInvitation{}, err
		}
		if hiringInvitation.InvitedUserID() == userQuery.Email {
			alreadyAcceptedInvitations[invitation.EventId.String()] = hiringInvitation
		}
	}

	if len(alreadyAcceptedInvitations) > 0 {
		for _, accepted := range alreadyAcceptedInvitations {
			for j, invitation := range myInvitations {
				if invitation.EnterpriseID() == accepted.EnterpriseID() ||
					invitation.ProposalPosition().String() == accepted.ProposedPosition().String() {
					delete(myInvitations, j)
				}
			}
		}
	}
	hiringInvitationsDTO := make([]dto.HiringInvitation, 0)
	for eventId, invitation := range myInvitations {
		enterprise, err := repositories.MapperRegistryInstance().Get(
			"Enterprise").(repositories.EnterpriseRepository).Get(
			ctx,
			invitation.EnterpriseID(),
		)
		if err != nil {
			return []dto.HiringInvitation{}, err
		}
		hiringInvitationsDTO = append(hiringInvitationsDTO, dto.NewHiringInvitation(
			eventId,
			false,
			*user,
			*enterprise,
			invitation,
		))
	}
	return hiringInvitationsDTO, nil
}
