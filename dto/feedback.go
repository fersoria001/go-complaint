package dto

import (
	"go-complaint/domain/model/feedback"
)

type Feedback struct {
	ComplaintID string
	ReviewerID  string
	ReplyReview []*ReplyReview
	Answer      []*Answer
}

func NewFeedback(
	domainObject *feedback.Feedback,
	answer []*feedback.Answer,
) *Feedback {
	var feedback *Feedback = new(Feedback)
	feedback.ComplaintID = domainObject.ComplaintID().String()
	feedback.ReviewerID = domainObject.ReviewerID()
	feedback.ReplyReview = make([]*ReplyReview, 0)
	for item := range domainObject.ReplyReview().Iter() {
		var replyReview *ReplyReview = new(ReplyReview)
		replyReview.ID = item.ID().String()
		replyReview.FeedbackID = item.FeedbackID().String()
		replyReview.SenderID = item.Reply().SenderID()
		replyReview.SenderIMG = item.Reply().SenderIMG()
		replyReview.SenderName = item.Reply().SenderName()
		replyReview.Body = item.Reply().Body()
		replyReview.CreatedAt = item.Reply().CreatedAt().StringRepresentation()
		replyReview.Read = item.Reply().Read()
		replyReview.ReadAt = item.Reply().ReadAt().StringRepresentation()
		replyReview.UpdatedAt = item.Reply().UpdatedAt().StringRepresentation()
		replyReview.ReviewerID = item.Review().ReviewerID()
		replyReview.ReviewerIMG = item.Review().ReviewerIMG()
		replyReview.ReviewerName = item.Review().ReviewerName()
		replyReview.ReviewedAt = item.Review().ReviewedAt().StringRepresentation()
		replyReview.Comment = item.Review().Comment()
		feedback.ReplyReview = append(feedback.ReplyReview, replyReview)
	}

	return feedback
}

type ReplyReview struct {
	ID           string
	FeedbackID   string
	SenderID     string
	SenderIMG    string
	SenderName   string
	Body         string
	CreatedAt    string
	Read         bool
	ReadAt       string
	UpdatedAt    string
	ReviewerID   string
	ReviewerIMG  string
	ReviewerName string
	ReviewedAt   string
	Comment      string
}

type Answer struct {
	ID         string
	FeedbackID string
	SenderID   string
	SenderIMG  string
	SenderName string
	Body       string
	CreatedAt  string
	Read       bool
	ReadAt     string
	UpdatedAt  string
}
