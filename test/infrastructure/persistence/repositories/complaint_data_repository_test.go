package repositories_test

import (
	"context"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/persistence/finders/find_all_complaint_data"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplaintDataRepository_Save(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("ComplaintData").(repositories.ComplaintDataRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewComplaintData {
		cd := complaint.NewComplaintData(v.Id, v.OwnerId, v.ComplaintId, v.OccurredOn, v.DataType)
		err := repository.Save(ctx, *cd)
		assert.Nil(t, err)
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewComplaintData {
			err := repository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestComplaintDataRepository_Get(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("ComplaintData").(repositories.ComplaintDataRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewComplaintData {
		cd := complaint.NewComplaintData(v.Id, v.OwnerId, v.ComplaintId, v.OccurredOn, v.DataType)
		err := repository.Save(ctx, *cd)
		assert.Nil(t, err)
		dbCd, err := repository.Get(ctx, v.Id)
		assert.Nil(t, err)
		assert.NotNil(t, dbCd)
		assert.Equal(t, v.Id, dbCd.Id())
		assert.Equal(t, v.OwnerId, dbCd.OwnerId())
		assert.Equal(t, v.ComplaintId, dbCd.ComplaintId())
		assert.Equal(t, v.DataType, dbCd.DataType())
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewComplaintData {
			err := repository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestComplaintDataRepository_Find(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("ComplaintData").(repositories.ComplaintDataRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewComplaintData {
		cd := complaint.NewComplaintData(v.Id, v.OwnerId, v.ComplaintId, v.OccurredOn, v.DataType)
		err := repository.Save(ctx, *cd)
		assert.Nil(t, err)
	}
	ownerId := mock_data.NewUsers["valid"].Id
	dbCd, err := repository.FindAll(ctx, find_all_complaint_data.ByOwnerId(ownerId))
	assert.Nil(t, err)
	assert.NotNil(t, dbCd)
	assert.Equal(t, len(dbCd), 3)
	t.Cleanup(func() {
		for _, v := range mock_data.NewComplaintData {
			err := repository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}
