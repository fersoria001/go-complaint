package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"slices"
)

/*
There must exists a struct
that represents a List of this ComplaintDTO
-or of every list of a DTO-
that wraps the slice of ComplaintDTO
along its size and offset or
left items to paginate
*/
type ComplaintPagination struct {
	Complaints []*Complaint `json:"complaints"`
	NextCursor int          `json:"nextCursor"`
	PrevCursor int          `json:"prevCursor"`
}

type Complaint struct {
	Id          string     `json:"id"`
	Author      *Recipient `json:"author"`
	Receiver    *Recipient `json:"receiver"`
	Status      string     `json:"status"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Rating      *Rating    `json:"rating"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
	Replies     []*Reply   `json:"replies"`
}

type Rating struct {
	Id      string `json:"id"`
	Rate    int    `json:"rate"`
	Comment string `json:"comment"`
}

func NewRating(obj complaint.Rating) *Rating {
	return &Rating{
		Id:      obj.Id().String(),
		Rate:    obj.Rate(),
		Comment: obj.Comment(),
	}
}

type Reply struct {
	Id          string     `json:"id"`
	ComplaintId string     `json:"complaintId"`
	Sender      *Recipient `json:"sender"`
	Body        string     `json:"body"`
	CreatedAt   string     `json:"createdAt"`
	Read        bool       `json:"read"`
	ReadAt      string     `json:"readAt"`
	UpdatedAt   string     `json:"updatedAt"`
}

func NewComplaint(
	domainComplaint complaint.Complaint,
) *Complaint {
	replyDTOSlice := []*Reply{}
	for reply := range domainComplaint.Replies().Iter() {
		replyDTOSlice = append(replyDTOSlice, NewReply(reply))
	}
	slices.SortStableFunc(replyDTOSlice, func(i, j *Reply) int {
		iCreatedAt, _ := common.ParseDate(i.CreatedAt)
		jCreatedAt, _ := common.ParseDate(j.CreatedAt)
		if iCreatedAt.Before(jCreatedAt) {
			return -1
		}
		if iCreatedAt.After(jCreatedAt) {
			return 1
		}
		return 0
	})
	return &Complaint{
		Id:          domainComplaint.Id().String(),
		Author:      NewRecipient(domainComplaint.Author()),
		Receiver:    NewRecipient(domainComplaint.Receiver()),
		Status:      domainComplaint.Status().String(),
		Title:       domainComplaint.Title(),
		Description: domainComplaint.Description(),
		Rating:      NewRating(domainComplaint.Rating()),
		CreatedAt:   domainComplaint.CreatedAt().StringRepresentation(),
		UpdatedAt:   domainComplaint.UpdatedAt().StringRepresentation(),
		Replies:     replyDTOSlice,
	}
}

func NewReply(
	domainReply complaint.Reply,
) *Reply {
	return &Reply{
		Id:          domainReply.ID().String(),
		ComplaintId: domainReply.ComplaintId().String(),
		Sender:      NewRecipient(domainReply.Sender()),
		Body:        domainReply.Body(),
		CreatedAt:   domainReply.CreatedAt().StringRepresentation(),
		Read:        domainReply.Read(),
		ReadAt:      domainReply.ReadAt().StringRepresentation(),
		UpdatedAt:   domainReply.UpdatedAt().StringRepresentation(),
	}
}
