package models

import (
	"go-complaint/domain/model/complaint"
	"go-complaint/erros"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type Reply struct {
	ID          uuid.UUID
	ComplaintID uuid.UUID
	SenderID    string
	SenderIMG   string
	SenderName  string
	Body        string
	Read        bool
	CreatedAt   string
	ReadAt      string
	UpdatedAt   string
}

// this return values, test
func NewReply(domain *complaint.Reply) Reply {
	return Reply{
		ID:          domain.ID(),
		ComplaintID: domain.ComplaintID(),
		SenderID:    domain.SenderID(),
		SenderIMG:   domain.SenderIMG(),
		SenderName:  domain.SenderName(),
		Body:        domain.Body(),
		CreatedAt:   domain.CreatedAt().StringRepresentation(),
		Read:        domain.Read(),
		ReadAt:      domain.ReadAt().StringRepresentation(),
		UpdatedAt:   domain.UpdatedAt().StringRepresentation(),
	}
}

func NewReplies(domains mapset.Set[*complaint.Reply]) (mapset.Set[Reply], error) {
	replies := mapset.NewSet[Reply]()
	for domain := range domains.Iter() {
		replies.Add(NewReply(domain))
	}
	if replies.Cardinality() != domains.Cardinality() {
		return nil, &erros.NullValueError{}
	}
	//domains.Clear() could be if the set was ptr, but is not and maybe its useless for it to be a *complaint.Reply set otherwise
	//in go you cant ptr to interface
	//if i clear at this point the super struct Complaint is still in memory, could end on nilpointer dereference
	return replies, nil
}

func (r *Reply) Columns() Columns {
	return Columns{
		"id",
		"complaint_id",
		"sender_id",
		"sender_img",
		"sender_name",
		"body",
		"read_status",
		"read_at",
		"created_at",
		"updated_at",
	}
}

// this has ptr receiver
func (r *Reply) Values() Values {
	return Values{
		&r.ID,
		&r.ComplaintID,
		&r.SenderID,
		&r.SenderIMG,
		&r.SenderName,
		&r.Body,
		&r.Read,
		&r.ReadAt,
		&r.CreatedAt,
		&r.UpdatedAt,
	}
}

func (r *Reply) Args() string {
	return "$1, $2, $3, $4, $5, $6, $7, $8, $9, $10"
}

func (r *Reply) Table() string {
	return "replies"
}
