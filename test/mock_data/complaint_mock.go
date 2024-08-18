package mock_data

import (
	"fmt"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"time"

	"github.com/google/uuid"
)

type RecipientMock struct {
	Id               uuid.UUID
	SubjectName      string
	SubjectThumbnail string
	SubjectEmail     string
	IsEnterprise     bool
}

type RatingMock struct {
	Id      uuid.UUID
	Rate    int
	Comment string
}

type ReplyMock struct {
	Id          uuid.UUID
	ComplaintId uuid.UUID
	Sender      *RecipientMock
	Body        string
	CreatedAt   common.Date
	Read        bool
	ReadAt      common.Date
	UpdatedAt   common.Date
}

type ComplaintMock struct {
	Id          uuid.UUID
	Author      *RecipientMock
	Receiver    *RecipientMock
	Title       string
	Description string
	Status      complaint.Status
	Rating      *RatingMock
	CreatedAt   common.Date
	UpdatedAt   common.Date
	Replies     []*ReplyMock
}

type ComplaintDataMock struct {
	Id          uuid.UUID
	OwnerId     uuid.UUID
	AuthorId    uuid.UUID
	ReceiverId  uuid.UUID
	ComplaintId uuid.UUID
	OccurredOn  time.Time
	DataType    complaint.ComplaintDataType
}

var NewComplaintData = []ComplaintDataMock{
	{
		Id:          uuid.MustParse("8c910eaf-b942-430f-b01d-21a38e271d0f"),
		OwnerId:     NewUsers["valid"].Id,
		AuthorId:    NewComplaints["valid"].Author.Id,
		ReceiverId:  NewComplaints["valid"].Receiver.Id,
		ComplaintId: NewComplaints["valid"].Id,
		OccurredOn:  time.Now(),
		DataType:    complaint.RECEIVED,
	},
	{
		Id:          uuid.MustParse("8c910eaf-b942-430f-b01d-21a38e271d1f"),
		OwnerId:     NewUsers["valid"].Id,
		AuthorId:    NewComplaints["valid"].Author.Id,
		ReceiverId:  NewComplaints["valid"].Receiver.Id,
		ComplaintId: NewComplaints["valid"].Id,
		OccurredOn:  time.Now(),
		DataType:    complaint.RESOLVED,
	},
	{
		Id:          uuid.MustParse("8c910eaf-b942-430f-b01d-21a38e271d2f"),
		OwnerId:     NewUsers["valid"].Id,
		AuthorId:    NewComplaints["valid"].Author.Id,
		ReceiverId:  NewComplaints["valid"].Receiver.Id,
		ComplaintId: NewComplaints["valid"].Id,
		OccurredOn:  time.Now(),
		DataType:    complaint.REVIEWED,
	},
}

var NewRecipients = map[string]*RecipientMock{
	"enterprise": {
		Id:               NewEnterprises["valid"].Id,
		SubjectName:      NewEnterprises["valid"].Name,
		SubjectThumbnail: NewEnterprises["valid"].LogoImg,
		SubjectEmail:     NewEnterprises["valid"].Email,
		IsEnterprise:     true,
	},
	"user": {
		Id:               NewUsers["valid"].Id,
		SubjectName:      fmt.Sprintf("%s %s", NewUsers["valid"].Person.FirstName, NewUsers["valid"].Person.LastName),
		SubjectThumbnail: NewUsers["valid"].Person.ProfileImg,
		SubjectEmail:     NewUsers["valid"].Person.Email,
		IsEnterprise:     false,
	},
	"enterprise1": {
		Id:               NewEnterprises["valid1"].Id,
		SubjectName:      NewEnterprises["valid1"].Name,
		SubjectThumbnail: NewEnterprises["valid1"].LogoImg,
		SubjectEmail:     NewEnterprises["valid1"].Email,
		IsEnterprise:     true,
	},
	"user1": {
		Id:               NewUsers["valid1"].Id,
		SubjectName:      NewUsers["valid1"].Person.FirstName + " " + NewUsers["valid1"].Person.LastName,
		SubjectThumbnail: NewUsers["valid1"].Person.ProfileImg,
		SubjectEmail:     NewUsers["valid1"].Person.Email,
		IsEnterprise:     false,
	},
}

