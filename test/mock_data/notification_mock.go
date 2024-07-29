package mock_data

import (
	"time"

	"github.com/google/uuid"
)

type NotificationMock struct {
	Id         uuid.UUID
	Owner      *RecipientMock
	Sender     *RecipientMock
	Title      string
	Content    string
	Link       string
	OccurredOn time.Time
	Seen       bool
}

var NewNotifications = map[string]NotificationMock{
	"valid": {
		Id:         uuid.MustParse("107a2be9-38e4-4429-87b1-c08d3b1f4823"),
		Owner:      NewRecipients["enterprise"],
		Sender:     NewRecipients["user"],
		Title:      "this is a notification title",
		Content:    "this is a notification content",
		Link:       "/something",
		OccurredOn: time.Now(),
		Seen:       false,
	},
	"valid1": {
		Id:         uuid.MustParse("107a2be9-38e4-4429-87b1-c08d3b1f4824"),
		Owner:      NewRecipients["user"],
		Sender:     NewRecipients["enterprise"],
		Title:      "this is a notification title",
		Content:    "this is a notification content",
		Link:       "/something",
		OccurredOn: time.Now(),
		Seen:       false,
	},
}
