package repositories_test

import (
	"context"
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/recipient"
	"go-complaint/infrastructure/persistence/finders/find_all_complaint_replies"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

func TestComplaintRepliesRepository_Setup(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewRecipients {
		r := recipient.NewRecipient(
			v.Id,
			v.SubjectName,
			v.SubjectThumbnail,
			v.SubjectEmail,
			v.IsEnterprise,
		)
		err := repository.Save(ctx, *r)
		assert.Nil(t, err)
	}
}

func TestComplaintRepliesRepository_Save(t *testing.T) {
	TestComplaintRepliesRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Reply").(repositories.ComplaintRepliesRepository)
	assert.True(t, ok)
	for i := range mock_data.NewReplies {
		for _, v := range mock_data.NewReplies[i] {
			author := recipient.NewRecipient(
				v.Sender.Id,
				v.Sender.SubjectName,
				v.Sender.SubjectThumbnail,
				v.Sender.SubjectEmail,
				v.Sender.IsEnterprise,
			)
			r, err := complaint.NewReply(
				v.Id,
				v.ComplaintId,
				*author,
				v.Body,
				v.Read,
				v.CreatedAt,
				v.ReadAt,
				v.UpdatedAt,
				author.IsEnterprise(),
				author.Id(),
			)
			assert.Nil(t, err)
			assert.NotNil(t, r)
			err = repository.Save(ctx, *r)
			assert.Nil(t, err)
		}
	}
	t.Cleanup(func() {
		replyRepository, ok := reg.Get("Reply").(repositories.ComplaintRepliesRepository)
		assert.True(t, ok)
		recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for i := range mock_data.NewReplies {
			for _, v := range mock_data.NewReplies[i] {
				err := replyRepository.Remove(ctx, v.Id)
				assert.Nil(t, err)
			}
		}
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestComplaintRepliesRepository_Get(t *testing.T) {
	TestComplaintRepliesRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Reply").(repositories.ComplaintRepliesRepository)
	assert.True(t, ok)
	for i := range mock_data.NewReplies {
		for _, v := range mock_data.NewReplies[i] {
			author := recipient.NewRecipient(
				v.Sender.Id,
				v.Sender.SubjectName,
				v.Sender.SubjectThumbnail,
				v.Sender.SubjectEmail,
				v.Sender.IsEnterprise,
			)
			r, err := complaint.NewReply(
				v.Id,
				v.ComplaintId,
				*author,
				v.Body,
				v.Read,
				v.CreatedAt,
				v.ReadAt,
				v.UpdatedAt,
				author.IsEnterprise(),
				author.Id(),
			)
			assert.Nil(t, err)
			assert.NotNil(t, r)
			err = repository.Save(ctx, *r)
			assert.Nil(t, err)
			dbR, err := repository.Get(ctx, v.Id)
			assert.Nil(t, err)
			assert.NotNil(t, dbR)
			assert.Equal(t, v.ComplaintId, dbR.ComplaintId())
			assert.Equal(t, author.Id(), dbR.Sender().Id())
			assert.Equal(t, author.SubjectName(), dbR.Sender().SubjectName())
			assert.Equal(t, author.SubjectThumbnail(), dbR.Sender().SubjectThumbnail())
			assert.Equal(t, author.IsEnterprise(), dbR.Sender().IsEnterprise())
			assert.Equal(t, v.Body, dbR.Body())
			assert.Equal(t, v.Read, dbR.Read())
			assert.Equal(t, v.CreatedAt.StringRepresentation(), dbR.CreatedAt().StringRepresentation())
			assert.Equal(t, v.ReadAt.StringRepresentation(), dbR.ReadAt().StringRepresentation())
			assert.Equal(t, v.UpdatedAt.StringRepresentation(), dbR.UpdatedAt().StringRepresentation())
		}
	}
	t.Cleanup(func() {
		replyRepository, ok := reg.Get("Reply").(repositories.ComplaintRepliesRepository)
		assert.True(t, ok)
		recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for i := range mock_data.NewReplies {
			for _, v := range mock_data.NewReplies[i] {
				err := replyRepository.Remove(ctx, v.Id)
				assert.Nil(t, err)
			}
		}
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestComplaintRepliesRepository_SaveAll(t *testing.T) {
	TestComplaintRepliesRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Reply").(repositories.ComplaintRepliesRepository)
	assert.True(t, ok)
	repliesSet := mapset.NewSet[complaint.Reply]()
	for i := range mock_data.NewReplies {
		for _, v := range mock_data.NewReplies[i] {
			author := recipient.NewRecipient(
				v.Sender.Id,
				v.Sender.SubjectName,
				v.Sender.SubjectThumbnail,
				v.Sender.SubjectEmail,
				v.Sender.IsEnterprise,
			)
			r, err := complaint.NewReply(
				v.Id,
				v.ComplaintId,
				*author,
				v.Body,
				v.Read,
				v.CreatedAt,
				v.ReadAt,
				v.UpdatedAt,
				author.IsEnterprise(),
				author.Id(),
			)
			assert.Nil(t, err)
			assert.NotNil(t, r)
			repliesSet.Add(*r)
		}
	}
	err := repository.SaveAll(ctx, repliesSet)
	assert.Nil(t, err)
	t.Cleanup(func() {
		replyRepository, ok := reg.Get("Reply").(repositories.ComplaintRepliesRepository)
		assert.True(t, ok)
		recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for k := range mock_data.NewReplies {
			err := replyRepository.DeleteAll(ctx, k)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestComplaintRepliesRepository_FindAll(t *testing.T) {
	TestComplaintRepliesRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Reply").(repositories.ComplaintRepliesRepository)
	assert.True(t, ok)
	repliesSet := mapset.NewSet[complaint.Reply]()
	for i := range mock_data.NewReplies {
		for _, v := range mock_data.NewReplies[i] {
			author := recipient.NewRecipient(
				v.Sender.Id,
				v.Sender.SubjectName,
				v.Sender.SubjectThumbnail,
				v.Sender.SubjectEmail,
				v.Sender.IsEnterprise,
			)
			r, err := complaint.NewReply(
				v.Id,
				v.ComplaintId,
				*author,
				v.Body,
				v.Read,
				v.CreatedAt,
				v.ReadAt,
				v.UpdatedAt,
				author.IsEnterprise(),
				author.Id(),
			)
			assert.Nil(t, err)
			assert.NotNil(t, r)
			repliesSet.Add(*r)
		}
	}
	err := repository.SaveAll(ctx, repliesSet)
	assert.Nil(t, err)
	for k := range mock_data.NewReplies {
		dbReplies, err := repository.FindAll(ctx, find_all_complaint_replies.ByComplaintID(k))
		assert.Nil(t, err)
		assert.Equal(t, 2, dbReplies.Cardinality())
	}
	t.Cleanup(func() {
		replyRepository, ok := reg.Get("Reply").(repositories.ComplaintRepliesRepository)
		assert.True(t, ok)
		recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for k := range mock_data.NewReplies {
			err := replyRepository.DeleteAll(ctx, k)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}
