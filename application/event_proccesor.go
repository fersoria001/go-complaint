package application

import (
	"context"
	"go-complaint/domain"
	"go-complaint/dto"
	"reflect"
)

// Package application

type EventProcessor struct {
}

func (ep *EventProcessor) HandleEvent(event domain.DomainEvent) error {
	var err error
	ctx := context.Background()
	eventStore := EventStore{}
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
	var interfacePtr *domain.DomainEvent
	interfaceType := reflect.TypeOf(interfacePtr)
	return interfaceType
}
