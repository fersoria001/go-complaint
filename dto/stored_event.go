package dto

import (
	"encoding/json"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"reflect"

	"github.com/google/uuid"
)

type StoredEvent struct {
	EventId    uuid.UUID `json:"event_id"`
	EventBody  []byte    `json:"event_body"`
	OccurredOn string    `json:"occurred_on"`
	TypeName   string    `json:"type_name"`
}

func NewStoredEvent(e domain.DomainEvent) (*StoredEvent, error) {
	eventID := uuid.New()
	date := common.StringDate(e.OccurredOn())
	typeName := reflect.TypeOf(e).String()
	body, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return &StoredEvent{
		EventId:    eventID,
		EventBody:  body,
		OccurredOn: date,
		TypeName:   typeName,
	}, nil
}
