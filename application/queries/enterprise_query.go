package queries

import (
	"context"
	"encoding/json"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_all_events"
	"go-complaint/infrastructure/persistence/finders/find_all_users"
	"go-complaint/infrastructure/persistence/repositories"
	"slices"
	"time"
)

type EnterpriseQuery struct {
	EnterpriseName string `json:"enterprise_name"`
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
}

func (enterpriseQuery EnterpriseQuery) IsEnterpriseNameAvailable(
	ctx context.Context,
) (bool, error) {
	mapper := repositories.MapperRegistryInstance().Get("Enterprise")
	enterpriseRepository, ok := mapper.(repositories.EnterpriseRepository)
	if !ok {
		return false, repositories.ErrWrongTypeAssertion
	}
	_, err := enterpriseRepository.Get(ctx, enterpriseQuery.EnterpriseName)
	if err != nil {
		return true, nil
	}
	return false, nil
}

func (enterpriseQuery EnterpriseQuery) Enterprise(
	ctx context.Context,
) (dto.Enterprise, error) {
	mapper := repositories.MapperRegistryInstance().Get("Enterprise")
	enterpriseRepository, ok := mapper.(repositories.EnterpriseRepository)
	if !ok {
		return dto.Enterprise{}, repositories.ErrWrongTypeAssertion
	}
	enterprise, err := enterpriseRepository.Get(ctx, enterpriseQuery.EnterpriseName)
	if err != nil {
		return dto.Enterprise{}, err
	}
	return dto.NewEnterprise(enterprise), nil
}

func (enterpriseQuery EnterpriseQuery) HiringInvitationsAccepted(
	ctx context.Context,
) ([]dto.PendingHires, error) {
	if enterpriseQuery.EnterpriseName == "" {
		return nil, ErrBadRequest
	}
	storedEvents, err := repositories.MapperRegistryInstance().Get(
		"StoredEvent",
	).(repositories.EventRepository).FindAll(
		ctx,
		find_all_events.By(),
	)
	if err != nil {
		return nil, err
	}
	myPendings := map[string]identity.HiringInvitationAccepted{}
	for storedEvent := range storedEvents.Iter() {
		if storedEvent.TypeName == "*identity.HiringInvitationAccepted" {
			var e identity.HiringInvitationAccepted
			err := json.Unmarshal(storedEvent.EventBody, &e)
			if err != nil {
				return nil, err
			}
			expired := time.Since(e.OccurredOn()) > 24*10*time.Hour
			if e.EnterpriseID() == enterpriseQuery.EnterpriseName &&
				!expired {
				myPendings[storedEvent.EventId.String()] = e
			}
		}
	}
	for storedEvent := range storedEvents.Iter() {
		switch storedEvent.TypeName {
		case "*enterprise.HiringProccessCanceled":
			var e enterprise.HiringProccessCanceled
			err := json.Unmarshal(storedEvent.EventBody, &e)
			if err != nil {
				return nil, err
			}
			if e.EnterpriseID() == enterpriseQuery.EnterpriseName {
				role, err := identity.ParseRole(e.Position().String())
				if err != nil {
					return nil, err
				}
				for k, v := range myPendings {
					if v.InvitedUserID() == e.CandidateID() ||
						v.ProposedPosition() == role {
						delete(myPendings, k)
					}
				}
			}
		case "*enterprise.EmployeeHired":
			var e enterprise.EmployeeHired
			err := json.Unmarshal(storedEvent.EventBody, &e)
			if err != nil {
				return nil, err
			}
			if e.EnterpriseName() == enterpriseQuery.EnterpriseName {
				role, err := identity.ParseRole(e.Position().String())
				if err != nil {
					return nil, err
				}
				for k, v := range myPendings {
					if v.InvitedUserID() == e.EmployeeEmail() ||
						v.ProposedPosition() == role {
						delete(myPendings, k)
					}
				}
			}
		default:
			continue
		}
	}
	result := make([]dto.PendingHires, 0, len(myPendings))
	for k, v := range myPendings {
		invitedUser, err := repositories.MapperRegistryInstance().Get(
			"User",
		).(repositories.UserRepository).Get(ctx, v.InvitedUserID())
		if err != nil {
			return nil, err
		}
		result = append(result, dto.PendingHires{
			EventID:    k,
			User:       dto.NewUser(*invitedUser),
			Position:   v.ProposedPosition().String(),
			OccurredOn: v.OccurredOn().String(),
		})
	}

	return result, nil
}

func (query EnterpriseQuery) UsersForHiring(
	ctx context.Context,
) ([]dto.User, error) {
	if query.EnterpriseName == "" {
		return nil, ErrBadRequest
	}
	dbEnterprise, err := repositories.MapperRegistryInstance().Get(
		"Enterprise",
	).(repositories.EnterpriseRepository).Get(ctx, query.EnterpriseName)
	if err != nil {
		return nil, err
	}
	users, err := repositories.MapperRegistryInstance().Get(
		"User",
	).(repositories.UserRepository).FindAll(
		ctx,
		find_all_users.ThatAreVerified(query.Limit, query.Offset),
	)
	if err != nil {
		return nil, err
	}
	for emp := range dbEnterprise.Employees().Iter() {
		for user := range users.Iter() {
			if user.Email() == emp.Email() {
				users.Remove(user)
			}
		}
	}
	hireables := users.ToSlice()
	slices.SortStableFunc(hireables, func(i, j identity.User) int {
		if i.RegisterDate().Date().Before(j.RegisterDate().Date()) {
			return -1
		}
		if i.RegisterDate().Date().After(j.RegisterDate().Date()) {
			return 1
		}
		return 0
	})
	result := dto.NewUsersForHiring(hireables, query.Limit, query.Offset)
	return result.Users, nil
}
