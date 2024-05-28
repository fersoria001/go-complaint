package feedback

import (
	"go-complaint/domain/model/common"
	"go-complaint/erros"
)

type Reply struct {
	senderID   string
	senderIMG  string
	senderName string
	body       string
	createdAt  common.Date
	read       bool
	readAt     common.Date
	updatedAt  common.Date
}

func NewReply(
	senderID string,
	senderIMG string,
	senderName string,
	body string,
	createdAt common.Date,
	read bool,
	readAt common.Date,
	updatedAt common.Date,
) (*Reply, error) {
	var reply *Reply = new(Reply)
	err := reply.setSenderID(senderID)
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
	err = reply.setRead(read)
	if err != nil {
		return nil, err
	}
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
	r.body = body
	return nil
}

func (r *Reply) setCreatedAt(createdAt common.Date) error {

	if createdAt == (common.Date{}) {
		return &erros.EmptyStructError{}
	}
	r.createdAt = createdAt
	return nil
}

func (r *Reply) setRead(read bool) error {
	r.read = read
	return nil
}

func (r *Reply) setReadAt(readAt common.Date) error {
	if readAt == (common.Date{}) {
		return &erros.EmptyStructError{}
	}
	r.readAt = readAt
	return nil
}

func (r *Reply) setUpdatedAt(updatedAt common.Date) error {
	if updatedAt == (common.Date{}) {
		return &erros.EmptyStructError{}
	}
	r.updatedAt = updatedAt
	return nil
}

func (r *Reply) SenderID() string {
	return r.senderID
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
