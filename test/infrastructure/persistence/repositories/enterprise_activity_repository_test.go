package repositories_test

import (
	"context"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/recipient"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnterpriseActivity_Save(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("EnterpriseActivity").(repositories.EnterpriseActivityRepository)
	assert.True(t, ok)
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewRecipients {
		err := recipientRepository.Save(ctx, *recipient.NewRecipient(
			v.Id, v.SubjectName, v.SubjectThumbnail, v.SubjectEmail, v.IsEnterprise,
		))
		assert.Nil(t, err)
	}
	t.Run("save enterprise activity", func(t *testing.T) {
		for _, v := range mock_data.NewEnterpriseActivity {
			user, err := recipientRepository.Get(ctx, v.UserId)
			assert.Nil(t, err)
			ea := enterprise.NewEnterpriseActivity(
				v.Id,
				v.ActivityId,
				v.EnterpriseId,
				*user,
				v.EnterpriseName,
				v.OccurredOn,
				v.ActivityType,
			)
			err = r.Save(ctx, *ea)
			assert.Nil(t, err)
		}
	})
	t.Run("get enterprise activity", func(t *testing.T) {
		for _, v := range mock_data.NewEnterpriseActivity {
			dbEa, err := r.Get(ctx, v.Id)
			assert.Nil(t, err)
			assert.NotNil(t, dbEa)
			assert.Equal(t, v.ActivityId, dbEa.ActivityId())
			assert.Equal(t, v.UserId, dbEa.User().Id())
			assert.Equal(t, v.EnterpriseId, dbEa.EnterpriseId())
			assert.Equal(t, v.EnterpriseName, dbEa.EnterpriseName())
			assert.Equal(t, v.ActivityType, dbEa.ActivityType())
		}
	})

	t.Cleanup(func() {
		for _, v := range mock_data.NewEnterpriseActivity {
			err := r.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}
