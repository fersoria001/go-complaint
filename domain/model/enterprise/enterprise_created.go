package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"go-complaint/erros"
	"time"
)

type EnterpriseCreated struct {
	enterpriseName string
	industryName   string
	ocurredOn      time.Time
}

func (event EnterpriseCreated) OccurredOn() time.Time {
	return event.ocurredOn
}

func NewEnterpriseCreated(enterpriseName, industryName string, occurredOn time.Time) (*EnterpriseCreated, error) {
	newEvent := &EnterpriseCreated{}
	err := newEvent.setEnterpriseName(enterpriseName)
	if err != nil {
		return nil, err
	}
	err = newEvent.setIndustryName(industryName)
	if err != nil {
		return nil, err
	}
	err = newEvent.setOccurredOn(occurredOn)
	if err != nil {
		return nil, err
	}
	return newEvent, nil
}

func (event *EnterpriseCreated) setEnterpriseName(enterpriseName string) error {
	if enterpriseName == "" {
		return &erros.NullValueError{}
	}
	event.enterpriseName = enterpriseName
	return nil
}

func (event *EnterpriseCreated) setIndustryName(industryName string) error {
	if industryName == "" {
		return &erros.NullValueError{}
	}
	event.industryName = industryName
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
	event.ocurredOn = occurredOn
	return nil
}

func (event *EnterpriseCreated) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		EnterpriseName string
		IndustryName   string
		OccurredOn     string
	}{
		EnterpriseName: event.enterpriseName,
		IndustryName:   event.industryName,
		OccurredOn:     common.StringDate(event.ocurredOn),
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
		EnterpriseName string
		IndustryName   string
		OccurredOn     string
	}{}
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}
	occurredOn, err := common.ParseDate(j.OccurredOn)
	if err != nil {
		return err
	}
	err = event.setEnterpriseName(j.EnterpriseName)
	if err != nil {
		return err
	}
	err = event.setIndustryName(j.IndustryName)
	if err != nil {
		return err
	}
	err = event.setOccurredOn(occurredOn)
	if err != nil {
		return err
	}

	return nil
}
