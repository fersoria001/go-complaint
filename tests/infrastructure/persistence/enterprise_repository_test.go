package persistence_test

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestEnterpriseRepository_EMPTY_EMPLOYEES(t *testing.T) {
	schema := datasource.EnterpriseSchema()

	ctx := context.Background()
	err := schema.Connect(ctx)
	if err != nil {
		t.Errorf("error at connect schema, got %v", err)
	}
	repository := repositories.NewEnterpriseRepository(schema)
	//The aggregate has the identity of the root entity
	ownerID := "owner@gmail.com"
	id := "enterpriseName"
	address, err := common.NewAddress("country", "county", "city")
	if err != nil {
		t.Errorf("error at create address, got %v", err)
	}
	sameDate := common.NewDate(time.Now())
	industry, err := enterprise.NewIndustry(0, "industry")
	if err != nil {
		t.Errorf("error at create industry, got %v", err)
	}
	employees := make(map[string]*enterprise.Employee)
	ddmmyy := time.Now().Format("02-01-2006")
	uuidSlice := strings.Split(uuid.New().String(), "-")[4]
	employeeID := id + "owner@gmail.com" + ddmmyy + uuidSlice
	today := common.NewDate(time.Now())
	manager, err := enterprise.NewEmployee(
		employeeID,
		"profile.jpg",
		"Brauford",
		"Vulkoff",
		30,
		"BrauforVulkoff@live.com",
		"0123456789",
		enterprise.MANAGER,
		today, true, today)
	if err != nil {
		t.Errorf("error at create employee, got %v", err)
	}
	employees[manager.Email()] = manager
	enterprise, err := enterprise.NewEnterprise(
		ownerID,
		id,
		"/banner.jpg",
		"logo.png",
		"awseomwebsite.com",
		"enteprrise@gmail.com",
		"01234567890",
		address,
		industry,
		sameDate,
		sameDate,
	)
	if err != nil {
		t.Errorf("error at create enterprise, got %v", err)
	}
	t.Run("Save", func(t *testing.T) {
		err := repository.Save(ctx, enterprise)
		if err != nil {
			t.Errorf("error at save enterprise, got %v", err)
		}
		result, err := repository.Get(ctx, id)
		if err != nil {
			t.Errorf("error at get enterprise, got %v", err)
		}
		results, err := repository.GetAll(ctx)
		if err != nil {
			t.Errorf("error at get all enterprises, got %v", err)
		}
		if results.Cardinality() == 0 {
			t.Errorf("error at get all enterprises, got %v", results.Cardinality())
		}

		if result == nil {
			t.Errorf("error at get enterprise, got %v", result)
		}
		err = repository.Remove(ctx, id)
		if err != nil {
			t.Errorf("error at remove enterprise, got %v", err)
		}
	})

}
