package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
)

type Chat struct {
	ID           string       `json:"id"`
	EnterpriseId string       `json:"enterpriseId"`
	RecipientOne *Recipient   `json:"recipientOne"`
	RecipientTwo *Recipient   `json:"recipientTwo"`
	Replies      []*ChatReply `json:"replies"`
}

func NewChat(domain enterprise.Chat) *Chat {
	replies := make([]*ChatReply, 0)
	for _, reply := range domain.Replies() {
		replies = append(replies, NewChatReply(*reply))
	}
	return &Chat{
		ID:           domain.Id().String(),
		EnterpriseId: domain.EnterpriseId().String(),
		RecipientOne: NewRecipient(domain.RecipientOne()),
		RecipientTwo: NewRecipient(domain.RecipientTwo()),
		Replies:      replies,
	}
}

type ChatReply struct {
	Id        string     `json:"id"`
	ChatId    string     `json:"chatId"`
	Sender    *Recipient `json:"sender"`
	Content   string     `json:"content"`
	Seen      bool       `json:"seen"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
}

func NewChatReply(domain enterprise.Reply) *ChatReply {
	return &ChatReply{
		Id:        domain.Id().String(),
		ChatId:    domain.ChatId().String(),
		Sender:    NewRecipient(domain.Sender()),
		Content:   domain.Content(),
		Seen:      domain.Seen(),
		CreatedAt: common.StringDate(domain.CreatedAt()),
		UpdatedAt: common.StringDate(domain.UpdatedAt()),
	}
}
