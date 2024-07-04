package queries

import (
	"context"
	"encoding/json"
	"go-complaint/application"
	"log"
	"slices"

	"go-complaint/application/application_services"
	"go-complaint/application/commands"
	"go-complaint/domain"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/dto"
	"go-complaint/erros"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/finders/find_all_events"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"
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
				log.Println("userSignedIn", user.Email(), user.FullName(), userSignedIn.ConfirmationCode())
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
	cache.InMemoryCacheInstance().Set(token.Token(), confirmation)
	return dto.JWTToken{
		Token: token.Token(),
	}, nil
}

func (userQuery UserQuery) UserDescriptor(
	ctx context.Context,
) (dto.UserDescriptor, error) {
	if userQuery.Email == "" {
		return dto.UserDescriptor{}, &erros.ValueNotFoundError{}
	}
	clientData := application_services.AuthorizationApplicationServiceInstance().ClientData(ctx)
	user, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, userQuery.Email)
	if err != nil {
		return dto.UserDescriptor{}, err
	}
	return dto.NewUserDescriptor(clientData, *user, false), nil
}

func (userQuery UserQuery) Login(
	ctx context.Context,
) (dto.JWTToken, error) {
	if userQuery.Token == "" {
		return dto.JWTToken{}, &erros.ValueNotFoundError{}
	}
	confirmation, ok := cache.InMemoryCacheInstance().Get(userQuery.Token)
	if !ok {
		return dto.JWTToken{}, ErrConfirmationNotFound
	}
	loginConfirmation, ok := confirmation.(*application.LoginConfirmation)
	if !ok {
		return dto.JWTToken{}, ErrWrongTypeAssertion
	}
	if loginConfirmation.IsConfirmed() {
		cache.InMemoryCacheInstance().Delete(userQuery.Token)
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
	_, swap := cache.InMemoryCacheInstance().Swap(userQuery.Token, loginConfirmation)
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
	storedEvents, err := repositories.MapperRegistryInstance().Get(
		"Event",
	).(repositories.EventRepository).FindAll(
		ctx,
		find_all_events.By(),
	)
	if err != nil {
		return []dto.HiringInvitation{}, err
	}
	myInvitations := map[string]map[string]*dto.HiringInvitation{}
	for storedEvent := range storedEvents.Iter() {
		if storedEvent.TypeName == "*enterprise.HiringInvitationSent" {
			var e enterprise.HiringInvitationSent
			err := json.Unmarshal(storedEvent.EventBody, &e)
			if err != nil {
				return []dto.HiringInvitation{}, err
			}
			if e.ProposedTo() == user.Email() {
				_, ok := myInvitations[e.ProposedTo()]
				if !ok {
					myInvitations[e.ProposedTo()] = make(map[string]*dto.HiringInvitation)
				}
				myInvitations[e.ProposedTo()][storedEvent.EventId.String()] = dto.NewHiringInvitation(storedEvent.EventId.String(), e)
				myInvitations[e.ProposedTo()][storedEvent.EventId.String()].SetStatus(dto.PENDING.String())
			}
		}
	}
	for e := range storedEvents.Iter() {
		switch e.TypeName {
		case "*identity.HiringInvitationAccepted":
			var event identity.HiringInvitationAccepted
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return []dto.HiringInvitation{}, err
			}
			inv, ok := myInvitations[event.InvitedUserID()]
			if ok {
				for k := range inv {
					if event.OccurredOn().After(inv[k].GetOcurredOn()) {
						inv[k].SetStatus(dto.ACCEPTED.String())
					}
				}
			}
		case "*identity.HiringInvitationRejected":
			var event identity.HiringInvitationRejected
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return []dto.HiringInvitation{}, err
			}
			inv, ok := myInvitations[event.InvitedUserID()]
			if ok {
				for k := range inv {
					if event.OccurredOn().After(inv[k].GetOcurredOn()) {
						inv[k].SetStatus(dto.REJECTED.String())
					}
				}
			}
		case "*enterprise.HiringProccessCanceled":
			var event enterprise.HiringProccessCanceled
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return []dto.HiringInvitation{}, err
			}
			inv, ok := myInvitations[event.CandidateID()]
			if ok {
				for k := range inv {
					if event.OccurredOn().After(inv[k].GetOcurredOn()) {
						inv[k].SetStatus(dto.CANCELED.String())
						inv[k].SetReason(event.Reason())
					}
				}
			}
		}
	}

	slice := []*dto.HiringInvitation{}
	for _, v := range myInvitations {
		for _, c := range v {
			slice = append(slice, c)
		}
	}
	slices.SortStableFunc(slice, func(i, j *dto.HiringInvitation) int {
		if i.GetOcurredOn().Before(j.GetOcurredOn()) {
			return 1
		}
		if i.GetOcurredOn().After(j.GetOcurredOn()) {
			return -1
		}
		return 0
	})
	slice = slices.CompactFunc(slice, func(i, j *dto.HiringInvitation) bool {
		return i.OwnerID == j.OwnerID
	})
	var hiringInvitationsDTO []dto.HiringInvitation
	for _, c := range slice {
		user, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, c.OwnerID)
		if err != nil {
			return []dto.HiringInvitation{}, err
		}
		ep, err := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Get(ctx, c.EnterpriseID)
		if err != nil {
			return []dto.HiringInvitation{}, err
		}
		c.SetUser(*user)
		c.SetEnterprise(*ep)
		hiringInvitationsDTO = append(hiringInvitationsDTO, *c)
	}

	return hiringInvitationsDTO, nil
}
