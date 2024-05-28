package domain

import (
	"reflect"
)

type DomainEventSubscriber struct {
	HandleEvent           func(event DomainEvent) error
	SubscribedToEventType func() reflect.Type
}
