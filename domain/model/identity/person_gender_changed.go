package identity

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

// Package identity
// <<Domain event>>
type PersonGenreChanged struct {
	userID     string
	oldValue   string
	newValue   string
	occurredOn time.Time
}

func NewPersonGenreChanged(
	userID string,
	oldValue string,
	newValue string,
) (*PersonGenreChanged, error) {
	var event = new(PersonGenreChanged)
	event.userID = userID
	event.oldValue = oldValue
	event.newValue = newValue
	event.occurredOn = time.Now()
	return event, nil
}

func (eu *PersonGenreChanged) OccurredOn() time.Time {
	return eu.occurredOn
}

func (eu *PersonGenreChanged) MarshalJSON() ([]byte, error) {
	commonDate := common.NewDate(eu.occurredOn)
	return json.Marshal(map[string]interface{}{
		"user_id":     eu.userID,
		"old_value":   eu.oldValue,
		"new_value":   eu.newValue,
		"occurred_on": commonDate.StringRepresentation(),
	})
}

func (eu *PersonGenreChanged) UnmarshalJSON(b []byte) error {
	var v struct {
		UserID     string `json:"user_id"`
		OldValue   string `json:"old_value"`
		NewValue   string `json:"new_value"`
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
