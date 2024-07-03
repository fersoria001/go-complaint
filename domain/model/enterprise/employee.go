package enterprise

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/identity"

	"github.com/google/uuid"
)

type Employee interface {
	ID() uuid.UUID
	EnterpriseID() string
	HiringDate() common.Date
	Email() string
	Position() Position
	ApprovedHiring() bool
	ApprovedHiringAt() common.Date
	GetUser() *identity.User
	SetPosition(position Position) error
	SetApprovedHiring(approvedHiring bool)
}
