package commands_test

import (
	"context"
	"go-complaint/application/commands"
	"go-complaint/domain/model/recipient"
	"go-complaint/infrastructure/persistence/finders/find_all_notifications"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkNotificationAsReadCommandTest_Setup(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewRecipients {
		r := recipient.NewRecipient(v.Id, v.SubjectName, v.SubjectThumbnail, v.SubjectEmail, v.IsEnterprise)
		err := repository.Save(ctx, *r)
		assert.Nil(t, err)
	}
}
func TestMarkNotificationAsReadCommandTest_Execute(t *testing.T) {
	TestMarkNotificationAsReadCommandTest_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	notificationRepository, ok := reg.Get("Notification").(repositories.NotificationRepository)
	assert.True(t, ok)
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
		dbn, err := notificationRepository.FindAll(ctx, find_all_notifications.ByOwnerId(v.Owner.Id))
		assert.Nil(t, err)
		assert.GreaterOrEqual(t, len(dbn), 1)
		for _, n := range dbn {
			assert.False(t, n.Seen())
			n.MarkAsRead()
			err = notificationRepository.Update(ctx, *n)
			assert.Nil(t, err)
			updatedN, err := notificationRepository.Get(ctx, n.ID())
			assert.Nil(t, err)
			assert.NotNil(t, updatedN)
			assert.True(t, updatedN.Seen())
		}

	}
	t.Cleanup(func() {
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
