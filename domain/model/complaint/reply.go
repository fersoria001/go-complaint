package complaint

import (
	"go-complaint/domain/model/common"
	"go-complaint/erros"
	"time"

	"github.com/google/uuid"
)

// Package complaint
// << Entity >>
// It represents the reply of the complaint
// Its complaintID is the id of the complaint that own this reply
type Reply struct {
	id          uuid.UUID
	complaintID uuid.UUID
	senderID    string
	senderIMG   string
	senderName  string
	body        string
	createdAt   common.Date
	read        bool
	readAt      common.Date
	updatedAt   common.Date
}

func (r *Reply) MarkAsRead() {
	r.read = true
	r.readAt = common.NewDate(time.Now())
}

func NewReply(id uuid.UUID,
	complaintID uuid.UUID,
	senderID string,
	senderIMG string,
	senderName string,
	body string,
	read bool,
	createdAt,
	readAt,
	updatedAt common.Date) (*Reply, error) {
	var reply *Reply = new(Reply)
	err := reply.setID(id)
	if err != nil {
		return nil, err
	}
	err = reply.setSurrogateID(complaintID)
	if err != nil {
		return nil, err
	}
	err = reply.setSenderID(senderID)
	if err != nil {
		return nil, err
	}
	err = reply.setSenderIMG(senderIMG)
	if err != nil {
		return nil, err
	}
	err = reply.setSenderName(senderName)
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

func (r *Reply) setSenderID(senderID string) error {
	if senderID == "" {
		return &erros.NullValueError{}
	}
	r.senderID = senderID
	return nil
}

// To validate it senderName needs to be previously validated
func (r *Reply) setSurrogateID(complaintID uuid.UUID) error {
	if complaintID == uuid.Nil {
		return &erros.NullValueError{}
	}
	r.complaintID = complaintID
	return nil
}

func (r *Reply) setID(id uuid.UUID) error {
	if id == uuid.Nil {
		return &erros.NullValueError{}
	}
	r.id = id
	return nil
}

func (r *Reply) setSenderIMG(senderIMG string) error {
	if senderIMG == "" {
		return &erros.NullValueError{}
	}
	r.senderIMG = senderIMG
	return nil
}

func (r *Reply) setSenderName(senderName string) error {
	if senderName == "" {
		return &erros.NullValueError{}
	}
	r.senderName = senderName
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

func (r *Reply) ID() uuid.UUID {
	return r.id
}

func (r *Reply) ComplaintID() uuid.UUID {
	return r.complaintID
}

func (r *Reply) SenderIMG() string {
	return r.senderIMG
}

func (r *Reply) SenderName() string {
	return r.senderName
}

func (r *Reply) Body() string {
	return r.body
}

func (r *Reply) CreatedAt() common.Date {
	return r.createdAt
}

func (r *Reply) Read() bool {
	return r.read
}

func (r *Reply) ReadAt() common.Date {
	return r.readAt
}

func (r *Reply) UpdatedAt() common.Date {
	return r.updatedAt
}

func (r *Reply) SenderID() string {
	return r.senderID
}
