package domain

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	id         uuid.UUID
	ownerID    string
	thumbnail  string
	title      string
	content    string
	link       string
	occurredOn time.Time
	seen       bool
}

func NewNotification(
	id uuid.UUID,
	ownerID string,
	thumbnail string,
	title string,
	content string,
	link string,
	occurredOn time.Time,
	seen bool,
) (*Notification, error) {
	return &Notification{
		id:         id,
		ownerID:    ownerID,
		thumbnail:  thumbnail,
		title:      title,
		content:    content,
		link:       link,
		occurredOn: occurredOn,
		seen:       seen,
	}, nil
}
func (notification *Notification) MarkAsRead() {
	notification.seen = true
}
func (notification *Notification) ID() uuid.UUID {
	return notification.id
}

func (notification *Notification) OwnerID() string {
	return notification.ownerID
}

func (notification *Notification) Thumbnail() string {
	return notification.thumbnail
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

func (notification *Notification) MarshalJSON() ([]byte, error) {
	milliSeconds := notification.occurredOn.UnixMilli()
	milliSecondsString := strconv.FormatInt(milliSeconds, 10)
	return json.Marshal(&struct {
		ID         uuid.UUID `json:"id"`
		OwnerID    string    `json:"owner_id"`
		Thumbnail  string    `json:"thumbnail"`
		Title      string    `json:"title"`
		Content    string    `json:"content"`
		Link       string    `json:"link"`
		OccurredOn string    `json:"occurred_on"`
		Seen       bool      `json:"seen"`
	}{
		ID:         notification.id,
		OwnerID:    notification.ownerID,
		Thumbnail:  notification.thumbnail,
		Title:      notification.title,
		Content:    notification.content,
		Link:       notification.link,
		OccurredOn: milliSecondsString,
		Seen:       notification.seen,
	})
}

func (notification *Notification) UnmarshalJSON(data []byte) error {
	aux := &struct {
		ID         uuid.UUID `json:"id"`
		OwnerID    string    `json:"owner_id"`
		Thumbnail  string    `json:"thumbnail"`
		Title      string    `json:"title"`
		Content    string    `json:"content"`
		Link       string    `json:"link"`
		OccurredOn string    `json:"occurred_on"`
		Seen       bool      `json:"seen"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	var err error
	integer, err := strconv.ParseInt(aux.OccurredOn, 10, 64)
	if err != nil {
		return err
	}
	notification.id = aux.ID
	notification.occurredOn = time.UnixMilli(integer)
	notification.ownerID = aux.OwnerID
	notification.thumbnail = aux.Thumbnail
	notification.title = aux.Title
	notification.content = aux.Content
	notification.link = aux.Link
	notification.seen = aux.Seen
	return nil
}
