package repositories_test

import (
	"context"
	"go-complaint/domain/model/recipient"
	"go-complaint/infrastructure/persistence/finders/find_all_recipients"
	"go-complaint/infrastructure/persistence/finders/find_recipient"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecipientRepository_Save(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewRecipients {
		r := recipient.NewRecipient(v.Id, v.SubjectName, v.SubjectThumbnail, v.SubjectEmail, v.IsEnterprise)
		err := repository.Save(ctx, *r)
		assert.Nil(t, err)
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewRecipients {
			err := repository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestRecipientRepository_Get(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewRecipients {
		r := recipient.NewRecipient(v.Id, v.SubjectName, v.SubjectThumbnail, v.SubjectEmail, v.IsEnterprise)
		err := repository.Save(ctx, *r)
		assert.Nil(t, err)
		dbR, err := repository.Get(ctx, v.Id)
		assert.Nil(t, err)
		assert.NotNil(t, dbR)
		assert.Equal(t, v.Id, dbR.Id())
		assert.Equal(t, v.SubjectName, dbR.SubjectName())
		assert.Equal(t, v.SubjectThumbnail, dbR.SubjectThumbnail())
		assert.Equal(t, v.IsEnterprise, dbR.IsEnterprise())
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewRecipients {
			err := repository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestRecipientRepository_Find_ByNameAndEmail(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewRecipients {
		r := recipient.NewRecipient(v.Id, v.SubjectName, v.SubjectThumbnail, v.SubjectEmail, v.IsEnterprise)
		err := repository.Save(ctx, *r)
		assert.Nil(t, err)
		dbR, err := repository.Find(ctx, find_recipient.ByNameAndEmail(v.SubjectName, v.SubjectEmail))
		assert.Nil(t, err)
		assert.NotNil(t, dbR)
		assert.Equal(t, v.Id, dbR.Id())
		assert.Equal(t, v.SubjectName, dbR.SubjectName())
		assert.Equal(t, v.SubjectThumbnail, dbR.SubjectThumbnail())
		assert.Equal(t, v.IsEnterprise, dbR.IsEnterprise())
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewRecipients {
			err := repository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestRecipientRepository_FindAll_ByNameLike(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewRecipients {
		r := recipient.NewRecipient(v.Id, v.SubjectName, v.SubjectThumbnail, v.SubjectEmail, v.IsEnterprise)
		err := repository.Save(ctx, *r)
		assert.Nil(t, err)
	}
	userMock := mock_data.NewRecipients["user"]
	query := userMock.SubjectName[:3]
	dbR, err := repository.FindAll(ctx, find_all_recipients.ByNameLike(query))
	assert.Nil(t, err)
	assert.NotNil(t, dbR)
	assert.GreaterOrEqual(t, len(dbR), 1)
	t.Cleanup(func() {
		for _, v := range mock_data.NewRecipients {
			err := repository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestRecipientRepository_Update(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewRecipients {
		r := recipient.NewRecipient(v.Id, v.SubjectName, v.SubjectThumbnail, v.SubjectEmail, v.IsEnterprise)
		err := repository.Save(ctx, *r)
		assert.Nil(t, err)
		dbR, err := repository.Get(ctx, v.Id)
		assert.Nil(t, err)
		assert.NotNil(t, dbR)
		err = dbR.SetSubjectName(v.SubjectName + "1")
		assert.Nil(t, err)
		err = dbR.SetSubjectThumbnail(v.SubjectThumbnail + "1")
		assert.Nil(t, err)
		err = repository.Update(ctx, *dbR)
		assert.Nil(t, err)
		updatedR, err := repository.Get(ctx, dbR.Id())
		assert.Nil(t, err)
		assert.NotNil(t, updatedR)
		assert.Equal(t, v.SubjectName+"1", updatedR.SubjectName())
		assert.Equal(t, v.SubjectThumbnail+"1", updatedR.SubjectThumbnail())
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewRecipients {
			err := repository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}
