package queries

import (
	"context"
	"encoding/json"
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/feedback"
	"go-complaint/dto"
	employeefindall "go-complaint/infrastructure/persistence/finders/employee_findall"
	"go-complaint/infrastructure/persistence/finders/find_all_events"
	"go-complaint/infrastructure/persistence/repositories"
	"net/mail"

	"github.com/google/uuid"
)

type EmployeeQuery struct {
	EmployeeID   string
	EnterpriseID string
}

func (query EmployeeQuery) Employee(ctx context.Context) (dto.Employee, error) {
	parsedID, err := uuid.Parse(query.EmployeeID)
	if err != nil {
		return dto.Employee{}, err
	}
	emp, err := repositories.MapperRegistryInstance().Get("Employee").(repositories.EmployeeRepository).Get(ctx, parsedID)
	if err != nil {
		return dto.Employee{}, err
	}
	employeeDto := dto.NewEmployee(*emp)
	storedEvents, err := repositories.MapperRegistryInstance().Get("Event").(repositories.EventRepository).FindAll(
		ctx,
		find_all_events.By(),
	)
	if err != nil {
		return dto.Employee{}, err
	}
	ratedsMap := make(map[string]bool)
	for e := range storedEvents.Iter() {
		if e.TypeName == "*complaint.ComplaintRated" {
			var event complaint.ComplaintRated
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {

				return dto.Employee{}, err
			}
			ratedsMap[event.ComplaintID().String()] = true
		}
	}
	for e := range storedEvents.Iter() {
		switch e.TypeName {
		case "*complaint.ComplaintSentForReview":
			var event complaint.ComplaintSentForReview
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return dto.Employee{}, err
			}

			if _, err := mail.ParseAddress(event.ReceiverID()); err != nil {

				if employeeDto.Email == event.TriggeredBy() {
					if _, ok := ratedsMap[event.ComplaintID().String()]; ok {
						employeeDto.SetComplaintsSolved(employeeDto.ComplaintsSolved + 1)
						employeeDto.AppendComplaintsSolvedIds(event.ComplaintID().String())
					}
				}

			}
		case "*complaint.ComplaintRated":
			var event complaint.ComplaintRated
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return dto.Employee{}, err
			}
			if employeeDto.Email == event.RatedBy() {
				employeeDto.SetComplaintsRated(employeeDto.ComplaintsRated + 1)
				employeeDto.AppendComplaintsRatedIDs(event.ComplaintID().String())
			}
		case "*enterprise.HiringInvitationSent":
			var event enterprise.HiringInvitationSent
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return dto.Employee{}, err
			}
			if employeeDto.Email == event.UserID() {
				employeeDto.SetHireInvitationsSent(employeeDto.HireInvitationsSent + 1)
			}

		case "*enterprise.EmployeeHired":
			var event enterprise.EmployeeHired
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return dto.Employee{}, err
			}
			if employeeDto.Email == event.EmitedBy() {
				employeeDto.SetEmployeesHired(employeeDto.EmployeesHired + 1)
			}
		case "*feedback.AddedFeedback":
			var event feedback.AddedFeedback
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return dto.Employee{}, err
			}
			if employeeDto.Email == event.ReviewedID() {
				employeeDto.SetFeedbackReceived(employeeDto.FeedbackReceived + 1)
				employeeDto.AppendFeedbackReceivedIDs(event.FeedbackID().String())
			} else if employeeDto.Email == event.ReviewerID() {
				employeeDto.SetComplaintsFeedbacked(employeeDto.ComplaintsFeedbacked + 1)
				employeeDto.AppendComplaintsFeedbackedIDs(event.ComplaintID().String())
			}
		}
	}
	return *employeeDto, nil
}

