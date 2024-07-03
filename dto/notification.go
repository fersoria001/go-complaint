package dto

import (
	"go-complaint/domain"
	"go-complaint/domain/model/common"
)

type Notification struct {
	ID         string `json:"id"`
	OwnerID    string `json:"ownerID"`
	Thumbnail  string `json:"thumbnail"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Link       string `json:"link"`
	Seen       bool   `json:"seen"`
	OccurredOn string `json:"occurredOn"`
}

func NewNotification(domainObj domain.Notification) Notification {
	ms := common.StringDate(domainObj.OccurredOn())
	return Notification{
		ID:         domainObj.ID().String(),
		OwnerID:    domainObj.OwnerID(),
		Thumbnail:  domainObj.Thumbnail(),
		Title:      domainObj.Title(),
		Content:    domainObj.Content(),
		Link:       domainObj.Link(),
		Seen:       domainObj.Seen(),
		OccurredOn: ms,
	}
}
