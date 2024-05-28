package persistence_test

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/tests"
	"testing"
	"time"

	"github.com/google/uuid"
)

const fixedID = "1b3ece75-4925-43e5-b8d0-61bca46b87d1"

func TestComplaintRepository(t *testing.T) {
	ctx := context.Background()
	messagge, err := complaint.NewMessage(tests.Repeat("a", 10), tests.Repeat("a", 30), tests.Repeat("a", 100))
	if err != nil {
		t.Errorf("error at create message, got %v", err)
	}
	genericDate := common.NewDate(time.Now())
	parseID, err := uuid.Parse(fixedID)
	if err != nil {
		t.Errorf("error at parse id, got %v", err)
	}
	complain, err := complaint.NewComplaint(
		parseID,
		"authorID@gmail.com",
		"receiverID@gmail.com",
		complaint.OPEN,
		messagge,
		genericDate,
		genericDate,
		complaint.Rating{})
	if err != nil {
		t.Errorf("error at create complaint, got %v", err)
	}

	schema := datasource.ComplaintSchema()

	repo := repositories.NewComplaintRepository(schema)

	t.Run("given a valid complaint, it should save it", func(t *testing.T) {
		err := repo.Save(ctx, complain)
		if err != nil {
			t.Errorf("error at save complaint, got %v", err)
		}
	})
	t.Run("given a valid and previously saved complaint, it should return it", func(t *testing.T) {
		complaint, err := repo.Get(ctx, complain.ID().String())
		if err != nil {
			t.Errorf("error at find complaint, got %v", err)
		}
		if complaint == nil {
			t.Errorf("expected complaint, got nil")
		}
	})
	t.Run("given a valid and previously saved complaint, it should delete it", func(t *testing.T) {
		err := repo.Remove(ctx, complain.ID().String())
		if err != nil {
			t.Errorf("error at delete complaint, got %v", err)
		}
	})

}
