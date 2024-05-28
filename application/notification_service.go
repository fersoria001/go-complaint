package application

import (
	"context"
	"encoding/json"
	"go-complaint/domain/model/enterprise"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"
)

type NotificationService struct {
	eventRepository *repositories.EventRepository
}

func NewNotificationService(
	eventRepository *repositories.EventRepository,
) *NotificationService {
	return &NotificationService{
		eventRepository: eventRepository,
	}
}
func (n *NotificationService) EnterpriseNotifications(
	ctx context.Context,
	enterpriseID string,
) (*dto.EnterpriseNotifications, error) {
	notifications := dto.EnterpriseNotifications{
		Count:                      0,
		EmployeeWaitingForApproval: make([]dto.EmployeeWaitingForApproval, 0),
	}
	employeeWaitingForApprovalType := reflect.TypeOf(&enterprise.EmployeeWaitingForApproval{}).String()
	events, err := n.eventRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for e := range events.Iter() {
		switch e.TypeName {
		case employeeWaitingForApprovalType:
			var event enterprise.EmployeeWaitingForApproval
			err = json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return nil, err
			}
			if event.EnterpriseName() == enterpriseID {
				notifications.Count++
				notifications.EmployeeWaitingForApproval = append(notifications.EmployeeWaitingForApproval,
					dto.NewEmployeeWaitingForApproval(e.EventId.String(), false, event),
				)
			}
		default:
			continue
		}
	}
	log, err := n.eventRepository.GetAllFromLog(ctx)
	if err != nil {
		return nil, err
	}
	for e := range log.Iter() {
		switch e.TypeName {
		case employeeWaitingForApprovalType:
			continue
		default:
			continue
		}
	}
	return &notifications, nil
}
func (n *NotificationService) UserNotifications(
	ctx context.Context,
	userID string,
) (*dto.UserNotifications, error) {
	notifications := dto.UserNotifications{
		Count:            0,
		HiringInvitation: make([]dto.HiringInvitation, 0),
	}
	hiringInvitationType := reflect.TypeOf(&enterprise.HiringInvitationSent{}).String()
	events, err := n.eventRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for e := range events.Iter() {
		switch e.TypeName {
		case hiringInvitationType:
			var event enterprise.HiringInvitationSent
			err = json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return nil, err
			}
			if event.Email() == userID {
				notifications.Count++
				notifications.HiringInvitation = append(
					notifications.HiringInvitation,
					dto.NewHiringInvitation(e.EventId.String(), false, event),
				)
			}
		default:
			continue
		}
	}
	log, err := n.eventRepository.GetAllFromLog(ctx)
	if err != nil {
		return nil, err
	}
	for e := range log.Iter() {
		switch e.TypeName {
		case hiringInvitationType:
			var event enterprise.HiringInvitationSent
			err = json.Unmarshal(e.EventBody, &event)
			if err != nil {
				return nil, err
			}
			if event.Email() == userID {
				notifications.Count++
				notifications.HiringInvitation = append(
					notifications.HiringInvitation,
					dto.NewHiringInvitation(e.EventId.String(), true, event),
				)
			}
		default:
			continue
		}
	}
	return &notifications, nil
}

func (n *NotificationService) PendingHires(
	ctx context.Context,
	enterpriseID string) ([]enterprise.EmployeeWaitingForApproval,
	error) {
	actualModel := &enterprise.EmployeeWaitingForApproval{}
	actualTypeName := reflect.TypeOf(actualModel).String()
	actualEvents, err := n.eventRepository.FindByTypeName(ctx, actualTypeName)
	if err != nil {
		return nil, err
	}
	results := make([]enterprise.EmployeeWaitingForApproval, 0)
	for e := range actualEvents.Iter() {
		var event enterprise.EmployeeWaitingForApproval
		err = json.Unmarshal(e.EventBody, &event)
		if err != nil {
			return nil, err
		}
		if event.EnterpriseName() == enterpriseID {
			results = append(results, event)
		}
	}
	return results, nil
}

// func (n *NotificationService) NotifyHiringInvitationAccepted(
// 	ctx context.Context,
// 	eventID string,
// ) error {
// 	j, err := n.eventRepository.Get(ctx, eventID)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	var event enterprise.HiringInvitationSent
// 	err = json.Unmarshal(j.EventBody, &event)
// 	if err != nil {
// 		return err
// 	}
// 	ep, err := n.enterpriseService.Enterprise(ctx, event.EnterpriseID())
// 	if err != nil {
// 		return err
// 	}

// 	err = n.enterpriseService.AcceptHiringInvitation(ctx,
// 		eventID,
// 		event.EnterpriseID(),
// 		ep.OwnerID,
// 		event.ProfileIMG(),
// 		event.FirstName(),
// 		event.LastName(),
// 		event.Email(),
// 		event.Phone(),
// 		event.Age(),
// 		event.ProposalPosition().String(),
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	err = n.eventRepository.SaveToLog(ctx, *j)
// 	if err != nil {
// 		return err
// 	}
// 	err = n.eventRepository.Remove(ctx, eventID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (n *NotificationService) MoveToLog(ctx context.Context, eventID string) error {
	j, err := n.eventRepository.Get(ctx, eventID)
	if err != nil {
		return err
	}
	err = n.eventRepository.SaveToLog(ctx, *j)
	if err != nil {
		return err
	}
	err = n.eventRepository.Remove(ctx, eventID)
	if err != nil {
		return err
	}
	return nil
}

func (n *NotificationService) ProvideFromStore(ctx context.Context, eventID string) (*dto.StoredEvent, error) {
	return n.eventRepository.Get(ctx, eventID)
}

// Dedicated to the manager created invitation
// func (n *NotificationService) NotifyJobSelectionInvitationAccepted(
// 	ctx context.Context,
// 	eventID string,
// ) error {
// 	j, err := n.eventRepository.Get(ctx, eventID)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	var event enterprise.JobSelectionInvitationSent
// 	err = json.Unmarshal(j.EventBody, &event)
// 	if err != nil {
// 		return err
// 	}
// 	err = n.employeesService.HireEmployee(
// 		ctx,
// 		event.EnterpriseID(),
// 		event.ManagerID(),
// 		event.ProfileIMG(),
// 		event.FirstName(),
// 		event.LastName(),
// 		event.Email(),
// 		event.Phone(),
// 		event.Age(),
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	err = n.eventRepository.SaveToLog(ctx, *j)
// 	if err != nil {
// 		return err
// 	}
// 	err = n.eventRepository.Remove(ctx, eventID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (s *NotificationService) Notification(ctx context.Context, id string) ([]byte, error) {
	notification, err := s.eventRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return notification.EventBody, nil
}
