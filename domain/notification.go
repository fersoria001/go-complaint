package domain

import (
	"go-complaint/domain/model/recipient"
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	id         uuid.UUID
	owner      recipient.Recipient
	sender     recipient.Recipient
	title      string
	content    string
	link       string
	occurredOn time.Time
	seen       bool
}

func NewNotification(
	id uuid.UUID,
	owner,
	sender recipient.Recipient,
	title string,
	content string,
	link string,
	occurredOn time.Time,
	seen bool,
) *Notification {
	return &Notification{
		id:         id,
		owner:      owner,
		sender:     sender,
		title:      title,
		content:    content,
		link:       link,
		occurredOn: occurredOn,
		seen:       seen,
	}
}
func (notification *Notification) MarkAsRead() {
	notification.seen = true
}
func (notification *Notification) ID() uuid.UUID {
	return notification.id
}

func (notification *Notification) Owner() recipient.Recipient {
	return notification.owner
}

func (notification *Notification) Sender() recipient.Recipient {
	return notification.sender
}

func (notification *Notification) Title() string {
	return notification.title
}

func (notification *Notification) Content() string {
	return notification.content
}

func (notification *Notification) Link() string {
	return notification.link
}

func (notification *Notification) OccurredOn() time.Time {
	return notification.occurredOn
}

func (notification *Notification) Seen() bool {
	return notification.seen
}
