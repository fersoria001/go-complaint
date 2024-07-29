package queries

import (
	"context"
	"go-complaint/domain/model/enterprise"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"
)

type EnterpriseQuery struct {
	OwnerId        string `json:"owner_id"`
	EnterpriseName string `json:"enterprise_name"`
	UserID         string `json:"user_id"`
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
	Term           string `json:"term"`
	EventID        string `json:"event_id"`
	ChatID         string `json:"chat_id"`
}

// func (enterpriseQuery EnterpriseQuery) IsEnterpriseNameAvailable(
// 	ctx context.Context,
// ) (bool, error) {
// 	mapper := repositories.MapperRegistryInstance().Get("Enterprise")
// 	enterpriseRepository, ok := mapper.(repositories.EnterpriseRepository)
// 	if !ok {
// 		return false, repositories.ErrWrongTypeAssertion
// 	}
// 	_, err := enterpriseRepository.Get(ctx, enterpriseQuery.EnterpriseName)
// 	if err != nil {
// 		return true, nil
// 	}
// 	return false, nil
// }

// func (enterpriseQuery EnterpriseQuery) Enterprise(
// 	ctx context.Context,
// ) (dto.Enterprise, error) {
// 	mapper := repositories.MapperRegistryInstance().Get("Enterprise")
// 	enterpriseRepository, ok := mapper.(repositories.EnterpriseRepository)
// 	if !ok {
// 		return dto.Enterprise{}, repositories.ErrWrongTypeAssertion
// 	}
// 	enterprise, err := enterpriseRepository.Get(ctx, enterpriseQuery.EnterpriseName)
// 	if err != nil {
// 		return dto.Enterprise{}, err
// 	}
// 	return dto.NewEnterprise(enterprise), nil
// }

// func (q EnterpriseQuery) EnterprisesByOwnerId(
// 	ctx context.Context,
// ) ([]dto.Enterprise, error) {
// 	mapper, ok := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository)
// 	if !ok {
// 		return nil, ErrWrongTypeAssertion
// 	}
// 	enterprises, err := mapper.FindAll(ctx,
// 		find_all_enterprises.ByOwnerId(q.OwnerId),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	s := enterprises.ToSlice()
// 	result := make([]dto.Enterprise, 0, len(s))
// 	for _, v := range s {
// 		result = append(result, dto.NewEnterpriseDTO(v))
// 	}
// 	slices.SortStableFunc(result, func(a, b dto.Enterprise) int {
// 		return strings.Compare(a.Name, b.Name)
// 	})
// 	return result, nil
// }

// func (enterpriseQuery EnterpriseQuery) HiringProcceses(
// 	ctx context.Context,
// ) (dto.HiringProccessList, error) {
// 	if enterpriseQuery.EnterpriseName == "" {
// 		return dto.HiringProccessList{}, ErrBadRequest
// 	}
// 	storedEvents, err := repositories.MapperRegistryInstance().Get(
// 		"Event",
// 	).(repositories.EventRepository).FindAll(
// 		ctx,
// 		find_all_events.By(),
// 	)
// 	if err != nil {
// 		return dto.HiringProccessList{}, err
// 	}
// 	invs := []*dto.HiringProccess{}
// 	for storedEvent := range storedEvents.Iter() {
// 		switch storedEvent.TypeName {
// 		case "*enterprise.HiringInvitationSent":
// 			var e enterprise.HiringInvitationSent
// 			err := json.Unmarshal(storedEvent.EventBody, &e)
// 			if err != nil {
// 				return dto.HiringProccessList{}, err
// 			}
// 			expired := time.Since(e.OccurredOn()) > 24*10*time.Hour
// 			if e.EnterpriseID() == enterpriseQuery.EnterpriseName &&
// 				!expired {
// 				user, err := repositories.MapperRegistryInstance().Get(
// 					"User",
// 				).(repositories.UserRepository).Get(ctx, e.ProposedTo())
// 				if err != nil {
// 					return dto.HiringProccessList{}, err
// 				}
// 				inv := dto.NewHiringProccess(
// 					e.OccurredOn(),
// 					*user,
// 					e.ProposalPosition(),
// 					e.UserID())
// 				inv.SetStatus(dto.PENDING)
// 				invs = append(invs, inv)
// 			}

