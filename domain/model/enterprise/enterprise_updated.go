package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

// Package enterprise
// <<Domain event>>
type EnterpriseUpdated struct {
	enterpriseName string
	industryName   string
	oldValues      map[string]string
	newValues      map[string]string
	occurredOn     time.Time
}

func NewEnterpriseUpdated(enterpriseName, industryName string,
	 oldValues map[string]string, newValues map[string]string) (*EnterpriseUpdated, error) {
	return &EnterpriseUpdated{
		enterpriseName: enterpriseName,
		industryName:   industryName,
		oldValues:      oldValues,
		newValues:      newValues,
		occurredOn:     time.Now(),
	}, nil
}

func (eu *EnterpriseUpdated) OccurredOn() time.Time {
	return eu.occurredOn
}


func (eu *EnterpriseUpdated) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		EnterpriseName string
		IndustryName   string
		OldValues      map[string]string
		NewValues      map[string]string

		OccurredOn     string
	}{
		EnterpriseName: eu.enterpriseName,
		IndustryName:   eu.industryName,
		OldValues:      eu.oldValues,
		NewValues:      eu.newValues,

		OccurredOn:     common.StringDate(eu.occurredOn),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (eu *EnterpriseUpdated) UnmarshalJSON(b []byte) error {
	var data struct {
		EnterpriseName string
		IndustryName   string
		OldValues      map[string]string
		NewValues      map[string]string

		OccurredOn     string
	}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	eu.enterpriseName = data.EnterpriseName
	eu.industryName = data.IndustryName
	eu.oldValues = data.OldValues
	eu.newValues = data.NewValues
	eu.occurredOn, err = common.ParseDate(data.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}
