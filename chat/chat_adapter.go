package chat

import (
	"sync"
)

type ChatAdapter struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mu         sync.Mutex
}

func NewChatAdapter() *ChatAdapter {
	return &ChatAdapter{
		clients:    make(map[*Client]bool, 0),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (cas *ChatAdapter) Register(client *Client) {
	cas.register <- client
}

func (cas *ChatAdapter) Run() {
	for {
		select {
		case client := <-cas.register:
			cas.mu.Lock()
			if _, ok := cas.clients[client]; !ok {
				cas.clients[client] = true
			}
			cas.mu.Unlock()
		case client := <-cas.unregister:
			cas.mu.Lock()
			if _, ok := cas.clients[client]; ok {
				delete(cas.clients, client)
				close(client.send)
			}
			cas.mu.Unlock()
		case message := <-cas.broadcast:
			for client := range cas.clients {
				select {
				case client.send <- message:
				default:
				}
			}
		}
	}
}
