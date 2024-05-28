package identity

import (
	"encoding/json"
	"time"
)

type PasswordChanged struct {
	userID      string
	oldPassword string
	occurredOn  time.Time
}

func NewPasswordChanged(userID, oldPassword string) *PasswordChanged {
	return &PasswordChanged{
		userID:      userID,
		oldPassword: oldPassword,
		occurredOn:  time.Now(),
	}
}

func (pc *PasswordChanged) OccurredOn() time.Time {
	return pc.occurredOn
}

func (pc *PasswordChanged) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserID      string    `json:"user_id"`
		OldPassword string    `json:"old_password"`
		OccurredOn  time.Time `json:"occurred_on"`
	}{
		UserID:      pc.userID,
		OldPassword: pc.oldPassword,
		OccurredOn:  pc.occurredOn,
	})
}

func (pc *PasswordChanged) UnmarshalJSON(data []byte) error {
	aux := new(struct {
		UserID      string    `json:"user_id"`
		OldPassword string    `json:"old_password"`
		OccurredOn  time.Time `json:"occurred_on"`
	})
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	pc.oldPassword = aux.OldPassword
	pc.userID = aux.UserID
	pc.occurredOn = aux.OccurredOn
	return nil
}
