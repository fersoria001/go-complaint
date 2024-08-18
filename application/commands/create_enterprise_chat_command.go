package commands

import (
	"context"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/finders/find_chats"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type CreateEnterpriseChatCommand struct {
	EnterpriseId string `json:"enterpriseId"`
	SenderId     string `json:"senderId"`
	ReceiverId   string `json:"receiverId"`
}

func NewCreateEnterpriseChatCommand(enterpriseId, senderId, receiverId string) *CreateEnterpriseChatCommand {
	return &CreateEnterpriseChatCommand{
		EnterpriseId: enterpriseId,
		SenderId:     senderId,
		ReceiverId:   receiverId,
	}
}

func (c CreateEnterpriseChatCommand) Execute(ctx context.Context) error {
	enterpriseId, err := uuid.Parse(c.EnterpriseId)
	if err != nil {
		return err
	}
	senderId, err := uuid.Parse(c.SenderId)
	if err != nil {
		return err
	}
	receiverId, err := uuid.Parse(c.ReceiverId)
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
	dbChat, _ := chatRepository.Find(ctx, find_chats.ByRecipientOneTwoAndEnterpriseId(
		senderId, receiverId, enterpriseId,
	))
	if dbChat != nil {
		return ErrChatAlreadyExists
	}
	sender, err := recipientRepository.Get(ctx, senderId)
	if err != nil {
		return err
	}
	receiver, err := recipientRepository.Get(ctx, receiverId)
	if err != nil {
		return err
	}
	chat := enterprise.NewChatEntity(uuid.New(), enterpriseId, *sender, *receiver)
	err = chatRepository.Save(ctx, chat)
	if err != nil {
		return err
	}
	return nil
}
