package repositories_test

import (
	"context"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRatingRepository_Save(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Rating").(repositories.RatingRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewComplaints {
		rating, err := complaint.NewRating(v.Rating.Id, v.Rating.Rate, v.Rating.Comment)
		assert.Nil(t, err)
		err = repository.Save(ctx, rating)
		assert.Nil(t, err)
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewComplaints {
			err := repository.Remove(ctx, v.Rating.Id)
			assert.Nil(t, err)
		}
	})
}

func TestRatingRepository_Get(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Rating").(repositories.RatingRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewComplaints {
		rating, err := complaint.NewRating(v.Rating.Id, v.Rating.Rate, v.Rating.Comment)
		assert.Nil(t, err)
		err = repository.Save(ctx, rating)
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
	})
}
