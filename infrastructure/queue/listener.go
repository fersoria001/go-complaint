package queue

import (
	"go-complaint/domain"
	"reflect"
)

// For now I will avoid generic struct because
// I can't use anonymous inline override of methods
type Listener struct {
	eventType reflect.Type
	callback  func(event domain.DomainEvent)
}

func NewListener(eventType reflect.Type, callback func(event domain.DomainEvent)) Listener {
	return Listener{
		eventType: eventType,
		callback:  callback,
	}
}

func (l *Listener) SubscribedToEventType() reflect.Type {
	return l.eventType
}

func (l *Listener) HandleEvent(event domain.DomainEvent) error {
	l.callback(event)
	return nil
}
