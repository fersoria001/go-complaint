package application

import (
	"log"
	"slices"
	"sync"
)

var applicationMessagePublisherInstance ApplicationMessagePublisher
var applicationMessagePublisherOnce sync.Once

func ApplicationMessagePublisherInstance() *ApplicationMessagePublisher {
	applicationMessagePublisherOnce.Do(func() {
		applicationMessagePublisherInstance = ApplicationMessagePublisher{
			subscribers:   make([]*Subscriber, 0),
			subscribeCh:   make(chan *Subscriber),
			unsubscribeCh: make(chan *Subscriber),
			publishCh:     make(chan ApplicationMessage),
		}
	})
	go applicationMessagePublisherInstance.Start()
	return &applicationMessagePublisherInstance
}

type ApplicationMessage struct {
	id       string
	dataType string
	value    any
}

func (m ApplicationMessage) Id() string {
	return m.id
}

func (m ApplicationMessage) DataType() string {
	return m.dataType
}

func (m ApplicationMessage) Value() any {
	return m.value
}

func NewApplicationMessage(id, dataType string, value any) ApplicationMessage {
	return ApplicationMessage{
		id:       id,
		dataType: dataType,
		value:    value,
	}
}

type Subscriber struct {
	Id     string
	UserId string
	Send   chan ApplicationMessage
}

type ApplicationMessagePublisher struct {
	subscribers   []*Subscriber
	subscribeCh   chan *Subscriber
	unsubscribeCh chan *Subscriber
	publishCh     chan ApplicationMessage
	mu            sync.Mutex
}

func (p *ApplicationMessagePublisher) Start() {
	for {
		select {
		case sub := <-p.subscribeCh:
			p.mu.Lock()
			p.subscribers = append(p.subscribers, sub)
			for _, v := range p.subscribers {
				log.Printf("send %s connection status to all subscriber", sub.UserId)
				v.Send <- NewApplicationMessage(v.Id, "subscriber_connected", sub)
				log.Printf("send all connected users connection status to %s", sub.UserId)
				sub.Send <- NewApplicationMessage(v.Id, "subscriber_connected", v)
			}
			p.mu.Unlock()
		case sub := <-p.unsubscribeCh:
			p.mu.Lock()
			p.subscribers = slices.DeleteFunc(p.subscribers, func(e *Subscriber) bool {
				if e.UserId != "" {
					return e.UserId == sub.UserId
				}
				return e.Id == sub.Id
			})
			for _, v := range p.subscribers {
				log.Printf("sending %s disconnected status to all remaining subscribers", sub.UserId)
				v.Send <- NewApplicationMessage(v.Id, "subscriber_disconnected", sub)
			}
			p.mu.Unlock()
		case m := <-p.publishCh:
			p.mu.Lock()
			for _, v := range p.subscribers {
				if v.Id == m.id {
					v.Send <- m
				}
			}
			p.mu.Unlock()
		}
	}
}

func (p *ApplicationMessagePublisher) Publish(m ApplicationMessage) {
	p.publishCh <- m
}

func (p *ApplicationMessagePublisher) Unsubscribe(sub *Subscriber) {
	p.unsubscribeCh <- sub
}

func (p *ApplicationMessagePublisher) Subscribe(subscriber *Subscriber) {
	p.subscribeCh <- subscriber
}
