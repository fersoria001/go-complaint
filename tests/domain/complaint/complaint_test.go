package complaint_test

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/tests"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestComplaint(t *testing.T) {
	author := tests.UserClient
	receiver := tests.Enterprise1
	id := uuid.New()
	status := complaint.OPEN
	msg, err := complaint.NewMessage(
		tests.RepeatString("t", 11),
		tests.RepeatString("d", 31),
		tests.RepeatString("b", 51),
	)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	commonDate := common.NewDate(time.Now())
	rating, err := complaint.NewRating(5, "good")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	emptySet := mapset.NewSet[*complaint.Reply]()
	newReply, err := complaint.NewReply(
		id,
		id,
		receiver.Email(),
		receiver.LogoIMG(),
		receiver.Name(),
		"reply body",
		false,
		commonDate,
		commonDate,
		commonDate,
		true,
		"FloatingPoint. Ltd",
	)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	newComplaint, err := complaint.NewComplaint(
		id,
		author.Email(),
		author.FullName(),
		author.ProfileIMG(),
		receiver.Email(),
		receiver.Name(),
		receiver.LogoIMG(),
		status,
		msg,
		commonDate,
		commonDate,
		rating,
		emptySet,
	)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	newComplaint.AddReply(newReply)
	assert.NotNil(t, newComplaint)
	assert.Equal(t, newComplaint.ID(), id)
	assert.Equal(t, newComplaint.AuthorID(), author.Email())
	assert.Equal(t, newComplaint.ReceiverID(), receiver.Email())
	assert.Equal(t, newComplaint.Status(), status)
	assert.Equal(t, newComplaint.Message(), msg)
	assert.Equal(t, newComplaint.CreatedAt(), commonDate)
	assert.Equal(t, newComplaint.UpdatedAt(), commonDate)
	assert.Equal(t, newComplaint.Rating(), rating)
	assert.Equal(t, newComplaint.Replies().Cardinality(), 1)
}
