package complaint_test

import (
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/recipient"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReply_New(t *testing.T) {
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
			)
			assert.Nil(t, err)
			assert.NotNil(t, r)
			assert.Equal(t, v.Id, r.ID())
			assert.Equal(t, v.ComplaintId, r.ComplaintId())
			assert.Equal(t, v.Sender.Id, r.Sender().Id())
			assert.Equal(t, v.Sender.SubjectName, r.Sender().SubjectName())
			assert.Equal(t, v.Sender.SubjectThumbnail, r.Sender().SubjectThumbnail())
			assert.Equal(t, v.Sender.IsEnterprise, r.Sender().IsEnterprise())
			assert.Equal(t, v.Body, r.Body())
			assert.Equal(t, v.Read, r.Read())
			assert.Equal(t, v.CreatedAt.StringRepresentation(), r.CreatedAt().StringRepresentation())
			assert.Equal(t, v.ReadAt.StringRepresentation(), r.ReadAt().StringRepresentation())
			assert.Equal(t, v.UpdatedAt.StringRepresentation(), r.UpdatedAt().StringRepresentation())
		}
	}
}
