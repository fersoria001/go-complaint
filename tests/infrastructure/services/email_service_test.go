package services_test

import (
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure"
	"reflect"
	"testing"
)

func TestEmailService_HandleEvent(t *testing.T) {
	service := infrastructure.NewEmailService(reflect.TypeOf(identity.UserCreated{}))
	subscribedTo := service.SubscribedToEventType()
	if subscribedTo.String() != "identity.UserCreated" {
		t.Errorf("Expected identity.UserCreated but got %s", subscribedTo.String())
	}
}
