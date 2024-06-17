package queries

import (
	"context"
	"encoding/json"
	"go-complaint/domain/model/complaint"
	"go-complaint/dto"
	employeefindall "go-complaint/infrastructure/persistence/finders/employee_findall"
	"go-complaint/infrastructure/persistence/finders/find_all_events"
	"go-complaint/infrastructure/persistence/repositories"

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
	return dto.NewEmployee(*emp), nil
}

func (query EmployeeQuery) Employees(ctx context.Context) ([]dto.Employee, error) {
	emp, err := repositories.MapperRegistryInstance().Get("Employee").(repositories.EmployeeRepository).FindAll(
		ctx,
		employeefindall.NewByEnterpriseID(query.EnterpriseID),
	)
	if err != nil {
		return nil, err
	}
	var employees []dto.Employee
	for e := range emp.Iter() {
		employees = append(employees, dto.NewEmployee(*e))
	}
	return employees, nil
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
