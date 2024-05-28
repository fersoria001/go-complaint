package graphql_

import (
	"errors"
	"fmt"
	"go-complaint/application"
	"go-complaint/cmd"
	"go-complaint/cmd/api/middleware"
	"go-complaint/dto"
	"go-complaint/erros"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"net/mail"

	"github.com/graphql-go/graphql"
)

/*
Front-end requirements:
- User should be able to register:
	- This requires a list of countries, counties, cities and phone codes to be available (done)
- User should be able to recover it's password 		(done)
- User should be able to update it's profile: 		(done)
	- Fields that can be updated: profile image, first name, last name, phone, country, county, city (done)
	- This requires a list of countries, counties, cities and phone codes to be available (done)
- User should be able to login and receive a JWTtoken (done)
- User should know beforehand if he/she owns any enterprise ????? maybe this is a design problem
- User should be able to send a complaint 		(done)
- User should be able to see his/her sent,draft and archived complaints (done)
- User should be able to rate a closed complaint (done)

- When user is an enterprise owner:
	- User should be able to create an enterprise: 		(done)
		- This requires a list of industries to be available (done)
		- This requires a list of countries, counties, cities and phone codes to be available (done)
	- User should be able to update the enterprise profile:
		- Fields that can be updated: logo image, website, email, phone,
		 country, county, city   												(done)
		- This requires a list of industries to be available (done)
		- This requires a list of countries, counties, cities and phone codes to be available (done)
	- User should be able to see the enterprise profile:
		- He/she should be able to hire,approve and fire employees (not in front-end yet)
		- He/she should be able to see the enterprise complaints
		- He/she should be able to see,review and manage the enterprise employees (done)
		- User should be able to create,see and manage enterprise complaints

- When user is an employee:
	- User should be able to see the enterprise employee profile
	- User should be able to see the enterprise complaints
	- User should be able to see the enterprise employees
	- User should be able to see and manage the enterprise complaints
	- If employeee is a manager:
		- User should be able to hire employees
		- User should be able to feedback assistants



All app services will be translated later to command/queries after the persistence layer update
*/

