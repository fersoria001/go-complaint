package dto

import (
	"go-complaint/domain"
	"go-complaint/domain/model/common"
)

type Notification struct {
	Id         string    `json:"id"`
	Owner      Recipient `json:"owner"`
	Sender     Recipient `json:"sender"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Link       string    `json:"link"`
	Seen       bool      `json:"seen"`
	OccurredOn string    `json:"occurredOn"`
}

func NewNotification(domainObj domain.Notification) *Notification {
	ms := common.StringDate(domainObj.OccurredOn())
	return &Notification{
		Id:         domainObj.ID().String(),
		Owner:      *NewRecipient(domainObj.Owner()),
		Sender:     *NewRecipient(domainObj.Sender()),
		Title:      domainObj.Title(),
		Content:    domainObj.Content(),
		Link:       domainObj.Link(),
		Seen:       domainObj.Seen(),
		OccurredOn: ms,
	}
}
