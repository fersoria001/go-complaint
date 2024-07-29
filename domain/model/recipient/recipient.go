package recipient

import (
	"net/mail"

	"github.com/google/uuid"
)

type Recipient struct {
	id               uuid.UUID
	subjectName      string
	subjectThumbnail string
	subjectEmail     string
	isEnterprise     bool
}

func NewRecipient(id uuid.UUID, subjectName, subjectThumbnail, subjectEmail string, isEnterprise bool) *Recipient {
	return &Recipient{
		id:               id,
		subjectName:      subjectName,
		subjectThumbnail: subjectThumbnail,
		subjectEmail:     subjectEmail,
		isEnterprise:     isEnterprise,
	}
}

func (r *Recipient) SetSubjectEmail(email string) error {
	if email == "" {
		return ValidationError{Message: "email cannot be nil"}
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}
	r.subjectEmail = email
	return nil
}

func (r *Recipient) SetSubjectThumbnail(thumbnail string) error {
	if thumbnail == "" {
		return ValidationError{Message: "name cannot be nil"}
	}
	r.subjectThumbnail = thumbnail
	return nil
}

func (r *Recipient) SetSubjectName(name string) error {
	if name == "" {
		return ValidationError{Message: "name cannot be nil"}
	}
	r.subjectName = name
	return nil
}

func (r Recipient) Id() uuid.UUID {
	return r.id
}

func (r Recipient) SubjectName() string {
	return r.subjectName
}

func (r Recipient) SubjectThumbnail() string {
	return r.subjectThumbnail
}

func (r Recipient) SubjectEmail() string {
	return r.subjectEmail
}

func (r Recipient) IsEnterprise() bool {
	return r.isEnterprise
}
