package repositories_test

import (
	"context"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/recipient"
	"go-complaint/infrastructure/persistence/finders/find_all_hiring_proccesses"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHiringProccessRepository_Setup(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewRecipients {
		recipient := recipient.NewRecipient(
			v.Id,
			v.SubjectName,
			v.SubjectThumbnail,
			v.SubjectThumbnail,
			v.IsEnterprise,
		)
		err := recipientRepository.Save(ctx, *recipient)
		assert.Nil(t, err)
	}
}

func TestHiringProccessRepository_Save(t *testing.T) {
	TestHiringProccessRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewHiringProccesses {
		e := recipient.NewRecipient(
			v.Enterprise.Id,
			v.Enterprise.SubjectName,
			v.Enterprise.SubjectThumbnail,
			v.Enterprise.SubjectThumbnail,
			v.Enterprise.IsEnterprise,
		)
		user := recipient.NewRecipient(
			v.User.Id,
			v.User.SubjectName,
			v.User.SubjectThumbnail,
			v.User.SubjectThumbnail,
			v.User.IsEnterprise,
		)
		emitedBy := recipient.NewRecipient(
			v.EmitedBy.Id,
			v.EmitedBy.SubjectName,
			v.EmitedBy.SubjectThumbnail,
			v.EmitedBy.SubjectThumbnail,
			v.EmitedBy.IsEnterprise,
		)
		updatedBy := recipient.NewRecipient(
			v.UpdatedBy.Id,
			v.UpdatedBy.SubjectName,
			v.UpdatedBy.SubjectThumbnail,
			v.UpdatedBy.SubjectThumbnail,
			v.UpdatedBy.IsEnterprise,
		)
		h := enterprise.NewHiringProccess(
			v.Id,
			*e,
			*user,
			v.Role,
			v.Status,
			v.Reason,
			*emitedBy,
			v.OccurredOn,
			v.LastUpdate,
			*updatedBy,
		)
		err := hiringProccessRepository.Save(ctx, *h)
		assert.Nil(t, err)
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewHiringProccesses {
			err := hiringProccessRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestHiringProccessRepository_Get(t *testing.T) {
	TestHiringProccessRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewHiringProccesses {
		e := recipient.NewRecipient(
			v.Enterprise.Id,
			v.Enterprise.SubjectName,
			v.Enterprise.SubjectThumbnail,
			v.Enterprise.SubjectThumbnail,
			v.Enterprise.IsEnterprise,
		)
		user := recipient.NewRecipient(
			v.User.Id,
			v.User.SubjectName,
			v.User.SubjectThumbnail,
			v.User.SubjectThumbnail,
			v.User.IsEnterprise,
		)
		emitedBy := recipient.NewRecipient(
			v.EmitedBy.Id,
			v.EmitedBy.SubjectName,
			v.EmitedBy.SubjectThumbnail,
			v.EmitedBy.SubjectThumbnail,
			v.EmitedBy.IsEnterprise,
		)
		updatedBy := recipient.NewRecipient(
			v.UpdatedBy.Id,
			v.UpdatedBy.SubjectName,
			v.UpdatedBy.SubjectThumbnail,
			v.UpdatedBy.SubjectThumbnail,
			v.UpdatedBy.IsEnterprise,
		)
		h := enterprise.NewHiringProccess(
			v.Id,
			*e,
			*user,
			v.Role,
			v.Status,
			v.Reason,
			*emitedBy,
			v.OccurredOn,
			v.LastUpdate,
			*updatedBy,
		)
		err := hiringProccessRepository.Save(ctx, *h)
		assert.Nil(t, err)
		dbH, err := hiringProccessRepository.Get(ctx, h.Id())
		assert.Nil(t, err)
		assert.NotNil(t, dbH)
		assert.Equal(t, v.Id, dbH.Id())
		assert.Equal(t, *user, dbH.User())
		assert.Equal(t, *e, dbH.Enterprise())
		assert.Equal(t, v.Role, dbH.Role())
		assert.Equal(t, v.Status, dbH.Status())
		assert.Equal(t, v.Reason, dbH.Reason())
		assert.Equal(t, *emitedBy, dbH.EmitedBy())
		assert.Equal(t, *updatedBy, dbH.UpdatedBy())
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewHiringProccesses {
			err := hiringProccessRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestHiringProccessRepository_FindAll_ByUserId(t *testing.T) {
	TestHiringProccessRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewHiringProccesses {
		e := recipient.NewRecipient(
			v.Enterprise.Id,
			v.Enterprise.SubjectName,
			v.Enterprise.SubjectThumbnail,
			v.Enterprise.SubjectThumbnail,
			v.Enterprise.IsEnterprise,
		)
		user := recipient.NewRecipient(
			v.User.Id,
			v.User.SubjectName,
			v.User.SubjectThumbnail,
			v.User.SubjectThumbnail,
			v.User.IsEnterprise,
		)
		emitedBy := recipient.NewRecipient(
			v.EmitedBy.Id,
			v.EmitedBy.SubjectName,
			v.EmitedBy.SubjectThumbnail,
			v.EmitedBy.SubjectThumbnail,
			v.EmitedBy.IsEnterprise,
		)
		updatedBy := recipient.NewRecipient(
			v.UpdatedBy.Id,
			v.UpdatedBy.SubjectName,
			v.UpdatedBy.SubjectThumbnail,
			v.UpdatedBy.SubjectThumbnail,
			v.UpdatedBy.IsEnterprise,
		)
		h := enterprise.NewHiringProccess(
			v.Id,
			*e,
			*user,
			v.Role,
			v.Status,
			v.Reason,
			*emitedBy,
			v.OccurredOn,
			v.LastUpdate,
			*updatedBy,
		)
		err := hiringProccessRepository.Save(ctx, *h)
		assert.Nil(t, err)
		dbH, err := hiringProccessRepository.FindAll(ctx, find_all_hiring_proccesses.ByUserId(user.Id()))
		assert.Nil(t, err)
		assert.NotNil(t, dbH)
		assert.GreaterOrEqual(t, len(dbH), 1)
	}

	t.Cleanup(func() {
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewHiringProccesses {
			err := hiringProccessRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestHiringProccessRepository_FindAll_ByEnterpriseId(t *testing.T) {
	TestHiringProccessRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewHiringProccesses {
		e := recipient.NewRecipient(
			v.Enterprise.Id,
			v.Enterprise.SubjectName,
			v.Enterprise.SubjectThumbnail,
			v.Enterprise.SubjectThumbnail,
			v.Enterprise.IsEnterprise,
		)
		user := recipient.NewRecipient(
			v.User.Id,
			v.User.SubjectName,
			v.User.SubjectThumbnail,
			v.User.SubjectThumbnail,
			v.User.IsEnterprise,
		)
		emitedBy := recipient.NewRecipient(
			v.EmitedBy.Id,
			v.EmitedBy.SubjectName,
			v.EmitedBy.SubjectThumbnail,
			v.EmitedBy.SubjectThumbnail,
			v.EmitedBy.IsEnterprise,
		)
		updatedBy := recipient.NewRecipient(
			v.UpdatedBy.Id,
			v.UpdatedBy.SubjectName,
			v.UpdatedBy.SubjectThumbnail,
			v.UpdatedBy.SubjectThumbnail,
			v.UpdatedBy.IsEnterprise,
		)
		h := enterprise.NewHiringProccess(
			v.Id,
			*e,
			*user,
			v.Role,
			v.Status,
			v.Reason,
			*emitedBy,
			v.OccurredOn,
			v.LastUpdate,
			*updatedBy,
		)
		err := hiringProccessRepository.Save(ctx, *h)
		assert.Nil(t, err)
		dbH, err := hiringProccessRepository.FindAll(ctx, find_all_hiring_proccesses.ByEnterpriseId(e.Id()))
		assert.Nil(t, err)
		assert.NotNil(t, dbH)
		assert.GreaterOrEqual(t, len(dbH), 1)
	}

	t.Cleanup(func() {
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewHiringProccesses {
			err := hiringProccessRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestHiringProccessRepository_Update(t *testing.T) {
	TestHiringProccessRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewHiringProccesses {
		e := recipient.NewRecipient(
			v.Enterprise.Id,
			v.Enterprise.SubjectName,
			v.Enterprise.SubjectThumbnail,
			v.Enterprise.SubjectThumbnail,
			v.Enterprise.IsEnterprise,
		)
		user := recipient.NewRecipient(
			v.User.Id,
			v.User.SubjectName,
			v.User.SubjectThumbnail,
			v.User.SubjectThumbnail,
			v.User.IsEnterprise,
		)
		emitedBy := recipient.NewRecipient(
			v.EmitedBy.Id,
			v.EmitedBy.SubjectName,
			v.EmitedBy.SubjectThumbnail,
			v.EmitedBy.SubjectThumbnail,
			v.EmitedBy.IsEnterprise,
		)
		updatedBy := recipient.NewRecipient(
			v.UpdatedBy.Id,
			v.UpdatedBy.SubjectName,
			v.UpdatedBy.SubjectThumbnail,
			v.UpdatedBy.SubjectThumbnail,
			v.UpdatedBy.IsEnterprise,
		)
		h := enterprise.NewHiringProccess(
			v.Id,
			*e,
			*user,
			v.Role,
			v.Status,
			v.Reason,
			*emitedBy,
			v.OccurredOn,
			v.LastUpdate,
			*updatedBy,
		)
		err := hiringProccessRepository.Save(ctx, *h)
		assert.Nil(t, err)
		dbH, err := hiringProccessRepository.Get(ctx, h.Id())
		assert.Nil(t, err)
		assert.NotNil(t, dbH)
		updUpdatedBy := recipient.NewRecipient(
			mock_data.NewRecipients["user"].Id,
			mock_data.NewRecipients["user"].SubjectName,
			mock_data.NewRecipients["user"].SubjectThumbnail,
			mock_data.NewRecipients["user"].SubjectThumbnail,
			mock_data.NewRecipients["user"].IsEnterprise,
		)
		dbH.ChangeStatus(ctx, enterprise.ACCEPTED, *updUpdatedBy)
		dbH.WriteAReason("this is a reason mostly used in rject, and this test", *updUpdatedBy)
		err = hiringProccessRepository.Update(ctx, *dbH)
		assert.Nil(t, err)
		updatedH, err := hiringProccessRepository.Get(ctx, dbH.Id())
		assert.Nil(t, err)
		assert.NotNil(t, updatedH)
		assert.NotNil(t, updatedH.Reason())
		assert.Equal(t, enterprise.ACCEPTED, updatedH.Status())
		assert.Equal(t, *updUpdatedBy, updatedH.UpdatedBy())
		assert.NotEqual(t, updatedH.EmitedBy(), updatedH.UpdatedBy())
		if updatedH.LastUpdate().Before(updatedH.OccurredOn()) {
			t.Errorf("last update should not be before occurred on")
		}
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewHiringProccesses {
			err := hiringProccessRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}
