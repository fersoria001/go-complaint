package graphql_types

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
		"gender": &graphql.Field{
			Type: graphql.String,
		},
		"pronoun": &graphql.Field{
			Type: graphql.String,
		},
		"loginDate": &graphql.Field{
			Type: graphql.String,
		},
		"ip": &graphql.Field{
			Type: graphql.String,
		},
		"device": &graphql.Field{
			Type: graphql.String,
		},
		"geolocation": &graphql.Field{
			Type: GeolocationType,
		},
		"grantedAuthorities": &graphql.Field{
			Type: graphql.NewList(GrantedAuthorityType),
		},
	},
})
var GrantedAuthorityType = graphql.NewObject(graphql.ObjectConfig{
	Name: "GrantedAuthority",
	Fields: graphql.Fields{
		"enterpriseID": &graphql.Field{
			Type: graphql.String,
		},
		"authority": &graphql.Field{
			Type: graphql.String,
		},
	},
})
var GeolocationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Geolocation",
	Fields: graphql.Fields{
		"latitude": &graphql.Field{
			Type: graphql.Float,
		},
		"longitude": &graphql.Field{
			Type: graphql.Float,
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
		"gender": &graphql.Field{
			Type: graphql.String,
		},
		"pronoun": &graphql.Field{
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
		"status": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var CountryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Country",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"phoneCode": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var CountyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "County",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
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
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"countryCode": &graphql.Field{
			Type: graphql.String,
		},
		"latitude": &graphql.Field{
			Type: graphql.Float,
		},
		"longitude": &graphql.Field{
			Type: graphql.Float,
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
		"employees": &graphql.Field{
			Type: graphql.NewList(EmployeeType),
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
		"id": &graphql.Field{
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
		"complaintsSolved": &graphql.Field{
			Type: graphql.Int,
		},
		"complaintsSolvedIds": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"complaintsRated": &graphql.Field{
			Type: graphql.Int,
		},
		"complaintsRatedIDs": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"complaintsFeedbacked": &graphql.Field{
			Type: graphql.Int,
		},
		"complaintsFeedbackedIDs": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"feedbackReceived": &graphql.Field{
			Type: graphql.Int,
		},
		"feedbackReceivedIDs": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"hireInvitationsSent": &graphql.Field{
			Type: graphql.Int,
		},
		"employeesHired": &graphql.Field{
			Type: graphql.Int,
		},
		"employeesFired": &graphql.Field{
			Type: graphql.Int,
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
		"enterprisePhone": &graphql.Field{
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
