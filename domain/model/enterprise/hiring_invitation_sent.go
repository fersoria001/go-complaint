package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

type HiringInvitationSent struct {
	enterpriseID     string
	employeeID       string
	proposalPosition Position
	profileIMG       string
	firstName        string
	lastName         string
	email            string
	phone            string
	age              int
	occurredOn       time.Time
}

func NewHiringInvitationSent(
	enterpriseID,
	employeeID,
	profileIMG,
	firstName,
	lastName,
	email,
	phone string,
	age int,
	proposalPosition Position) *HiringInvitationSent {
	return &HiringInvitationSent{
		enterpriseID:     enterpriseID,
		employeeID:       employeeID,
		profileIMG:       profileIMG,
		proposalPosition: proposalPosition,
		firstName:        firstName,
		lastName:         lastName,
		email:            email,
		phone:            phone,
		age:              age,
		occurredOn:       time.Now(),
	}
}

func (h *HiringInvitationSent) OccurredOn() time.Time {
	return h.occurredOn
}

func (h *HiringInvitationSent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		EnterpriseID     string `json:"enterprise_id"`
		EmployeeID       string `json:"employee_id"`
		ProfileIMG       string `json:"profile_img"`
		ProposalPosition string `json:"position_proposal"`
		FirstName        string `json:"first_name"`
		LastName         string `json:"last_name"`
		Email            string `json:"email"`
		Phone            string `json:"phone"`
		Age              int    `json:"age"`
		OccurredOn       string `json:"occurred_on"`
	}{
		EnterpriseID:     h.enterpriseID,
		EmployeeID:       h.employeeID,
		ProfileIMG:       h.profileIMG,
		ProposalPosition: h.proposalPosition.String(),
		FirstName:        h.firstName,
		LastName:         h.lastName,
		Email:            h.email,
		Phone:            h.phone,
		Age:              h.age,
		OccurredOn:       common.StringDate(h.occurredOn),
	})
}

func (h *HiringInvitationSent) UnmarshalJSON(data []byte) error {
	var alias struct {
		EnterpriseID     string `json:"enterprise_id"`
		EmployeeID       string `json:"employee_id"`
		ProfileIMG       string `json:"profile_img"`
		ProposalPosition string `json:"position_proposal"`
		FirstName        string `json:"first_name"`
		LastName         string `json:"last_name"`
		Email            string `json:"email"`
		Phone            string `json:"phone"`
		Age              int    `json:"age"`
		OccurredOn       string `json:"occurred_on"`
	}
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}
	position, err := ParsePosition(alias.ProposalPosition)
	if err != nil {
		return err
	}
	h.profileIMG = alias.ProfileIMG
	h.proposalPosition = position
	h.enterpriseID = alias.EnterpriseID
	h.employeeID = alias.EmployeeID
	h.firstName = alias.FirstName
	h.lastName = alias.LastName
	h.email = alias.Email
	h.phone = alias.Phone
	h.age = alias.Age
	h.occurredOn, err = common.ParseDate(alias.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}

func (h *HiringInvitationSent) EnterpriseID() string {
	return h.enterpriseID
}

func (h *HiringInvitationSent) EmployeeID() string {
	return h.employeeID
}

func (h *HiringInvitationSent) ProfileIMG() string {
	return h.profileIMG
}

func (h *HiringInvitationSent) ProposalPosition() Position {
	return h.proposalPosition
}

func (h *HiringInvitationSent) FirstName() string {
	return h.firstName
}

func (h *HiringInvitationSent) LastName() string {
	return h.lastName
}

func (h *HiringInvitationSent) Email() string {
	return h.email
}

func (h *HiringInvitationSent) Phone() string {
	return h.phone
}

func (h *HiringInvitationSent) Age() int {
	return h.age
}
