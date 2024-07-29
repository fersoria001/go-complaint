package repositories_test

import (
	"context"
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/recipient"
	"go-complaint/infrastructure/persistence/counters/count_complaints"
	"go-complaint/infrastructure/persistence/finders/find_all_complaints"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

func TestComplaintRepository_Setup(t *testing.T) {
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

func TestComplaintRepository_Save(t *testing.T) {
	TestComplaintRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewComplaints {
		author := recipient.NewRecipient(
			v.Author.Id,
			v.Author.SubjectName,
			v.Author.SubjectThumbnail,
			v.Author.SubjectEmail,
			v.Author.IsEnterprise,
		)
		receiver := recipient.NewRecipient(
			v.Receiver.Id,
			v.Receiver.SubjectName,
			v.Receiver.SubjectThumbnail,
			v.Author.SubjectEmail,
			v.Receiver.IsEnterprise,
		)
		newComplaint, err := complaint.CreateNew(
			ctx,
			v.Id,
			*author,
			*receiver,
		)
		assert.Nil(t, err)
		assert.NotNil(t, newComplaint)
		err = repository.Save(ctx, newComplaint)
		assert.Nil(t, err)
	}
	t.Cleanup(func() {
		complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
		assert.True(t, ok)
		recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for _, v := range mock_data.NewComplaints {
			err := complaintRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestComplaintRepository_Get(t *testing.T) {
	TestComplaintRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewComplaints {
		author := recipient.NewRecipient(
			v.Author.Id,
			v.Author.SubjectName,
			v.Author.SubjectThumbnail,
			v.Author.SubjectEmail,
			v.Author.IsEnterprise,
		)
		receiver := recipient.NewRecipient(
			v.Receiver.Id,
			v.Receiver.SubjectName,
			v.Receiver.SubjectThumbnail,
			v.Receiver.SubjectEmail,
			v.Receiver.IsEnterprise,
		)
		newComplaint, err := complaint.CreateNew(
			ctx,
			v.Id,
			*author,
			*receiver,
		)
		assert.Nil(t, err)
		assert.NotNil(t, newComplaint)
		err = repository.Save(ctx, newComplaint)
		assert.Nil(t, err)
		dbC, err := repository.Get(ctx, newComplaint.Id())
		assert.Nil(t, err)
		assert.NotNil(t, dbC)
		assert.Equal(t, author.Id(), dbC.Author().Id())
		assert.Equal(t, author.SubjectName(), dbC.Author().SubjectName())
		assert.Equal(t, author.SubjectThumbnail(), dbC.Author().SubjectThumbnail())
		assert.Equal(t, author.IsEnterprise(), dbC.Author().IsEnterprise())
		assert.Equal(t, receiver.Id(), dbC.Receiver().Id())
		assert.Equal(t, receiver.SubjectName(), dbC.Receiver().SubjectName())
		assert.Equal(t, receiver.SubjectThumbnail(), dbC.Receiver().SubjectThumbnail())
		assert.Equal(t, complaint.WRITING, dbC.Status())
		assert.Equal(t, 0, dbC.Replies().Cardinality())
	}
	t.Cleanup(func() {
		complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
		assert.True(t, ok)
		recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for _, v := range mock_data.NewComplaints {
			err := complaintRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestComplaintRepository_Update(t *testing.T) {
	TestComplaintRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewComplaints {
		author := recipient.NewRecipient(
			v.Author.Id,
			v.Author.SubjectName,
			v.Author.SubjectThumbnail,
			v.Author.SubjectEmail,
			v.Author.IsEnterprise,
		)
		receiver := recipient.NewRecipient(
			v.Receiver.Id,
			v.Receiver.SubjectName,
			v.Receiver.SubjectThumbnail,
			v.Receiver.SubjectEmail,
			v.Receiver.IsEnterprise,
		)
		newComplaint, err := complaint.CreateNew(
			ctx,
			v.Id,
			*author,
			*receiver,
		)
		assert.Nil(t, err)
		assert.NotNil(t, newComplaint)
		err = repository.Save(ctx, newComplaint)
		assert.Nil(t, err)
		dbC, err := repository.Get(ctx, newComplaint.Id())
		assert.Nil(t, err)
		assert.NotNil(t, dbC)
		assert.Greater(t, len(v.Replies), 0)
		mockBody := v.Replies[0].Body
		err = dbC.SetTitle(ctx, v.Title)
		assert.Nil(t, err)
		err = dbC.SetDescription(ctx, v.Description)
		assert.Nil(t, err)
		err = dbC.SetBody(ctx, mockBody)
		assert.Nil(t, err)
		err = repository.Update(ctx, dbC)
		assert.Nil(t, err)
		updatedC, err := repository.Get(ctx, dbC.Id())
		assert.Nil(t, err)
		assert.NotNil(t, updatedC)
		assert.Equal(t, v.Title, updatedC.Title())
		assert.Equal(t, v.Description, updatedC.Description())
		assert.Greater(t, updatedC.Replies().Cardinality(), 0)
		assert.Equal(t, mockBody, updatedC.Body())
	}
	t.Cleanup(func() {
		complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
		assert.True(t, ok)
		recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for _, v := range mock_data.NewComplaints {
			err := complaintRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestComplaintRepository_Get_Complete(t *testing.T) {
	TestComplaintRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewComplaints {
		author := recipient.NewRecipient(
			v.Author.Id,
			v.Author.SubjectName,
			v.Author.SubjectThumbnail,
			v.Author.SubjectEmail,
			v.Author.IsEnterprise,
		)
		receiver := recipient.NewRecipient(
			v.Receiver.Id,
			v.Receiver.SubjectName,
			v.Receiver.SubjectThumbnail,
			v.Receiver.SubjectEmail,
			v.Receiver.IsEnterprise,
		)
		rating, err := complaint.NewRating(v.Id, v.Rating.Rate, v.Rating.Comment)
		assert.Nil(t, err)
		assert.Greater(t, len(v.Replies), 0)
		mockReply := v.Replies[0]
		reply, err := complaint.NewReply(
			v.Id,
			v.Id,
			*author,
			mockReply.Body,
			false,
			mockReply.CreatedAt,
			mockReply.ReadAt,
			mockReply.UpdatedAt,
		)
		assert.Nil(t, err)
		replies := mapset.NewSet(reply)
		newComplaint, err := complaint.NewComplaint(
			v.Id,
			*author,
			*receiver,
			v.Status,
			v.Title,
			v.Description,
			v.CreatedAt,
			v.UpdatedAt,
			rating,
			replies,
		)
		assert.Nil(t, err)
		assert.NotNil(t, newComplaint)
		err = repository.Save(ctx, newComplaint)
		assert.Nil(t, err)
		err = repository.Update(ctx, newComplaint)
		assert.Nil(t, err)
		dbC, err := repository.Get(ctx, newComplaint.Id())
		assert.Nil(t, err)
		assert.NotNil(t, dbC)
		assert.Equal(t, v.Id, dbC.Id())
		assert.Equal(t, v.Author.Id, dbC.Author().Id())
		assert.Equal(t, v.Author.SubjectName, dbC.Author().SubjectName())
		assert.Equal(t, v.Author.SubjectThumbnail, dbC.Author().SubjectThumbnail())
		assert.Equal(t, v.Author.IsEnterprise, dbC.Author().IsEnterprise())
		assert.Equal(t, v.Receiver.Id, dbC.Receiver().Id())
		assert.Equal(t, v.Receiver.SubjectName, dbC.Receiver().SubjectName())
		assert.Equal(t, v.Receiver.SubjectThumbnail, dbC.Receiver().SubjectThumbnail())
		assert.Equal(t, v.Receiver.IsEnterprise, dbC.Receiver().IsEnterprise())
		assert.Equal(t, v.Status, dbC.Status())
		assert.Equal(t, v.Title, dbC.Title())
		assert.Equal(t, v.Description, dbC.Description())
		assert.Equal(t, v.CreatedAt.StringRepresentation(), dbC.CreatedAt().StringRepresentation())
		assert.Equal(t, v.UpdatedAt.StringRepresentation(), dbC.UpdatedAt().StringRepresentation())
		assert.Equal(t, rating.Id(), dbC.Rating().Id())
		assert.Equal(t, rating.Rate(), dbC.Rating().Rate())
		assert.Equal(t, rating.Comment(), dbC.Rating().Comment())
		assert.Equal(t, 1, dbC.Replies().Cardinality())
		assert.Equal(t, mockReply.Body, dbC.Body())
	}
	t.Cleanup(func() {
		complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
		assert.True(t, ok)
		recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for _, v := range mock_data.NewComplaints {
			err := complaintRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestComplaintRepository_FindAll_ByAuthorOrReceiver(t *testing.T) {
	TestComplaintRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewComplaints {
		author := recipient.NewRecipient(
			v.Author.Id,
			v.Author.SubjectName,
			v.Author.SubjectThumbnail,
			v.Author.SubjectEmail,
			v.Author.IsEnterprise,
		)
		receiver := recipient.NewRecipient(
			v.Receiver.Id,
			v.Receiver.SubjectName,
			v.Receiver.SubjectThumbnail,
			v.Receiver.SubjectEmail,
			v.Receiver.IsEnterprise,
		)
		rating, err := complaint.NewRating(v.Id, v.Rating.Rate, v.Rating.Comment)
		assert.Nil(t, err)
		assert.Greater(t, len(v.Replies), 0)
		mockReply := v.Replies[0]
		reply, err := complaint.NewReply(
			v.Id,
			v.Id,
			*author,
			mockReply.Body,
			false,
			mockReply.CreatedAt,
			mockReply.ReadAt,
			mockReply.UpdatedAt,
		)
		assert.Nil(t, err)
		replies := mapset.NewSet(reply)
		newComplaint, err := complaint.NewComplaint(
			v.Id,
			*author,
			*receiver,
			v.Status,
			v.Title,
			v.Description,
			v.CreatedAt,
			v.UpdatedAt,
			rating,
			replies,
		)
		assert.Nil(t, err)
		assert.NotNil(t, newComplaint)
		err = repository.Save(ctx, newComplaint)
		assert.Nil(t, err)
		err = repository.Update(ctx, newComplaint)
		assert.Nil(t, err)
	}
	receiver := mock_data.NewRecipients["user"]
	dbCs, err := repository.FindAll(ctx, find_all_complaints.ByAuthorOrReceiver(
		receiver.Id,
		[]string{complaint.OPEN.String(), complaint.STARTED.String(), complaint.IN_DISCUSSION.String()}))
	assert.Nil(t, err)
	assert.NotNil(t, dbCs)
	assert.Greater(t, len(dbCs), 0)
	t.Cleanup(func() {
		complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
		assert.True(t, ok)
		recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for _, v := range mock_data.NewComplaints {
			err := complaintRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestComplaintRepository_FindAll_ByAuthorIdAndStatusIn(t *testing.T) {
	TestComplaintRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewComplaints {
		author := recipient.NewRecipient(
			v.Author.Id,
			v.Author.SubjectName,
			v.Author.SubjectThumbnail,
			v.Author.SubjectEmail,
			v.Author.IsEnterprise,
		)
		receiver := recipient.NewRecipient(
			v.Receiver.Id,
			v.Receiver.SubjectName,
			v.Receiver.SubjectThumbnail,
			v.Receiver.SubjectEmail,
			v.Receiver.IsEnterprise,
		)
		rating, err := complaint.NewRating(v.Id, v.Rating.Rate, v.Rating.Comment)
		assert.Nil(t, err)
		assert.Greater(t, len(v.Replies), 0)
		mockReply := v.Replies[0]
		reply, err := complaint.NewReply(
			v.Id,
			v.Id,
			*author,
			mockReply.Body,
			false,
			mockReply.CreatedAt,
			mockReply.ReadAt,
			mockReply.UpdatedAt,
		)
		assert.Nil(t, err)
		replies := mapset.NewSet(reply)
		newComplaint, err := complaint.NewComplaint(
			v.Id,
			*author,
			*receiver,
			complaint.IN_REVIEW,
			v.Title,
			v.Description,
			v.CreatedAt,
			v.UpdatedAt,
			rating,
			replies,
		)
		assert.Nil(t, err)
		assert.NotNil(t, newComplaint)
		err = repository.Save(ctx, newComplaint)
		assert.Nil(t, err)
		err = repository.Update(ctx, newComplaint)
		assert.Nil(t, err)
	}
	receiver := mock_data.NewRecipients["user"]
	dbCs, err := repository.FindAll(ctx, find_all_complaints.ByAuthorAndStatusIn(
		receiver.Id,
		[]string{complaint.IN_REVIEW.String()}))
	assert.Nil(t, err)
	assert.NotNil(t, dbCs)
	assert.Greater(t, len(dbCs), 0)
	t.Cleanup(func() {
		complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
		assert.True(t, ok)
		recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for _, v := range mock_data.NewComplaints {
			err := complaintRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestComplaintRepository_Count_WhereAuthorOrReceiver(t *testing.T) {
	TestComplaintRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewComplaints {
		author := recipient.NewRecipient(
			v.Author.Id,
			v.Author.SubjectName,
			v.Author.SubjectThumbnail,
			v.Author.SubjectEmail,
			v.Author.IsEnterprise,
		)
		receiver := recipient.NewRecipient(
			v.Receiver.Id,
			v.Receiver.SubjectName,
			v.Receiver.SubjectThumbnail,
			v.Receiver.SubjectEmail,
			v.Receiver.IsEnterprise,
		)
		rating, err := complaint.NewRating(v.Id, v.Rating.Rate, v.Rating.Comment)
		assert.Nil(t, err)
		assert.Greater(t, len(v.Replies), 0)
		mockReply := v.Replies[0]
		reply, err := complaint.NewReply(
			v.Id,
			v.Id,
			*author,
			mockReply.Body,
			false,
			mockReply.CreatedAt,
			mockReply.ReadAt,
			mockReply.UpdatedAt,
		)
		assert.Nil(t, err)
		replies := mapset.NewSet(reply)
		newComplaint, err := complaint.NewComplaint(
			v.Id,
			*author,
			*receiver,
			v.Status,
			v.Title,
			v.Description,
			v.CreatedAt,
			v.UpdatedAt,
			rating,
			replies,
		)
		assert.Nil(t, err)
		assert.NotNil(t, newComplaint)
		err = repository.Save(ctx, newComplaint)
		assert.Nil(t, err)
		err = repository.Update(ctx, newComplaint)
		assert.Nil(t, err)
	}
	receiver := mock_data.NewRecipients["user"]
	count, err := repository.Count(ctx, count_complaints.WhereAuthorOrReceiverId(
		receiver.Id,
		[]string{complaint.OPEN.String(), complaint.STARTED.String(), complaint.IN_DISCUSSION.String()}))
	assert.Nil(t, err)
	assert.NotNil(t, count)
	assert.Equal(t, count, 2)
	t.Cleanup(func() {
		complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
		assert.True(t, ok)
		recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for _, v := range mock_data.NewComplaints {
			err := complaintRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}
