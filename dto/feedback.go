package dto

import "go-complaint/domain/model/feedback"

type Feedback struct {
	ID             string           `json:"id"`
	ComplaintID    string           `json:"complaint_id"`
	ReviewedID     string           `json:"reviewed_id"`
	ReplyReview    []ReplyReview    `json:"reply_review"`
	FeedbackAnswer []FeedbackAnswer `json:"feedback_answer"`
}

func NewFeedbackDTO(domainObject feedback.Feedback) Feedback {
	var replyReviews []ReplyReview
	for _, replyReview := range domainObject.ReplyReview().ToSlice() {
		replyReviews = append(replyReviews, NewReplyReviewDTO(replyReview))
	}
	var feedbackAnswers []FeedbackAnswer
	for _, feedbackAnswer := range domainObject.FeedbackAnswers().ToSlice() {
		feedbackAnswers = append(feedbackAnswers, NewFeedbackAnswerDTO(feedbackAnswer))
	}
	return Feedback{
		ID:             domainObject.ID().String(),
		ComplaintID:    domainObject.ComplaintID().String(),
		ReviewedID:     domainObject.ReviewedID(),
		ReplyReview:    replyReviews,
		FeedbackAnswer: feedbackAnswers,
	}
}

type FeedbackAnswer struct {
	ID           string `json:"id"`
	FeedbackID   string `json:"feedback_id"`
	SenderID     string `json:"sender_id"`
	SenderIMG    string `json:"sender_img"`
	SenderName   string `json:"sender_name"`
	Body         string `json:"body"`
	CreatedAt    string `json:"created_at"`
	Read         bool   `json:"read"`
	ReadAt       string `json:"read_at"`
	UpdatedAt    string `json:"updated_at"`
	IsEnterprise bool   `json:"is_enterprise"`
	EnterpriseID string `json:"enterprise_id"`
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
	FeedbackID string     `json:"feedback_id"`
	Replies    []ReplyDTO `json:"replies"`
	Review     ReviewDTO  `json:"review"`
	Color      string     `json:"color"`
}

func NewReplyReviewDTO(domainObject feedback.ReplyReview) ReplyReview {
	var replies []ReplyDTO
	for _, reply := range domainObject.Replies().ToSlice() {
		replies = append(replies, NewReplyDTO(reply, "CLOSED"))
	}
	return ReplyReview{
		ID:         domainObject.ID().String(),
		FeedbackID: domainObject.FeedbackID().String(),
		Replies:    replies,
		Review:     NewReviewDto(domainObject.Review()),
		Color:      domainObject.Color(),
	}
}

type ReviewDTO struct {
	ReplyReviewID string `json:"reply_review_id"`
	ReviewerID    string `json:"reviewer_id"`
	ReviewedAt    string `json:"reviewed_at"`
	Comment       string `json:"comment"`
}

func NewReviewDto(domainObject feedback.Review) ReviewDTO {
	return ReviewDTO{
		ReplyReviewID: domainObject.ReplyReviewID().String(),
		ReviewerID:    domainObject.ReviewerID(),
		ReviewedAt:    domainObject.ReviewedAt().StringRepresentation(),
		Comment:       domainObject.Comment(),
	}
}
