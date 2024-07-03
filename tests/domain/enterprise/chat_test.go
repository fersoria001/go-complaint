package enterprise_test

import (
	"go-complaint/domain/model/enterprise"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplyChat(t *testing.T) {
	id, err := enterprise.NewChatID("chat:fernandosoria1379@gmail.com#bercho001@gmail.com")
	assert.Nil(t, err)
	assert.Equal(t, "chat:fernandosoria1379@gmail.com#bercho001@gmail.com", id.String())
	id1, err := enterprise.NewChatID("chat:Go-Complaint#fernandosoria1379@gmail.com")
	assert.Nil(t, err)
	assert.Equal(t, "chat:Go-Complaint#fernandosoria1379@gmail.com", id1.String())
}
