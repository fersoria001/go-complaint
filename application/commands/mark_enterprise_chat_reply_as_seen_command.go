package commands

import (
	"context"
	"go-complaint/application"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type MarkEnterpriseChatReplyAsSeenCommand struct {
	RepliesIds []string `json:"repliesIds"`
	ChatId     string   `json:"chatId"`
}

func NewMarkEnterpriseChatReplyAsSeenCommand(chatId string, repliesIds []string) *MarkEnterpriseChatReplyAsSeenCommand {
	return &MarkEnterpriseChatReplyAsSeenCommand{
		ChatId:     chatId,
		RepliesIds: repliesIds,
	}
}

func (c MarkEnterpriseChatReplyAsSeenCommand) Execute(ctx context.Context) error {
	parseIds := mapset.NewSet[uuid.UUID]()
	for i := range c.RepliesIds {
		replyId, err := uuid.Parse(c.RepliesIds[i])
		if err != nil {
			return err
		}
		parseIds.Add(replyId)
	}
	chatId, err := uuid.Parse(c.ChatId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Chat").(repositories.ChatRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	dbChat, err := r.Get(ctx, chatId)
	if err != nil {
		return err
	}
	dbChat.MarkAsSeen(parseIds)
	err = r.Update(ctx, dbChat)
	if err != nil {
		return err
	}
	svc := application.ApplicationMessagePublisherInstance()
	replies := dbChat.Replies()
	for i := range replies {
		if parseIds.Contains(replies[i].Id()) {
			svc.Publish(application.NewApplicationMessage(
				dbChat.RecipientOne().Id().String(),
				"enterprise.Reply",
				*dto.NewChatReply(*replies[i]),
			))
			svc.Publish(application.NewApplicationMessage(
				dbChat.RecipientTwo().Id().String(),
				"enterprise.Reply",
				*dto.NewChatReply(*replies[i]),
			))
		}
	}
	return nil
}
