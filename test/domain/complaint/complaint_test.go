package complaint_test

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/recipient"
	"go-complaint/test/mock_data"
	"reflect"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

func TestComplaint_New(t *testing.T) {
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
		rating := &complaint.Rating{}
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
			mockReply.Sender.IsEnterprise,
			mockReply.Sender.Id,
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
		assert.Equal(t, v.Id, newComplaint.Id())
		assert.Equal(t, v.Author.Id, newComplaint.Author().Id())
		assert.Equal(t, v.Author.SubjectName, newComplaint.Author().SubjectName())
		assert.Equal(t, v.Author.SubjectThumbnail, newComplaint.Author().SubjectThumbnail())
		assert.Equal(t, v.Author.IsEnterprise, newComplaint.Author().IsEnterprise())
		assert.Equal(t, v.Receiver.Id, newComplaint.Receiver().Id())
		assert.Equal(t, v.Receiver.SubjectName, newComplaint.Receiver().SubjectName())
		assert.Equal(t, v.Receiver.SubjectThumbnail, newComplaint.Receiver().SubjectThumbnail())
		assert.Equal(t, v.Receiver.IsEnterprise, newComplaint.Receiver().IsEnterprise())
		assert.Equal(t, v.Status, newComplaint.Status())
		assert.Equal(t, v.Title, newComplaint.Title())
		assert.Equal(t, v.Description, newComplaint.Description())
		assert.Equal(t, v.CreatedAt, newComplaint.CreatedAt())
		assert.Equal(t, v.UpdatedAt, newComplaint.UpdatedAt())
		assert.Equal(t, rating.Id(), newComplaint.Rating().Id())
		assert.Equal(t, rating.Rate(), newComplaint.Rating().Rate())
		assert.Equal(t, rating.Comment(), newComplaint.Rating().Comment())
		assert.Equal(t, 1, newComplaint.Replies().Cardinality())
		assert.Equal(t, mockReply.Body, newComplaint.Body())
	}
}

func TestComplaint_CreateNew(t *testing.T) {
	ctx := context.Background()
	for _, v := range mock_data.NewComplaints {
		c := 0
		domain.DomainEventPublisherInstance().Subscribe(
			domain.DomainEventSubscriber{
				HandleEvent: func(event domain.DomainEvent) error {
					if _, ok := event.(*complaint.ComplaintCreated); !ok {
						t.Fatalf("Incorrect EventType in Subscriber")
					}
					c++
					return nil
				},
				SubscribedToEventType: func() reflect.Type {
					return reflect.TypeOf(&complaint.ComplaintCreated{})
				},
			},
		)
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
		assert.Equal(t, 1, c)
		assert.Nil(t, err)
		assert.NotNil(t, newComplaint)
		assert.Equal(t, v.Id, newComplaint.Id())
		assert.Equal(t, v.Author.Id, newComplaint.Author().Id())
		assert.Equal(t, v.Author.SubjectName, newComplaint.Author().SubjectName())
		assert.Equal(t, v.Author.SubjectThumbnail, newComplaint.Author().SubjectThumbnail())
		assert.Equal(t, v.Author.IsEnterprise, newComplaint.Author().IsEnterprise())
		assert.Equal(t, v.Receiver.Id, newComplaint.Receiver().Id())
		assert.Equal(t, v.Receiver.SubjectName, newComplaint.Receiver().SubjectName())
		assert.Equal(t, v.Receiver.SubjectThumbnail, newComplaint.Receiver().SubjectThumbnail())
		assert.Equal(t, v.Receiver.IsEnterprise, newComplaint.Receiver().IsEnterprise())
	}
}

func TestComplaint_SetTitle(t *testing.T) {
	ctx := context.Background()
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
		c := 0
		domain.DomainEventPublisherInstance().Subscribe(
			domain.DomainEventSubscriber{
				HandleEvent: func(event domain.DomainEvent) error {
					if _, ok := event.(*complaint.ComplaintTitleSet); !ok {
						t.Fatalf("Incorrect EventType in Subscriber")
					}
					c++
					return nil
				},
				SubscribedToEventType: func() reflect.Type {
					return reflect.TypeOf(&complaint.ComplaintTitleSet{})
				},
			},
		)
		err = newComplaint.SetTitle(ctx, v.Title)
		assert.Equal(t, 1, c)
		assert.Nil(t, err)
		assert.Equal(t, v.Title, newComplaint.Title())
	}
}

func TestComplaint_SetDescription(t *testing.T) {
	ctx := context.Background()
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
		err = newComplaint.SetTitle(ctx, v.Title)
		assert.Nil(t, err)
		c := 0
		domain.DomainEventPublisherInstance().Subscribe(
			domain.DomainEventSubscriber{
				HandleEvent: func(event domain.DomainEvent) error {
					if _, ok := event.(*complaint.ComplaintDescriptionSet); !ok {
						t.Fatalf("Incorrect EventType in Subscriber")
					}
					c++
					return nil
				},
				SubscribedToEventType: func() reflect.Type {
					return reflect.TypeOf(&complaint.ComplaintDescriptionSet{})
				},
			},
		)
		err = newComplaint.SetDescription(ctx, v.Description)
		assert.Nil(t, err)
		assert.Equal(t, 1, c)
		assert.Equal(t, v.Description, newComplaint.Description())
	}
}

func TestComplaint_SetBody(t *testing.T) {
	ctx := context.Background()
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
		err = newComplaint.SetTitle(ctx, v.Title)
		assert.Nil(t, err)
		err = newComplaint.SetDescription(ctx, v.Description)
		assert.Nil(t, err)
		c := 0
		domain.DomainEventPublisherInstance().Subscribe(
			domain.DomainEventSubscriber{
				HandleEvent: func(event domain.DomainEvent) error {
					if _, ok := event.(*complaint.ComplaintBodySet); !ok {
						t.Fatalf("Incorrect EventType in Subscriber")
					}
					c++
					return nil
				},
				SubscribedToEventType: func() reflect.Type {
					return reflect.TypeOf(&complaint.ComplaintBodySet{})
				},
			},
		)
		assert.Greater(t, len(v.Replies), 0)
		mockBody := v.Replies[len(v.Replies)-1].Body
		assert.NotNil(t, mockBody)
		err = newComplaint.SetBody(ctx, mockBody)
		assert.Nil(t, err)
		assert.Equal(t, 1, c)
		assert.Equal(t, mockBody, newComplaint.Body())
	}
}