var NewReplies = map[uuid.UUID][]*ReplyMock{
	uuid.MustParse("f41298ea-7c1b-441a-8a0b-32d01ca27e2b"): {
		{
			Id:          uuid.MustParse("f41298ea-7c1b-441a-8a0b-32d01ca27e3b"),
			ComplaintId: uuid.MustParse("f41298ea-7c1b-441a-8a0b-32d01ca27e2b"),
			Sender:      NewRecipients["user"],
			Body:        "This is the second reply",
			CreatedAt:   CommonDate,
			Read:        false,
			ReadAt:      CommonDate,
			UpdatedAt:   CommonDate,
		},
		{
			Id:          uuid.MustParse("f41298ea-7c1b-441a-8a0b-32d01ca27e4b"),
			ComplaintId: uuid.MustParse("f41298ea-7c1b-441a-8a0b-32d01ca27e2b"),
			Sender:      NewRecipients["user"],
			Body:        "This is the third reply",
			CreatedAt:   CommonDate,
			Read:        false,
			ReadAt:      CommonDate,
			UpdatedAt:   CommonDate,
		},
		{
			Id:          uuid.MustParse("f41298ea-7c1b-441a-8a0b-32d01ca27a2c"),
			ComplaintId: uuid.MustParse("f41298ea-7c1b-441a-8a0b-32d01ca27e2b"),
			Sender:      NewRecipients["user"],
			Body:        "This is the fourth reply",
			CreatedAt:   CommonDate,
			Read:        false,
			ReadAt:      CommonDate,
			UpdatedAt:   CommonDate,
		},
	},
}

var NewComplaints = map[string]*ComplaintMock{
	"valid": {
		Id:          uuid.MustParse("f41298ea-7c1b-441a-8a0b-32d01ca27e2b"),
		Author:      NewRecipients["user"],
		Receiver:    NewRecipients["enterprise"],
		Title:       "This is a complaint.",
		Description: "This is a complaint description.",
		Status:      complaint.OPEN,
		Rating: &RatingMock{
			Id:      uuid.MustParse("f41298ea-7c1b-441a-8a0b-32d01ca27e2b"),
			Rate:    5,
			Comment: "This is a rating comment.",
		},
		CreatedAt: CommonDate,
		UpdatedAt: CommonDate,
		Replies: []*ReplyMock{
			{
				Id:          uuid.MustParse("f41298ea-7c1b-441a-8a0b-32d01ca27e2b"),
				ComplaintId: uuid.MustParse("f41298ea-7c1b-441a-8a0b-32d01ca27e2b"),
				Sender:      NewRecipients["user"],
				Body:        "This is the body of the complaint and behaves as the first reply of it",
				CreatedAt:   CommonDate,
				Read:        false,
				ReadAt:      CommonDate,
				UpdatedAt:   CommonDate,
			},
		},
	},
	"valid1": {
		Id:          uuid.MustParse("f51298ea-7b1a-532a-8a0b-35d02ab28c3e"),
		Author:      NewRecipients["user1"],
		Receiver:    NewRecipients["user"],
		Title:       "This is a complaint.",
		Description: "This is a complaint description.",
		Status:      complaint.OPEN,
		Rating: &RatingMock{
			Id:      uuid.MustParse("f51298ea-7b1a-532a-8a0b-35d02ab28c3e"),
			Rate:    5,
			Comment: "This is a rating comment.",
		},
		CreatedAt: CommonDate,
		UpdatedAt: CommonDate,
		Replies: []*ReplyMock{
			{
				Id:          uuid.MustParse("f51298ea-7b1a-532a-8a0b-35d02ab28c3e"),
				ComplaintId: uuid.MustParse("f51298ea-7b1a-532a-8a0b-35d02ab28c3e"),
				Sender:      NewRecipients["user"],
				Body:        "This is the body of the complaint and behaves as the first reply of it",
				CreatedAt:   CommonDate,
				Read:        false,
				ReadAt:      CommonDate,
				UpdatedAt:   CommonDate,
			},
		},
	},
}