// 		}
// 	}

// 	slice := make([]*dto.HiringProccess, 0, len(invs))
// 	slice = append(slice, invs...)
// 	slices.SortStableFunc(slice, func(i, j *dto.HiringProccess) int {
// 		if i.GetOcurredOn().Before(j.GetOcurredOn()) {
// 			return 1
// 		}
// 		if i.GetOcurredOn().After(j.GetOcurredOn()) {
// 			return -1
// 		}
// 		return 0
// 	})
// 	slice = slices.CompactFunc(slice, func(i, j *dto.HiringProccess) bool {
// 		return i.User.Email == j.User.Email
// 	})
// 	accepted := map[string]*dto.HiringProccess{}
// 	for _, v := range slice {
// 		accepted[v.User.Email] = v
// 	}
// 	for storedEvent := range storedEvents.Iter() {
// 		switch storedEvent.TypeName {
// 		case "*identity.HiringInvitationRejected":
// 			var e identity.HiringInvitationRejected
// 			err := json.Unmarshal(storedEvent.EventBody, &e)
// 			if err != nil {
// 				return dto.HiringProccessList{}, err
// 			}
// 			inv, ok := accepted[e.InvitedUserID()]
// 			if ok && e.OccurredOn().After(inv.GetOcurredOn()) {
// 				inv.SetStatus(dto.REJECTED)
// 				inv.SetLastUpdate(e.OccurredOn())
// 				inv.SetReason(e.RejectionReason())
// 				inv.SetEventID(storedEvent.EventId.String())
// 			}
// 		case "*enterprise.HiringProccessCanceled":
// 			var e enterprise.HiringProccessCanceled
// 			err := json.Unmarshal(storedEvent.EventBody, &e)
// 			if err != nil {
// 				return dto.HiringProccessList{}, err
// 			}
// 			inv, ok := accepted[e.CandidateID()]
// 			if ok && e.OccurredOn().After(inv.GetOcurredOn()) {
// 				inv.SetStatus(dto.CANCELED)
// 				inv.SetLastUpdate(e.OccurredOn())
// 				inv.SetReason(e.Reason())
// 				inv.SetEmitedByID(e.EmitedBy())
// 				inv.SetEventID(storedEvent.EventId.String())
// 			}
// 		case "*enterprise.EmployeeHired":
// 			var e enterprise.EmployeeHired
// 			err := json.Unmarshal(storedEvent.EventBody, &e)
// 			if err != nil {
// 				return dto.HiringProccessList{}, err
// 			}
// 			inv, ok := accepted[e.EmployeeEmail()]
// 			if ok && e.OccurredOn().After(inv.GetOcurredOn()) {
// 				inv.SetStatus(dto.HIRED)
// 				inv.SetLastUpdate(e.OccurredOn())
// 				inv.SetEmitedByID(e.EmitedBy())
// 				inv.SetEventID(storedEvent.EventId.String())
// 			}
// 		case "*identity.HiringInvitationAccepted":
// 			var e identity.HiringInvitationAccepted
// 			err := json.Unmarshal(storedEvent.EventBody, &e)
// 			if err != nil {
// 				return dto.HiringProccessList{}, err
// 			}
// 			inv, ok := accepted[e.InvitedUserID()]
// 			if ok && e.OccurredOn().After(inv.GetOcurredOn()) {
// 				inv.SetStatus(dto.USER_ACCEPTED)
// 				inv.SetLastUpdate(e.OccurredOn())
// 				inv.SetEventID(storedEvent.EventId.String())
// 			}
// 		case "*enterprise.EmployeeLeaved":
// 			var e enterprise.EmployeeLeaved

