package repositories_test

import (
	"context"
	"go-complaint/domain/model/feedback"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFeedbackRepository_Save(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Feedback").(repositories.FeedbackRepository)
	require.True(t, ok)
	for _, v := range mock_data.NewFeedbacks {
		f := feedback.NewFeedbackEntity(
			v.Id,
			v.ComplaintId,
			v.EnterpriseId,
		)
		err := r.Save(ctx, f)
		require.NoError(t, err)
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewFeedbacks {
			err := r.Remove(ctx, v.Id)
			require.NoError(t, err)
		}
	})
}

func TestFeedbackRepository_Get(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Feedback").(repositories.FeedbackRepository)
	require.True(t, ok)
	for _, v := range mock_data.NewFeedbacks {
		f := feedback.NewFeedbackEntity(
			v.Id,
			v.ComplaintId,
			v.EnterpriseId,
		)
		err := r.Save(ctx, f)
		require.NoError(t, err)
		f, err = r.Get(ctx, v.Id)
		require.NoError(t, err)
		require.NotNil(t, f)
		require.Equal(t, v.ComplaintId, f.ComplaintId())
		require.Equal(t, v.EnterpriseId, f.EnterpriseId())
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewFeedbacks {
			err := r.Remove(ctx, v.Id)
			require.NoError(t, err)
		}
	})
}
