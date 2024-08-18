package complaint

import (
	"time"

	"github.com/google/uuid"
)

type ComplaintData struct {
	id          uuid.UUID
	ownerId     uuid.UUID
	authorId    uuid.UUID
	receiverId  uuid.UUID
	complaintId uuid.UUID
	occurredOn  time.Time
	dataType    ComplaintDataType
}

func NewComplaintData(id, ownerId, authorId, receiverId, complaintId uuid.UUID, occurredOn time.Time, dataType ComplaintDataType) *ComplaintData {
	return &ComplaintData{
		id:          id,
		ownerId:     ownerId,
		authorId:    authorId,
		receiverId:  receiverId,
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

func (cd ComplaintData) AuthorId() uuid.UUID {
	return cd.authorId
}

func (cd ComplaintData) ReceiverId() uuid.UUID {
	return cd.receiverId
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
