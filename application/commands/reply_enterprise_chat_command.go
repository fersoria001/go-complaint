package commands

import (
	"context"
	"go-complaint/domain/model/enterprise"
	"go-complaint/dto"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type ReplyEnterpriseChatCommand struct {
	ChatId   string `json:"chatId"`
	SenderId string `json:"senderId"`
	Content  string `json:"content"`
}

func (c ReplyEnterpriseChatCommand) Execute(ctx context.Context) error {
	chatId, err := uuid.Parse(c.ChatId)
	if err != nil {
		return err
	}
	senderId, err := uuid.Parse(c.SenderId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	chatRepository, ok := reg.Get("Chat").(repositories.ChatRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	sender, err := recipientRepository.Get(ctx, senderId)
	if err != nil {
		return err
	}
	dbChat, err := chatRepository.Get(ctx, chatId)
	if err != nil {
		return err
	}
	reply := enterprise.NewReplyEntity(uuid.New(), dbChat.Id(), *sender, c.Content)
	err = dbChat.Reply(reply)
	if err != nil {
		return err
	}
	err = chatRepository.Update(ctx, dbChat)
	if err != nil {
		return err
	}
	cache.InMemoryInstance().Set(c.ChatId, *dto.NewChatReply(*reply))
	return nil
}
