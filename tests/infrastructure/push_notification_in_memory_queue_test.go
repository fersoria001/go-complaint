package infrastructure_test

import (
	"go-complaint/infrastructure"
	"testing"
)

func TestPushNotificationInMemoryQueueInstance(t *testing.T) {
	queue := infrastructure.PushNotificationInMemoryQueueInstance()
	if queue == nil {
		t.Errorf("Error: %v", queue)
	}
}
