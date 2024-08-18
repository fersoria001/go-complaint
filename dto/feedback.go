package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/feedback"
)

type Feedback struct {
	Id             string           `json:"id"`
	ComplaintId    string           `json:"complaintId"`
	EnterpriseId   string           `json:"enterpriseId"`
	ReplyReview    []ReplyReview    `json:"replyReview"`
	FeedbackAnswer []FeedbackAnswer `json:"feedbackAnswers"`
	ReviewedAt     string           `json:"reviewedAt"`
	UpdatedAt      string           `json:"updatedAt"`
	IsDone         bool             `json:"isDone"`
}

func NewFeedbackDTO(domainObject feedback.Feedback) *Feedback {
	var replyReviews []ReplyReview
	for _, replyReview := range domainObject.ReplyReviews().ToSlice() {
		replyReviews = append(replyReviews, NewReplyReviewDTO(replyReview))
	}
	var feedbackAnswers []FeedbackAnswer
	for _, feedbackAnswer := range domainObject.FeedbackAnswers().ToSlice() {
		feedbackAnswers = append(feedbackAnswers, NewFeedbackAnswerDTO(feedbackAnswer))
	}
	return &Feedback{
		Id:             domainObject.Id().String(),
		ComplaintId:    domainObject.ComplaintId().String(),
		EnterpriseId:   domainObject.EnterpriseId().String(),
		ReplyReview:    replyReviews,
		FeedbackAnswer: feedbackAnswers,
		ReviewedAt:     common.StringDate(domainObject.ReviewedAt()),
		UpdatedAt:      common.StringDate(domainObject.UpdatedAt()),
		IsDone:         domainObject.IsDone(),
	}
}

type FeedbackAnswer struct {
	Id           string `json:"id"`
	FeedbackId   string `json:"feedbackId"`
	SenderId     string `json:"senderId"`
	SenderImg    string `json:"senderImg"`
	SenderName   string `json:"senderName"`
	Body         string `json:"body"`
	CreatedAt    string `json:"createdAt"`
	Read         bool   `json:"read"`
	ReadAt       string `json:"readAt"`
	UpdatedAt    string `json:"updatedAt"`
	IsEnterprise bool   `json:"isEnterprise"`
	EnterpriseId string `json:"enterpriseId"`
}

func NewFeedbackAnswerDTO(domainObject feedback.Answer) FeedbackAnswer {
	return FeedbackAnswer{
		Id:           domainObject.Id().String(),
		FeedbackId:   domainObject.FeedbackId().String(),
		SenderId:     domainObject.SenderId().String(),
		SenderImg:    domainObject.SenderImg(),
		SenderName:   domainObject.SenderName(),
		Body:         domainObject.Body(),
		CreatedAt:    domainObject.CreatedAt().StringRepresentation(),
		Read:         domainObject.Read(),
		ReadAt:       domainObject.ReadAt().StringRepresentation(),
		UpdatedAt:    domainObject.UpdatedAt().StringRepresentation(),
		IsEnterprise: domainObject.IsEnterprise(),
		EnterpriseId: domainObject.EnterpriseId(),
	}
}

type ReplyReview struct {
	ID         string    `json:"id"`
	FeedbackID string    `json:"feedbackID"`
	Reviewer   *User     `json:"reviewer"`
	Replies    []*Reply  `json:"replies"`
	Review     ReviewDTO `json:"review"`
	Color      string    `json:"color"`
	CreatedAt  string    `json:"createdAt"`
}

func NewReplyReviewDTO(domainObject feedback.ReplyReview) ReplyReview {
	var replies []*Reply
	for _, reply := range domainObject.Replies().ToSlice() {
		replies = append(replies, NewReply(reply))
	}
	reviewer := domainObject.Reviewer()
	return ReplyReview{
		ID:         domainObject.ID().String(),
		FeedbackID: domainObject.FeedbackId().String(),
		Reviewer:   NewUser(&reviewer),
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
