package application

import "sync"

var applicationMessagePublisherInstance ApplicationMessagePublisher
var applicationMessagePublisherOnce sync.Once

func ApplicationMessagePublisherInstance() ApplicationMessagePublisher {
	applicationMessagePublisherOnce.Do(func() {
		applicationMessagePublisherInstance = ApplicationMessagePublisher{
			subscribers:   make(map[string]*Subscriber),
			subscribeCh:   make(chan *Subscriber),
			unsubscribeCh: make(chan string),
			publishCh:     make(chan ApplicationMessage),
		}
	})
	go applicationMessagePublisherInstance.Start()
	return applicationMessagePublisherInstance
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
	Id   string
	Send chan ApplicationMessage
}

type ApplicationMessagePublisher struct {
	subscribers   map[string]*Subscriber
	subscribeCh   chan *Subscriber
	unsubscribeCh chan string
	publishCh     chan ApplicationMessage
}

func (p *ApplicationMessagePublisher) Start() {
	for {
		select {
		case v := <-p.subscribeCh:
			p.subscribers[v.Id] = v
		case v := <-p.unsubscribeCh:
			delete(p.subscribers, v)
		case m := <-p.publishCh:
			for k, v := range p.subscribers {
				if k == m.id {
					v.Send <- m
				}
			}
		}
	}
}

func (p *ApplicationMessagePublisher) Publish(m ApplicationMessage) {
	p.publishCh <- m
}

func (p *ApplicationMessagePublisher) Unsubscribe(id string) {
	p.unsubscribeCh <- id
}

func (p *ApplicationMessagePublisher) Subscribe(subscriber *Subscriber) {
	p.subscribeCh <- subscriber
}
