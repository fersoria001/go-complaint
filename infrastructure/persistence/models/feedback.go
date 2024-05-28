package models

import (
	"go-complaint/domain/model/feedback"

	"github.com/google/uuid"
)

type Feedback struct {
	ID              uuid.UUID
	FeedbackID      uuid.UUID
	ReviewedID      string
	ReplySenderID   string
	ReplySenderIMG  string
	ReplySenderName string
	ReplyBody       string
	ReplyCreatedAt  string
	ReplyRead       bool
	ReplyReadAt     string
	ReplyUpdatedAt  string
	ReviewerID      string
	ReviewerIMG     string
	ReviewerName    string
	ReviewedAt      string
	ReviewComment   string
}

func NewFeedback(
	rootID uuid.UUID,
	reviewerID string,
	reviewedID string,
	replyReview *feedback.ReplyReview,
) *Feedback {
	return &Feedback{
		ID:              replyReview.ID(),
		FeedbackID:      rootID,
		ReviewedID:      reviewedID,
		ReplySenderID:   replyReview.Reply().SenderID(),
		ReplySenderIMG:  replyReview.Reply().SenderIMG(),
		ReplySenderName: replyReview.Reply().SenderName(),
		ReplyBody:       replyReview.Reply().Body(),
		ReplyCreatedAt:  replyReview.Reply().CreatedAt().StringRepresentation(),
		ReplyRead:       replyReview.Reply().Read(),
		ReplyReadAt:     replyReview.Reply().ReadAt().StringRepresentation(),
		ReplyUpdatedAt:  replyReview.Reply().UpdatedAt().StringRepresentation(),
		ReviewerID:      replyReview.Review().ReviewerID(),
		ReviewerIMG:     replyReview.Review().ReviewerIMG(),
		ReviewerName:    replyReview.Review().ReviewerName(),
		ReviewedAt:      replyReview.Review().ReviewedAt().StringRepresentation(),
		ReviewComment:   replyReview.Review().Comment(),
	}
}

func (rr *Feedback) Columns() Columns {
	return Columns{
		"id",
		"feedback_id",
		"reviewed_id",
		"reply_sender_id",
		"reply_sender_img",
		"reply_sender_name",
		"reply_body",
		"reply_created_at",
		"reply_read",
		"reply_read_at",
		"reply_updated_at",
		"reviewer_id",
		"reviewer_img",
		"reviewer_name",
		"reviewed_at",
		"review_comment",
	}
}

func (rr *Feedback) Values() Values {
	return Values{
		&rr.ID,
		&rr.FeedbackID,
		&rr.ReviewedID,
		&rr.ReplySenderID,
		&rr.ReplySenderIMG,
		&rr.ReplySenderName,
		&rr.ReplyBody,
		&rr.ReplyCreatedAt,
		&rr.ReplyRead,
		&rr.ReplyReadAt,
		&rr.ReplyUpdatedAt,
		&rr.ReviewerID,
		&rr.ReviewerIMG,
		&rr.ReviewerName,
		&rr.ReviewedAt,
		&rr.ReviewComment,
	}
}

func (rr *Feedback) Table() string {
	return "feedbacks"
}

func (rr *Feedback) Args() string {
	return "$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16"
}
