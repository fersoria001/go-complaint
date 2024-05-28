package identity

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

type UserUpdated struct {
	userID     string
	oldValues  map[string]string
	newValues  map[string]string
	occurredOn time.Time
}

func NewUserUpdated(
	userID string,
	oldValues map[string]string,
	newValues map[string]string,
) *UserUpdated {
	return &UserUpdated{
		userID:     userID,
		oldValues:  oldValues,
		newValues:  newValues,
		occurredOn: time.Now(),
	}
}

func (uu *UserUpdated) OccurredOn() time.Time {
	return uu.occurredOn
}

func (uu *UserUpdated) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserID     string
		OldValues  map[string]string
		NewValues  map[string]string
		OccurredOn string
	}{
		UserID:     uu.userID,
		OldValues:  uu.oldValues,
		NewValues:  uu.newValues,
		OccurredOn: common.StringDate(uu.occurredOn),
	})
}

func (uu *UserUpdated) UnmarshalJSON(b []byte) error {
	var data struct {
		UserID     string
		OldValues  map[string]string
		NewValues  map[string]string
		OccurredOn string
	}
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	occurredOn, err := common.ParseDate(data.OccurredOn)
	if err != nil {
		return err
	}

	uu.userID = data.UserID
	uu.oldValues = data.OldValues
	uu.newValues = data.NewValues
	uu.occurredOn = occurredOn

	return nil
}
