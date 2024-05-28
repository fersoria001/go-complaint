package dto

import (
	"go-complaint/domain/model/complaint"

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
	Complaints    []*ComplaintDTO `json:"complaints"`
	Count         int             `json:"count"`
	CurrentLimit  int             `json:"current_limit"`
	CurrentOffset int             `json:"current_offset"`
}
type ComplaintDTO struct {
	ID                 uuid.UUID   `json:"id"`
	AuthorID           string      `json:"author_id"`
	AuthorFullName     string      `json:"author_full_name"`
	AuthorProfileIMG   string      `json:"author_profile_img"`
	ReceiverID         string      `json:"receiver_id"`
	ReceiverFullName   string      `json:"receiver_full_name"`
	ReceiverProfileIMG string      `json:"receiver_profile_img"`
	Status             string      `json:"status"`
	Message            MessageDTO  `json:"message"`
	Rating             RatingDTO   `json:"rating"`
	CreatedAt          string      `json:"created_at"`
	UpdatedAt          string      `json:"updated_at"`
	Replies            []*ReplyDTO `json:"replies"`
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
	ID          string `json:"id"`
	ComplaintID string `json:"complaint_id"`
	SenderID    string `json:"sender_id"`
	SenderIMG   string `json:"sender_img"`
	SenderName  string `json:"sender_name"`
	Body        string `json:"body"`
	CreatedAt   string `json:"created_at"`
	Read        bool   `json:"read"`
	ReadAt      string `json:"read_at"`
	UpdatedAt   string `json:"updated_at"`
}

func NewComplaintDTO(
	domainComplaint *complaint.Complaint,
	replies []*complaint.Reply,
) *ComplaintDTO {
	replyDTOSlice := []*ReplyDTO{}
	for _, reply := range replies {
		replyDTOSlice = append(replyDTOSlice, NewReplyDTO(reply))
	}
	return &ComplaintDTO{
		ID:         domainComplaint.ID(),
		AuthorID:   domainComplaint.AuthorID(),
		ReceiverID: domainComplaint.ReceiverID(),
		Status:     domainComplaint.Status().String(),
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
	domainReply *complaint.Reply,
) *ReplyDTO {
	return &ReplyDTO{
		ID:          domainReply.ID().String(),
		ComplaintID: domainReply.ComplaintID().String(),
		SenderID:    domainReply.SenderID(),
		SenderIMG:   domainReply.SenderIMG(),
		SenderName:  domainReply.SenderName(),
		Body:        domainReply.Body(),
		CreatedAt:   domainReply.CreatedAt().StringRepresentation(),
		Read:        domainReply.Read(),
		ReadAt:      domainReply.ReadAt().StringRepresentation(),
		UpdatedAt:   domainReply.UpdatedAt().StringRepresentation(),
	}
}
