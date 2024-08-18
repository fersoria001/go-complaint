package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_chats"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type EnterpriseChatByRecipientsIdAndEnterpriseIdQuery struct {
	EnterpriseId   string `json:"enterpriseId"`
	RecipientOneId string `json:"recipientOneId"`
	RecipientTwoId string `json:"recipientTwoId"`
}

func NewEnterpriseChatByRecipientsIdAndEnterpriseIdQuery(enterpriseId, recipientOneId, recipientTwoId string) *EnterpriseChatByRecipientsIdAndEnterpriseIdQuery {
	return &EnterpriseChatByRecipientsIdAndEnterpriseIdQuery{
		EnterpriseId:   enterpriseId,
		RecipientOneId: recipientOneId,
		RecipientTwoId: recipientTwoId,
	}
}

func (q EnterpriseChatByRecipientsIdAndEnterpriseIdQuery) Execute(ctx context.Context) (*dto.Chat, error) {
	enterpriseId, err := uuid.Parse(q.EnterpriseId)
	if err != nil {
		return nil, err
	}
	recipientOneId, err := uuid.Parse(q.RecipientOneId)
	if err != nil {
		return nil, err
	}
	recipientTwoId, err := uuid.Parse(q.RecipientTwoId)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Chat").(repositories.ChatRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	chat, err := r.Find(ctx, find_chats.ByRecipientOneTwoAndEnterpriseId(
		recipientOneId,
		recipientTwoId,
		enterpriseId,
	))
	if err != nil {
		return nil, err
	}
	return dto.NewChat(*chat), nil
}
