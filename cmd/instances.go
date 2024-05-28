package cmd

import (
	"go-complaint/application"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"sync"
)

var ComplaintService *application.ComplaintService
var EmployeesService *application.EmployeesService
var EnterpriseService *application.EnterpriseService
var FeedbackService *application.FeedbackService
var IdentityService *application.IdentityService
var JWTService *application.JWTService
var NotificationService *application.NotificationService

var instance1 sync.Once
var instance2 sync.Once
var instance3 sync.Once
var instance4 sync.Once
var instance5 sync.Once
var instance6 sync.Once
var instance7 sync.Once

func ComplaintServiceInstance() *application.ComplaintService {
	complaintRepository := repositories.NewComplaintRepository(datasource.ComplaintSchema())
	repliesRepository := repositories.NewReplyRepository(datasource.ComplaintSchema())
	instance1.Do(func() {
		ComplaintService = application.NewComplaintService(
			complaintRepository,
			repliesRepository,
		)
	})
	return ComplaintService
}

func IdentityServiceInstance() *application.IdentityService {
	userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
	instance3.Do(func() {
		IdentityService = application.NewIdentityService(
			userRepository,
		)
	})
	return IdentityService
}

func EmployeesServiceInstance() *application.EmployeesService {
	employeeRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
	instance2.Do(func() {
		EmployeesService = application.NewEmployeesService(
			employeeRepository,
			IdentityServiceInstance(),
		)
	})
	return EmployeesService
}

func NotificationServiceInstance() *application.NotificationService {
	notificationRepository := repositories.NewEventRepository(datasource.EventSchema())
	instance6.Do(func() {
		NotificationService = application.NewNotificationService(
			notificationRepository,
		)
	})
	return NotificationService
}

func EnterpriseServiceInstance() *application.EnterpriseService {
	enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
	employeeRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
	instance4.Do(func() {
		EnterpriseService = application.NewEnterpriseService(
			enterpriseRepository,
			IdentityServiceInstance(),
			employeeRepository,
			NotificationServiceInstance(),
		)
	})
	return EnterpriseService
}

func FeedbackServiceInstance() *application.FeedbackService {
	feedbackRepository := repositories.NewFeedbackRepository(datasource.FeedbackSchema())
	answerRepository := repositories.NewAnswerRepository(datasource.FeedbackSchema())
	instance5.Do(func() {
		FeedbackService = application.NewFeedbackService(
			feedbackRepository,
			answerRepository,
		)
	})
	return FeedbackService
}

func JWTServiceInstance() *application.JWTService {
	instance7.Do(func() {
		JWTService = application.NewJWTService()
	})
	return JWTService
}
