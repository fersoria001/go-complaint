package application_test

import (
	"go-complaint/application"
	"go-complaint/domain"
	"go-complaint/tests"
	"reflect"
	"testing"
)

func TestEventProcessor(t *testing.T) {
	eventProccesor := application.NewEventProcessor()
	if eventProccesor == nil {
		t.Error("Expected eventProccesor to not be nil")
	}
	t.Logf("EventProcessor.subscribedTo: %v", eventProccesor.SubscribedToEventType())
	var interfacePtr *domain.DomainEvent
	var interfaceType = reflect.TypeOf(interfacePtr)
	var fakeEvent domain.DomainEvent = tests.NewFakeEvent()
	t.Logf("reflect.TypeOf(fakeEvent): %v", reflect.TypeOf(fakeEvent))
	t.Logf("reflect.TypeOf(&domain.DomainEvent): %v", interfaceType)
	if eventProccesor.SubscribedToEventType() == interfaceType {
		t.Logf("EventProcessor.subscribedTo: %v", eventProccesor.SubscribedToEventType())
		t.Logf("reflect.TypeOf(&domain.DomainEvent): %v", interfaceType)
	} else {
		t.Errorf("Expected eventProccesor.SubscribedToEventType() to be equal to reflect.TypeOf(&domain.DomainEvent)")
	}
	proccesorKind := eventProccesor.SubscribedToEventType().Kind()
	t.Logf("EventProcessor.subscribedTo.Kind(): %v", proccesorKind)
}