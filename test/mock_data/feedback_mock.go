package mock_data

import (
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type ReviewMock struct {
	FeedbackId    uuid.UUID
	ReplyReviewId uuid.UUID
	Comment       string
}

type ReplyReviewMock struct {
	Id         uuid.UUID
	FeedbackId uuid.UUID
	Replies    []*ReplyMock
	Review     *ReviewMock
	Reviewer   *UserMock
	Color      string
	CreatedAt  time.Time
}

type AnswerMock struct {
	Id           uuid.UUID
	FeedbackId   uuid.UUID
	SenderId     uuid.UUID
	SenderImg    string
	SenderName   string
	Body         string
	CreatedAt    common.Date
	Read         bool
	ReadAt       common.Date
	UpdatedAt    common.Date
	IsEnterprise bool
	EnterpriseId uuid.UUID
}

type FeedbackMock struct {
	Id              uuid.UUID
	ComplaintId     uuid.UUID
	EnterpriseId    uuid.UUID
	ReplyReview     []*ReplyReviewMock
	FeedbackAnswers []*AnswerMock
	ReviewedAt      time.Time
	UpdatedAt       time.Time
	IsDone          bool
}

var NewReplyReviews = map[uuid.UUID]*ReplyReviewMock{
	NewFeedbacks["valid"].Id: {
		Id:         uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1a"),
		FeedbackId: NewFeedbacks["valid"].Id,
		Reviewer:   NewUsers["valid"],
		Color:      "#431202",
		Replies:    NewReplies[NewComplaints["valid"].Id],
		Review: &ReviewMock{
			FeedbackId:    uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1e"),
			ReplyReviewId: uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1a"),
			Comment:       "comment",
		},
	},
}

var NewFeedbacks = map[string]*FeedbackMock{
	"valid": {
		Id:           uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1e"),
		ComplaintId:  NewComplaints["valid"].Id,
		EnterpriseId: NewEnterprises["valid"].Id,
		ReplyReview: []*ReplyReviewMock{
			{
				Id:         uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1a"),
				FeedbackId: uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1e"),
				Reviewer:   NewUsers["valid"],
				Color:      "#431202",
				Replies:    NewReplies[NewComplaints["valid"].Id][0:1],
				Review: &ReviewMock{
					FeedbackId:    uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1e"),
					ReplyReviewId: uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1a"),
					Comment:       "comment",
				},
			},
			{
				Id:         uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1a"),
				FeedbackId: uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1e"),
				Reviewer:   NewUsers["valid"],
				Color:      "#451202",
				Replies:    NewReplies[NewComplaints["valid"].Id][1:2],
				Review: &ReviewMock{
					FeedbackId:    uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1e"),
					ReplyReviewId: uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1a"),
					Comment:       "comment",
				},
			},
			{
				Id:         uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1a"),
				FeedbackId: uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1e"),
				Reviewer:   NewUsers["valid"],
				Color:      "#421202",
				Replies:    NewReplies[NewComplaints["valid"].Id][2:3],
				Review: &ReviewMock{
					FeedbackId:    uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1e"),
					ReplyReviewId: uuid.MustParse("3f5bff4a-382b-252c-f432-63b123f54b1a"),
					Comment:       "comment",
				},
			},
		},
	},
}
