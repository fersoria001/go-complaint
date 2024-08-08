package application

import "sync"

var applicationMessagePublisherInstance ApplicationMessagePublisher
var applicationMessagePublisherOnce sync.Once

func ApplicationMessagePublisherInstance() *ApplicationMessagePublisher {
	applicationMessagePublisherOnce.Do(func() {
		applicationMessagePublisherInstance = ApplicationMessagePublisher{
			subscribers:   make(map[string]*Subscriber),
			subscribeCh:   make(chan *Subscriber),
			unsubscribeCh: make(chan string),
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
	Id   string
	Send chan ApplicationMessage
}

type ApplicationMessagePublisher struct {
	subscribers   map[string]*Subscriber
	subscribeCh   chan *Subscriber
	unsubscribeCh chan string
	publishCh     chan ApplicationMessage
	mu            sync.Mutex
}

func (p *ApplicationMessagePublisher) Start() {
	for {
		select {
		case sub := <-p.subscribeCh:
			p.mu.Lock()
			p.subscribers[sub.Id] = sub
			p.mu.Unlock()
			for k, sub := range p.subscribers {
				sub.Send <- NewApplicationMessage(k, "subscriber_connected", sub.Id)
			}
		case unsubId := <-p.unsubscribeCh:
			p.mu.Lock()
			delete(p.subscribers, unsubId)
			p.mu.Unlock()
			for k, sub := range p.subscribers {
				sub.Send <- NewApplicationMessage(k, "subscriber_disconnected", unsubId)
			}
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

func (p *ApplicationMessagePublisher) ApplicationSubscribers() []string {
	ids := make([]string, 0, len(p.subscribers))
	for k := range p.subscribers {
		ids = append(ids, k)
	}
	return ids
}
