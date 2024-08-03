package commands_test

import (
	"context"
	"go-complaint/application/commands"
	"go-complaint/infrastructure/persistence/finders/find_feedback"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFeedbackCommandTest(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Feedback").(repositories.FeedbackRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewFeedbacks {
		c := commands.NewCreateFeedbackCommand(v.ComplaintId.String(),
			v.EnterpriseId.String())
		err := c.Execute(ctx)
		assert.Nil(t, err)
		dbf, err := r.Find(ctx, find_feedback.ByComplaintId(v.ComplaintId))
		assert.Nil(t, err)
		assert.NotNil(t, dbf)
	}

	t.Cleanup(func() {
		for _, v := range mock_data.NewFeedbacks {
			dbf, err := r.Find(ctx, find_feedback.ByComplaintId(v.ComplaintId))
			assert.Nil(t, err)
			err = r.Remove(ctx, dbf.Id())
			assert.Nil(t, err)
		}
	})
}
