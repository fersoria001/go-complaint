package tests

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func NewEmployeeID(enterpriseName, managerEmail, userID string) string {
	return enterpriseName + "-" + managerEmail + "-" + time.Now().Format("02/01/2006") + "-" +
		strings.Split(uuid.New().String(), "-")[4]
}
