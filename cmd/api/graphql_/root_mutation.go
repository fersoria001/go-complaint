package graphql_

import (
	"go-complaint/cmd/api/graphql_/graphql_arguments"
	"go-complaint/cmd/api/graphql_/resolvers/complaint_resolvers"
	"go-complaint/cmd/api/graphql_/resolvers/employee_resolvers"
	"go-complaint/cmd/api/graphql_/resolvers/enterprise_resolvers"
	"go-complaint/cmd/api/graphql_/resolvers/feedback_resolvers"
	"go-complaint/cmd/api/graphql_/resolvers/public_resolvers"
	"go-complaint/cmd/api/graphql_/resolvers/user_resolvers"

	"github.com/graphql-go/graphql"
)

// root mutation
var mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"Contact": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Send a contact email",
			Args:        graphql_arguments.ContactArgument,
			Resolve:     public_resolvers.ContactResolver,
		},
		//################# USER SCHEMA ####################
		"CreateUser": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Create new user",
			Args:        graphql_arguments.CreateUser,
			Resolve:     user_resolvers.CreateUserResolver,
		},
		"VerifyEmail": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Verify the email from the link sent by email",
			Args:        graphql_arguments.StringID,
			Resolve:     user_resolvers.VerifyEmailResolver,
		},
		"RecoverPassword": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Set a new random password for the user and send it by email",
			Args:        graphql_arguments.StringID,
			Resolve:     user_resolvers.RecoverPasswordResolver,
		},
		"ChangePassword": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Set a new random password for the user and send it by email",
			Args:        graphql_arguments.ChangePassword,
			Resolve:     user_resolvers.ChangePasswordResolver,
		},
		"UpdateUser": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Update the user personal information",
			Args:        graphql_arguments.UpdateUserProfile,
			Resolve:     user_resolvers.UpdateUserResolver,
		},
		"AcceptEnterpriseInvitation": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Accept the invitation to join the enterprise",
			Args:        graphql_arguments.StringID,
			Resolve:     user_resolvers.AcceptEnterpriseInvitationResolver,
		},
		"RejectEnterpriseInvitation": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Reject the invitation to join the enterprise",
			Args:        graphql_arguments.RejectHiringInvitationArgument,
			Resolve:     user_resolvers.RejectEnterpriseInvitationResolver,
		},
		//################# ENTERPRISE SCHEMA ####################
		"CreateEnterprise": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Create a new enterprise",
			Args:        graphql_arguments.CreateAnEnterprise,
			Resolve:     enterprise_resolvers.CreateEnterpriseResolver,
		},
		"UpdateEnterprise": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Update the enterprise",
			Args:        graphql_arguments.UpdateEnterprise,
			Resolve:     enterprise_resolvers.UpdateEnterpriseResolver,
		},
		"InviteToEnterprise": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Invite a user to join the enterprise",
			Args:        graphql_arguments.InviteToEnterpriseArgument,
			Resolve:     enterprise_resolvers.InviteToEnterpriseResolver,
		},
		"HireEmployee": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Hire an employee to the enterprise",
			Args:        graphql_arguments.EnterpriseEventArgument,
			Resolve:     enterprise_resolvers.HireEmployeeResolver,
		},
		"CancelHiringProccess": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Cancel the hiring proccess",
			Args:        graphql_arguments.EnterpriseEventArgument,
			Resolve:     enterprise_resolvers.CancelHiringProccessResolver,
		},
		"FireEmployee": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Fire an employee from the enterprise",
			Args:        graphql_arguments.EmployeeActionArgument,
			Resolve:     enterprise_resolvers.FireEmployeeResolver,
		},
		"ReplyChat": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Reply a chat",
			Args:        graphql_arguments.ChatReplyArgument,
			Resolve:     enterprise_resolvers.ChatReplyResolver,
		},
		"MarkReplyChatAsSeen": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Mark an enterprise chat reply as seen",
			Args:        graphql_arguments.MarkChatReplyAsSeenArgument,
			Resolve:     enterprise_resolvers.MarkChatReplyAsSeenResolver,
		},
		"LeaveEnterprise": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Leave the enterprise",
			Args:        graphql_arguments.EmployeeActionArgument,
			Resolve:     employee_resolvers.LeaveEnterpriseResolver,
		},
		"PromoteEmployee": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Promote an employee",
			Args:        graphql_arguments.PromoteEmployeeArgument,
			Resolve:     enterprise_resolvers.PromoteEmployeeResolver,
		},
		"SendComplaint": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Send a new complaint",
			Args:        graphql_arguments.SendComplaintArgument,
			Resolve:     complaint_resolvers.SendComplaintResolver,
		},
		"ReplyComplaint": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Reply a complaint",
			Args:        graphql_arguments.ReplyComplaintArgument,
			Resolve:     complaint_resolvers.ReplyComplaintResolver,
		},
		"MarkAsSeen": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Mark a complaint reply as seen",
			Args:        graphql_arguments.MarkReplyAsSeen,
			Resolve:     complaint_resolvers.MarkAsSeenResolver,
		},
		"SendForReviewing": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Send a complaint for reviewing",
			Args:        graphql_arguments.StringID,
			Resolve:     complaint_resolvers.SendForReviewingResolver,
		},
		"RateComplaint": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Rate a complaint",
			Args:        graphql_arguments.RateComplaintArgument,
			Resolve:     complaint_resolvers.RateComplaintResolver,
		},
		"CreateFeedback": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Create a feedback",
			Args:        graphql_arguments.FeedbackArgument,
			Resolve:     feedback_resolvers.CreateFeedbackResolver,
		},
		"AddReply": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Add a review to a reply",
			Args:        graphql_arguments.FeedbackArgument,
			Resolve:     feedback_resolvers.AddReplyResolver,
		},
		"RemoveReply": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Remove a review from a reply",
			Args:        graphql_arguments.FeedbackArgument,
			Resolve:     feedback_resolvers.RemoveReplyResolver,
		},
		"AddComment": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Add a comment to a feedback",
			Args:        graphql_arguments.FeedbackArgument,
			Resolve:     feedback_resolvers.AddCommentResolver,
		},
		"DeleteComment": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Delete a comment from a feedback",
			Args:        graphql_arguments.FeedbackArgument,
			Resolve:     feedback_resolvers.DeleteCommentResolver,
		},
		"EndFeedback": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "End the feedback",
			Args:        graphql_arguments.FeedbackArgument,
			Resolve:     feedback_resolvers.EndFeedbackResolver,
		},
		"AnswerFeedback": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Answer a feedback",
			Args:        graphql_arguments.AnswerFeedbackArgument,
			Resolve:     feedback_resolvers.AnswerFeedbackResolver,
		},
		"MarkNotificationAsRead": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Mark a notification as read",
			Args:        graphql_arguments.StringID,
			Resolve:     user_resolvers.MarkNotificationAsReadResolver,
		},
	},
})
