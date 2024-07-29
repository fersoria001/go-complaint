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
	"go-complaint/infrastructure/trie"
	"log"
	"math"
	"slices"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type UsersForHiringQuery struct {
	EnterpriseId string `json:"enterpriseId"`
	Term         string `json:"term"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
}

func NewUsersForHiringQuery(
	enterpriseId,
	term string,
	limit,
	offset int,
) *UsersForHiringQuery {
	return &UsersForHiringQuery{
		EnterpriseId: enterpriseId,
		Term:         term,
		Limit:        limit,
		Offset:       offset,
	}
}

func (q UsersForHiringQuery) Execute(ctx context.Context) (*dto.UserTypeList, error) {
	id, err := uuid.Parse(q.EnterpriseId)
	if err != nil {
		return nil, err
	}
	r := repositories.MapperRegistryInstance()
	enterpriseRepository, ok := r.Get("Enterprise").(repositories.EnterpriseRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	userRepository, ok := r.Get("User").(repositories.UserRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	eventsRepository, ok := r.Get("Event").(repositories.EventRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	dbEnterprise, err := enterpriseRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	users, err := userRepository.FindAll(ctx, find_all_users.ThatAreVerified(q.Limit, q.Offset))
	if err != nil {
		return nil, err
	}
	storedEvents, err := eventsRepository.FindAll(ctx, find_all_events.By())
	if err != nil {
		return nil, err
	}
	employeesIds := mapset.NewSet[uuid.UUID]()
	for emp := range dbEnterprise.Employees().Iter() {
		employeesIds.Add(emp.Id())
	}
	idMap := map[uuid.UUID]identity.User{}
	for user := range users.Iter() {
		if user.Id() != dbEnterprise.OwnerId() &&
			!employeesIds.Contains(user.Id()) {
			idMap[user.Id()] = user
		}
	}
	previousEmployees := mapset.NewSet[uuid.UUID]()
	for _, storedEvent := range storedEvents {
		if storedEvent.TypeName == "*enterprise.EmployeeFired" {
			var e enterprise.EmployeeFired

			err := json.Unmarshal(storedEvent.EventBody, &e)

			if err != nil {
				log.Println(err)
				return nil, err
			}

			if e.EnterpriseId() == id {
				previousEmployees.Add(e.UserId())
			}
		}
	}
	for _, storedEvent := range storedEvents {
		switch storedEvent.TypeName {
		case "*identity.HiringInvitationAccepted":
			var e identity.HiringInvitationAccepted
			err := json.Unmarshal(storedEvent.EventBody, &e)
			if err != nil {
				return nil, err
			}
			if e.EnterpriseId() == id && !previousEmployees.Contains(e.InvitedUserId()) {
				delete(idMap, e.InvitedUserId())
			}
		}
	}

	users.Clear()
	for _, user := range idMap {
		users.Add(user)
	}
	if q.Term != "" {

		tree := trie.NewTrie()
		for user := range users.Iter() {
			id := user.Email()
			tree.InsertText(id, user.FullName(), " ")
			tree.InsertText(id, user.Email(), " ")
			tree.InsertText(id, user.Address().City().Name(), " ")
			tree.InsertText(id, user.Address().Country().Name(), " ")
			tree.InsertText(id, user.Address().CountryState().Name(), " ")
			tree.InsertText(id, user.Genre(), " ")
			tree.InsertText(id, user.Pronoun(), " ")
		}

		ids := tree.Search(q.Term)
		s := users.ToSlice()
		for i := range s {
			if !ids.Contains(s[i].Email()) {
				users.Remove(s[i])
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
	result := make([]*dto.User, 0, len(hireables))
	for _, user := range hireables {
		result = append(result, dto.NewUser(&user))
	}
	offset := q.Offset
	limit := q.Limit
	length := len(result)
	offsetLimit := offset + limit
	nextCursor := math.Floor(float64(offset) / float64(limit))
	//offset: 0 | < len | > len
	//limit: 10 | < len | > len

	//bad request
	if offset > length {
		return &dto.UserTypeList{
			Users:         []*dto.User{},
			Count:         0,
			CurrentLimit:  0,
			CurrentOffset: 0,
			NextCursor:    -1,
		}, nil
	}

	//last page
	if offset+limit > length {
		offsetLimit = offset + (length - offset)
		nextCursor = -1
	}

	result = result[offset:offsetLimit]
	return &dto.UserTypeList{
		Users:         result,
		Count:         len(result),
		CurrentLimit:  q.Limit,
		CurrentOffset: q.Offset,
		NextCursor:    int(nextCursor),
	}, nil
}
