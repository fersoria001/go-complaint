package queue

import (
	"go-complaint/domain"
	"reflect"
)

type QueuedEvent struct {
	typeName reflect.Type
	event    domain.DomainEvent
}

func NewQueuedEvent(typeName reflect.Type, event domain.DomainEvent) *QueuedEvent {
	return &QueuedEvent{
		typeName: typeName,
		event:    event,
	}
}
