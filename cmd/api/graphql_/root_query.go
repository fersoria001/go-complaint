package graphql_

import (
	"go-complaint/cmd/api/graphql_/graphql_arguments"
	"go-complaint/cmd/api/graphql_/graphql_types"
	"go-complaint/cmd/api/graphql_/resolvers/address_resolvers"
	"go-complaint/cmd/api/graphql_/resolvers/complaint_resolvers"
	"go-complaint/cmd/api/graphql_/resolvers/employee_resolvers"
	"go-complaint/cmd/api/graphql_/resolvers/enterprise_resolvers"
	"go-complaint/cmd/api/graphql_/resolvers/feedback_resolvers"
	"go-complaint/cmd/api/graphql_/resolvers/industry_resolvers"
	"go-complaint/cmd/api/graphql_/resolvers/user_resolvers"

	"github.com/graphql-go/graphql"
)

var query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"SignIn": &graphql.Field{
			Type:        graphql_types.JwtTokenType,
			Description: "Get the token for the authenticated user or error",
			Args:        graphql_arguments.LoginUser,
			Resolve:     user_resolvers.SignInResolver,
		},
		"Login": &graphql.Field{
			Type:        graphql.String,
			Description: "Authenticate the user with the token and confirmation code it got the token from the request header",
			Args: graphql.FieldConfigArgument{
				"confirmationCode": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				}},
			Resolve: user_resolvers.LoginResolver,
		},
		"UserDescriptor": &graphql.Field{
			Type:        graphql_types.UserDescriptorType,
			Description: "Get the user descriptor for the current session",
			Resolve:     user_resolvers.UserDescriptorResolver,
		},
		"User": &graphql.Field{
			Type:        graphql_types.UserType,
			Description: "Get a user without private information by it's ID",
			Args:        graphql_arguments.StringID,
			Resolve:     user_resolvers.UserResolver,
		},
		"Enterprise": &graphql.Field{
			Type:        graphql_types.EnterpriseType,
			Description: "Return the enterprise info",
			Args:        graphql_arguments.StringID,
			Resolve:     enterprise_resolvers.EnterpriseResolver,
		},
		"IsEnterpriseNameAvailable": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Check if the enterprise name is available",
			Args:        graphql_arguments.StringID,
			Resolve:     enterprise_resolvers.IsEnterpriseNameAvailableResolver,
		},
		"Employee": &graphql.Field{
			Type:        graphql_types.EmployeeType,
			Description: "Get the employee by it's ID, enterpriseID required for auth for authorization",
			Args:        graphql_arguments.EmployeeActionArgument,
			Resolve:     employee_resolvers.EmployeeResolver,
		},
		"Employees": &graphql.Field{
			Type:        graphql.NewList(graphql_types.EmployeeType),
			Description: "Get the list of employees for the enterprise, enterpriseID required for authorization",
			Args:        graphql_arguments.EmployeeActionArgument,
			Resolve:     employee_resolvers.EmployeesResolver,
		},

		"HiringInvitations": &graphql.Field{
			Type:        graphql.NewList(graphql_types.HiringInvitationType),
			Description: "Get the list of hiring invitations",
			Resolve:     user_resolvers.HiringInvitationsResolver,
		},
		"UsersForHiring": &graphql.Field{
			Type:        graphql.NewList(graphql_types.UserType),
			Description: "Get the list of users for hiring",
			Resolve:     user_resolvers.UsersForHiringResolver,
		},
		"Countries": &graphql.Field{
			Type:        graphql.NewList(graphql_types.CountryType),
			Description: "Get the list of countries",
			Resolve:     address_resolvers.CountriesResolver,
		},
		"CountryStates": &graphql.Field{
			Type:        graphql.NewList(graphql_types.CountyType),
			Description: "find all counties by country ID",
			Args:        graphql_arguments.IntegerID,
			Resolve:     address_resolvers.CountryStatesResolver,
		},
		"Cities": &graphql.Field{
			Type:        graphql.NewList(graphql_types.CityType),
			Description: "find all cities by county ID",
			Args:        graphql_arguments.IntegerID,
			Resolve:     address_resolvers.CitiesResolver,
		},

		"Industries": &graphql.Field{
			Type:        graphql.NewList(graphql_types.IndustryType),
			Description: "Get the list of industries",
			Resolve:     industry_resolvers.IndustriesResolver,
		},

		"FindComplaintReceivers": &graphql.Field{
			Type:        graphql.NewList(graphql_types.ComplaintReceiverType),
			Description: "Find the receivers for a complaint",
			Args:        graphql_arguments.SearchTermArgument,
			Resolve:     complaint_resolvers.FindComplaintReceiversResolver,
		},
		"Complaint": &graphql.Field{
			Type:        graphql_types.ComplaintType,
			Description: "Get a complaint by it's ID",
			Args:        graphql_arguments.StringID,
			Resolve:     complaint_resolvers.ComplaintResolver,
		},
		"PendingComplaintReviews": &graphql.Field{
			Type:        graphql.NewList(graphql_types.PendingReviewType),
			Description: "Get the list of complaints waiting for review",
			Args:        graphql_arguments.StringID,
			Resolve:     complaint_resolvers.PendingComplaintReviewsResolver,
		},
		"ComplaintInbox": &graphql.Field{
			Type:        graphql_types.ComplaintListType,
			Description: "Get the list of inbox complaints for the current user",
			Args:        graphql_arguments.ComplaintListArgument,
			Resolve:     complaint_resolvers.ComplaintInboxResolver,
		},
		"ComplaintsSent": &graphql.Field{
			Type:        graphql_types.ComplaintListType,
			Description: "Get the list of sent complaints for the current user",
			Args:        graphql_arguments.ComplaintListArgument,
			Resolve:     complaint_resolvers.ComplaintsSentResolver,
		},
		"ComplaintHistory": &graphql.Field{
			Type:        graphql_types.ComplaintListType,
			Description: "Get the list of inbox complaints for the current user",
			Args:        graphql_arguments.ComplaintListArgument,
			Resolve:     complaint_resolvers.ComplaintHistoryResolver,
		},
		"SolvedComplaints": &graphql.Field{
			Type:        graphql.NewList(graphql_types.ComplaintType),
			Description: "Get the list of solved complaints for the current user",
			Args:        graphql_arguments.StringID,
			Resolve:     employee_resolvers.SolvedComplaintsResolver,
		},
		"FeedbackByComplaintID": &graphql.Field{
			Type:        graphql.NewList(graphql_types.FeedbackType),
			Description: "Get the feedback by complaint ID",
			Args:        graphql_arguments.StringID,
			Resolve:     feedback_resolvers.FeedbackByComplaintIDResolver,
		},
		"FeedbackByReviewerID": &graphql.Field{
			Type:        graphql.NewList(graphql_types.FeedbackType),
			Description: "Get the reviews by reviewer ID",
			Args:        graphql_arguments.StringID,
			Resolve:     feedback_resolvers.FeedbackByReviewerIDResolver,
		},
		"FeedbackByRevieeID": &graphql.Field{
			Type:        graphql.NewList(graphql_types.FeedbackType),
			Description: "Get the reviews by reviee ID",
			Args:        graphql_arguments.StringID,
			Resolve:     feedback_resolvers.FeedbackByRevieweeIDResolver,
		},
	},
})
