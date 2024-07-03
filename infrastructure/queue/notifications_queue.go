package queue

import (
	"go-complaint/dto"
	"go-complaint/infrastructure/cache"
)

type NotificationList struct {
	*LinkedList[dto.Notification]
}

func NewNotificationList(operationID string) (*NotificationList, error) {
	notificationList := NewLinkedList[dto.Notification]()
	cache.RequestChannel <- cache.Request{
		Type:    cache.WRITE,
		Payload: notificationList,
		Key:     operationID,
	}
	return &NotificationList{
		LinkedList: notificationList,
	}, nil
}
