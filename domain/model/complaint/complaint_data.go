package complaint

import (
	"time"

	"github.com/google/uuid"
)

type ComplaintData struct {
	id          uuid.UUID
	ownerId     uuid.UUID
	complaintId uuid.UUID
	occurredOn  time.Time
	dataType    ComplaintDataType
}

func NewComplaintData(id, ownerId, complaintId uuid.UUID, occurredOn time.Time, dataType ComplaintDataType) *ComplaintData {
	return &ComplaintData{
		id:          id,
		ownerId:     ownerId,
		complaintId: complaintId,
		occurredOn:  occurredOn,
		dataType:    dataType,
	}
}

func (cd ComplaintData) Id() uuid.UUID {
	return cd.id
}

func (cd ComplaintData) OwnerId() uuid.UUID {
	return cd.ownerId
}

func (cd ComplaintData) ComplaintId() uuid.UUID {
	return cd.complaintId
}

func (cd ComplaintData) OccurredOn() time.Time {
	return cd.occurredOn
}

func (cd ComplaintData) DataType() ComplaintDataType {
	return cd.dataType
}
