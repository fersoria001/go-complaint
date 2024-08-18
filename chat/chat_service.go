package chat

import (
	"sync"
)

type ChatService struct {
	chats map[string]*ChatAdapter
	mu    sync.Mutex
}

func (c *ChatService) GetChatOnlineClientsIds(chatId string) []string {
	ids := make([]string, 0)
	c.mu.Lock()
	svc, ok := c.chats[chatId]
	if ok {
		for k := range svc.clients {
			if k.id != "" {
				ids = append(ids, k.id)
			}
		}
	}
	c.mu.Unlock()
	return ids
}

func (c *ChatService) GetChat(chatId string) *ChatAdapter {
	c.mu.Lock()
	svc, ok := c.chats[chatId]
	if !ok {
		newInstance := NewChatAdapter()
		go newInstance.Run()
		c.chats[chatId] = newInstance
		svc = newInstance
	}
	c.mu.Unlock()
	return svc
}

func ChatServiceInstance() *ChatService {
	chatsOnce.Do(func() {
		chatsInstance = ChatService{
			chats: make(map[string]*ChatAdapter),
			mu:    sync.Mutex{},
		}
	})
	return &chatsInstance
}

var chatsInstance ChatService
var chatsOnce sync.Once
