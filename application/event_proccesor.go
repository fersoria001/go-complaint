package application

import (
	"context"
	"go-complaint/domain"
	"go-complaint/dto"
	"reflect"
)

// Package application
// Find a way to use it as an aspect @beforeeach //todo()//
type EventProcessor struct {
	subscribedTo reflect.Type
}

func NewEventProcessor() *EventProcessor {
	var interfacePtr *domain.DomainEvent
	interfaceType := reflect.TypeOf(interfacePtr)
	return &EventProcessor{
		subscribedTo: interfaceType,
	}
}
func (ep *EventProcessor) Subscriber() domain.DomainEventSubscriber {
	return domain.DomainEventSubscriber{
		HandleEvent:           ep.HandleEvent,
		SubscribedToEventType: ep.SubscribedToEventType,
	}
}

func (ep *EventProcessor) HandleEvent(event domain.DomainEvent) error {
	var err error
	ctx := context.Background()
	eventStore, err := EventStoreInstance()
	if err != nil {
		return err
	}
	storedEvent, err := dto.NewStoredEvent(event)
	if err != nil {
		return err
	}
	err = eventStore.Save(ctx, *storedEvent)
	if err != nil {
		return err
	}
	return nil
}

func (ep *EventProcessor) SubscribedToEventType() reflect.Type {
	return ep.subscribedTo
}
