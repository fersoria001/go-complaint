package tests

import (
	"encoding/json"
	"go-complaint/domain/model/common"

	"time"
)

type FakeEvent struct {
	occurredOn time.Time
}

func NewFakeEvent() *FakeEvent {
	fe := &FakeEvent{}
	fe.occurredOn = time.Now()
	return fe
}

func (e *FakeEvent) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		OccurredOn string
	}{
		OccurredOn: common.StringDate(e.occurredOn),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (e *FakeEvent) UnmarshalJSON(data []byte) error {
	aux := struct {
		OccurredOn string
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	e.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}

func (e FakeEvent) OccurredOn() time.Time {
	return e.occurredOn
}
