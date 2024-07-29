package queries

import (
	"context"

	"go-complaint/application/application_services"
	"go-complaint/dto"
	"go-complaint/erros"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
)

type UserQuery struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	RememberMe       bool   `json:"rememberMe"`
	Token            string `json:"token"`
	ConfirmationCode int    `json:"confirmation_code"`
	EventID          string `json:"event_id"`
}

// avoid using this
func (userQuery UserQuery) UserDescriptor(
	ctx context.Context,
) (*dto.UserDescriptor, error) {
	if userQuery.Email == "" {
		return nil, &erros.ValueNotFoundError{}
	}
	clientData := application_services.AuthorizationApplicationServiceInstance().ClientData(ctx)
	user, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Find(ctx, find_user.ByUsername(userQuery.Email))
	if err != nil {
		return nil, err
	}
	return dto.NewUserDescriptor(clientData, *user), nil
}

// func (userQuery UserQuery) User(
// 	ctx context.Context,
// ) (dto.User, error) {
// 	if userQuery.Email == "" {
// 		return dto.User{}, ErrNilValue
// 	}
// 	mapper := repositories.MapperRegistryInstance().Get("User")
// 	userRepository, ok := mapper.(repositories.UserRepository)
// 	if !ok {
// 		return dto.User{}, repositories.ErrWrongTypeAssertion
// 	}
// 	user, err := userRepository.Find(ctx, find_user.ByUsername(userQuery.Email))
// 	if err != nil {
// 		return dto.User{}, err
// 	}
// 	if !user.IsConfirmed() {
// 		return dto.User{}, ErrUserNotConfirmed
// 	}
// 	return dto.NewUser(*user), nil
// }

// func (userQuery UserQuery) HiringInvitations(
// 	ctx context.Context,
// ) ([]dto.HiringInvitation, error) {
// 	if userQuery.Email == "" {
// 		return []dto.HiringInvitation{}, ErrBadRequest
// 	}
// 	user, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Find(ctx, find_user.ByUsername(userQuery.Email))
// 	if err != nil {
// 		return []dto.HiringInvitation{}, err
// 	}
// 	storedEvents, err := repositories.MapperRegistryInstance().Get(
// 		"Event",
// 	).(repositories.EventRepository).FindAll(
// 		ctx,
// 		find_all_events.By(),
// 	)
// 	if err != nil {
// 		return []dto.HiringInvitation{}, err
// 	}
// 	myInvitations := map[uuid.UUID]map[string]*dto.HiringInvitation{}
// 	for storedEvent := range storedEvents.Iter() {
// 		if storedEvent.TypeName == "*enterprise.HiringInvitationSent" {
// 			var e enterprise.HiringInvitationSent
// 			err := json.Unmarshal(storedEvent.EventBody, &e)
// 			if err != nil {
// 				return []dto.HiringInvitation{}, err
// 			}
// 			if e.ProposedTo() == user.Id() {
// 				_, ok := myInvitations[e.ProposedTo()]
// 				if !ok {
// 					myInvitations[e.ProposedTo()] = make(map[string]*dto.HiringInvitation)
// 				}
// 				myInvitations[e.ProposedTo()][storedEvent.EventId.String()] = dto.NewHiringInvitation(storedEvent.EventId.String(), e)
// 				myInvitations[e.ProposedTo()][storedEvent.EventId.String()].SetStatus(dto.PENDING.String())
// 			}
// 		}
// 	}
// 	for e := range storedEvents.Iter() {
// 		switch e.TypeName {
// 		case "*identity.HiringInvitationAccepted":
// 			var event identity.HiringInvitationAccepted
// 			err := json.Unmarshal(e.EventBody, &event)
// 			if err != nil {
// 				return []dto.HiringInvitation{}, err
// 			}
// 			inv, ok := myInvitations[event.InvitedUserId()]
// 			if ok {
// 				for k := range inv {
// 					if event.OccurredOn().After(inv[k].GetOcurredOn()) {
// 						inv[k].SetStatus(dto.ACCEPTED.String())
// 					}
// 				}
// 			}
// 		case "*identity.HiringInvitationRejected":
// 			var event identity.HiringInvitationRejected
// 			err := json.Unmarshal(e.EventBody, &event)
// 			if err != nil {
// 				return []dto.HiringInvitation{}, err
// 			}
// 			inv, ok := myInvitations[event.InvitedUserId()]
// 			if ok {
// 				for k := range inv {
// 					if event.OccurredOn().After(inv[k].GetOcurredOn()) {
// 						inv[k].SetStatus(dto.REJECTED.String())
// 					}
// 				}
// 			}
// 		case "*enterprise.HiringProccessCanceled":
// 			var event enterprise.HiringProccessCanceled
// 			err := json.Unmarshal(e.EventBody, &event)
// 			if err != nil {
// 				return []dto.HiringInvitation{}, err
// 			}
// 			inv, ok := myInvitations[event.CandidateId()]
// 			if ok {
// 				for k := range inv {
// 					if event.OccurredOn().After(inv[k].GetOcurredOn()) {
// 						inv[k].SetStatus(dto.CANCELED.String())
// 						inv[k].SetReason(event.Reason())
// 					}
// 				}
// 			}
// 		}
// 	}

// 	slice := []*dto.HiringInvitation{}
// 	for _, v := range myInvitations {
// 		for _, c := range v {
// 			slice = append(slice, c)
// 		}
// 	}
// 	slices.SortStableFunc(slice, func(i, j *dto.HiringInvitation) int {
// 		if i.GetOcurredOn().Before(j.GetOcurredOn()) {
// 			return 1
// 		}
// 		if i.GetOcurredOn().After(j.GetOcurredOn()) {
// 			return -1
// 		}
// 		return 0
// 	})
// 	slice = slices.CompactFunc(slice, func(i, j *dto.HiringInvitation) bool {
// 		return i.OwnerID == j.OwnerID
// 	})
// 	var hiringInvitationsDTO []dto.HiringInvitation
// 	for _, c := range slice {
// 		ownerId, err := uuid.Parse(c.OwnerID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		user, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, ownerId)
// 		if err != nil {
// 			return []dto.HiringInvitation{}, err
// 		}
// 		ep, err := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Get(ctx, c.EnterpriseID)
// 		if err != nil {
// 			return []dto.HiringInvitation{}, err
// 		}
// 		c.SetUser(*user)
// 		c.SetEnterprise(*ep)
// 		hiringInvitationsDTO = append(hiringInvitationsDTO, *c)
// 	}
// 	slices.SortStableFunc(hiringInvitationsDTO, func(i, j dto.HiringInvitation) int {
// 		if i.GetOcurredOn().Before(j.GetOcurredOn()) {
// 			return 1
// 		}
// 		if i.GetOcurredOn().After(j.GetOcurredOn()) {
// 			return -1
// 		}
// 		return 0
// 	})
// 	return hiringInvitationsDTO, nil
// }
