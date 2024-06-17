package identity

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

// Package identity
// <<Domain event>>
type PersonCountryStateChanged struct {
	userID     string
	oldValue   int
	newValue   int
	occurredOn time.Time
}

func NewPersonCountryStateChanged(
	userID string,
	oldValue int,
	newValue int,
) (*PersonCountryStateChanged, error) {
	var event = new(PersonCountryStateChanged)
	event.userID = userID
	event.oldValue = oldValue
	event.newValue = newValue
	event.occurredOn = time.Now()
	return event, nil
}

func (eu *PersonCountryStateChanged) OccurredOn() time.Time {
	return eu.occurredOn
}

func (eu *PersonCountryStateChanged) MarshalJSON() ([]byte, error) {
	commonDate := common.NewDate(eu.occurredOn)
	return json.Marshal(map[string]interface{}{
		"user_id":     eu.userID,
		"old_value":   eu.oldValue,
		"new_value":   eu.newValue,
		"occurred_on": commonDate.StringRepresentation(),
	})
}

func (eu *PersonCountryStateChanged) UnmarshalJSON(b []byte) error {
	var v struct {
		UserID     string `json:"user_id"`
		OldValue   int    `json:"old_value"`
		NewValue   int    `json:"new_value"`
		OccurredOn string `json:"occurred_on"`
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	eu.userID = v.UserID
	eu.oldValue = v.OldValue
	eu.newValue = v.NewValue
	parsedDate, err := common.ParseDate(v.OccurredOn)
	if err != nil {
		return err
	}
	eu.occurredOn = parsedDate
	return nil
}
