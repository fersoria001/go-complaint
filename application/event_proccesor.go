package application

import (
	"context"
	"go-complaint/domain"
	"go-complaint/dto"
	"reflect"
	"sync"
)

// Package application
var eventProcesorInstance *EventProcessor
var eventProcessorOnce sync.Once

func EventProcessorInstance() *EventProcessor {
	eventProcessorOnce.Do(func() {
		eventProcesorInstance = NewEventProcessor()
	})
	return eventProcesorInstance
}

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

func (ep *EventProcessor) ResetDomainEventPublisher() {
	domain.DomainEventPublisherInstance().Reset()
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent:           eventProcesorInstance.HandleEvent,
			SubscribedToEventType: eventProcesorInstance.SubscribedToEventType,
		},
	)
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
	err = eventStore.Save(ctx, storedEvent)

	if err != nil {
		return err
	}
	return nil
}

func (ep *EventProcessor) SubscribedToEventType() reflect.Type {
	return ep.subscribedTo
}
