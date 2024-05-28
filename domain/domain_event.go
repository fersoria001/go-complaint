package domain

import (
	"encoding/json"
	"time"
)

// Package events
// All dates from events should be in milliseconds after
// serialization
type DomainEvent interface {
	OccurredOn() time.Time
	json.Unmarshaler
	json.Marshaler
}
