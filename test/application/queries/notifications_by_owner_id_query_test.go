package queries_test

import (
	"context"
	"go-complaint/application/commands"
	"go-complaint/application/queries"
	"go-complaint/domain/model/recipient"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotificationsByOwnerIdQuery_Setup(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewRecipients {
		r := recipient.NewRecipient(v.Id, v.SubjectName, v.SubjectThumbnail, v.SubjectEmail, v.IsEnterprise)
		err := recipientRepository.Save(ctx, *r)
		assert.Nil(t, err)
	}
	for _, v := range mock_data.NewNotifications {
		c := commands.NewSendNotificationCommand(
			v.Owner.Id.String(),
			v.Sender.Id.String(),
			v.Title,
			v.Content,
			v.Link,
		)
		err := c.Execute(ctx)
		assert.Nil(t, err)
	}
}

func TestNotificationsByOwnerIdQuery_Execute(t *testing.T) {
	TestNotificationsByOwnerIdQuery_Setup(t)
	ctx := context.Background()
	for _, v := range mock_data.NewNotifications {
		q := queries.NewNotificationsByOwnerIdQuery(v.Owner.Id.String())
		dbn, err := q.Execute(ctx)
		assert.Nil(t, err)
		assert.GreaterOrEqual(t, len(dbn), 1)
	}
	t.Cleanup(func() {
		reg := repositories.MapperRegistryInstance()
		notificationRepository, ok := reg.Get("Notification").(repositories.NotificationRepository)
		assert.True(t, ok)
		recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewNotifications {
			err := notificationRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}
