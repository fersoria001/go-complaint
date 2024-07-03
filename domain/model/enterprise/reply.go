package enterprise

import (
	"go-complaint/domain/model/identity"
	"time"

	"github.com/google/uuid"
)

type Reply struct {
	id        uuid.UUID
	chatID    ChatID
	user      identity.User
	content   string
	seen      bool
	createdAt time.Time
	updatedAt time.Time
}

func NewReply(id uuid.UUID, chatID ChatID, user identity.User, content string,
	seen bool, createdAt, updatedAt time.Time) *Reply {
	return &Reply{
		id:        id,
		chatID:    chatID,
		user:      user,
		content:   content,
		seen:      seen,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func NewReplyEntity(id uuid.UUID, chatID ChatID, user identity.User, content string) *Reply {
	return &Reply{
		id:        id,
		chatID:    chatID,
		user:      user,
		content:   content,
		seen:      false,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func (r Reply) ID() uuid.UUID {
	return r.id
}

func (r Reply) ChatID() ChatID {
	return r.chatID
}

func (r Reply) User() identity.User {
	return r.user
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
