package commands_test

import (
	"context"
	"go-complaint/application/commands"
	"go-complaint/domain"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/persistence/finders/find_recipient"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestReplyComplaintCommand_Setup(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	userRepository, ok := reg.Get("User").(repositories.UserRepository)
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

}

func TestReplyComplaintCommand_Execute(t *testing.T) {
	TestReplyComplaintCommand_Setup(t)
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
	for _, v := range mock_data.NewComplaints {
		author, err := recipientsRepository.Find(ctx, find_recipient.ByNameAndEmail(v.Author.SubjectName, v.Author.SubjectEmail))
		assert.Nil(t, err)
		receiver, err := recipientsRepository.Find(ctx, find_recipient.ByNameAndEmail(v.Receiver.SubjectName, v.Receiver.SubjectEmail))
		assert.Nil(t, err)
		var complaintId uuid.UUID
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
		c2 := commands.NewSendComplaintCommand(complaintId.String(), mockBody)
		err = c2.Execute(ctx)
		assert.Nil(t, err)
		for _, repliesMock := range mock_data.NewReplies {
			for _, replyMock := range repliesMock {
				c3 := commands.NewReplyComplaintCommand(author.Id().String(), complaintId.String(), replyMock.Body)
				err := c3.Execute(ctx)
				assert.Nil(t, err)
			}
		}
		dbC, err := complaintRepository.Get(ctx, complaintId)
		assert.Nil(t, err)
		assert.Equal(t, 3, dbC.Replies().Cardinality())
	}

	t.Cleanup(func() {
		for _, v := range mock_data.NewComplaints {
			err := complaintRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewRegisterEnterprises {
			err := enterpriseRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
			err = recipientsRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewUsers {
			err := userRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
			err = recipientsRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}
