package domain

import (
	"context"
	"errors"
	"reflect"
	"sync"
)

// Package domain
type DomainEventPublisher struct {
	subscribers []DomainEventSubscriber
	publishing  bool
	mu          sync.Mutex
}

func (dep *DomainEventPublisher) Subscribe(subscriber DomainEventSubscriber) {
	dep.mu.Lock()
	if dep.publishing {
		return
	}
	if dep.subscribers == nil {
		var subscribers = make([]DomainEventSubscriber, 0)
		dep.subscribers = subscribers
	}
	dep.subscribers = append(dep.subscribers, subscriber)
	dep.mu.Unlock()
}

func (dep *DomainEventPublisher) Publish(ctx context.Context, event DomainEvent) error {
	if dep.publishing {
		return nil
	}
	var (
		errs          error
		err           error
		eventType     = reflect.TypeOf(event)
		interfacePtr  *DomainEvent
		interfaceType = reflect.TypeOf(interfacePtr)
	)
	dep.publishing = true
	defer func() {
		dep.publishing = false
	}()
	for _, subscriber := range dep.subscribers {
		if subscriber.SubscribedToEventType() == eventType ||
			subscriber.SubscribedToEventType() == interfaceType {
			err = subscriber.HandleEvent(event)
			if err != nil {
				errs = errors.Join(errs, err)
			}
		}
	}
	if errs != nil {
		return errs
	}
	return nil
}

// Reset resets the domain event publisher.
func (dep *DomainEventPublisher) Reset() *DomainEventPublisher {
	dep.mu.Lock()
	if !dep.publishing {
		dep.subscribers = nil
	}
	dep.mu.Unlock()
	return dep
}

var domainEventPublisher *DomainEventPublisher
var once sync.Once

func DomainEventPublisherInstance() *DomainEventPublisher {
	once.Do(func() {
		domainEventPublisher = &DomainEventPublisher{}
	})
	return domainEventPublisher
}
