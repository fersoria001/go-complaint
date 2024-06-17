package employee_test

import (
	"go-complaint/domain/model/employee"
	"go-complaint/domain/model/enterprise"
	"go-complaint/tests"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEmployee(t *testing.T) {
	//Arrange
	newID := uuid.New()
	//Act
	newEmployee, err := employee.NewEmployee(
		newID,
		"FloatingPoint. Ltd",
		tests.UserAssistant,
		enterprise.ASSISTANT,
		tests.CommonDate,
		false,
		tests.CommonDate,
	)
	//Assert
	assert.Nil(t, err)
	assert.NotNil(t, newEmployee)
}
