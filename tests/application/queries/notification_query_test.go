package queries_test

import (
	"context"
	"go-complaint/application/queries"
	"testing"
)

func TestNotificationQuery(t *testing.T) {
	ctx := context.Background()
	q := queries.NotificationQuery{
		OwnerID: "notifications:bercho001@gmail.com?notifications:Go-Complaint",
	}
	notifications, err := q.Notifications(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(notifications) == 0 {
		t.Fatal("expected notifications, got none")
	}
	t.Logf("notifications: %v", notifications)
}
