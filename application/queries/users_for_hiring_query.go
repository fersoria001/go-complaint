package queries

import (
	"context"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_all_hiring_proccesses"
	"go-complaint/infrastructure/persistence/finders/find_all_users"
	"go-complaint/infrastructure/persistence/finders/find_enterprise"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/infrastructure/trie"
	"math"
	"slices"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type UsersForHiringQuery struct {
	EnterpriseName string `json:"enterpriseName"`
	Term           string `json:"term"`
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
}

func NewUsersForHiringQuery(
	enterpriseName,
	term string,
	limit,
	offset int,
) *UsersForHiringQuery {
	return &UsersForHiringQuery{
		EnterpriseName: enterpriseName,
		Term:           term,
		Limit:          limit,
		Offset:         offset,
	}
}

func (q UsersForHiringQuery) Execute(ctx context.Context) (*dto.UserTypeList, error) {
	r := repositories.MapperRegistryInstance()
	enterpriseRepository, ok := r.Get("Enterprise").(repositories.EnterpriseRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	userRepository, ok := r.Get("User").(repositories.UserRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	hiringProcessRepository, ok := r.Get("HiringProccess").(repositories.HiringProccessRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	dbEnterprise, err := enterpriseRepository.Find(ctx, find_enterprise.ByName(q.EnterpriseName))
	if err != nil {
		return nil, err
	}
	users, err := userRepository.FindAll(ctx, find_all_users.ThatAreVerified(q.Limit, q.Offset))
	if err != nil {
		return nil, err
	}
	hiringProcesses, err := hiringProcessRepository.FindAll(ctx, find_all_hiring_proccesses.ByEnterpriseId(dbEnterprise.Id()))
	if err != nil {
		return nil, err
	}
	employeesIds := mapset.NewSet[uuid.UUID]()
	for _, emp := range dbEnterprise.Employees().ToSlice() {
		employeesIds.Add(emp.Id())
	}
	userSlice := users.ToSlice()
	userSlice = slices.DeleteFunc(userSlice, func(e identity.User) bool {
		return employeesIds.Contains(e.Id()) || e.Id() == dbEnterprise.OwnerId()
	})
	userSlice = slices.DeleteFunc(userSlice, func(e identity.User) bool {

		return slices.ContainsFunc(hiringProcesses, func(e1 *enterprise.HiringProccess) bool {
			if e1.User().Id() == e.Id() {
				if e1.Status() == enterprise.REJECTED {
					return false
				}
				if e1.Status() == enterprise.CANCELED {
					return false
				}
				if e1.Status() == enterprise.FIRED {
					return false
				}
				if e1.Status() == enterprise.LEAVED {
					return false
				}
				return true
			}
			return false
		})

	})

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
		userSlice = slices.DeleteFunc(userSlice, func(e identity.User) bool {
			return ids.Contains(e.Email())
		})
	}
	slices.SortStableFunc(userSlice, func(i, j identity.User) int {
		if i.RegisterDate().Date().Before(j.RegisterDate().Date()) {
			return -1
		}
		if i.RegisterDate().Date().After(j.RegisterDate().Date()) {
			return 1
		}
		return 0
	})
	result := make([]*dto.User, 0, len(userSlice))
	for _, user := range userSlice {
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