// 			err := json.Unmarshal(storedEvent.EventBody, &e)
// 			if err != nil {
// 				log.Println(err)
// 				return dto.HiringProccessList{}, err
// 			}
// 			log.Println("employeLEaved", e)
// 			inv, ok := accepted[e.UserID()]

// 			if ok && e.OccurredOn().After(inv.GetOcurredOn()) {
// 				inv.SetStatus(dto.LEAVED)
// 				inv.SetLastUpdate(e.OccurredOn())
// 				inv.SetEmitedByID(e.UserID())
// 				inv.SetEventID(storedEvent.EventId.String())
// 			}
// 		case "*enterprise.EmployeeFired":
// 			var e enterprise.EmployeeFired
// 			err := json.Unmarshal(storedEvent.EventBody, &e)
// 			if err != nil {
// 				return dto.HiringProccessList{}, err
// 			}
// 			inv, ok := accepted[e.UserID()]
// 			if ok && e.OccurredOn().After(inv.GetOcurredOn()) {
// 				inv.SetStatus(dto.FIRED)
// 				inv.SetLastUpdate(e.OccurredOn())
// 				inv.SetEmitedByID(e.UserID())
// 				inv.SetEventID(storedEvent.EventId.String())
// 			}
// 		default:
// 			continue
// 		}
// 	}

// 	result := make([]dto.HiringProccess, 0, len(accepted))
// 	for _, v := range accepted {
// 		emitedBy, err := repositories.MapperRegistryInstance().Get(
// 			"User",
// 		).(repositories.UserRepository).Get(ctx, v.EmitedByID)
// 		if err != nil {
// 			return dto.HiringProccessList{}, err
// 		}
// 		v.SetEmitedBy(*emitedBy)
// 		result = append(result, *v)
// 	}
// 	if enterpriseQuery.Term != "" {
// 		tree := trie.NewTrie()
// 		for _, v := range result {
// 			tree.InsertText(v.User.Email, v.User.FirstName, " ")
// 			tree.InsertText(v.User.Email, v.User.LastName, " ")
// 			tree.InsertText(v.User.Email, v.User.Email, " ")
// 			tree.InsertText(v.User.Email, v.User.Address.City, " ")
// 			tree.InsertText(v.User.Email, v.User.Address.Country, " ")
// 			tree.InsertText(v.User.Email, v.User.Address.County, " ")
// 			tree.InsertText(v.User.Email, v.User.Gender, " ")
// 			tree.InsertText(v.User.Email, v.Status, "_")
// 			tree.InsertText(v.User.Email, v.EmitedBy.FirstName, " ")
// 			tree.InsertText(v.User.Email, v.EmitedBy.LastName, " ")
// 			tree.InsertText(v.User.Email, v.EmitedBy.Email, " ")
// 			tree.InsertText(v.User.Email, v.EmitedBy.Address.City, " ")
// 			tree.InsertText(v.User.Email, v.EmitedBy.Address.Country, " ")
// 			tree.InsertText(v.User.Email, v.EmitedBy.Address.County, " ")
// 			tree.InsertText(v.User.Email, v.EmitedBy.Gender, " ")
// 		}
// 		ids := tree.Search(enterpriseQuery.Term)
// 		result = slices.DeleteFunc(result, func(i dto.HiringProccess) bool {
// 			return !ids.Contains(i.User.Email)
// 		})
// 	}
// 	count := len(result)
// 	offset := enterpriseQuery.Offset
// 	limit := enterpriseQuery.Limit
// 	length := len(result)
// 	offsetLimit := offset + limit
// 	//offset: 0 | < len | > len
// 	//limit: 10 | < len | > len
// 	if offset > length {
// 		return dto.HiringProccessList{}, fmt.Errorf("offset is greater than the length of the list")
// 	}
// 	if offset+limit > length {
// 		offsetLimit = offset + (length - offset)
// 	}
// 	result = result[offset:offsetLimit]
// 	return dto.HiringProccessList{
// 		HiringProccesses: result,
// 		Count:            count,
// 		CurrentLimit:     limit,
// 		CurrentOffset:    offset,
// 	}, nil
// }

