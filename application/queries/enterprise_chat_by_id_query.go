package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type EnterpriseChatByIdQuery struct {
	ChatId string `json:"chatId"`
}

func NewEnterpriseChatByIdQuery(chatId string) *EnterpriseChatByIdQuery {
	return &EnterpriseChatByIdQuery{
		ChatId: chatId,
	}
}

func (q EnterpriseChatByIdQuery) Execute(ctx context.Context) (*dto.Chat, error) {
	chatId, err := uuid.Parse(q.ChatId)
	if err != nil {
		return nil, err
	}

	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Chat").(repositories.ChatRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	chat, err := r.Get(ctx, chatId)
	if err != nil {
		return nil, err
	}
	return dto.NewChat(*chat), nil
}
