package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"go-complaint/erros"
	"time"

	"github.com/google/uuid"
)

type EnterpriseCreated struct {
	enterpriseId uuid.UUID
	industryId   int
	occurredOn   time.Time
}

func (event EnterpriseCreated) OccurredOn() time.Time {
	return event.occurredOn
}

func NewEnterpriseCreated(
	enterpriseId uuid.UUID,
	industryId int,
	occurredOn time.Time,
) (*EnterpriseCreated, error) {
	newEvent := new(EnterpriseCreated)
	err := newEvent.setEnterpriseId(enterpriseId)
	if err != nil {
		return nil, err
	}
	err = newEvent.setIndustryId(industryId)
	if err != nil {
		return nil, err
	}
	err = newEvent.setOccurredOn(occurredOn)
	if err != nil {
		return nil, err
	}
	return newEvent, nil
}

func (event *EnterpriseCreated) setEnterpriseId(enterpriseId uuid.UUID) error {
	if enterpriseId == uuid.Nil {
		return ErrNilPointer
	}
	event.enterpriseId = enterpriseId
	return nil
}

func (event *EnterpriseCreated) setIndustryId(industryId int) error {
	event.industryId = industryId
	return nil
}

func (event *EnterpriseCreated) setOccurredOn(occurredOn time.Time) error {
	if occurredOn.IsZero() {
		return &erros.NullValueError{}
	}
	if occurredOn.After(time.Now()) {
		return &erros.InvalidDateError{}
	}
	if occurredOn == (time.Time{}) {
		return &erros.EmptyStructError{}
	}
	event.occurredOn = occurredOn
	return nil
}

func (event *EnterpriseCreated) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		EnterpriseId uuid.UUID
		IndustryId   int
		OccurredOn   string
	}{
		EnterpriseId: event.enterpriseId,
		IndustryId:   event.industryId,
		OccurredOn:   common.StringDate(event.occurredOn),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (event *EnterpriseCreated) UnmarshalJSON(data []byte) error {
	if event == nil {
		return &erros.NullValueError{}
	}
	j := struct {
		EnterpriseId uuid.UUID
		IndustryId   int
		OccurredOn   string
	}{}
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}
	occurredOn, err := common.ParseDate(j.OccurredOn)
	if err != nil {
		return err
	}
	err = event.setEnterpriseId(j.EnterpriseId)
	if err != nil {
		return err
	}
	err = event.setIndustryId(j.IndustryId)
	if err != nil {
		return err
	}
	err = event.setOccurredOn(occurredOn)
	if err != nil {
		return err
	}

	return nil
}