// func (query EnterpriseQuery) UsersForHiring(
// 	ctx context.Context,
// ) (dto.UserTypeList, error) {
// 	if query.EnterpriseName == "" {
// 		return dto.UserTypeList{}, ErrBadRequest
// 	}
// 	dbEnterprise, err := repositories.MapperRegistryInstance().Get(
// 		"Enterprise",
// 	).(repositories.EnterpriseRepository).Get(ctx, query.EnterpriseName)
// 	if err != nil {
// 		return dto.UserTypeList{}, err
// 	}
// 	users, err := repositories.MapperRegistryInstance().Get(
// 		"User",
// 	).(repositories.UserRepository).FindAll(
// 		ctx,
// 		find_all_users.ThatAreVerified(query.Limit, query.Offset),
// 	)
// 	if err != nil {
// 		return dto.UserTypeList{}, err
// 	}
// 	storedEvents, err := repositories.MapperRegistryInstance().Get(
// 		"Event",
// 	).(repositories.EventRepository).FindAll(
// 		ctx,
// 		find_all_events.By(),
// 	)
// 	if err != nil {
// 		return dto.UserTypeList{}, err
// 	}
// 	employeesIDs := mapset.NewSet[string]()
// 	for emp := range dbEnterprise.Employees().Iter() {
// 		employeesIDs.Add(emp.Email())
// 	}
// 	idMap := map[string]identity.User{}
// 	for user := range users.Iter() {
// 		if user.Email() != dbEnterprise.Owner() &&
// 			!employeesIDs.Contains(user.Email()) {
// 			idMap[user.Email()] = user
// 		}
// 	}
// 	previousEmployees := mapset.NewSet[string]()
// 	for storedEvent := range storedEvents.Iter() {
// 		if storedEvent.TypeName == "*enterprise.EmployeeFired" {
// 			var e enterprise.EmployeeFired

// 			err := json.Unmarshal(storedEvent.EventBody, &e)

// 			if err != nil {
// 				log.Println(err)
// 				return dto.UserTypeList{}, err
// 			}

// 			if e.EnterpriseID() == query.EnterpriseName {
// 				previousEmployees.Add(e.UserID())
// 			}
// 		}
// 	}
// 	for storedEvent := range storedEvents.Iter() {
// 		switch storedEvent.TypeName {
// 		case "*identity.HiringInvitationAccepted":
// 			var e identity.HiringInvitationAccepted
// 			err := json.Unmarshal(storedEvent.EventBody, &e)
// 			if err != nil {
// 				return dto.UserTypeList{}, err
// 			}
// 			if e.EnterpriseID() == query.EnterpriseName && !previousEmployees.Contains(e.InvitedUserID()) {
// 				delete(idMap, e.InvitedUserID())
// 			}
// 		}
// 	}

// 	users.Clear()
// 	for _, user := range idMap {
// 		users.Add(user)
// 	}
// 	if query.Term != "" {

// 		tree := trie.NewTrie()
// 		for user := range users.Iter() {
// 			id := user.Email()
// 			tree.InsertText(id, user.FullName(), " ")
// 			tree.InsertText(id, user.Email(), " ")
// 			tree.InsertText(id, user.Address().City().Name(), " ")
// 			tree.InsertText(id, user.Address().Country().Name(), " ")
// 			tree.InsertText(id, user.Address().CountryState().Name(), " ")
// 			tree.InsertText(id, user.Gender(), " ")
// 			tree.InsertText(id, user.Pronoun(), " ")
// 		}

