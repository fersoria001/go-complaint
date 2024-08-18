package enterprise

import (
	"go-complaint/domain/model/recipient"
	"time"

	"github.com/google/uuid"
)

type Reply struct {
	id        uuid.UUID
	chatId    uuid.UUID
	sender    recipient.Recipient
	content   string
	seen      bool
	createdAt time.Time
	updatedAt time.Time
}

func NewReply(id uuid.UUID, chatId uuid.UUID, sender recipient.Recipient, content string,
	seen bool, createdAt, updatedAt time.Time) *Reply {
	return &Reply{
		id:        id,
		chatId:    chatId,
		sender:    sender,
		content:   content,
		seen:      seen,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func NewReplyEntity(id uuid.UUID, chatId uuid.UUID, sender recipient.Recipient, content string) *Reply {
	return &Reply{
		id:        id,
		chatId:    chatId,
		sender:    sender,
		content:   content,
		seen:      false,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func (r Reply) Id() uuid.UUID {
	return r.id
}

func (r Reply) ChatId() uuid.UUID {
	return r.chatId
}

func (r Reply) Sender() recipient.Recipient {
	return r.sender
}

func (r Reply) Content() string {
	return r.content
}

func (r Reply) Seen() bool {
	return r.seen
}

func (r Reply) CreatedAt() time.Time {
	return r.createdAt
}

func (r Reply) UpdatedAt() time.Time {
	return r.updatedAt
}
