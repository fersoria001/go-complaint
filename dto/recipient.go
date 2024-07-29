package dto

import "go-complaint/domain/model/recipient"

type Recipient struct {
	Id               string `json:"id"`
	SubjectName      string `json:"subjectName"`
	SubjectThumbnail string `json:"subjectThumbnail"`
	IsEnterprise     bool   `json:"isEnterprise"`
}

func NewRecipient(obj recipient.Recipient) *Recipient {
	return &Recipient{
		Id:           obj.Id().String(),
		SubjectName:  obj.SubjectThumbnail(),
		IsEnterprise: obj.IsEnterprise(),
	}
}
