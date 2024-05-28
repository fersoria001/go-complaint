package graphql_

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
	"birthDate": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"phone": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"country": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"county": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"city": &graphql.ArgumentConfig{
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
	"ID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var StringIDWPaginationAndQuery = graphql.FieldConfigArgument{
	"ID": &graphql.ArgumentConfig{
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
var GetComplaint = graphql.FieldConfigArgument{
	"ID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"status": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"limit": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"offset": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}

var SearchTerm = graphql.FieldConfigArgument{
	"term": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var UpdateUserProfile = graphql.FieldConfigArgument{
	"profileIMG": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"firstName": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"lastName": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"phone": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"country": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"county": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"city": &graphql.ArgumentConfig{
		Type: graphql.String,
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

var RateAComplaint = graphql.FieldConfigArgument{
	"complaintID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"rate": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"comment": &graphql.ArgumentConfig{
		Type: graphql.String,
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
	"phone": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"country": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"county": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"city": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"industry": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"foundationDate": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

var UpdateEnterprise = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"logoIMG": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"website": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"email": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"phone": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"country": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"county": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"city": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

var InviteToProject = graphql.FieldConfigArgument{
	"enterpriseName": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"userEmail": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"position": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
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
	"employeeID": &graphql.ArgumentConfig{
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
	"ID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
}
