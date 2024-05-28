package queue

import (
	"go-complaint/domain"
	"reflect"
)

type InmediateForwardingSubscriber struct {
	subscribedTo reflect.Type
}

func NewInmediateForwardingSubscriber(subscribeTo reflect.Type) *InmediateForwardingSubscriber {
	return &InmediateForwardingSubscriber{
		subscribedTo: subscribeTo,
	}
}

func (ifs *InmediateForwardingSubscriber) SubscribedToEventType() reflect.Type {
	return ifs.subscribedTo
}

func (ifs *InmediateForwardingSubscriber) HandleEvent(event domain.DomainEvent) error {
	queuedEvent := NewQueuedEvent(ifs.SubscribedToEventType(), event)
	InMemoryQueueAdapterInstance().Enqueue(queuedEvent)
	return nil
}
