package enterprise

import (
	"go-complaint/domain/model/recipient"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type Chat struct {
	id           uuid.UUID
	enterpriseId uuid.UUID
	recipientOne recipient.Recipient
	recipientTwo recipient.Recipient
	replies      []*Reply
}

func NewChat(
	id uuid.UUID,
	enterpriseId uuid.UUID,
	recipientOne,
	recipientTwo recipient.Recipient,
	replies []*Reply) *Chat {
	return &Chat{
		id:           id,
		enterpriseId: enterpriseId,
		recipientOne: recipientOne,
		recipientTwo: recipientTwo,
		replies:      replies,
	}
}

func NewChatEntity(id, enterpriseId uuid.UUID, recipientOne, recipientTwo recipient.Recipient) *Chat {
	return &Chat{
		id:           id,
		enterpriseId: enterpriseId,
		recipientOne: recipientOne,
		recipientTwo: recipientTwo,
		replies:      make([]*Reply, 0),
	}
}

func (chat Chat) Id() uuid.UUID {
	return chat.id
}

func (chat Chat) EnterpriseId() uuid.UUID {
	return chat.enterpriseId
}

func (chat Chat) RecipientOne() recipient.Recipient {
	return chat.recipientOne
}

func (chat Chat) RecipientTwo() recipient.Recipient {
	return chat.recipientTwo
}

func (chat Chat) Replies() []*Reply {
	return chat.replies
}

func (chat Chat) MarkAsSeen(ids mapset.Set[uuid.UUID]) {
	for _, reply := range chat.replies {
		if ids.Contains(reply.Id()) {
			reply.seen = true
			reply.updatedAt = time.Now()
		}
	}
}

func (chat *Chat) Reply(reply *Reply) error {
	chat.replies = append(chat.replies, reply)
	return nil
}
