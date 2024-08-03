package repositories_test

import (
	"context"
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/feedback"
	"go-complaint/domain/model/identity"
	"go-complaint/domain/model/recipient"
	"go-complaint/infrastructure/persistence/finders/find_all_feedback_reply_review"
	"go-complaint/infrastructure/persistence/finders/find_recipient"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFeedbackReplyReviewRepository_Setup(t *testing.T) {
	ctx := context.Background()
	userRepository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	recipientRepository, ok := repositories.MapperRegistryInstance().Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewUsers {
		person, err := identity.NewPerson(
			v.Person.Id,
			v.Person.Email,
			v.Person.ProfileImg,
			v.Person.Genre,
			v.Person.Pronoun,
			v.Person.FirstName,
			v.Person.LastName,
			v.Person.Phone,
			v.Person.BirthDate,
			v.Person.Address,
		)
		assert.Nil(t, err)
		user, err := identity.NewUser(
			v.Id,
			v.UserName,
			v.Password,
			v.RegisterDate,
			person,
			v.IsConfirmed,
			v.UserRoles,
		)
		assert.Nil(t, err)
		err = userRepository.Save(ctx, user)
		assert.Nil(t, err)
		userRecipient := recipient.NewRecipient(user.Id(), user.FullName(), user.ProfileIMG(), user.UserName(), false)
		err = recipientRepository.Save(ctx, *userRecipient)
		assert.Nil(t, err)
	}
}
func TestFeedbackReplyReviewRepository_SaveAll(t *testing.T) {
	TestFeedbackReplyReviewRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	userRepository, ok := reg.Get("User").(repositories.UserRepository)
	require.True(t, ok)
	replyReviewRepository, ok := reg.Get("ReplyReview").(repositories.FeedbackReplyReviewRepository)
	require.True(t, ok)
	for _, v := range mock_data.NewReplyReviews {
		user, err := userRepository.Find(ctx, find_user.ByUsername(v.Reviewer.UserName))
		require.NoError(t, err)
		replyReview := feedback.NewReplyReviewEntity(
			v.Id,
			v.FeedbackId,
			*user,
			v.Color,
		)
		replyReviewsSet := mapset.NewSet(*replyReview)
		err = replyReviewRepository.SaveAll(ctx, replyReviewsSet)
		require.NoError(t, err)
	}

	t.Cleanup(func() {
		for k := range mock_data.NewReplyReviews {
			err := replyReviewRepository.DeleteAll(ctx, k)
			require.NoError(t, err)
		}
		for _, v := range mock_data.NewUsers {
			user, err := userRepository.Find(ctx, find_user.ByUsername(v.UserName))
			require.NoError(t, err)
			err = userRepository.Remove(ctx, user.Id())
			require.NoError(t, err)
		}
	})

}

func TestFeedbackReplyReviewRepository_FindAll(t *testing.T) {
	TestFeedbackReplyReviewRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	userRepository, ok := reg.Get("User").(repositories.UserRepository)
	require.True(t, ok)
	replyReviewRepository, ok := reg.Get("ReplyReview").(repositories.FeedbackReplyReviewRepository)
	require.True(t, ok)
	for _, v := range mock_data.NewReplyReviews {
		user, err := userRepository.Find(ctx, find_user.ByUsername(v.Reviewer.UserName))
		require.NoError(t, err)
		replyReview := feedback.NewReplyReviewEntity(
			v.Id,
			v.FeedbackId,
			*user,
			v.Color,
		)
		replyReviewsSet := mapset.NewSet(*replyReview)
		err = replyReviewRepository.SaveAll(ctx, replyReviewsSet)
		require.NoError(t, err)
	}
	for k := range mock_data.NewReplyReviews {
		replyReviews, err := replyReviewRepository.FindAll(ctx, find_all_feedback_reply_review.ByFeedbackID(k))
		require.NoError(t, err)
		require.Equal(t, replyReviews.Cardinality(), 1)
	}
	t.Cleanup(func() {
		for k := range mock_data.NewReplyReviews {
			err := replyReviewRepository.DeleteAll(ctx, k)
			require.NoError(t, err)
		}
		for _, v := range mock_data.NewUsers {
			user, err := userRepository.Find(ctx, find_user.ByUsername(v.UserName))
			require.NoError(t, err)
			err = userRepository.Remove(ctx, user.Id())
			require.NoError(t, err)
		}
	})

}

func TestFeedbackReplyReviewRepository_Update(t *testing.T) {
	TestFeedbackReplyReviewRepository_Setup(t)
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	userRepository, ok := reg.Get("User").(repositories.UserRepository)
	require.True(t, ok)
	replyReviewRepository, ok := reg.Get("ReplyReview").(repositories.FeedbackReplyReviewRepository)
	require.True(t, ok)
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	require.True(t, ok)
	for _, v := range mock_data.NewReplyReviews {
		user, err := userRepository.Find(ctx, find_user.ByUsername(v.Reviewer.UserName))
		require.NoError(t, err)
		replyReview := feedback.NewReplyReviewEntity(
			v.Id,
			v.FeedbackId,
			*user,
			v.Color,
		)
		replyReviewsSet := mapset.NewSet(*replyReview)
		err = replyReviewRepository.SaveAll(ctx, replyReviewsSet)
		require.NoError(t, err)
	}
	replies := make([]*complaint.Reply, 0)
	for _, v := range mock_data.NewReplyReviews {
		for _, reply := range v.Replies {
			sender, err := recipientRepository.Find(ctx, find_recipient.ByNameAndEmail(reply.Sender.SubjectName, reply.Sender.SubjectEmail))
			require.NoError(t, err)
			newReply, err := complaint.NewReply(
				reply.Id,
				reply.ComplaintId,
				*sender,
				reply.Body,
				reply.Read,
				reply.CreatedAt,
				reply.ReadAt,
				reply.UpdatedAt,
			)
			require.NoError(t, err)
			replies = append(replies, newReply)
		}
	}

	for k := range mock_data.NewReplyReviews {
		replyReviews, err := replyReviewRepository.FindAll(ctx, find_all_feedback_reply_review.ByFeedbackID(k))
		require.NoError(t, err)
		require.Equal(t, replyReviews.Cardinality(), 1)
		replyReview, ok := replyReviews.Pop()
		require.True(t, ok)
		for _, v := range replies {
			replyReview.AddReply(*v)
		}
		newMap := mapset.NewSet(*replyReview)
		err = replyReviewRepository.DeleteAll(ctx, k)
		require.NoError(t, err)
		err = replyReviewRepository.SaveAll(ctx, newMap)
		require.NoError(t, err)
	}
	t.Cleanup(func() {

		for k := range mock_data.NewReplyReviews {
			err := replyReviewRepository.DeleteAll(ctx, k)
			require.NoError(t, err)
		}

		for _, v := range mock_data.NewUsers {
			user, err := userRepository.Find(ctx, find_user.ByUsername(v.UserName))
			require.NoError(t, err)
			err = userRepository.Remove(ctx, user.Id())
			require.NoError(t, err)
			err = recipientRepository.Remove(ctx, user.Id())
			require.NoError(t, err)
		}
	})

}
