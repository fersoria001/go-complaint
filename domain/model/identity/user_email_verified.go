package identity

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

type UserEmailVerified struct {
	userID     string
	occurredOn time.Time
}

func NewUserEmailVerified(userID string) *UserEmailVerified {
	return &UserEmailVerified{
		userID:     userID,
		occurredOn: time.Now(),
	}
}

func (e UserEmailVerified) UserID() string {
	return e.userID
}

func (e UserEmailVerified) OccurredOn() time.Time {
	return e.occurredOn
}

func (e *UserEmailVerified) MarshalJSON() ([]byte, error) {
	commonDate := common.NewDate(e.occurredOn)
	return json.Marshal(map[string]interface{}{
		"user_id":     e.userID,
		"occurred_on": commonDate,
	})
}

func (e *UserEmailVerified) UnmarshalJSON(data []byte) error {
	var v struct {
		UserID     string `json:"user_id"`
		OccurredOn string `json:"occurred_on"`
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	e.userID = v.UserID
	parsedDate, err := common.ParseDate(v.OccurredOn)
	if err != nil {
		return err
	}
	e.occurredOn = parsedDate
	return nil
}
