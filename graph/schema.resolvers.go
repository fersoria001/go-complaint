package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"go-complaint/application/application_services"
	"go-complaint/application/commands"
	"go-complaint/application/queries"
	"go-complaint/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*model.User, error) {
	c := commands.UserCommand{
		Email:          input.UserName,
		Password:       input.Password,
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		Gender:         input.Genre,
		Pronoun:        input.Pronoun,
		BirthDate:      input.BirthDate,
		Phone:          input.PhoneNumber,
		CountryID:      input.CountryID,
		CountryStateID: input.CountryStateID,
		CityID:         input.CityID,
	}
	err := c.Register(ctx)
	if err != nil {
		return nil, err
	}
	q := queries.UserQuery{
		Email: input.UserName,
	}
	user, err := q.User(ctx)
	if err != nil {
		return nil, err
	}
	return &model.User{
		UserName: user.Email,
		Person: &model.Person{
			ProfileImg:  user.ProfileIMG,
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Genre:       user.Gender,
			Pronoun:     user.Pronoun,
			Age:         user.Age,
			PhoneNumber: user.Phone,
			Address: &model.Address{
				Country:      user.Address.Country,
				CountryState: user.Address.County,
				City:         user.Address.City,
			},
		},
		Status: model.UserStatusOffline,
	}, nil
}

// CreateEnterprise is the resolver for the createEnterprise field.
func (r *mutationResolver) CreateEnterprise(ctx context.Context, input model.CreateEnterprise) (*model.Enterprise, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(ctx)
	if err != nil {
		return nil, err
	}
	c := commands.EnterpriseCommand{
		OwnerID:        currentUser.Email,
		Name:           input.Name,
		Website:        input.Website,
		Email:          input.Email,
		Phone:          input.PhoneNumber,
		CountryID:      input.CountryID,
		CountryStateID: input.CountryStateID,
		CityID:         input.CityID,
		IndustryID:     input.IndustryID,
		FoundationDate: input.FoundationDate,
	}
	err = c.Register(ctx)
	if err != nil {
		return nil, err
	}
	q := queries.EnterpriseQuery{
		EnterpriseName: input.Name,
	}
	enterprise, err := q.Enterprise(ctx)
	if err != nil {
		return nil, err
	}
	employees := make([]*model.Employee, 0, len(enterprise.Employees))
	for _, v := range enterprise.Employees {
		employees = append(employees, &model.Employee{
			ID:           v.ID.String(),
			EnterpriseID: v.EnterpriseID,
			UserID:       v.UserID,
			User: &model.User{
				UserName: v.Email,
				Person: &model.Person{
					ProfileImg:  v.ProfileIMG,
					Email:       v.Email,
					FirstName:   v.FirstName,
					LastName:    v.LastName,
					Genre:       "",
					Pronoun:     "",
					Age:         v.Age,
					PhoneNumber: v.Phone,
					Address:     &model.Address{},
				},
				Status: model.UserStatusOffline,
			},
			HiringDate:         v.HiringDate,
			ApprovedHiring:     v.ApprovedHiring,
			ApprovedHiringAt:   v.ApprovedHiringAt,
			EnterprisePosition: v.Position,
		})
	}
	return &model.Enterprise{
		Name:        enterprise.Name,
		LogoImg:     enterprise.LogoIMG,
		BannerImg:   enterprise.BannerIMG,
		Website:     enterprise.Website,
		Email:       enterprise.Email,
		PhoneNumber: enterprise.Phone,
		Address: &model.Address{
			Country:      enterprise.Address.Country,
			CountryState: enterprise.Address.County,
			City:         enterprise.Address.City,
		},
		Industry:       enterprise.Industry,
		FoundationDate: enterprise.FoundationDate,
		OwnerID:        enterprise.OwnerID,
		Employees:      employees,
	}, nil
}

