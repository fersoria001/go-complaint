package repositories

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/tests"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNotificationRepository_Create(t *testing.T) {
	// Setup
	ctx := context.Background()
	newNotification, err := domain.NewNotification(
		uuid.New(),
		"owner_id",
		tests.UserRegisterAndVerifyCommands["1"].Email,
		"title",
		"content",
		"link",
		time.Now(),
		false,
	)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	err = repositories.MapperRegistryInstance().Get("Notification").(repositories.NotificationRepository).Save(ctx, newNotification)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	n, err := repositories.MapperRegistryInstance().Get("Notification").(repositories.NotificationRepository).Get(ctx, newNotification.ID())
	if err != nil {

		t.Fatalf("Error: %v", err)
	}
	if n.ID() != newNotification.ID() {
		t.Fatalf("Error: ID not match")
	}
	if n.OwnerID() != newNotification.OwnerID() {
		t.Fatalf("Error: OwnerID not match")
	}
	if n.Thumbnail() != "/default.jpg" {

		t.Fatalf("Error: Thumbnail not match")
	}
	if n.Title() != newNotification.Title() {
		t.Fatalf("Error: Title not match")
	}
	if n.Content() != newNotification.Content() {
		t.Fatalf("Error: Content not match")
	}
	if n.Link() != newNotification.Link() {
		t.Fatalf("Error: Link not match")
	}
	if common.StringDate(n.OccurredOn()) != common.StringDate(newNotification.OccurredOn()) {
		t.Fatalf("Error: CreatedAt not match")
	}
	if n.Seen() != newNotification.Seen() {
		t.Fatalf("Error: Seen not match")
	}
}
