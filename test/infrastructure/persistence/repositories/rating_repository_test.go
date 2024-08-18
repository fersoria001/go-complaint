package repositories_test

import (
	"context"
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/recipient"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRatingRepository_Save(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Rating").(repositories.RatingRepository)
	assert.True(t, ok)
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewRecipients {
		err := recipientRepository.Save(ctx, *recipient.NewRecipient(
			v.Id,
			v.SubjectName,
			v.SubjectThumbnail,
			v.SubjectEmail,
			v.IsEnterprise,
		))
		assert.Nil(t, err)
	}
	for _, v := range mock_data.NewComplaints {
		author, err := recipientRepository.Get(ctx, v.Author.Id)
		assert.Nil(t, err)
		receiver, err := recipientRepository.Get(ctx, v.Receiver.Id)
		assert.Nil(t, err)
		rating := complaint.NewRating(v.Rating.Id, *receiver, *author, v.Rating.Rate, v.Rating.Comment, time.Now(), time.Now())
		err = repository.Save(ctx, *rating)
		assert.Nil(t, err)
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewComplaints {
			err := repository.Remove(ctx, v.Rating.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestRatingRepository_Get(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Rating").(repositories.RatingRepository)
	assert.True(t, ok)
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewRecipients {
		err := recipientRepository.Save(ctx, *recipient.NewRecipient(
			v.Id,
			v.SubjectName,
			v.SubjectThumbnail,
			v.SubjectEmail,
			v.IsEnterprise,
		))
		assert.Nil(t, err)
	}

	for _, v := range mock_data.NewComplaints {
		author, err := recipientRepository.Get(ctx, v.Author.Id)
		assert.Nil(t, err)
		receiver, err := recipientRepository.Get(ctx, v.Receiver.Id)
		assert.Nil(t, err)
		rating := complaint.NewRating(v.Rating.Id, *receiver, *author, v.Rating.Rate, v.Rating.Comment, time.Now(), time.Now())
		err = repository.Save(ctx, *rating)
		assert.Nil(t, err)
		err = repository.Update(ctx, *rating)
		assert.Nil(t, err)
		dbR, err := repository.Get(ctx, v.Rating.Id)
		assert.Nil(t, err)
		assert.Equal(t, v.Rating.Rate, dbR.Rate())
		assert.Equal(t, v.Rating.Comment, dbR.Comment())
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewComplaints {
			err := repository.Remove(ctx, v.Rating.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewRecipients {
			err := recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}
