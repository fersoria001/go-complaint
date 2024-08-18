package commands

import (
	"context"
	"go-complaint/application/commands"
	"go-complaint/domain"
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/recipient"
	"go-complaint/infrastructure/persistence/finders/find_all_complaint_data"
	"go-complaint/infrastructure/persistence/finders/find_recipient"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSendComplaintCommand_Setup(t *testing.T) {
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

func TestSendComplaintCommand_Execute(t *testing.T) {
	TestSendComplaintCommand_Setup(t)
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
	complaintDataRepository, ok := reg.Get("ComplaintData").(repositories.ComplaintDataRepository)
	assert.True(t, ok)
	complaintDatas := make([]*complaint.ComplaintData, 0)
	recipients := make([]*recipient.Recipient, 0)
	complaints := make([]*complaint.Complaint, 0)
	for _, v := range mock_data.NewComplaints {
		author, err := recipientsRepository.Find(ctx, find_recipient.ByNameAndEmail(v.Author.SubjectName, v.Author.SubjectEmail))
		assert.Nil(t, err)
		receiver, err := recipientsRepository.Find(ctx, find_recipient.ByNameAndEmail(v.Receiver.SubjectName, v.Receiver.SubjectEmail))
		assert.Nil(t, err)
		recipients = append(recipients, author)
		recipients = append(recipients, receiver)
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
		c2 := commands.NewSendComplaintCommand(complaintId.String(), author.Id().String(), mockBody)
		err = c2.Execute(ctx)
		assert.Nil(t, err)
		updatedC, err := complaintRepository.Get(ctx, complaintId)
		assert.Nil(t, err)
		complaints = append(complaints, updatedC)
		assert.Equal(t, v.Title, updatedC.Title())
		assert.Equal(t, v.Description, updatedC.Description())
		assert.Equal(t, mockBody, updatedC.Body())
		assert.Equal(t, complaint.OPEN, updatedC.Status())
		complaintData, err := complaintDataRepository.FindAll(ctx, find_all_complaint_data.ByOwnerIdAndDataOwnership(author.Id()))
		assert.Nil(t, err)
		assert.NotNil(t, complaintData)
		assert.Equal(t, len(complaintData), 1)
		complaintData1, err := complaintDataRepository.FindAll(ctx, find_all_complaint_data.ByOwnerIdAndDataOwnership(receiver.Id()))
		assert.Nil(t, err)
		assert.NotNil(t, complaintData1)
		assert.Equal(t, len(complaintData1), 1)
		complaintDatas = append(complaintDatas, complaintData...)
		complaintDatas = append(complaintDatas, complaintData1...)
	}
	t.Cleanup(func() {
		for _, v := range complaintDatas {
			err := complaintDataRepository.Remove(ctx, v.Id())
			assert.Nil(t, err)
		}
		for _, v := range complaints {
			err := complaintRepository.Remove(ctx, v.Id())
			assert.Nil(t, err)
		}
		for _, v := range recipients {
			if v.IsEnterprise() {
				err := enterpriseRepository.Remove(ctx, v.Id())
				assert.Nil(t, err)
			} else {
				err := userRepository.Remove(ctx, v.Id())
				assert.Nil(t, err)
			}
			err := recipientsRepository.Remove(ctx, v.Id())
			assert.Nil(t, err)
		}
	})
}
