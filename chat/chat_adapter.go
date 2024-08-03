package chat

type ChatAdapter struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
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

func (cas *ChatAdapter) registerClient(client *Client) {
	if _, ok := cas.clients[client]; !ok {
		cas.clients[client] = true
	}
}

func (cas *ChatAdapter) unregisterClient(client *Client) {
	if _, ok := cas.clients[client]; !ok {
		delete(cas.clients, client)
		close(client.send)
	}
}

func (cas *ChatAdapter) Run() {
	for {
		select {
		case client := <-cas.register:
			cas.registerClient(client)
		case client := <-cas.unregister:
			cas.unregisterClient(client)
		case message := <-cas.broadcast:
			for client := range cas.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(cas.clients, client)
				}
			}
		}
	}
}
