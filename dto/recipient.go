package dto

import "go-complaint/domain/model/recipient"

type Recipient struct {
	Id               string `json:"id"`
	SubjectName      string `json:"subjectName"`
	SubjectThumbnail string `json:"subjectThumbnail"`
	SubjectEmail     string `json:"subjectEmail"`
	IsEnterprise     bool   `json:"isEnterprise"`
	IsOnline         bool   `json:"isOnline"`
}

func NewRecipient(obj recipient.Recipient) *Recipient {
	return &Recipient{
		Id:               obj.Id().String(),
		SubjectName:      obj.SubjectName(),
		SubjectThumbnail: obj.SubjectThumbnail(),
		SubjectEmail:     obj.SubjectEmail(),
		IsEnterprise:     obj.IsEnterprise(),
	}
}
