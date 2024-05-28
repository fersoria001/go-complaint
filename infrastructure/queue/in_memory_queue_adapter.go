package queue

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

// This should implement Publish() to external bounded context or message broker if necessary
type InMemoryQueueAdapter struct {
	listeners map[reflect.Type][]Listener
	queue     LinkedQueue[QueuedEvent]
	consuming bool
}

// Listener(typeName, callback(domainEvent))
func (imqa *InMemoryQueueAdapter) AddListener(listener Listener) {
	if imqa.consuming {
		return
	}
	if imqa.listeners == nil {
		imqa.listeners = make(map[reflect.Type][]Listener)
	}
	if _, ok := imqa.listeners[listener.SubscribedToEventType()]; !ok {
		imqa.listeners[listener.SubscribedToEventType()] = make([]Listener, 0)
	}
	imqa.listeners[listener.SubscribedToEventType()] = append(imqa.listeners[listener.SubscribedToEventType()], listener)
}
func (imqa *InMemoryQueueAdapter) Enqueue(event *QueuedEvent) {
	imqa.queue.Enqueue(event)
}
func (imqa *InMemoryQueueAdapter) Connect() {
	fmt.Println("Connected to in memory queue: in memory events length: ", imqa.queue.Length())
}
func (imqa *InMemoryQueueAdapter) StartConsuming() error {
	var errs error
	for imqa.queue.Length() > 0 {
		queuedEvent, _ := imqa.queue.Dequeue()
		fmt.Println("Consuming event: ", queuedEvent.typeName)
		for _, listener := range imqa.listeners[queuedEvent.typeName] {
			//in this case i should attach all the listeners in the cmd/api/main.go
			err := listener.HandleEvent(queuedEvent.event)
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
func (imqa *InMemoryQueueAdapter) Disconnect() {
	fmt.Println("Disconnected from in memory queue: in memory events length: ", imqa.queue.Length())
	fmt.Println("Queue will be cleared")
	imqa.queue.Clear()
	fmt.Println("Queue is cleared, in memory events length: ", imqa.queue.Length())
}

func (imqa *InMemoryQueueAdapter) Reset() *InMemoryQueueAdapter {
	if !imqa.consuming {
		imqa.queue.Clear()
		imqa.listeners = nil
	}
	return imqa
}

var inMemoryQueueAdapter *InMemoryQueueAdapter
var onceQueue sync.Once

func InMemoryQueueAdapterInstance() *InMemoryQueueAdapter {
	onceQueue.Do(func() {
		inMemoryQueueAdapter = &InMemoryQueueAdapter{
			queue:     *NewLinkedQueue[QueuedEvent](),
			listeners: nil,
			consuming: false,
		}
	})
	return inMemoryQueueAdapter
}