// 		ids := tree.Search(query.Term)
// 		s := users.ToSlice()
// 		for i := range s {
// 			if !ids.Contains(s[i].Email()) {
// 				users.Remove(s[i])
// 			}
// 		}
// 	}
// 	hireables := users.ToSlice()
// 	slices.SortStableFunc(hireables, func(i, j identity.User) int {
// 		if i.RegisterDate().Date().Before(j.RegisterDate().Date()) {
// 			return -1
// 		}
// 		if i.RegisterDate().Date().After(j.RegisterDate().Date()) {
// 			return 1
// 		}
// 		return 0
// 	})
// 	result := make([]dto.User, 0, len(hireables))
// 	for _, user := range hireables {
// 		result = append(result, dto.NewUser(user))
// 	}
// 	offset := query.Offset
// 	limit := query.Limit
// 	length := len(result)
// 	offsetLimit := offset + limit
// 	nextCursor := math.Floor(float64(offset) / float64(limit))
// 	//offset: 0 | < len | > len
// 	//limit: 10 | < len | > len

// 	//bad request
// 	if offset > length {
// 		return dto.UserTypeList{
// 			Users:         []dto.User{},
// 			Count:         0,
// 			CurrentLimit:  0,
// 			CurrentOffset: 0,
// 			NextCursor:    -1,
// 		}, nil
// 	}

// 	//last page
// 	if offset+limit > length {
// 		offsetLimit = offset + (length - offset)
// 		nextCursor = -1
// 	}

// 	result = result[offset:offsetLimit]
// 	return dto.UserTypeList{
// 		Users:         result,
// 		Count:         len(result),
// 		CurrentLimit:  query.Limit,
// 		CurrentOffset: query.Offset,
// 		NextCursor:    int(nextCursor),
// 	}, nil
// }

// func (query EnterpriseQuery) OnlineUsers(
// 	ctx context.Context,
// ) ([]*dto.User, error) {
// 	if query.EnterpriseName == "" || query.UserID == "" {
// 		return nil, ErrBadRequest
// 	}
// 	ep, err := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Get(ctx, query.EnterpriseName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	employees := ep.Employees().ToSlice()
// 	users := make([]*dto.User, 0, len(employees))
// 	for _, v := range employees {
// 		if v.Email() != query.UserID {
// 			userDto := dto.NewUserPtr(*v.GetUser())
// 			userDto.SetStatus(dto.OFFLINE)
// 			users = append(users, userDto)
// 		}
// 	}
// 	if query.UserID != ep.Owner() {
// 		owner, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, ep.Owner())
// 		if err != nil {
// 			return nil, err
// 		}
// 		ownerDto := dto.NewUserPtr(*owner)
// 		ownerDto.SetStatus(dto.OFFLINE)
// 		users = append(users, ownerDto)
// 	}
// 	out := make(chan cache.Request, 1)
// 	cache.RequestChannel2 <- cache.Request{
// 		Type: cache.READ,
// 		Key:  query.EnterpriseName,
// 		Out:  out,
// 	}
// 	req := <-out
// 	if req.Payload != nil {
// 		p, ok := req.Payload.([]string)
// 		if !ok {
// 			return nil, ErrWrongTypeAssertion
// 		}
// 		p = slices.DeleteFunc(p, func(i string) bool {
// 			return query.UserID == i
// 		})
// 		for _, v := range users {
// 			if slices.ContainsFunc(p, func(i string) bool { return i == v.Email }) {
// 				v.SetStatus(dto.ONLINE)
// 			}
// 		}
// 	}
// 	return users, nil
// }

func (query EnterpriseQuery) EnterpriseChat(
	ctx context.Context,
) (dto.Chat, error) {
	if query.EnterpriseName == "" || query.ChatID == "" {
		return dto.Chat{}, ErrBadRequest
	}
	chatID, err := enterprise.NewChatID(query.ChatID)
	if err != nil {
		return dto.Chat{}, err
	}
	chat, err := repositories.MapperRegistryInstance().Get("Chat").(repositories.ChatRepository).Get(ctx, *chatID)
	if err != nil {
		return dto.Chat{}, err
	}
	return *dto.NewChat(*chat), nil
}
