package repositories_test

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/recipient"
	"go-complaint/infrastructure/persistence/finders/find_all_notifications"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotificationRepository_Setup(t *testing.T) {
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

func TestNotificationRepository_Save(t *testing.T) {
	TestNotificationRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	notificationRepository, ok := reg.Get("Notification").(repositories.NotificationRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewNotifications {
		owner := recipient.NewRecipient(v.Owner.Id, v.Owner.SubjectName, v.Owner.SubjectEmail, v.Owner.SubjectThumbnail, v.Owner.IsEnterprise)
		sender := recipient.NewRecipient(v.Sender.Id, v.Owner.SubjectName, v.Sender.SubjectEmail, v.Sender.SubjectEmail, v.Sender.IsEnterprise)
		n := domain.NewNotification(v.Id, *owner, *sender, v.Title, v.Content, v.Link, v.OccurredOn, v.Seen)
		err := notificationRepository.Save(ctx, *n)
		assert.Nil(t, err)
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

func TestNotificationRepository_Get(t *testing.T) {
	TestNotificationRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	notificationRepository, ok := reg.Get("Notification").(repositories.NotificationRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewNotifications {
		owner := recipient.NewRecipient(v.Owner.Id, v.Owner.SubjectName, v.Owner.SubjectEmail, v.Owner.SubjectThumbnail, v.Owner.IsEnterprise)
		sender := recipient.NewRecipient(v.Sender.Id, v.Owner.SubjectName, v.Sender.SubjectEmail, v.Sender.SubjectEmail, v.Sender.IsEnterprise)
		n := domain.NewNotification(v.Id, *owner, *sender, v.Title, v.Content, v.Link, v.OccurredOn, v.Seen)
		err := notificationRepository.Save(ctx, *n)
		assert.Nil(t, err)
		dbN, err := notificationRepository.Get(ctx, v.Id)
		assert.Nil(t, err)
		assert.NotNil(t, dbN)
		assert.Equal(t, owner.Id(), dbN.Owner().Id())
		assert.Equal(t, sender.Id(), dbN.Sender().Id())
		assert.Equal(t, v.Title, dbN.Title())
		assert.Equal(t, v.Content, dbN.Content())
		assert.Equal(t, v.Link, dbN.Link())
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

func TestNotificationRepository_Update(t *testing.T) {
	TestNotificationRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	notificationRepository, ok := reg.Get("Notification").(repositories.NotificationRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewNotifications {
		owner := recipient.NewRecipient(v.Owner.Id, v.Owner.SubjectName, v.Owner.SubjectEmail, v.Owner.SubjectThumbnail, v.Owner.IsEnterprise)
		sender := recipient.NewRecipient(v.Sender.Id, v.Owner.SubjectName, v.Sender.SubjectEmail, v.Sender.SubjectEmail, v.Sender.IsEnterprise)
		n := domain.NewNotification(v.Id, *owner, *sender, v.Title, v.Content, v.Link, v.OccurredOn, v.Seen)
		err := notificationRepository.Save(ctx, *n)
		assert.Nil(t, err)
		dbN, err := notificationRepository.Get(ctx, v.Id)
		assert.Nil(t, err)
		assert.NotNil(t, dbN)
		assert.False(t, dbN.Seen())
		dbN.MarkAsRead()
		err = notificationRepository.Update(ctx, *dbN)
		assert.Nil(t, err)
		updatedN, err := notificationRepository.Get(ctx, dbN.ID())
		assert.Nil(t, err)
		assert.NotNil(t, updatedN)
		assert.True(t, updatedN.Seen())
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

func TestNotificationRepository_FindAll(t *testing.T) {
	TestNotificationRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	notificationRepository, ok := reg.Get("Notification").(repositories.NotificationRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewNotifications {
		owner := recipient.NewRecipient(v.Owner.Id, v.Owner.SubjectName, v.Owner.SubjectEmail, v.Owner.SubjectThumbnail, v.Owner.IsEnterprise)
		sender := recipient.NewRecipient(v.Sender.Id, v.Owner.SubjectName, v.Sender.SubjectEmail, v.Sender.SubjectEmail, v.Sender.IsEnterprise)
		n := domain.NewNotification(v.Id, *owner, *sender, v.Title, v.Content, v.Link, v.OccurredOn, v.Seen)
		err := notificationRepository.Save(ctx, *n)
		assert.Nil(t, err)
	}
	for _, v := range mock_data.NewNotifications {
		dbN, err := notificationRepository.FindAll(ctx, find_all_notifications.ByOwnerId(v.Owner.Id))
		assert.Nil(t, err)
		assert.GreaterOrEqual(t, len(dbN), 1)
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
