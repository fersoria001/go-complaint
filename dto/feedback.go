package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/feedback"
)

type Feedback struct {
	ID             string           `json:"id"`
	ComplaintID    string           `json:"complaintID"`
	EnterpriseID   string           `json:"enterpriseID"`
	ReplyReview    []ReplyReview    `json:"replyReview"`
	FeedbackAnswer []FeedbackAnswer `json:"feedbackAnswers"`
	ReviewedAt     string           `json:"reviewedAt"`
	UpdatedAt      string           `json:"updatedAt"`
	IsDone         bool             `json:"isDone"`
}

func NewFeedbackDTO(domainObject feedback.Feedback) Feedback {
	var replyReviews []ReplyReview
	for _, replyReview := range domainObject.ReplyReviews().ToSlice() {
		replyReviews = append(replyReviews, NewReplyReviewDTO(replyReview))
	}
	var feedbackAnswers []FeedbackAnswer
	for _, feedbackAnswer := range domainObject.FeedbackAnswers().ToSlice() {
		feedbackAnswers = append(feedbackAnswers, NewFeedbackAnswerDTO(feedbackAnswer))
	}
	return Feedback{
		ID:             domainObject.ID().String(),
		ComplaintID:    domainObject.ComplaintID().String(),
		EnterpriseID:   domainObject.EnterpriseID(),
		ReplyReview:    replyReviews,
		FeedbackAnswer: feedbackAnswers,
		ReviewedAt:     common.StringDate(domainObject.ReviewedAt()),
		UpdatedAt:      common.StringDate(domainObject.UpdatedAt()),
		IsDone:         domainObject.IsDone(),
	}
}

type FeedbackAnswer struct {
	ID           string `json:"id"`
	FeedbackID   string `json:"feedbackID"`
	SenderID     string `json:"senderID"`
	SenderIMG    string `json:"senderIMG"`
	SenderName   string `json:"senderName"`
	Body         string `json:"body"`
	CreatedAt    string `json:"createdAt"`
	Read         bool   `json:"read"`
	ReadAt       string `json:"readAt"`
	UpdatedAt    string `json:"updatedAt"`
	IsEnterprise bool   `json:"isEnterprise"`
	EnterpriseID string `json:"enterpriseID"`
}

func NewFeedbackAnswerDTO(domainObject feedback.Answer) FeedbackAnswer {
	return FeedbackAnswer{
		ID:           domainObject.ID().String(),
		FeedbackID:   domainObject.FeedbackID().String(),
		SenderID:     domainObject.SenderID(),
		SenderIMG:    domainObject.SenderIMG(),
		SenderName:   domainObject.SenderName(),
		Body:         domainObject.Body(),
		CreatedAt:    domainObject.CreatedAt().StringRepresentation(),
		Read:         domainObject.Read(),
		ReadAt:       domainObject.ReadAt().StringRepresentation(),
		UpdatedAt:    domainObject.UpdatedAt().StringRepresentation(),
		IsEnterprise: domainObject.IsEnterprise(),
		EnterpriseID: domainObject.EnterpriseID(),
	}
}

type ReplyReview struct {
	ID         string     `json:"id"`
	FeedbackID string     `json:"feedbackID"`
	Reviewer   User       `json:"reviewer"`
	Replies    []ReplyDTO `json:"replies"`
	Review     ReviewDTO  `json:"review"`
	Color      string     `json:"color"`
	CreatedAt  string     `json:"createdAt"`
}

func NewReplyReviewDTO(domainObject feedback.ReplyReview) ReplyReview {
	var replies []ReplyDTO
	for _, reply := range domainObject.Replies().ToSlice() {
		replies = append(replies, NewReplyDTO(reply, "CLOSED"))
	}
	return ReplyReview{
		ID:         domainObject.ID().String(),
		FeedbackID: domainObject.FeedbackID().String(),
		Reviewer:   NewUser(domainObject.Reviewer()),
		Replies:    replies,
		Review:     NewReviewDto(domainObject.Review()),
		Color:      domainObject.Color(),
		CreatedAt:  common.StringDate(domainObject.CreatedAt()),
	}
}

type ReviewDTO struct {
	ReplyReviewID string `json:"replyReviewID"`
	Comment       string `json:"comment"`
}

func NewReviewDto(domainObject feedback.Review) ReviewDTO {
	return ReviewDTO{
		ReplyReviewID: domainObject.ReplyReviewID().String(),
		Comment:       domainObject.Comment(),
	}
}
