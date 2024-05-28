package graphql_

import "github.com/graphql-go/graphql"

var UserDescriptorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserDescriptor",
	Fields: graphql.Fields{
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"fullName": &graphql.Field{
			Type: graphql.String,
		},
		"profileIMG": &graphql.Field{
			Type: graphql.String,
		},
		"ip": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UsersForHiringType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UsersForHiring",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type: graphql.NewList(UserType),
		},
		"count": &graphql.Field{
			Type: graphql.Int,
		},
		"currentLimit": &graphql.Field{
			Type: graphql.Int,
		},
		"currentOffset": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"profileIMG": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"firstName": &graphql.Field{
			Type: graphql.String,
		},
		"lastName": &graphql.Field{
			Type: graphql.String,
		},
		"age": &graphql.Field{
			Type: graphql.Int,
		},
		"phone": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: AddressType,
		},
	},
})

var CountryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Country",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var CountyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "County",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var CityType = graphql.NewObject(graphql.ObjectConfig{
	Name: "City",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var PhoneCodeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PhoneCode",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"code": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var JwtTokenType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JwtToken",
	Fields: graphql.Fields{
		"token": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})
var AddressType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Address",
	Fields: graphql.Fields{
		"country": &graphql.Field{
			Type: graphql.String,
		},
		"county": &graphql.Field{
			Type: graphql.String,
		},
		"city": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var EnterpriseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Enterprise",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"bannerIMG": &graphql.Field{
			Type: graphql.String,
		},
		"logoIMG": &graphql.Field{
			Type: graphql.String,
		},
		"website": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"phone": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: AddressType,
		},
		"industry": &graphql.Field{
			Type: graphql.String,
		},
		"foundationDate": &graphql.Field{
			Type: graphql.String,
		},
		"ownerID": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var IndustryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Industry",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.String,
		},
		"profileIMG": &graphql.Field{
			Type: graphql.String,
		},
		"firstName": &graphql.Field{
			Type: graphql.String,
		},
		"lastName": &graphql.Field{
			Type: graphql.String,
		},
		"age": &graphql.Field{
			Type: graphql.Int,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"phone": &graphql.Field{
			Type: graphql.String,
		},
		"hiringDate": &graphql.Field{
			Type: graphql.String,
		},
		"approvedHiring": &graphql.Field{
			Type: graphql.Boolean,
		},
		"approvedHiringAt": &graphql.Field{
			Type: graphql.String,
		},
		"position": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var OfficeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Office",
	Fields: graphql.Fields{
		"employeeID": &graphql.Field{
			Type: graphql.String,
		},
		"employeeFirstName": &graphql.Field{
			Type: graphql.String,
		},
		"employeePosition": &graphql.Field{
			Type: graphql.String,
		},
		"enterpriseLogoIMG": &graphql.Field{
			Type: graphql.String,
		},
		"enterpriseName": &graphql.Field{
			Type: graphql.String,
		},
		"enterpriseWebsite": &graphql.Field{
			Type: graphql.String,
		},
		"enterpriseEmail": &graphql.Field{
			Type: graphql.String,
		},
		"enterpriseIndustry": &graphql.Field{
			Type: graphql.String,
		},
		"enterpriseAddress": &graphql.Field{
			Type: AddressType,
		},
		"ownerFullName": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var Receiver = graphql.NewObject(graphql.ObjectConfig{
	Name: "Receiver",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.String,
		},
		"fullName": &graphql.Field{
			Type: graphql.String,
		},
		"IMG": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var ComplaintTypeList = graphql.NewObject(graphql.ObjectConfig{
	Name: "ComplaintList",
	Fields: graphql.Fields{
		"complaints": &graphql.Field{
			Type: graphql.NewList(ComplaintType),
		},
		"count": &graphql.Field{
			Type: graphql.Int,
		},
		"currentLimit": &graphql.Field{
			Type: graphql.Int,
		},
		"currentOffset": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var ComplaintType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Complaint",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"authorID": &graphql.Field{
			Type: graphql.String,
		},
		"authorFullName": &graphql.Field{
			Type: graphql.String,
		},
		"authorProfileIMG": &graphql.Field{
			Type: graphql.String,
		},
		"receiverID": &graphql.Field{
			Type: graphql.String,
		},
		"receiverFullName": &graphql.Field{
			Type: graphql.String,
		},
		"receiverProfileIMG": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: MessageType,
		},
		"rating": &graphql.Field{
			Type: RatingType,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
		"updatedAt": &graphql.Field{
			Type: graphql.String,
		},
		"replies": &graphql.Field{
			Type: graphql.NewList(ReplyType),
		},
	},
})

var MessageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Message",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"body": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var RatingType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Rating",
	Fields: graphql.Fields{
		"rate": &graphql.Field{
			Type: graphql.Int,
		},
		"comment": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var ReplyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Reply",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"complaintID": &graphql.Field{
			Type: graphql.ID,
		},
		"senderIMG": &graphql.Field{
			Type: graphql.String,
		},
		"senderName": &graphql.Field{
			Type: graphql.String,
		},
		"body": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
		"read": &graphql.Field{
			Type: graphql.Boolean,
		},
		"readAt": &graphql.Field{
			Type: graphql.String,
		},
		"updatedAt": &graphql.Field{
			Type: graphql.String,
		},
	},
})