func (query EmployeeQuery) Employees(ctx context.Context) ([]dto.Employee, error) {
	emp, err := repositories.MapperRegistryInstance().Get("Employee").(repositories.EmployeeRepository).FindAll(
		ctx,
		employeefindall.NewByEnterpriseID(query.EnterpriseID),
	)
	if err != nil {
		return nil, err
	}
	employees := make(map[string]*dto.Employee, 0)
	for i := range emp.Iter() {
		employees[i.Email()] = dto.NewEmployee(*i)
	}

	storedEvents, err := repositories.MapperRegistryInstance().Get("Event").(repositories.EventRepository).FindAll(
		ctx,
		find_all_events.By(),
	)
	if err != nil {
		return nil, err
	}
	ratedsMap := make(map[string]bool)
	for e := range storedEvents.Iter() {
		if e.TypeName == "*complaint.ComplaintRated" {
			var event complaint.ComplaintRated
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return nil, err
			}

			ratedsMap[event.ComplaintID().String()] = true
		}
	}
	for e := range storedEvents.Iter() {
		switch e.TypeName {
		case "*complaint.ComplaintSentForReview":
			var event complaint.ComplaintSentForReview
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return nil, err
			}
			if _, err := mail.ParseAddress(event.ReceiverID()); err != nil {
				if v, ok := employees[event.TriggeredBy()]; ok {

					if _, ok := ratedsMap[event.ComplaintID().String()]; ok {

						v.SetComplaintsSolved(v.ComplaintsSolved + 1)
						v.AppendComplaintsSolvedIds(event.ComplaintID().String())
					}

				}
			}
		case "*complaint.ComplaintRated":
			var event complaint.ComplaintRated
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return nil, err
			}
			if v, ok := employees[event.RatedBy()]; ok {
				v.SetComplaintsRated(v.ComplaintsRated + 1)
				v.AppendComplaintsRatedIDs(event.ComplaintID().String())
			}
		case "*enterprise.HiringInvitationSent":
			var event enterprise.HiringInvitationSent
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return nil, err
			}
			if v, ok := employees[event.UserID()]; ok {
				v.SetHireInvitationsSent(v.HireInvitationsSent + 1)
			}
		case "*enterprise.EmployeeHired":
			var event enterprise.EmployeeHired
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return nil, err
			}
			if v, ok := employees[event.EmitedBy()]; ok {
				v.SetEmployeesHired(v.EmployeesHired + 1)
			}
		case "*enterprise.EmployeeFired":
			var event enterprise.EmployeeFired
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return nil, err
			}
			if v, ok := employees[event.EmitedBy()]; ok {
				v.SetEmployeesFired(v.EmployeesFired + 1)
			}
		case "*feedback.AddedFeedback":
			var event feedback.AddedFeedback
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return nil, err
			}
			if v, ok := employees[event.ReviewedID()]; ok {
				v.SetFeedbackReceived(v.FeedbackReceived + 1)
				v.AppendFeedbackReceivedIDs(event.FeedbackID().String())
			} else if v, ok := employees[event.ReviewerID()]; ok {
				v.SetComplaintsFeedbacked(v.ComplaintsFeedbacked + 1)
				v.AppendComplaintsFeedbackedIDs(event.ComplaintID().String())
			}
		}
	}
	result := make([]dto.Employee, 0)
	for _, v := range employees {
		result = append(result, *v)
	}
	return result, nil
}

func (query EmployeeQuery) SolvedComplaints(
	ctx context.Context,
) ([]dto.ComplaintDTO, error) {
	if query.EmployeeID == "" {
		return nil, ErrBadRequest
	}
	storedEvents, err := repositories.MapperRegistryInstance().Get("Event").(repositories.EventRepository).FindAll(
		ctx,
		find_all_events.By(),
	)
	if err != nil {
		return nil, err
	}
	var sentToReview = make(map[string]map[string]complaint.ComplaintSentForReview)
	for e := range storedEvents.Iter() {
		if e.TypeName == "*complaint.ComplaintSentForReview" {
			var event complaint.ComplaintSentForReview
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return nil, err
			}
			if event.TriggeredBy() == query.EmployeeID {
				sentToReview[string(event.ComplaintID().String())] = make(map[string]complaint.ComplaintSentForReview)
				sentToReview[string(event.ComplaintID().String())][e.EventId.String()] = event
			}
		}
	}
	for e := range storedEvents.Iter() {
		switch e.TypeName {
		case "*complaint.ComplaintRated":
			var event complaint.ComplaintRated
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return nil, err
			}
			delete(sentToReview[string(event.ComplaintID().String())], event.ComplaintID().String())
		case "*complaint.ComplaintSentToHistory":
			var event complaint.ComplaintSentToHistory
			err := json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return nil, err
			}
			delete(sentToReview[string(event.ComplaintID().String())], event.ComplaintID().String())
		}
	}
	var complaints []dto.ComplaintDTO
	for _, v := range sentToReview {
		for _, c := range v {
			dbComplaint, err := repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).Get(ctx,
				c.ComplaintID())
			if err != nil {
				return nil, err
			}
			complaints = append(complaints, dto.NewComplaintDTO(*dbComplaint))
		}
	}
	return complaints, nil

}
