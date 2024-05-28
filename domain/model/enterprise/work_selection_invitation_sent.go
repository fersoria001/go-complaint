package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

type JobSelectionInvitationSent struct {
	enterpriseID string
	managerID    string
	profileIMG   string
	firstName    string
	lastName     string
	email        string
	phone        string
	age          int
	occurredOn   time.Time
}

func NewJobSelectionInvitationSent(
	enterpriseID,
	managerID,
	profileIMG,
	firstName,
	lastName,
	email,
	phone string,
	age int) *JobSelectionInvitationSent {
	return &JobSelectionInvitationSent{
		enterpriseID: enterpriseID,
		managerID:    managerID,
		profileIMG:   profileIMG,
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		phone:        phone,
		age:          age,
		occurredOn:   time.Now(),
	}
}

func (h *JobSelectionInvitationSent) OccurredOn() time.Time {
	return h.occurredOn
}

func (h *JobSelectionInvitationSent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		EnterpriseID string `json:"enterprise_id"`
		ManagerID    string `json:"manager_id"`
		ProfileIMG   string `json:"profile_img"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Email        string `json:"email"`
		Phone        string `json:"phone"`
		Age          int    `json:"age"`
		OccurredOn   string `json:"occurred_on"`
	}{
		EnterpriseID: h.enterpriseID,
		ManagerID:    h.managerID,
		ProfileIMG:   h.profileIMG,
		FirstName:    h.firstName,
		LastName:     h.lastName,
		Email:        h.email,
		Phone:        h.phone,
		Age:          h.age,
		OccurredOn:   common.StringDate(h.occurredOn),
	})
}

func (h *JobSelectionInvitationSent) UnmarshalJSON(data []byte) error {
	var alias struct {
		EnterpriseID string `json:"enterprise_id"`
		ManagerID    string `json:"manager_id"`
		ProfileIMG   string `json:"profile_img"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Email        string `json:"email"`
		Phone        string `json:"phone"`
		Age          int    `json:"age"`
		OccurredOn   string `json:"occurred_on"`
	}
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}
	h.profileIMG = alias.ProfileIMG
	h.enterpriseID = alias.EnterpriseID
	h.managerID = alias.ManagerID
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

func (h *JobSelectionInvitationSent) EnterpriseID() string {
	return h.enterpriseID
}

func (h *JobSelectionInvitationSent) ManagerID() string {
	return h.managerID
}

func (h *JobSelectionInvitationSent) FirstName() string {
	return h.firstName
}

func (h *JobSelectionInvitationSent) LastName() string {
	return h.lastName
}

func (h *JobSelectionInvitationSent) Email() string {
	return h.email
}

func (h *JobSelectionInvitationSent) Phone() string {
	return h.phone
}

func (h *JobSelectionInvitationSent) Age() int {
	return h.age
}

func (h *JobSelectionInvitationSent) ProfileIMG() string {
	return h.profileIMG
}