// UserDescriptor is the resolver for the userDescriptor field.
func (r *queryResolver) UserDescriptor(ctx context.Context) (*model.UserDescriptor, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(ctx)
	if err != nil {
		return nil, err
	}
	authorities := make([]*model.GrantedAuthority, 0, len(currentUser.GrantedAuthorities))
	for _, v := range currentUser.GrantedAuthorities {
		authorities = append(authorities, &model.GrantedAuthority{
			EnterpriseID: v.EnterpriseID,
			Authority:    v.Authority,
		})
	}

	return &model.UserDescriptor{
		UserName:    currentUser.Email,
		FullName:    currentUser.FullName,
		ProfileImg:  currentUser.ProfileIMG,
		Genre:       currentUser.Gender,
		Pronoun:     currentUser.Pronoun,
		Authorities: authorities,
	}, nil
}

// Countries is the resolver for the countries field.
func (r *queryResolver) Countries(ctx context.Context) ([]*model.Country, error) {
	q := queries.AddressQuery{}
	countries, err := q.AllCountries(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Country, 0, len(countries))
	for _, v := range countries {
		result = append(result, &model.Country{
			ID:        v.ID,
			Name:      v.Name,
			PhoneCode: v.PhoneCode,
		})
	}
	return result, nil
}

// CountryStates is the resolver for the countryStates field.
func (r *queryResolver) CountryStates(ctx context.Context, id int) ([]*model.CountryState, error) {
	q := queries.AddressQuery{
		CountryID: id,
	}
	countryStates, err := q.ProvideCountryStateByCountryID(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.CountryState, 0, len(countryStates))
	for _, v := range countryStates {
		result = append(result, &model.CountryState{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return result, nil
}

// Cities is the resolver for the cities field.
func (r *queryResolver) Cities(ctx context.Context, id int) ([]*model.City, error) {
	q := queries.AddressQuery{
		CountryStateID: id,
	}
	cities, err := q.ProvideStateCitiesByStateID(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.City, 0, len(cities))
	for _, v := range cities {
		result = append(result, &model.City{
			ID:          v.ID,
			Name:        v.Name,
			CountryCode: v.CountryCode,
			Latitude:    v.Latitude,
			Longitude:   v.Longitude,
		})
	}
	return result, nil
}

// Industries is the resolver for the industries field.
func (r *queryResolver) Industries(ctx context.Context) ([]*model.Industry, error) {
	q := queries.IndustryQuery{}
	industries, err := q.AllIndustries(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Industry, 0, len(industries))
	for _, v := range industries {
		result = append(result, &model.Industry{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return result, nil
}

// ComplaintsReceivedInfo is the resolver for the complaintsReceivedInfo field.
func (r *queryResolver) ComplaintsReceivedInfo(ctx context.Context, id string) (*model.ComplaintInfo, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		ctx,
		"rid",
		id,
		application_services.READ,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return nil, err
	}
	q := queries.ComplaintQuery{
		ID: id,
	}
	c, err := q.ComplaintsReceivedInfo(ctx)
	if err != nil {
		return nil, err
	}
	total := c.ComplaintsReceived + c.ComplaintsResolved + c.ComplaintsReviewed + c.ComplaintsPending
	return &model.ComplaintInfo{
		Received:  c.ComplaintsReceived,
		Resolved:  c.ComplaintsResolved,
		Reviewed:  c.ComplaintsReviewed,
		Pending:   c.ComplaintsPending,
		AvgRating: c.AverageRating,
		Total:     total,
	}, nil
}

// EnterpriseByID is the resolver for the enterpriseById field.
func (r *queryResolver) EnterpriseByID(ctx context.Context, id string) (*model.Enterprise, error) {
	q := queries.EnterpriseQuery{
		EnterpriseName: id,
	}
	enterprise, err := q.Enterprise(ctx)
	if err != nil {
		return nil, err
	}
	employees := make([]*model.Employee, 0, len(enterprise.Employees))
	for _, v := range enterprise.Employees {
		employees = append(employees, &model.Employee{
			ID:           v.ID.String(),
			EnterpriseID: v.EnterpriseID,
			UserID:       v.UserID,
			User: &model.User{
				UserName: v.Email,
				Person: &model.Person{
					ProfileImg:  v.ProfileIMG,
					Email:       v.Email,
					FirstName:   v.FirstName,
					LastName:    v.LastName,
					Genre:       "",
					Pronoun:     "",
					Age:         v.Age,
					PhoneNumber: v.Phone,
					Address:     &model.Address{},
				},
				Status: model.UserStatusOffline,
			},
			HiringDate:         v.HiringDate,
			ApprovedHiring:     v.ApprovedHiring,
			ApprovedHiringAt:   v.ApprovedHiringAt,
			EnterprisePosition: v.Position,
		})
	}
	return &model.Enterprise{
		Name:        enterprise.Name,
		LogoImg:     enterprise.LogoIMG,
		BannerImg:   enterprise.BannerIMG,
		Website:     enterprise.Website,
		Email:       enterprise.Email,
		PhoneNumber: enterprise.Phone,
		Address: &model.Address{
			Country:      enterprise.Address.Country,
			CountryState: enterprise.Address.County,
			City:         enterprise.Address.City,
		},
		Industry:       enterprise.Industry,
		FoundationDate: enterprise.FoundationDate,
		OwnerID:        enterprise.OwnerID,
		Employees:      employees,
	}, nil
}

// EnterprisesByAuthenticatedUser is the resolver for the enterprisesByAuthenticatedUser field.
func (r *queryResolver) EnterprisesByAuthenticatedUser(ctx context.Context) (*model.EnterprisesByAuthenticatedUserResult, error) {
	user, err := application_services.AuthorizationApplicationServiceInstance().Credentials(ctx)
	if err != nil {
		return nil, err
	}
	owned := make([]*model.EnterpriseByAuthenticatedUser, 0)
	employed := make([]*model.EnterpriseByAuthenticatedUser, 0)
	for _, v := range user.GrantedAuthorities {
		q := queries.EnterpriseQuery{
			EnterpriseName: v.EnterpriseID,
		}
		enterprise, err := q.Enterprise(ctx)
		if err != nil {
			return nil, err
		}
		employees := make([]*model.Employee, 0, len(enterprise.Employees))
		for _, v := range enterprise.Employees {
			employees = append(employees, &model.Employee{
				ID:           v.ID.String(),
				EnterpriseID: v.EnterpriseID,
				UserID:       v.UserID,
				User: &model.User{
					UserName: v.Email,
					Person: &model.Person{
						ProfileImg:  v.ProfileIMG,
						Email:       v.Email,
						FirstName:   v.FirstName,
						LastName:    v.LastName,
						Genre:       "",
						Pronoun:     "",
						Age:         v.Age,
						PhoneNumber: v.Phone,
						Address:     &model.Address{},
					},
					Status: model.UserStatusOffline,
				},
				HiringDate:         v.HiringDate,
				ApprovedHiring:     v.ApprovedHiring,
				ApprovedHiringAt:   v.ApprovedHiringAt,
				EnterprisePosition: v.Position,
			})
		}
		model := &model.EnterpriseByAuthenticatedUser{
			Authority: &model.GrantedAuthority{
				Authority:    v.Authority,
				EnterpriseID: v.EnterpriseID,
			},
			Enterprise: &model.Enterprise{
				Name:        enterprise.Name,
				LogoImg:     enterprise.LogoIMG,
				BannerImg:   enterprise.BannerIMG,
				Website:     enterprise.Website,
				Email:       enterprise.Email,
				PhoneNumber: enterprise.Phone,
				Address: &model.Address{
					Country:      enterprise.Address.Country,
					CountryState: enterprise.Address.County,
					City:         enterprise.Address.City,
				},
				Industry:       enterprise.Industry,
				FoundationDate: enterprise.FoundationDate,
				OwnerID:        enterprise.OwnerID,
				Employees:      employees,
			},
		}
		if v.Authority == "OWNER" {
			owned = append(owned, model)
		} else {
			employed = append(employed, model)
		}
	}
	return &model.EnterprisesByAuthenticatedUserResult{
		Enterprises: owned,
		Offices:     employed,
	}, nil
}

// UsersForHiring is the resolver for the usersForHiring field.
func (r *queryResolver) UsersForHiring(ctx context.Context, input model.SearchWithPagination) (*model.UsersForHiringResult, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		ctx,
		"Enterprise",
		input.ID,
		application_services.READ,
		"MANAGER", "OWNER",
	)
	if err != nil {
		return nil, err
	}
	q := queries.EnterpriseQuery{
		EnterpriseName: input.ID,
		Limit:          input.Limit,
		Offset:         input.Offset,
		Term:           input.Query,
	}
	users, err := q.UsersForHiring(ctx)
	if err != nil {
		return nil, err
	}
	results := make([]*model.User, 0, len(users.Users))
	for _, v := range users.Users {
		results = append(results, &model.User{
			UserName: v.Email,
			Person: &model.Person{
				ProfileImg:  v.ProfileIMG,
				Email:       v.Email,
				FirstName:   v.FirstName,
				LastName:    v.LastName,
				Genre:       v.Gender,
				Pronoun:     v.Pronoun,
				Age:         v.Age,
				PhoneNumber: v.Phone,
				Address: &model.Address{
					Country:      v.Address.Country,
					CountryState: v.Address.County,
					City:         v.Address.City,
				},
			},
			Status: model.UserStatusOffline,
		})
	}
	return &model.UsersForHiringResult{
		Users:      results,
		Count:      users.Count,
		Limit:      users.CurrentLimit,
		Offset:     users.CurrentOffset,
		NextCursor: users.NextCursor,
		PrevCursor: users.NextCursor - 1,
	}, nil
}

// UserByID is the resolver for the userById field.
func (r *queryResolver) UserByID(ctx context.Context, id string) (*model.User, error) {
	q := queries.UserQuery{
		Email: id,
	}
	user, err := q.User(ctx)
	if err != nil {
		return nil, err
	}
	return &model.User{
		UserName: user.Email,
		Person: &model.Person{
			ProfileImg:  user.ProfileIMG,
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Genre:       user.Gender,
			Pronoun:     user.Pronoun,
			Age:         user.Age,
			PhoneNumber: user.Phone,
			Address: &model.Address{
				Country:      user.Address.Country,
				CountryState: user.Address.County,
				City:         user.Address.City,
			},
		},
		Status: model.UserStatusOffline,
	}, nil
}

// HiringInvitationsByAuthenticatedUser is the resolver for the hiringInvitationsByAuthenticatedUser field.
func (r *queryResolver) HiringInvitationsByAuthenticatedUser(ctx context.Context) ([]*model.HiringInvitation, error) {
	user, err := application_services.AuthorizationApplicationServiceInstance().Credentials(ctx)
	if err != nil {
		return nil, err
	}
	q := queries.UserQuery{
		Email: user.Email,
	}
	hiringInvitations, err := q.HiringInvitations(ctx)
	if err != nil {
		return nil, err
	}
	results := make([]*model.HiringInvitation, 0, len(hiringInvitations))
	for _, v := range hiringInvitations {
		results = append(results, &model.HiringInvitation{
			EventID:           v.EventID,
			EnterpriseID:      v.EnterpriseID,
			ProposedPosition:  v.ProposedPosition,
			OwnerID:           v.OwnerID,
			FullName:          v.FullName,
			EnterpriseEmail:   v.EnterpriseEmail,
			EnterprisePhone:   v.EnterprisePhone,
			EnterpriseLogoImg: v.EnterpriseLogoIMG,
			OccurredOn:        v.OccurredOn,
			Seen:              v.Seen,
			Status:            model.HiringProccessState(v.Status),
			Reason:            v.Reason,
		})
	}
	return results, nil
}

// Notifications is the resolver for the notifications field.
func (r *subscriptionResolver) Notifications(ctx context.Context, id string) (<-chan *model.Notification, error) {
	in := commands.NotificationsChannel
	ch := make(chan *model.Notification)
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case n := <-in:
				//this will not work, when you read the message it is consumed from the channel,
				//if the id is not equal the message will be lost
				if n.ID == id {
					ch <- n
				}
			}
		}
	}()
	return ch, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }