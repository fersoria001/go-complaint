package dto

import (
	"go-complaint/domain/model/enterprise"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

// EmployeeWaitingForApproval is a struct type
type Employee struct {
	ID                      uuid.UUID `json:"id"`
	EnterpriseID            string    `json:"enterprise_id"`
	UserID                  string    `json:"user_id"`
	ProfileIMG              string    `json:"profile_img"`
	FirstName               string    `json:"first_name"`
	LastName                string    `json:"last_name"`
	Age                     int       `json:"age"`
	Email                   string    `json:"email"`
	Phone                   string    `json:"phone"`
	HiringDate              string    `json:"hiring_date"`
	ApprovedHiring          bool      `json:"approved_hiring"`
	ApprovedHiringAt        string    `json:"approved_hiring_at"`
	Position                string    `json:"position"`
	ComplaintsSolved        int       `json:"complaint_solved"`
	ComplaintsSolvedIds     []string  `json:"complaint_solved_ids"`
	ComplaintsRated         int       `json:"complaint_rated"`
	ComplaintsRatedIDs      []string  `json:"complaint_rated_ids"`
	ComplaintsFeedbacked    int       `json:"complaint_feedbacked"`
	ComplaintsFeedbackedIDs []string  `json:"complaint_feedbacked_ids"`
	FeedbackReceived        int       `json:"feedback_received"`
	FeedbackReceivedIDs     []string  `json:"feedback_received_ids"`
	HireInvitationsSent     int       `json:"hire_invitations_sent"`
	EmployeesHired          int       `json:"employees_hired"`
	EmployeesFired          int       `json:"employees_fired"`
}

func NewEmployee(domainObj *enterprise.Employee) *Employee {
	return &Employee{
		ID:               domainObj.ID(),
		ProfileIMG:       domainObj.ProfileIMG(),
		FirstName:        domainObj.FirstName(),
		LastName:         domainObj.LastName(),
		Age:              domainObj.Age(),
		Email:            domainObj.Email(),
		Phone:            domainObj.Phone(),
		HiringDate:       domainObj.HiringDate().StringRepresentation(),
		ApprovedHiring:   domainObj.ApprovedHiring(),
		ApprovedHiringAt: domainObj.ApprovedHiringAt().StringRepresentation(),
		Position:         domainObj.Position().String(),
	}
}

func (e *Employee) SetComplaintsSolved(complaintsSolved int) {
	e.ComplaintsSolved = complaintsSolved
}
func (e *Employee) AppendComplaintsSolvedIds(complaintsSolvedIds string) {
	if e.ComplaintsSolvedIds == nil {
		e.ComplaintsSolvedIds = []string{}
	}
	e.ComplaintsSolvedIds = append(e.ComplaintsSolvedIds, complaintsSolvedIds)
}
func (e *Employee) SetComplaintsRated(complaintsRated int) {
	e.ComplaintsRated = complaintsRated
}
func (e *Employee) AppendComplaintsRatedIDs(complaintsRatedIDs string) {
	if e.ComplaintsRatedIDs == nil {
		e.ComplaintsRatedIDs = []string{}
	}
	e.ComplaintsRatedIDs = append(e.ComplaintsRatedIDs, complaintsRatedIDs)
}
func (e *Employee) SetComplaintsFeedbacked(complaintsFeedbacked int) {
	e.ComplaintsFeedbacked = complaintsFeedbacked
}
func (e *Employee) AppendComplaintsFeedbackedIDs(complaintsFeedbackedIDs string) {
	if e.ComplaintsFeedbackedIDs == nil {
		e.ComplaintsFeedbackedIDs = []string{}
	}
	e.ComplaintsFeedbackedIDs = append(e.ComplaintsFeedbackedIDs, complaintsFeedbackedIDs)
}

func (e *Employee) SetFeedbackReceived(feedbackReceived int) {
	e.FeedbackReceived = feedbackReceived
}
func (e *Employee) AppendFeedbackReceivedIDs(feedbackReceivedIDs string) {
	if e.FeedbackReceivedIDs == nil {
		e.FeedbackReceivedIDs = []string{}
	}
	e.FeedbackReceivedIDs = append(e.FeedbackReceivedIDs, feedbackReceivedIDs)
}

func (e *Employee) SetHireInvitationsSent(hireInvitationsSent int) {
	e.HireInvitationsSent = hireInvitationsSent
}

func (e *Employee) SetEmployeesHired(employeesHired int) {
	e.EmployeesHired = employeesHired
}

func (e *Employee) SetEmployeesFired(employeesFired int) {
	e.EmployeesFired = employeesFired
}

func NewEmployeeList(domainObjs mapset.Set[enterprise.Employee]) []*Employee {
	employees := []*Employee{}
	for domainObj := range domainObjs.Iter() {
		employee := NewEmployee(&domainObj)
		employees = append(employees, employee)
	}
	return employees
}
