package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
)

type Chat struct {
	ID      string       `json:"id"`
	Replies []*ChatReply `json:"replies"`
}

func NewChat(domain enterprise.Chat) *Chat {
	replies := make([]*ChatReply, 0)
	for _, reply := range domain.Replies() {
		replies = append(replies, NewChatReply(*reply))
	}
	return &Chat{
		ID:      domain.ID().String(),
		Replies: replies,
	}
}

type ChatReply struct {
	ID        string `json:"id"`
	ChatID    string `json:"chatID"`
	User      User   `json:"user"`
	Content   string `json:"content"`
	Seen      bool   `json:"seen"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func NewChatReply(domain enterprise.Reply) *ChatReply {
	return &ChatReply{
		ID:        domain.ID().String(),
		ChatID:    domain.ChatID().String(),
		User:      NewUser(domain.User()),
		Content:   domain.Content(),
		Seen:      domain.Seen(),
		CreatedAt: common.StringDate(domain.CreatedAt()),
		UpdatedAt: common.StringDate(domain.UpdatedAt()),
	}
}
