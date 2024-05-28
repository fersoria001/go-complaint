package identity

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"go-complaint/erros"
	"time"
)

// UserCreated is a domain event
// Implements the DomainEvent interface
type UserCreated struct {
	email      string
	occurredOn time.Time
}

func NewUserCreated(email string, ocurredOn time.Time) (*UserCreated, error) {
	newEvent := &UserCreated{
		email:      email,
		occurredOn: ocurredOn,
	}
	return newEvent, nil
}

func (event UserCreated) OccurredOn() time.Time {
	return event.occurredOn
}

func (event *UserCreated) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Email      string
		OccurredOn string
	}{
		Email:      event.email,
		OccurredOn: common.StringDate(event.occurredOn),
	})

	if err != nil {
		return nil, err
	}
	return j, nil
}
func (event *UserCreated) UnmarshalJSON(data []byte) error {
	if event == nil {
		return &erros.NullValueError{}
	}
	j := struct {
		Email      string
		OccurredOn string
	}{}
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}
	event.email = j.Email
	occurredOn, err := common.ParseDate(j.OccurredOn)
	if err != nil {
		return err
	}
	event.occurredOn = occurredOn

	return nil
}
