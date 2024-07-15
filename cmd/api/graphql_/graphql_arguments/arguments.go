package graphql_arguments

import "github.com/graphql-go/graphql"

var CreateUser = graphql.FieldConfigArgument{
	"email": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"password": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"firstName": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"lastName": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"gender": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"pronoun": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"birthDate": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"phone": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"countryId": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"countryStateId": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"cityId": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
}

var ChangePassword = graphql.FieldConfigArgument{
	"oldPassword": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"newPassword": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var LoginUser = graphql.FieldConfigArgument{
	"email": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"password": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"rememberMe": &graphql.ArgumentConfig{
		Type: graphql.Boolean,
	},
}

var StringID = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}
var MarkReplyAsSeen = graphql.FieldConfigArgument{
	"complaintID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"ids": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

var StringIDWithPaginationAndQuery = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"query": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"limit": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"offset": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}

var UpdateUserProfile = graphql.FieldConfigArgument{
	"updateType": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"value": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"numberValue": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}

var CreateAComplaint = graphql.FieldConfigArgument{
	"authorID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"receiverID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"title": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"description": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"content": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var ReplyToComplaint = graphql.FieldConfigArgument{
	"complaintID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"profileIMG": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"fullName": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"senderID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"body": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}
var MarkAsReviewable = graphql.FieldConfigArgument{
	"complaintID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"enterpriseID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"assistantID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}
var CloseComplaint = graphql.FieldConfigArgument{
	"complaintID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"closeRequesterID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}
var CreateAnEnterprise = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"website": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"email": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"phoneCode": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"phone": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"countryID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"countryStateID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"cityID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"industryID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"foundationDate": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var UpdateEnterprise = graphql.FieldConfigArgument{
	"updateType": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"enterpriseID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"value": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"numberValue": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}

var HireFromEvent = graphql.FieldConfigArgument{
	"eventID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var SendHiringInvitation = graphql.FieldConfigArgument{
	"managerID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"userEmail": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var EndHiringProcess = graphql.FieldConfigArgument{
	"pendingEventID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"enterpriseName": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"accepted": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Boolean),
	},
}

var FireEmployee = graphql.FieldConfigArgument{
	"enterpriseName": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"employeeID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var CreateAFeedback = graphql.FieldConfigArgument{
	"complaintID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"reviewerID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"reviewedID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"reviewerName": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"replyID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"senderID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"senderIMG": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"senderName": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"body": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"createdAt": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"read": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Boolean),
	},
	"readAt": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"updatedAt": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"comment": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"isEnterprise": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Boolean),
	},
	"enterpriseID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var AnswerAFeedback = graphql.FieldConfigArgument{
	"complaintID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"senderID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"senderIMG": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"senderName": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"body": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"createdAt": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"read": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Boolean),
	},
	"readAt": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"updatedAt": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var MarkAnswerAsRead = graphql.FieldConfigArgument{
	"answerID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var IntegerID = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
}
