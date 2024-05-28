package models

import (
	"go-complaint/dto"

	"github.com/google/uuid"
)

type Event struct {
	EventId    uuid.UUID
	EventBody  []byte
	OccurredOn string
	TypeName   string
}

func NewEvent(e dto.StoredEvent) (*Event, error) {
	return &Event{
		EventId:    e.EventId,
		EventBody:  e.EventBody,
		OccurredOn: e.OccurredOn,
		TypeName:   e.TypeName,
	}, nil
}

func (e *Event) Columns() Columns {
	return Columns{
		"event_id",
		"event_body",
		"occurred_on",
		"type_name",
	}
}

func (e *Event) Values() Values {
	return Values{
		&e.EventId,
		&e.EventBody,
		&e.OccurredOn,
		&e.TypeName,
	}
}

func (e *Event) Table() string {
	return "events"
}

func (e *Event) Args() string {
	return "$1, $2, $3, $4"
}
