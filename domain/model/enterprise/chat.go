package enterprise

import (
	"go-complaint/domain/model/identity"
	"net/mail"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type ChatID struct {
	one   string
	two   string
	three string
}

func (c ChatID) Reverse() ChatID {
	return ChatID{one: c.one, two: c.three, three: c.two}
}
func (c ChatID) Partial() string {
	return "chat:" + c.one + "=" + c.two
}
func (c ChatID) ReversePartial() string {
	return "chat:" + c.one + "=" + c.three
}

// chat:enterpriseId=sender#receiver || chat:enterpriseId=receiver#sender
func NewChatID(id string) (*ChatID, error) {
	if !strings.HasPrefix(id, "chat") {
		return nil, ValidationError{Message: "chat ID must start with chat:"}
	}
	_, after, found := strings.Cut(id, ":")
	if !found {
		return nil, ValidationError{Message: "chat ID must have a colon after chat:"}
	}
	s := strings.Split(after, "=")
	if len(s) != 2 {
		return nil, ValidationError{Message: "chat ID must have two segments separated by ="}
	}
	segments := strings.Split(s[1], "#")
	if len(segments) != 2 {
		return nil, ValidationError{Message: "chat ID must have two segments separated by #"}
	}
	if _, err := mail.ParseAddress(segments[1]); err != nil {
		return nil, ValidationError{Message: "second segment of chat ID must be a valid email address"}
	}
	return &ChatID{
		one:   s[0],
		two:   segments[0],
		three: segments[1],
	}, nil
}

func (c ChatID) One() string {
	return c.one
}

func (c ChatID) Two() string {
	return c.two
}

func (c ChatID) Three() string {
	return c.three
}

func (c ChatID) String() string {
	return "chat:" + c.one + "=" + c.two + "#" + c.three
}

type Chat struct {
	id      ChatID
	replies []*Reply
}

func NewChat(id ChatID, replies []*Reply) *Chat {
	v := &Chat{}
	v.id = id
	v.replies = replies
	return v
}

func NewChatEntity(id ChatID) *Chat {
	v := &Chat{}
	v.id = id
	v.replies = []*Reply{}
	return v
}

func (chat Chat) ID() ChatID {
	return chat.id
}

func (chat Chat) Replies() []*Reply {
	return chat.replies
}

func (chat Chat) MarkAsSeen(ids mapset.Set[uuid.UUID]) {
	for _, reply := range chat.replies {
		if ids.Contains(reply.ID()) {
			reply.seen = true
			reply.updatedAt = time.Now()
		}
	}
}

func (chat *Chat) Reply(id uuid.UUID, user identity.User, content string) Reply {
	reply := NewReplyEntity(id, chat.id, user, content)
	chat.replies = append(chat.replies, reply)
	return *reply
}
