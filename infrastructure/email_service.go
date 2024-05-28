package infrastructure

import (
	"fmt"
	"go-complaint/domain"
	"go-complaint/infrastructure/persistence/models"
	"reflect"
)

// Package infrastructure
// EmailService struct Implementation of DomainEventSubscriber
type EmailService struct {
	//if subscribed to more than one I should use a Mediator pattern
	subscribedTo reflect.Type
}

func NewEmailService(subscribeTo reflect.Type) *EmailService {
	eis := &EmailService{}
	//this always have to be a value
	eis.subscribedTo = subscribeTo
	return eis
}

func (es *EmailService) HandleEvent(e domain.DomainEvent) error {
	fmt.Println("Handling event: ")
	//here you can dispatch the correct type of body to deserialize
	return nil
}

func (eis *EmailService) SubscribedToEventType() reflect.Type {
	return eis.subscribedTo
}

// here you can dispatch the correct type of body to deserialize
func (es *EmailService) SendUserEmail(user models.Event) error {
	fmt.Println("Sending email to: ")
	return nil
}

/*
In the case of the HiringInvitationSent
You've got to get all events and equals with the HiringInvitationSent
to build the email template ?link=eventID.String()

*/