// root mutation
var mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		//################# USER SCHEMA ####################
		"CreateUser": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Create new user",
			Args:        CreateUser,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())

				identityService := application.NewIdentityService(userRepository)
				locationsRepository := repositories.NewLocationsRepository(datasource.PublicSchema())
				a := p.Args
				country, err := locationsRepository.FindCountryByID(p.Context, a["country"].(int))
				if err != nil {
					return false, err
				}
				county, err := locationsRepository.FindCountyByID(p.Context, a["county"].(int))
				if err != nil {
					return false, err
				}
				city, err := locationsRepository.FindCityByID(p.Context, a["city"].(int))
				if err != nil {
					return false, err
				}
				err = identityService.RegisterUser(p.Context,
					"/default.jpg",
					a["email"].(string),
					a["password"].(string),
					a["firstName"].(string),
					a["lastName"].(string),
					a["birthDate"].(string),
					a["phone"].(string),
					country.Name,
					county.Name,
					city.Name,
				)
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		"RecoverPassword": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Set a new random password for the user and send it by email",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				identityService := application.NewIdentityService(userRepository)
				userID, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return false, err
				}
				err = identityService.RecoverPassword(params.Context, userID)
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		"ChangePassword": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Set a new random password for the user and send it by email",
			Args:        ChangePassword,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				identityService := application.NewIdentityService(userRepository)
				userID, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return false, err
				}
				err = identityService.ChangePassword(
					params.Context,
					userID,
					params.Args["oldPassword"].(string),
					params.Args["newPassword"].(string),
				)
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		"UpdateUserProfile": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Update the user profile mutable fields are: profile image, first name, last name, phone, country, county, city",
			Args:        UpdateUserProfile,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var (
					keys = []string{"profileIMG",
						"firstName",
						"lastName",
						"phone",
						"country",
						"county",
						"city"}
					profileIMG = ""
					firstName  = ""
					lastName   = ""
					phone      = ""
					country    = ""
					county     = ""
					city       = ""
				)
				for _, k := range keys {
					if value, ok := params.Args[k]; ok {
						switch k {
						case "profileIMG":
							profileIMG = value.(string)
						case "firstName":
							firstName = value.(string)
						case "lastName":
							lastName = value.(string)
						case "phone":
							phone = value.(string)
						case "country":
							country = value.(string)
						case "county":
							county = value.(string)
						case "city":
							city = value.(string)
						}
					}
				}
				userID, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return false, err
				}
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				identityService := application.NewIdentityService(userRepository)
				err = identityService.UpdateUserProfile(params.Context,
					userID,
					profileIMG, firstName, lastName, phone, country, county, city)
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		//################# ENTERPRISE SCHEMA ####################
		"CreateEnterprise": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Create a new enterprise",
			Args:        CreateAnEnterprise,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				ownerID, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return false, err
				}
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
				identityService := application.NewIdentityService(userRepository)
				eventStore := repositories.NewEventRepository(datasource.EventSchema())
				eventService := application.NewNotificationService(eventStore)
				enterpriseService := application.NewEnterpriseService(
					enterpriseRepository,
					identityService,
					employeesRepository,
					eventService,
				)
				locationsRepository := repositories.NewLocationsRepository(datasource.PublicSchema())
				a := params.Args
				country, err := locationsRepository.FindCountryByID(params.Context, a["country"].(int))
				if err != nil {
					return false, err
				}
				county, err := locationsRepository.FindCountyByID(params.Context, a["county"].(int))
				if err != nil {
					return false, err
				}
				city, err := locationsRepository.FindCityByID(params.Context, a["city"].(int))
				if err != nil {
					return false, err
				}
				//Add  banner image later
				err = enterpriseService.CreateEnterprise(params.Context,
					ownerID,
					params.Args["name"].(string),
					params.Args["website"].(string),
					params.Args["email"].(string),
					params.Args["phone"].(string),
					country.Name,
					county.Name,
					city.Name,
					params.Args["industry"].(string),
					params.Args["foundationDate"].(string))
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		"UpdateEnterprise": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Update the enterprise profile mutable fields are: logo image, website, email, phone, country, county, city, industry",
			Args:        UpdateEnterprise,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var (
					keys = []string{
						"logoIMG",
						"website",
						"email",
						"phone",
						"country",
						"county",
						"city",
						"industry"}
					logoIMG = ""
					website = ""
					email   = ""
					phone   = ""
					country = ""
					county  = ""
					city    = ""
				)
				for _, k := range keys {
					if value, ok := params.Args[k]; ok {
						switch k {
						case "logoIMG":
							logoIMG = value.(string)
						case "website":
							website = value.(string)
						case "email":
							email = value.(string)
						case "phone":
							phone = value.(string)
						case "country":
							country = value.(string)
						case "county":
							county = value.(string)
						case "city":
							city = value.(string)
						}
					}
				}
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
				identityService := application.NewIdentityService(userRepository)
				eventStore := repositories.NewEventRepository(datasource.EventSchema())
				eventService := application.NewNotificationService(eventStore)
				enterpriseService := application.NewEnterpriseService(
					enterpriseRepository,
					identityService,
					employeesRepository,
					eventService,
				)
				err := enterpriseService.UpdateEnterprise(params.Context,
					params.Args["id"].(string),
					logoIMG, website, email, phone, country, county, city)
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		//################# OWNER SCHEMA ####################
		"InviteToProject": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Invite a user to be part of a project in the enterprise",
			Args:        InviteToProject,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
				identityService := application.NewIdentityService(userRepository)
				eventStore := repositories.NewEventRepository(datasource.EventSchema())
				eventService := application.NewNotificationService(eventStore)
				enterpriseService := application.NewEnterpriseService(
					enterpriseRepository,
					identityService,
					employeesRepository,
					eventService,
				)
				userID, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return false, err
				}

				err = enterpriseService.InviteToProject(
					params.Context,
					params.Args["enterpriseName"].(string),
					userID,
					params.Args["userEmail"].(string),
					params.Args["position"].(string),
				)
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		// "HireFromEvent": &graphql.Field{
		// 	Type:        graphql.Boolean,
		// 	Description: "Invite a user to be part of a project in the enterprise",
		// 	Args:        HireFromEvent,
		// 	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// 		enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
		// 		userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
		// 		employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
		// 		identityService := application.NewIdentityService(userRepository)
		// 		eventStore := repositories.NewEventRepository(datasource.EventSchema())
		// 		enterpriseService := application.NewEnterpriseService(
		// 			enterpriseRepository,
		// 			identityService,
		// 			employeesRepository,
		// 			eventStore,
		// 		)
		// 		userID, err := middleware.GetContextPersonID(params.Context)
		// 		if err != nil {
		// 			return false, err
		// 		}

		// 		err = enterpriseService.HireInvitedUser(
		// 			params.Context,
		// 			userID,
		// 			params.Args["eventID"].(string),
		// 		)
		// 		if err != nil {
		// 			return false, err
		// 		}
		// 		return true, nil
		// 	},
		// },
		"EndHiringProcess": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Hire an employee who is in the hiring process",
			Args:        EndHiringProcess,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				accepted, ok := params.Args["accepted"].(bool)
				if !ok {
					return false, &erros.ValidationError{Expected: "accepted field is required"}
				}
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
				identityService := application.NewIdentityService(userRepository)
				eventStore := repositories.NewEventRepository(datasource.EventSchema())
				eventService := application.NewNotificationService(eventStore)
				enterpriseService := application.NewEnterpriseService(
					enterpriseRepository,
					identityService,
					employeesRepository,
					eventService,
				)
				notificationService := application.NewNotificationService(eventStore)
				userID, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return false, err
				}
				currentUser, err := identityService.User(params.Context, userID)
				if err != nil {
					return false, err
				}
				authorities := currentUser.GetAuthorities(params.Args["enterpriseName"].(string))
				var auth bool = false
				for _, a := range authorities {
					if a.Authority() == "OWNER" {
						auth = true
					}
				}
				if !auth {
					return false, &erros.UnauthorizedError{}
				}
				if accepted {
					err = enterpriseService.ApproveEmployee(
						params.Context,
						params.Args["enterpriseName"].(string),
						params.Args["employeeID"].(string),
					)
					if err != nil {
						return false, err
					}
				} else {
					err = enterpriseService.CancelHiring(
						params.Context,
						params.Args["enterpriseName"].(string),
						params.Args["employeeID"].(string),
					)
					if err != nil {
						return false, err
					}
				}
				err = notificationService.MoveToLog(params.Context, params.Args["pendingEventID"].(string))
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		"FireEmployee": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Hire an employee who is in the hiring process",
			Args:        FireEmployee,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
				identityService := application.NewIdentityService(userRepository)
				eventStore := repositories.NewEventRepository(datasource.EventSchema())
				eventService := application.NewNotificationService(eventStore)
				enterpriseService := application.NewEnterpriseService(
					enterpriseRepository,
					identityService,
					employeesRepository,
					eventService,
				)
				userID, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return false, err
				}
				currentUser, err := identityService.User(params.Context, userID)
				if err != nil {
					return false, err
				}
				authorities := currentUser.GetAuthorities(params.Args["enterpriseName"].(string))
				var auth bool = false
				for _, a := range authorities {
					if a.Authority() == "OWNER" {
						auth = true
					}
				}
				if !auth {
					return false, &erros.UnauthorizedError{}
				}
				err = enterpriseService.FireEmployee(
					params.Context,
					params.Args["enterpriseName"].(string),
					params.Args["employeeID"].(string),
				)
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		//################# EMPLOYEE SCHEMA ####################
		"SendHiringProcessInvitation": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Invite a user to be part of a project in the enterprise",
			Args:        SendHiringInvitation,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
				identityService := application.NewIdentityService(userRepository)
				employeeService := application.NewEmployeesService(
					employeesRepository,
					identityService,
				)
				err := employeeService.SendHiringInvitation(
					params.Context,
					params.Args["managerID"].(string),
					params.Args["userEmail"].(string),
				)
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		"AcceptHiringInvitation": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Let a user accept an invitation and trigger an event to notify the enterprise owner",
			Args:        StringID,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
				identityService := application.NewIdentityService(userRepository)
				eventStore := repositories.NewEventRepository(datasource.EventSchema())
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				eventService := application.NewNotificationService(eventStore)
				enterpriseService := application.NewEnterpriseService(
					enterpriseRepository,
					identityService,
					employeesRepository,
					eventService,
				)
				err := enterpriseService.AcceptHiringInvitation(params.Context, params.Args["ID"].(string))
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		//################# COMPLAINT SCHEMA ####################
		"SendComplaint": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Send a complaint to a user or enterprise",
			Args:        CreateAComplaint,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				complaintRepository := repositories.NewComplaintRepository(datasource.ComplaintSchema())
				repliesRepository := repositories.NewReplyRepository(datasource.ComplaintSchema())
				complaintService := application.NewComplaintService(complaintRepository, repliesRepository)
				//user context id -> (ucid, senderID) => isEnterpriseID? isOwnerByID(ucid, senderID) :
				//(ucid === senderID) ? send() : return error
				err := complaintService.CreateComplaint(params.Context,
					params.Args["authorID"].(string),
					params.Args["receiverID"].(string),
					params.Args["title"].(string),
					params.Args["description"].(string),
					params.Args["content"].(string))
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		"RateComplaint": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Rate a complaint",
			Args:        RateAComplaint,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				userID, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return false, err
				}
				complaintRepository := repositories.NewComplaintRepository(datasource.ComplaintSchema())
				repliesRepository := repositories.NewReplyRepository(datasource.ComplaintSchema())
				complaintService := application.NewComplaintService(complaintRepository, repliesRepository)
				err = complaintService.RateComplaint(
					params.Context,
					params.Args["ID"].(string),
					userID,
					params.Args["rating"].(int),
					params.Args["comment"].(string),
				)
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		"ReplyToComplaint": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Reply to a complaint",
			Args:        ReplyToComplaint,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				complaintRepository := repositories.NewComplaintRepository(datasource.ComplaintSchema())
				repliesRepository := repositories.NewReplyRepository(datasource.ComplaintSchema())
				complaintService := application.NewComplaintService(complaintRepository, repliesRepository)
				userID, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return false, err
				}
				err = complaintService.ReplyComplaint(
					params.Context,
					params.Args["complaintID"].(string),
					params.Args["profileIMG"].(string),
					params.Args["fullName"].(string),
					userID,
					params.Args["body"].(string),
				)
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		"MarkAsReviewable": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Mark a complaint as reviewable",
			Args:        MarkAsReviewable,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				complaintService := cmd.ComplaintServiceInstance()
				loggedUser, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return false, err
				}
				enterpriseID, ok := params.Args["enterpriseID"].(string)
				if !ok {
					return false, errors.New("enterpriseID is required")
				}
				assistantID, ok := params.Args["assistantID"].(string)
				if !ok || assistantID == "" {
					assistantID = ""
					eps, err := cmd.EnterpriseServiceInstance().OwnerOfEnterprise(params.Context, loggedUser)
					if err != nil {
						return false, err
					}
					for e := range eps.Iter() {
						if e.Name() == enterpriseID {
							assistantID = cmd.EnterpriseServiceInstance().NewEmployeeID(
								enterpriseID,
								loggedUser, loggedUser,
							)
						}
					}
					if assistantID == "" {
						return false, errors.New("assistantID can't be obtained from logged user")
					}
				}
				err = complaintService.SendForReviewing(
					params.Context,
					params.Args["complaintID"].(string),
					assistantID,
				)
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		"CloseComplaint": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Mark a complaint as reviewable",
			Args:        MarkAsReviewable,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				complaintRepository := repositories.NewComplaintRepository(datasource.ComplaintSchema())
				repliesRepository := repositories.NewReplyRepository(datasource.ComplaintSchema())
				complaintService := application.NewComplaintService(complaintRepository, repliesRepository)
				_, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return false, err
				}
				err = complaintService.Close(
					params.Context,
					params.Args["complaintID"].(string),
					params.Args["closeRequesterID"].(string),
				)
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		//################# FEEDBACK SCHEMA ####################
		"CreateAFeedback": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Create a new single feedback for a complaint",
			Args:        CreateAFeedback,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				feedbackRepository := repositories.NewFeedbackRepository(datasource.FeedbackSchema())
				answerRepository := repositories.NewAnswerRepository(datasource.FeedbackSchema())
				feedbackService := application.NewFeedbackService(feedbackRepository, answerRepository)
				_, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return false, err
				}
				err = feedbackService.CreateAFeedback(
					params.Context,
					params.Args["complaintID"].(string),
					params.Args["reviewerID"].(string),
					params.Args["reviewedID"].(string),
					params.Args["reviewerIMG"].(string),
					params.Args["reviewerName"].(string),
					params.Args["senderID"].(string),
					params.Args["senderIMG"].(string),
					params.Args["senderName"].(string),
					params.Args["body"].(string),
					params.Args["createdAt"].(string),
					params.Args["read"].(bool),
					params.Args["readAt"].(string),
					params.Args["updatedAt"].(string),
					params.Args["comment"].(string),
				)
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		"AnswerAFeedback": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Create a new single feedback for a complaint",
			Args:        AnswerAFeedback,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				feedbackRepository := repositories.NewFeedbackRepository(datasource.FeedbackSchema())
				answerRepository := repositories.NewAnswerRepository(datasource.FeedbackSchema())
				feedbackService := application.NewFeedbackService(feedbackRepository, answerRepository)
				_, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return false, err
				}
				err = feedbackService.AnswerAFeedback(
					params.Context,
					params.Args["complaintID"].(string),
					params.Args["senderID"].(string),
					params.Args["senderIMG"].(string),
					params.Args["senderName"].(string),
					params.Args["body"].(string),
					params.Args["createdAt"].(string),
					params.Args["read"].(bool),
					params.Args["readAt"].(string),
					params.Args["updatedAt"].(string),
				)
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		"MarkAnswerAsRead": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Mark a feedback answer as read",
			Args:        MarkAnswerAsRead,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				feedbackRepository := repositories.NewFeedbackRepository(datasource.FeedbackSchema())
				answerRepository := repositories.NewAnswerRepository(datasource.FeedbackSchema())
				feedbackService := application.NewFeedbackService(feedbackRepository, answerRepository)
				_, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return false, err
				}
				err = feedbackService.MarkAnswerAsRead(
					params.Context,
					params.Args["answerID"].(string),
				)
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
	},
})

