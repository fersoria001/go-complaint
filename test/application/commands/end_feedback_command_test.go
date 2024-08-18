package commands_test

import (
	"context"
	"go-complaint/application/commands"
	"go-complaint/domain"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/finders/find_enterprise"
	"go-complaint/infrastructure/persistence/finders/find_feedback"
	"go-complaint/infrastructure/persistence/finders/find_recipient"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"reflect"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndFeedbackCommand_Setup(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	userRepository, ok := reg.Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	//complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	//assert.True(t, ok)
	recipientsRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)

	for _, v := range mock_data.NewUsers {
		domain.DomainEventPublisherInstance().Reset()
		c := commands.NewRegisterUserCommand(
			v.UserName,
			v.Password,
			v.Person.FirstName,
			v.Person.LastName,
			v.Person.Genre,
			v.Person.Pronoun,
			v.Person.BirthDate.StringRepresentation(),
			v.Person.Phone,
			v.Person.ProfileImg,
			v.Person.Address.Country().ID(),
			v.Person.Address.CountryState().ID(),
			v.Person.Address.City().ID(),
		)
		err := c.Execute(ctx)
		assert.Nil(t, err)
	}
	for _, v := range mock_data.NewRegisterEnterprises {
		domain.DomainEventPublisherInstance().Reset()
		owner, err := userRepository.Find(ctx, find_user.ByUsername(mock_data.NewUsers["valid"].UserName))
		assert.Nil(t, err)
		c := commands.NewRegisterEnterpriseCommand(
			owner.Id().String(),
			v.Name,
			v.LogoImg,
			v.BannerImg,
			v.Website,
			v.Email,
			v.Phone,
			v.FoundationDate.StringRepresentation(),
			v.Industry.Id,
			v.Address.Country().ID(),
			v.Address.CountryState().ID(),
			v.Address.City().ID(),
		)
		err = c.Execute(ctx)
		assert.Nil(t, err)
	}
	for _, v := range mock_data.NewComplaints {
		author, err := recipientsRepository.Find(ctx, find_recipient.ByNameAndEmail(v.Author.SubjectName, v.Author.SubjectEmail))
		assert.Nil(t, err)
		receiver, err := recipientsRepository.Find(ctx, find_recipient.ByNameAndEmail(v.Receiver.SubjectName, v.Receiver.SubjectEmail))
		assert.Nil(t, err)

		domain.DomainEventPublisherInstance().Subscribe(
			domain.DomainEventSubscriber{
				HandleEvent: func(event domain.DomainEvent) error {
					if ev, ok := event.(*complaint.ComplaintCreated); ok {
						complaintId = ev.Id()
					}
					return nil
				},
				SubscribedToEventType: func() reflect.Type {
					return reflect.TypeOf(&complaint.ComplaintCreated{})
				},
			},
		)
		c := commands.NewCreateNewComplaintCommand(author.Id().String(), receiver.Id().String())
		err = c.Execute(ctx)
		assert.Nil(t, err)
		c1 := commands.NewDescribeComplaintCommand(complaintId.String(), v.Title, v.Description)
		err = c1.Execute(ctx)
		assert.Nil(t, err)
		assert.Greater(t, len(v.Replies), 0)
		mockBody := v.Replies[0].Body
		c2 := commands.NewSendComplaintCommand(complaintId.String(), author.Id().String(), mockBody)
		err = c2.Execute(ctx)
		assert.Nil(t, err)
		for _, repliesMock := range mock_data.NewReplies {
			for _, replyMock := range repliesMock {
				c3 := commands.NewReplyComplaintCommand(author.Id().String(), author.Id().String(), complaintId.String(), replyMock.Body)
				err := c3.Execute(ctx)
				assert.Nil(t, err)
				replyId, ok := cache.InMemoryInstance().Get(complaintId.String())
				assert.True(t, ok)
				assert.NotNil(t, replyId)
			}
		}
		c3 := commands.NewSendComplaintToReviewCommand(receiver.Id().String(), author.Id().String(), complaintId.String())
		err = c3.Execute(ctx)
		assert.Nil(t, err)

	}
}

func TestEndFeedbackCommand_Execute(t *testing.T) {
	TestEndFeedbackCommand_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	userRepository, ok := reg.Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	enterpriseRepository, ok := reg.Get("Enterprise").(repositories.EnterpriseRepository)
	assert.True(t, ok)
	complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	assert.True(t, ok)
	recipientsRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	feedbackRepository, ok := reg.Get("Feedback").(repositories.FeedbackRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewFeedbacks {
		c := commands.NewCreateFeedbackCommand(complaintId.String(),
			v.EnterpriseId.String())
		err := c.Execute(ctx)
		assert.Nil(t, err)
		enterpriseId = v.EnterpriseId
		dbf, err := feedbackRepository.Find(ctx, find_feedback.ByComplaintIdAndEnterpriseId(complaintId, v.EnterpriseId))
		assert.Nil(t, err)
		assert.NotNil(t, dbf)
		dbComplaint, err := complaintRepository.Get(ctx, complaintId)
		assert.Nil(t, err)
		repliesIds := make([]string, 0)
		for _, reply := range dbComplaint.Replies().ToSlice() {
			repliesIds = append(repliesIds, reply.ID().String())

		}
		for _, replyReview := range v.ReplyReview {
			reviewer, err := userRepository.Find(ctx, find_user.ByUsername(replyReview.Reviewer.UserName))
			assert.Nil(t, err)
			c1 := commands.NewAddFeedbackReplyCommand(dbf.Id().String(),
				reviewer.Id().String(), replyReview.Color, []string{repliesIds[len(repliesIds)-1]},
			)
			err = c1.Execute(ctx)
			assert.Nil(t, err)
			c2 := commands.NewAddFeedbackCommentCommand(dbf.Id().String(), replyReview.Color, replyReview.Review.Comment)
			err = c2.Execute(ctx)
			assert.Nil(t, err)
			repliesIds = slices.Delete(repliesIds, len(repliesIds)-1, len(repliesIds))
		}
		reviewerUsername := v.ReplyReview[len(v.ReplyReview)-1].Reviewer.UserName
		reviewer, err := userRepository.Find(ctx, find_user.ByUsername(reviewerUsername))
		assert.Nil(t, err)
		c3 := commands.NewEndFeedbackCommand(dbf.Id().String(), reviewer.Id().String())
		err = c3.Execute(ctx)
		assert.Nil(t, err)
		dbf, err = feedbackRepository.Find(ctx, find_feedback.ByComplaintIdAndEnterpriseId(complaintId, v.EnterpriseId))
		assert.Nil(t, err)
		assert.NotNil(t, dbf)
		assert.True(t, dbf.IsDone())
		dbc, err := complaintRepository.Get(ctx, complaintId)
		assert.Nil(t, err)
		assert.Equal(t, complaint.IN_HISTORY, dbc.Status())
	}
	t.Cleanup(func() {
		dbf, err := feedbackRepository.Find(ctx, find_feedback.ByComplaintIdAndEnterpriseId(complaintId, enterpriseId))
		assert.Nil(t, err)
		err = feedbackRepository.Remove(ctx, dbf.Id())
		assert.Nil(t, err)
		err = complaintRepository.Remove(ctx, complaintId)
		assert.Nil(t, err)
		for _, v := range mock_data.NewRegisterEnterprises {
			dbe, err := enterpriseRepository.Find(ctx, find_enterprise.ByName(v.Name))
			assert.Nil(t, err)
			err = enterpriseRepository.Remove(ctx, dbe.Id())
			assert.Nil(t, err)
			err = recipientsRepository.Remove(ctx, dbe.Id())
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewUsers {
			user, err := userRepository.Find(ctx, find_user.ByUsername(v.UserName))
			assert.Nil(t, err)
			err = userRepository.Remove(ctx, user.Id())
			assert.Nil(t, err)
			err = recipientsRepository.Remove(ctx, user.Id())
			assert.Nil(t, err)
		}
	})

}
