package enterprise

import (
	"go-complaint/domain/model/recipient"
	"time"

	"github.com/google/uuid"
)

type EnterpriseActivityType int

const (
	FeedbacksStarted EnterpriseActivityType = iota
	FeedbacksReceived
	JobProposalsSent
	EmployeesHired
	EmployeesFired
	ComplaintSent
	ComplaintResolved
	ComplaintReviewed
	ComplaintReplied
)

func (eat EnterpriseActivityType) String() string {
	switch eat {
	case FeedbacksStarted:
		return "FEEDBACKS_STARTED"
	case FeedbacksReceived:
		return "FEEDBACKS_RECEIVED"
	case JobProposalsSent:
		return "JOB_PROPOSALS_SENT"
	case EmployeesHired:
		return "EMPLOYEES_HIRED"
	case EmployeesFired:
		return "EMPLOYEES_FIRED"
	case ComplaintSent:
		return "COMPLAINT_SENT"
	case ComplaintResolved:
		return "COMPLAINT_RESOLVED"
	case ComplaintReviewed:
		return "COMPLAINT_REVIEWED"
	case ComplaintReplied:
		return "COMPLAINT_REPLIED"
	default:
		return ""
	}
}

func ParseActivityType(v string) EnterpriseActivityType {
	switch v {
	case "FEEDBACKS_STARTED":
		return FeedbacksStarted
	case "FEEDBACKS_RECEIVED":
		return FeedbacksReceived
	case "JOB_PROPOSALS_SENT":
		return JobProposalsSent
	case "EMPLOYEES_HIRED":
		return EmployeesHired
	case "EMPLOYEES_FIRED":
		return EmployeesFired
	case "COMPLAINT_SENT":
		return ComplaintSent
	case "COMPLAINT_RESOLVED":
		return ComplaintResolved
	case "COMPLAINT_REVIEWED":
		return ComplaintReviewed
	case "COMPLAINT_REPLIED":
		return ComplaintReplied
	default:
		return -1
	}
}

type EnterpriseActivity struct {
	id             uuid.UUID
	user           recipient.Recipient
	activityId     uuid.UUID
	enterpriseId   uuid.UUID
	enterpriseName string
	occurredOn     time.Time
	activityType   EnterpriseActivityType
}

func NewEnterpriseActivity(id, activityId, enterpriseId uuid.UUID, user recipient.Recipient,
	enterpriseName string, occurredOn time.Time, activityType EnterpriseActivityType) *EnterpriseActivity {
	return &EnterpriseActivity{
		id:             id,
		user:           user,
		activityId:     activityId,
		enterpriseId:   enterpriseId,
		enterpriseName: enterpriseName,
		occurredOn:     occurredOn,
		activityType:   activityType,
	}
}

func (ea EnterpriseActivity) Id() uuid.UUID {
	return ea.id
}

func (ea EnterpriseActivity) User() recipient.Recipient {
	return ea.user
}

func (ea EnterpriseActivity) ActivityId() uuid.UUID {
	return ea.activityId
}

func (ea EnterpriseActivity) EnterpriseId() uuid.UUID {
	return ea.enterpriseId
}

func (ea EnterpriseActivity) EnterpriseName() string {
	return ea.enterpriseName
}

func (ea EnterpriseActivity) OccurredOn() time.Time {
	return ea.occurredOn
}

func (ea EnterpriseActivity) ActivityType() EnterpriseActivityType {
	return ea.activityType
}
