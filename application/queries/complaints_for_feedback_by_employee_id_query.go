package queries

import (
	"context"
	"errors"
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/enterprise"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_all_complaints"
	"go-complaint/infrastructure/persistence/finders/find_all_enterprise_activity"
	"go-complaint/infrastructure/persistence/repositories"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ComplaintsForFeedbackByEmployeeIdQuery struct {
	EmployeeId string `json:"employeeId"`
}

func NewComplaintsForFeedbackByEmployeeIdQuery(id string) *ComplaintsForFeedbackByEmployeeIdQuery {
	return &ComplaintsForFeedbackByEmployeeIdQuery{
		EmployeeId: id,
	}
}

func (q ComplaintsForFeedbackByEmployeeIdQuery) Execute(ctx context.Context) ([]*dto.Complaint, error) {
	employeeId, err := uuid.Parse(q.EmployeeId)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	employeeRepository, ok := reg.Get("Employee").(repositories.EmployeeRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	enterpriseActivityRepository, ok := reg.Get("EnterpriseActivity").(repositories.EnterpriseActivityRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	complaintsRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	employee, err := employeeRepository.Get(ctx, employeeId)
	if err != nil {
		return nil, err
	}
	dbEa, err := enterpriseActivityRepository.FindAll(ctx, find_all_enterprise_activity.ByUserId(employee.GetUser().Id()))
	if err != nil {
		return nil, err
	}
	complaintIds := mapset.NewSet[uuid.UUID]()
	for _, v := range dbEa {
		switch v.ActivityType() {
		case enterprise.ComplaintReplied:
			complaintIds.Add(v.ActivityId())
		default:
			continue
		}
	}
	results := make([]*dto.Complaint, 0)
	dbC, err := complaintsRepository.FindAll(ctx, find_all_complaints.ByIdsWithStatus(
		complaintIds.ToSlice(),
		[]string{complaint.CLOSED.String()},
	))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return results, nil
		}
		return nil, err
	}
	for _, v := range dbC {
		results = append(results, dto.NewComplaint(*v))
	}
	return results, nil
}
