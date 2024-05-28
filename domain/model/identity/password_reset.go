package identity

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

// Package identity
// <<Domain event>> implements the domain.DomainEvent interface
type PasswordReset struct {
	email             string
	occurredOn        time.Time
	temporaryPassword string
}

func NewPasswordReset(email, temporaryPassword string) (*PasswordReset, error) {
	return &PasswordReset{
		email:             email,
		temporaryPassword: temporaryPassword,
		occurredOn:        time.Now(),
	}, nil
}

func (pr *PasswordReset) OccurredOn() time.Time {
	return pr.occurredOn
}

func (pr *PasswordReset) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Email             string `json:"email"`
		TemporaryPassword string `json:"temporary_password"`
		OccurredOn        string `json:"occurredOn"`
	}{
		Email:             pr.email,
		TemporaryPassword: pr.temporaryPassword,
		OccurredOn:        common.StringDate(pr.occurredOn),
	})
}

func (pr *PasswordReset) UnmarshalJSON(data []byte) error {
	var v struct {
		Email             string `json:"email"`
		TemporaryPassword string `json:"temporary_password"`
		OccurredOn        string `json:"occurredOn"`
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	pr.email = v.Email
	pr.temporaryPassword = v.TemporaryPassword
	date, err := common.ParseDate(v.OccurredOn)
	if err != nil {
		return err
	}
	pr.occurredOn = date
	return nil
}
