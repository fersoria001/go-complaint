package complaint

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/recipient"
	"go-complaint/erros"
	"time"

	"github.com/google/uuid"
)

// Package complaint
// << Entity >>
// It represents the reply of the complaint
// Its complaintID is the id of the complaint that own this reply
type Reply struct {
	id           uuid.UUID
	complaintId  uuid.UUID
	sender       recipient.Recipient
	body         string
	createdAt    common.Date
	read         bool
	readAt       common.Date
	updatedAt    common.Date
	isEnterprise bool
	enterpriseId uuid.UUID
}

func (r *Reply) MarkAsRead() {
	r.read = true
	r.readAt = common.NewDate(time.Now())
}

func CreateReply(
	id uuid.UUID,
	complaintId uuid.UUID,
	sender recipient.Recipient,
	body string,
	isEnterprise bool,
	enterpriseId uuid.UUID,
) *Reply {
	newCommonDate := common.NewDate(time.Now())
	r := &Reply{
		id:           id,
		complaintId:  complaintId,
		sender:       sender,
		body:         body,
		read:         false,
		readAt:       newCommonDate,
		createdAt:    newCommonDate,
		updatedAt:    newCommonDate,
		isEnterprise: isEnterprise,
		enterpriseId: enterpriseId,
	}
	return r
}

func NewReply(
	id,
	complaintId uuid.UUID,
	sender recipient.Recipient,
	body string,
	read bool,
	createdAt,
	readAt,
	updatedAt common.Date,
	isEnterprise bool,
	enterpriseId uuid.UUID,
) (*Reply, error) {
	var reply *Reply = new(Reply)
	reply.isEnterprise = isEnterprise
	reply.enterpriseId = enterpriseId
	reply.sender = sender
	reply.complaintId = complaintId
	err := reply.setID(id)
	if err != nil {
		return nil, err
	}
	err = reply.setBody(body)
	if err != nil {
		return nil, err
	}
	err = reply.setCreatedAt(createdAt)
	if err != nil {
		return nil, err
	}
	reply.setRead(read)
	err = reply.setReadAt(readAt)
	if err != nil {
		return nil, err
	}
	err = reply.setUpdatedAt(updatedAt)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (r *Reply) setID(id uuid.UUID) error {
	if id == uuid.Nil {
		return &erros.NullValueError{}
	}
	r.id = id
	return nil
}

func (r *Reply) setBody(body string) error {
	if body == "" {
		return &erros.NullValueError{}
	}
	if len(body) > 120 {
		return &erros.InvalidLengthError{
			AttributeName: "body",
			MinLength:     1,
			MaxLength:     120,
			CurrentLength: len(body),
		}
	}
	r.body = body
	return nil
}
func (r *Reply) setRead(read bool) {
	r.read = read
}

func (r *Reply) setReadAt(readAt common.Date) error {
	if readAt == (common.Date{}) {
		return &erros.NullValueError{}
	}
	r.readAt = readAt
	return nil
}

func (r *Reply) setCreatedAt(createdAt common.Date) error {
	if createdAt == (common.Date{}) {
		return &erros.EmptyStructError{}
	}
	r.createdAt = createdAt
	return nil
}

func (r *Reply) setUpdatedAt(updatedAt common.Date) error {
	if updatedAt == (common.Date{}) {
		return &erros.EmptyStructError{}
	}
	r.updatedAt = updatedAt
	return nil
}

func (r Reply) ID() uuid.UUID {
	return r.id
}

func (r Reply) ComplaintId() uuid.UUID {
	return r.complaintId
}

func (r Reply) Sender() recipient.Recipient {
	return r.sender
}

func (r Reply) Body() string {
	return r.body
}

func (r Reply) CreatedAt() common.Date {
	return r.createdAt
}

func (r Reply) Read() bool {
	return r.read
}

func (r Reply) ReadAt() common.Date {
	return r.readAt
}

func (r Reply) UpdatedAt() common.Date {
	return r.updatedAt
}

func (r Reply) IsEnterprise() bool {
	return r.isEnterprise
}

func (r Reply) EnterpriseId() uuid.UUID {
	return r.enterpriseId
}
