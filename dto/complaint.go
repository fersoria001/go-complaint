package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"slices"

	"github.com/google/uuid"
)

/*
There must exists a struct
that represents a List of this ComplaintDTO
-or of every list of a DTO-
that wraps the slice of ComplaintDTO
along its size and offset or
left items to paginate
*/
type ComplaintListDTO struct {
	Complaints    []ComplaintDTO `json:"complaints"`
	Count         int            `json:"count"`
	CurrentLimit  int            `json:"current_limit"`
	CurrentOffset int            `json:"current_offset"`
}
type ComplaintDTO struct {
	ID                 uuid.UUID  `json:"id"`
	AuthorID           string     `json:"author_id"`
	AuthorFullName     string     `json:"author_full_name"`
	AuthorProfileIMG   string     `json:"author_profile_img"`
	AuthorPronoun      string     `json:"author_pronoun"`
	ReceiverID         string     `json:"receiver_id"`
	ReceiverFullName   string     `json:"receiver_full_name"`
	ReceiverProfileIMG string     `json:"receiver_profile_img"`
	ReceiverPronoun    string     `json:"receiver_pronoun"`
	Status             string     `json:"status"`
	Message            MessageDTO `json:"message"`
	Rating             RatingDTO  `json:"rating"`
	CreatedAt          string     `json:"created_at"`
	UpdatedAt          string     `json:"updated_at"`
	Replies            []ReplyDTO `json:"replies"`
}

type MessageDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}
type RatingDTO struct {
	Rate    int    `json:"rate"`
	Comment string `json:"comment"`
}
type ReplyDTO struct {
	ID              string `json:"id"`
	ComplaintID     string `json:"complaintID"`
	SenderID        string `json:"senderID"`
	SenderIMG       string `json:"senderIMG"`
	SenderName      string `json:"senderName"`
	Body            string `json:"body"`
	CreatedAt       string `json:"createdAt"`
	Read            bool   `json:"read"`
	ReadAt          string `json:"readAt"`
	UpdatedAt       string `json:"updatedAt"`
	IsEnterprise    bool   `json:"isEnterprise"`
	EnterpriseID    string `json:"enterpriseID"`
	ComplaintStatus string `json:"complaintStatus"`
}

func NewComplaintDTO(
	domainComplaint complaint.Complaint,
) ComplaintDTO {
	replyDTOSlice := []ReplyDTO{}
	for reply := range domainComplaint.Replies().Iter() {
		replyDTOSlice = append(replyDTOSlice, NewReplyDTO(reply, domainComplaint.Status().String()))
	}
	slices.SortStableFunc(replyDTOSlice, func(i, j ReplyDTO) int {
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
	return ComplaintDTO{
		ID:                 domainComplaint.ID(),
		AuthorID:           domainComplaint.AuthorID(),
		AuthorFullName:     domainComplaint.AuthorFullName(),
		AuthorProfileIMG:   domainComplaint.AuthorProfileIMG(),
		ReceiverID:         domainComplaint.ReceiverID(),
		ReceiverFullName:   domainComplaint.ReceiverFullName(),
		ReceiverProfileIMG: domainComplaint.ReceiverProfileIMG(),
		Status:             domainComplaint.Status().String(),
		Message: MessageDTO{
			Title:       domainComplaint.Message().Title(),
			Description: domainComplaint.Message().Description(),
			Body:        domainComplaint.Message().Body(),
		},
		Rating: RatingDTO{
			Rate:    domainComplaint.Rating().Rate(),
			Comment: domainComplaint.Rating().Comment(),
		},
		CreatedAt: domainComplaint.CreatedAt().StringRepresentation(),
		UpdatedAt: domainComplaint.UpdatedAt().StringRepresentation(),
		Replies:   replyDTOSlice,
	}
}

func NewReplyDTO(
	domainReply complaint.Reply,
	complaintStatus string,
) ReplyDTO {
	return ReplyDTO{
		ID:              domainReply.ID().String(),
		ComplaintID:     domainReply.ComplaintID().String(),
		SenderID:        domainReply.SenderID(),
		SenderIMG:       domainReply.SenderIMG(),
		SenderName:      domainReply.SenderName(),
		Body:            domainReply.Body(),
		CreatedAt:       domainReply.CreatedAt().StringRepresentation(),
		Read:            domainReply.Read(),
		ReadAt:          domainReply.ReadAt().StringRepresentation(),
		UpdatedAt:       domainReply.UpdatedAt().StringRepresentation(),
		IsEnterprise:    domainReply.IsEnterprise(),
		EnterpriseID:    domainReply.EnterpriseID(),
		ComplaintStatus: complaintStatus,
	}
}

type NewUnreadReply struct {
	ComplaintID string `json:"complaint_id"`
	ReplyID     string `json:"reply_id"`
	SenderID    string `json:"sender_id"`
	ReceiverID  string `json:"receiver_id"`
}