var query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		//################# USER SCHEMA ####################
		"Login": &graphql.Field{
			Type:        JwtTokenType,
			Description: "Get the token for the authenticated user or error",
			Args:        LoginUser,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				rememberMe := false
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				identityService := application.NewIdentityService(userRepository)
				if params.Args["rememberMe"] != nil {
					rememberMe = params.Args["rememberMe"].(bool)
				}
				token, err := identityService.Login(params.Context,
					params.Args["email"].(string),
					params.Args["password"].(string), rememberMe)
				if err != nil {
					return nil, err
				}
				return struct {
					Token string `json:"token"`
				}{
					token,
				}, nil
			},
		},
		"UserDescriptor": &graphql.Field{
			Type:        UserDescriptorType,
			Description: "Get the user descriptor for the current session",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				token, err := middleware.GetContextToken(params.Context)
				if err != nil {
					return nil, err
				}
				userDescriptor, err := application.NewJWTService().ParseUserDescriptor(token)
				if err != nil {
					return nil, err
				}
				return struct {
					Email      string `json:"email"`
					FullName   string `json:"fullName"`
					ProfileIMG string `json:"profileIMG"`
					IP         string `json:"ip"`
				}{
					Email:      userDescriptor.Email,
					FullName:   userDescriptor.FullName,
					ProfileIMG: userDescriptor.ProfileIMG,
					IP:         userDescriptor.IP,
				}, nil
			},
		},
		"User": &graphql.Field{
			Type:        UserType,
			Description: "Get a user without private information by it's ID",
			Args:        StringID,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				identityService := application.NewIdentityService(userRepository)
				user, err := identityService.User(params.Context, params.Args["ID"].(string))
				if err != nil {
					return nil, err
				}
				result := dto.NewUser(user)
				return result, nil
			},
		},

		//################# MISC SCHEMA ####################
		"Countries": &graphql.Field{
			Type:        graphql.NewList(CountryType),
			Description: "Get the list of countries",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				locationsRepository := repositories.NewLocationsRepository(datasource.PublicSchema())
				countries, err := locationsRepository.FindAllCountries(params.Context)
				if err != nil {
					return nil, err
				}
				return countries, nil
			},
		},
		"Counties": &graphql.Field{
			Type:        graphql.NewList(CountyType),
			Description: "find all counties by country ID",
			Args:        IntegerID,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				locationsRepository := repositories.NewLocationsRepository(datasource.PublicSchema())
				counties, err := locationsRepository.FindCountyByCountryID(params.Context, params.Args["ID"].(int))
				if err != nil {
					return nil, err
				}
				return counties, nil
			},
		},
		"Cities": &graphql.Field{
			Type:        graphql.NewList(CityType),
			Description: "find all cities by county ID",
			Args:        IntegerID,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				locationsRepository := repositories.NewLocationsRepository(datasource.PublicSchema())
				cities, err := locationsRepository.FindCityByCountyID(params.Context, params.Args["ID"].(int))
				if err != nil {
					return nil, err
				}
				return cities, nil
			},
		},
		"PhoneCode": &graphql.Field{
			Type:        PhoneCodeType,
			Description: "Get the single phone code for the given country",
			Args:        IntegerID,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				locationsRepository := repositories.NewLocationsRepository(datasource.PublicSchema())
				phonecode, err := locationsRepository.FindPhoneCodeByCountryID(params.Context, params.Args["ID"].(int))
				if err != nil {
					return nil, err
				}
				return phonecode, nil
			},
		},
		"IsEnterpriseNameAvailable": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Check if the enterprise name is available",
			Args:        StringID,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
				identityService := application.NewIdentityService(userRepository)
				eventStore := repositories.NewEventRepository(datasource.EventSchema())
				eventService := application.NewNotificationService(eventStore)
				enterpriseService := application.NewEnterpriseService(
					enterpriseRepository,
					identityService,
					employeesRepository,
					eventService,
				)
				err := enterpriseService.IsAvailable(params.Context, params.Args["ID"].(string))
				if err != nil {
					return false, nil
				}
				return true, nil
			},
		},
		"Industries": &graphql.Field{
			Type:        graphql.NewList(IndustryType),
			Description: "Get the list of industries",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				industries, err := enterpriseRepository.GetIndustriesSlice(params.Context)
				if err != nil {
					return nil, err
				}
				return industries, nil
			},
		},
		//################# COMPLAINT SCHEMA ####################
		"FindReceiver": &graphql.Field{
			Type:        graphql.NewList(Receiver),
			Description: "Return a concat array of enterprises and user possible complaint receivers",
			Args:        SearchTerm,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				enterprisesReceivers, err := enterpriseRepository.FindByName(params.Context, params.Args["term"].(string))
				if err != nil {
					return nil, err
				}
				usersReceivers, err := userRepository.FindByName(params.Context, params.Args["term"].(string))
				if err != nil {
					return nil, err
				}
				result := append(enterprisesReceivers, usersReceivers...)
				return result, nil
			},
		},
		"Draft": &graphql.Field{
			Type:        ComplaintTypeList,
			Description: "Get the list of draft complaints",
			Args:        GetComplaint,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				complaintRepository := repositories.NewComplaintRepository(datasource.ComplaintSchema())
				repliesRepository := repositories.NewReplyRepository(datasource.ComplaintSchema())
				complaintService := application.NewComplaintService(complaintRepository, repliesRepository)
				id, ok := params.Args["ID"].(string)
				if !ok {
					return nil, fmt.Errorf("ID is required")
				}
				status, ok := params.Args["status"].(string)
				if !ok {
					status = ""
				}
				limit, ok := params.Args["limit"].(int)
				if !ok {
					limit = 10
				}
				offset, ok := params.Args["offset"].(int)
				if !ok {
					offset = 0
				}
				listWithInfo, err := complaintService.GetComplaintsTo(
					params.Context,
					id, status, limit, offset,
				)
				if err != nil {
					expected := &erros.ValueNotFoundError{}
					if errors.As(err, &expected) {
						return &dto.ComplaintListDTO{
							Complaints:    []*dto.ComplaintDTO{},
							Count:         0,
							CurrentLimit:  10,
							CurrentOffset: 0,
						}, nil
					}
					return nil, err
				}
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				for _, c := range listWithInfo.Complaints {
					if _, err = mail.ParseAddress(c.AuthorID); err != nil {
						enterprise, err := enterpriseRepository.Get(params.Context, c.AuthorID)
						if err != nil {
							return nil, err
						}
						c.AuthorFullName = enterprise.Name()
						c.AuthorProfileIMG = enterprise.LogoIMG()
					} else {
						user, err := userRepository.Get(params.Context, c.AuthorID)
						if err != nil {
							return nil, err
						}
						c.AuthorFullName = user.Person().FullName()
						c.AuthorProfileIMG = user.ProfileIMG()
					}
					if _, err = mail.ParseAddress(c.ReceiverID); err != nil {
						enterprise, err := enterpriseRepository.Get(params.Context, c.ReceiverID)
						if err != nil {
							return nil, err
						}
						c.ReceiverFullName = enterprise.Name()
						c.ReceiverProfileIMG = enterprise.LogoIMG()
					} else {
						user, err := userRepository.Get(params.Context, c.ReceiverID)
						if err != nil {
							return nil, err
						}
						c.ReceiverFullName = user.Person().FullName()
						c.ReceiverProfileIMG = user.ProfileIMG()
					}
				}
				return listWithInfo, nil
			},
		},
		"Sent": &graphql.Field{
			Type:        ComplaintTypeList,
			Description: "Get the list of sent complaints",
			Args:        GetComplaint,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				complaintRepository := repositories.NewComplaintRepository(datasource.ComplaintSchema())
				repliesRepository := repositories.NewReplyRepository(datasource.ComplaintSchema())
				complaintService := application.NewComplaintService(complaintRepository, repliesRepository)
				id, ok := params.Args["ID"].(string)
				if !ok {
					return nil, fmt.Errorf("ID is required")
				}
				status, ok := params.Args["status"].(string)
				if !ok {
					status = ""
				}
				limit, ok := params.Args["limit"].(int)
				if !ok {
					limit = 10
				}
				offset, ok := params.Args["offset"].(int)
				if !ok {
					offset = 0
				}
				listWithInfo, err := complaintService.GetComplaintsFrom(
					params.Context,
					id, status, limit, offset,
				)
				if err != nil {
					expected := &erros.ValueNotFoundError{}
					if errors.As(err, &expected) {
						return &dto.ComplaintListDTO{
							Complaints:    []*dto.ComplaintDTO{},
							Count:         0,
							CurrentLimit:  10,
							CurrentOffset: 0,
						}, nil
					}
					return nil, err
				}
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				for _, c := range listWithInfo.Complaints {
					if _, err = mail.ParseAddress(c.AuthorID); err != nil {
						enterprise, err := enterpriseRepository.Get(params.Context, c.AuthorID)
						if err != nil {
							return nil, err
						}
						c.AuthorFullName = enterprise.Name()
						c.AuthorProfileIMG = enterprise.LogoIMG()
					} else {
						user, err := userRepository.Get(params.Context, c.AuthorID)
						if err != nil {
							return nil, err
						}
						c.AuthorFullName = user.Person().FullName()
						c.AuthorProfileIMG = user.ProfileIMG()
					}
					if _, err = mail.ParseAddress(c.ReceiverID); err != nil {
						enterprise, err := enterpriseRepository.Get(params.Context, c.ReceiverID)
						if err != nil {
							return nil, err
						}
						c.ReceiverFullName = enterprise.Name()
						c.ReceiverProfileIMG = enterprise.LogoIMG()
					} else {
						user, err := userRepository.Get(params.Context, c.ReceiverID)
						if err != nil {
							return nil, err
						}
						c.ReceiverFullName = user.Person().FullName()
						c.ReceiverProfileIMG = user.ProfileIMG()
					}
				}
				return listWithInfo, nil
			},
		},
		"History": &graphql.Field{
			Type:        graphql.NewList(ComplaintType),
			Description: "Get the list of archived complaints",
			Args:        GetComplaint,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				complaintRepository := repositories.NewComplaintRepository(datasource.ComplaintSchema())
				repliesRepository := repositories.NewReplyRepository(datasource.ComplaintSchema())
				complaintService := application.NewComplaintService(complaintRepository, repliesRepository)
				return complaintService.GetComplaintsTo(
					params.Context,
					params.Args["ID"].(string),
					"IN_HISTORY",
					params.Args["limit"].(int),
					params.Args["offset"].(int),
				)
			},
		},
		"Complaint": &graphql.Field{
			Type:        ComplaintType,
			Description: "Get a complaint by ID with all it's details",
			Args:        StringID,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				complaintRepository := repositories.NewComplaintRepository(datasource.ComplaintSchema())
				repliesRepository := repositories.NewReplyRepository(datasource.ComplaintSchema())
				complaintService := application.NewComplaintService(complaintRepository, repliesRepository)
				c, err := complaintService.Complaint(params.Context, params.Args["ID"].(string))
				if err != nil {
					return nil, err
				}
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())

				if _, err = mail.ParseAddress(c.AuthorID); err != nil {
					enterprise, err := enterpriseRepository.Get(params.Context, c.AuthorID)
					if err != nil {
						return nil, err
					}
					c.AuthorFullName = enterprise.Name()
					c.AuthorProfileIMG = enterprise.LogoIMG()
				} else {
					user, err := userRepository.Get(params.Context, c.AuthorID)
					if err != nil {
						return nil, err
					}
					c.AuthorFullName = user.Person().FullName()
					c.AuthorProfileIMG = user.ProfileIMG()
				}
				if _, err = mail.ParseAddress(c.ReceiverID); err != nil {
					enterprise, err := enterpriseRepository.Get(params.Context, c.ReceiverID)
					if err != nil {
						return nil, err
					}
					c.ReceiverFullName = enterprise.Name()
					c.ReceiverProfileIMG = enterprise.LogoIMG()
				} else {
					user, err := userRepository.Get(params.Context, c.ReceiverID)
					if err != nil {
						return nil, err
					}
					c.ReceiverFullName = user.Person().FullName()
					c.ReceiverProfileIMG = user.ProfileIMG()
				}
				return *c, nil
			},
		},
		//################# ENTERPRISE SCHEMA #########################
		"OwnerEnterprises": &graphql.Field{
			Type:        graphql.NewList(EnterpriseType),
			Description: "Return all the enterprises where the subject is the owner",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
				identityService := application.NewIdentityService(userRepository)
				eventStore := repositories.NewEventRepository(datasource.EventSchema())
				eventService := application.NewNotificationService(eventStore)
				enterpriseService := application.NewEnterpriseService(
					enterpriseRepository,
					identityService,
					employeesRepository,
					eventService,
				)
				userID, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return nil, err
				}
				enterprises, err := enterpriseService.ProvideOwnerEnterprises(params.Context, userID)
				if err != nil {

					return enterprises, err
				}
				result := enterprises.ToSlice()

				return result, nil
			},
		},
		"Enterprise": &graphql.Field{
			Type:        EnterpriseType,
			Description: "Return the enterprise info",
			Args:        StringID,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
				identityService := application.NewIdentityService(userRepository)
				eventStore := repositories.NewEventRepository(datasource.EventSchema())
				eventService := application.NewNotificationService(eventStore)
				enterpriseService := application.NewEnterpriseService(
					enterpriseRepository,
					identityService,
					employeesRepository,
					eventService,
				)
				ep, err := enterpriseService.Enterprise(params.Context, params.Args["ID"].(string))
				if err != nil {
					return nil, err
				}

				return ep, nil
			},
		},
		"Offices": &graphql.Field{
			Type:        graphql.NewList(OfficeType),
			Description: "Return all the offices where the subject is an employee",
			Args:        StringID,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
				identityService := application.NewIdentityService(userRepository)
				eventStore := repositories.NewEventRepository(datasource.EventSchema())
				eventService := application.NewNotificationService(eventStore)
				enterpriseService := application.NewEnterpriseService(
					enterpriseRepository,
					identityService,
					employeesRepository,
					eventService,
				)
				offices, err := enterpriseService.ProvideEmployeeOffices(params.Context, params.Args["ID"].(string))
				if err != nil {
					return nil, err
				}
				return offices, nil
			},
		},

		"IsValidReceiver": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Return true if the id belongs to a valid receiver or false if not",
			Args:        StringID,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				_, err := mail.ParseAddress(params.Args["ID"].(string))
				if err != nil {
					_, err = enterpriseRepository.Get(params.Context, params.Args["ID"].(string))
					if err != nil {
						return false, err
					}
					return true, nil
				}
				_, err = userRepository.Get(params.Context, params.Args["ID"].(string))
				if err != nil {
					return false, err
				}
				return true, nil
			},
		},
		"UsersForHiring": &graphql.Field{
			Type:        UsersForHiringType,
			Description: "Return a list of users that are not employed by the enterprise",
			Args:        StringIDWPaginationAndQuery,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
				identityService := application.NewIdentityService(userRepository)
				eventStore := repositories.NewEventRepository(datasource.EventSchema())
				eventService := application.NewNotificationService(eventStore)
				enterpriseService := application.NewEnterpriseService(
					enterpriseRepository,
					identityService,
					employeesRepository,
					eventService,
				)
				userID, err := middleware.GetContextPersonID(params.Context)
				if err != nil {
					return nil, err
				}
				limit, ok := params.Args["limit"].(int)
				if !ok {
					limit = 10
				}
				offset, ok := params.Args["offset"].(int)
				if !ok {
					offset = 0
				}
				query, ok := params.Args["query"].(string)
				if !ok {
					query = ""
				}
				users, count, err := enterpriseService.AvailableUsers(params.Context, userID, query, params.Args["ID"].(string), limit, offset)
				if err != nil {
					return nil, err
				}
				result := dto.UsersForHiring{
					Users:         users,
					Count:         count,
					CurrentLimit:  params.Args["limit"].(int),
					CurrentOffset: params.Args["offset"].(int),
				}
				return result, nil
			},
		},
		//################# EMPLOYEE SCHEMA ####################
		"Employee": &graphql.Field{
			Type:        EmployeeType,
			Description: "Return the employee for the given 4 segments employeeID",
			Args:        StringID,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
				identityService := application.NewIdentityService(userRepository)
				employeeService := application.NewEmployeesService(employeesRepository, identityService)
				employee, err := employeeService.Employee(params.Context, params.Args["ID"].(string))
				if err != nil {
					return nil, err
				}
				return employee, nil
			},
		},
		"Employees": &graphql.Field{
			Type:        graphql.NewList(EmployeeType),
			Description: "Return a list of all employees for the given enterprise",
			Args:        StringID,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
				employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
				identityService := application.NewIdentityService(userRepository)
				employeeService := application.NewEmployeesService(employeesRepository, identityService)
				employee, err := employeeService.Employees(params.Context, params.Args["ID"].(string))
				if err != nil {
					return nil, err
				}
				return employee, nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    query,
	Mutation: mutation,
})
