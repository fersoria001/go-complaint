package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"go-complaint/application/application_services"
	"go-complaint/domain"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/persistence/repositories"
	"net/mail"
	"reflect"

	"github.com/google/uuid"
)

type ComplaintCommand struct {
	ID                string `json:"id"`
	UserID            string `json:"user_id"`
	AuthorID          string `json:"author_id"`
	ReceiverID        string `json:"receiver_id"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	Content           string `json:"content"`
	ReplyAuthorID     string `json:"reply_author_id"`
	ReplyBody         string `json:"reply_body"`
	ReplyEnterpriseID string `json:"reply_enterprise_id"`
	Rating            int    `json:"rating"`
	Comment           string `json:"comment"`
	EventID           string `json:"event_id"`
}

/*
Package commands
<< Command >>
@params: context.Context, context.Context, AuthorID,
ReceiverID, Title, Description, Content
@returns: error
*/
func (command ComplaintCommand) SendNew(
	ctx context.Context,
) error {
	if command.AuthorID == "" || command.ReceiverID == "" ||
		command.Title == "" || command.Description == "" || command.Content == "" {
		return ErrBadRequest
	}
	newID := uuid.New()
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*complaint.ComplaintSent); ok {
					NotificationCommand{
						OwnerID:   command.ReceiverID,
						Thumbnail: command.AuthorID,
						Title:     fmt.Sprintf("New Complaint from %s", command.AuthorID),
						Content:   fmt.Sprintf("You have a new complaint from %s", command.AuthorID),
						Link:      fmt.Sprintf("/complaint/%s", newID.String()),
					}.SaveNew(ctx)
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&complaint.ComplaintSent{})
			},
		})
	newComplaint, err := complaint.Send(
		ctx,
		newID,
		command.AuthorID,
		command.ReceiverID,
		command.Title,
		command.Description,
		command.Content,
	)
	if err != nil {
		return err
	}
	err = repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).Save(
		ctx, newComplaint)
	if err != nil {
		return err
	}
	return nil
}

/*
Package commands
<< Command >>
@params: context.Context, context.Context, ReplyAuthorID, ReplyBody, ReplyEnterpriseID
@returns: error
*/
func (command ComplaintCommand) Reply(
	ctx context.Context,
) error {
	if command.ReplyAuthorID == "" ||
		command.ReplyBody == "" ||
		command.ID == "" {
		return ErrBadRequest
	}

	parsedID, err := uuid.Parse(command.ID)
	if err != nil {
		return ErrBadRequest
	}
	mapper := repositories.MapperRegistryInstance().Get("Complaint")
	if mapper == nil {
		return ErrBadRequest
	}
	complaintMapper, ok := mapper.(repositories.ComplaintRepository)
	if !ok {
		return ErrBadRequest
	}
	dbComplaint, err := complaintMapper.Get(ctx, parsedID)

	if err != nil {
		return err
	}
	_, err = dbComplaint.Reply(
		ctx,
		uuid.New(),
		command.ReplyAuthorID,
		command.ReplyBody,
		command.ReplyEnterpriseID,
	)
	if err != nil {
		return err
	}

	infrastructure.PushNotificationInMemoryQueueInstance().QueueNotification(
		infrastructure.Operation{
			ID:        fmt.Sprintf("complaintLastReply:%s", dbComplaint.ID().String()),
			Operation: dbComplaint,
		})
	err = repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).Update(
		ctx, dbComplaint)
	if err != nil {
		return err
	}
	return nil
}

func (command ComplaintCommand) SendForReviewing(
	ctx context.Context,
) error {
	if command.ID == "" || command.UserID == "" {
		return ErrBadRequest
	}
	parsedID, err := uuid.Parse(command.ID)
	if err != nil {
		return ErrBadRequest
	}
	dbComplaint, err := repositories.MapperRegistryInstance().Get(
		"Complaint",
	).(repositories.ComplaintRepository).Get(
		ctx,
		parsedID,
	)
	if err != nil {
		return err
	}
	triggeredBy, err := repositories.MapperRegistryInstance().Get(
		"User",
	).(repositories.UserRepository).Get(
		ctx,
		command.UserID,
	)
	if err != nil {
		return err
	}
	notificationContent := fmt.Sprintf("%s has been asked for you to review %s attention to your complaint %s",
		triggeredBy.FullName(), triggeredBy.Pronoun(), dbComplaint.Message().Title())
	if _, err := mail.ParseAddress(dbComplaint.ReceiverID()); err != nil {
		receiver, err := repositories.MapperRegistryInstance().Get(
			"Enterprise").(repositories.EnterpriseRepository).Get(
			ctx,
			dbComplaint.ReceiverID(),
		)
		if err != nil {
			return err
		}
		notificationContent = fmt.Sprintf("%s has been asked for you to review %s attention to your complaint %s you  made to %s",
			triggeredBy.FullName(), triggeredBy.Pronoun(), dbComplaint.Message().Title(), receiver.Name())
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*complaint.ComplaintSentForReview); ok {
					NotificationCommand{
						OwnerID:   command.AuthorID,
						Thumbnail: command.ReceiverID,
						Title: fmt.Sprintf("%s has been ask you to review %s attention",
							triggeredBy.FullName(), triggeredBy.Pronoun()),
						Content: notificationContent,
						Link:    "/reviews",
					}.SaveNew(ctx)
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&complaint.ComplaintSentForReview{})
			},
		})
	err = dbComplaint.MarkAsReviewable(
		ctx,
		command.UserID,
	)
	if err != nil {
		return err
	}
	infrastructure.PushNotificationInMemoryQueueInstance().QueueNotification(
		infrastructure.Operation{
			ID:        fmt.Sprintf("complaintLastReply:%s", dbComplaint.ID().String()),
			Operation: dbComplaint,
		})
	return repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).Update(
		ctx,
		dbComplaint,
	)
}

func (command ComplaintCommand) RateComplaint(
	ctx context.Context,
) error {
	if command.ID == "" ||
		command.EventID == "" ||
		command.UserID == "" ||
		command.Comment == "" {
		return ErrBadRequest
	}
	parsedID, err := uuid.Parse(command.ID)
	if err != nil {
		return ErrBadRequest
	}
	dbComplaint, err := repositories.MapperRegistryInstance().Get(
		"Complaint",
	).(repositories.ComplaintRepository).Get(
		ctx,
		parsedID,
	)
	if err != nil {
		return err
	}
	parsedEventID, err := uuid.Parse(command.EventID)
	if err != nil {
		return ErrBadRequest
	}
	storedEvent, err := repositories.MapperRegistryInstance().Get("Event").(repositories.EventRepository).Get(
		ctx,
		parsedEventID,
	)
	if err != nil {
		return err
	}
	var e complaint.ComplaintSentForReview
	err = json.Unmarshal(storedEvent.EventBody, &e)
	if err != nil {
		return err
	}
	triggeredBy, err := repositories.MapperRegistryInstance().Get(
		"User",
	).(repositories.UserRepository).Get(
		ctx,
		command.UserID,
	)
	if err != nil {
		return err
	}
	credentials, err := application_services.AuthorizationApplicationServiceInstance().Credentials(ctx)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*complaint.ComplaintRated); ok {
					NotificationCommand{
						OwnerID:   dbComplaint.ReceiverID(),
						Thumbnail: command.ReceiverID,
						Title:     fmt.Sprintf("%s reviewed your attention", credentials.FullName),
						Content: fmt.Sprintf("%s reviewed the attention %s received from %s",
							credentials.FullName, credentials.Pronoun, triggeredBy.FullName()),
						Link: "/reviews",
					}.SaveNew(ctx)
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&complaint.ComplaintRated{})
			},
		})
	err = dbComplaint.Rate(
		ctx,
		command.UserID,
		command.Rating,
		command.Comment,
	)
	if err != nil {
		return err
	}
	return repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).Update(
		ctx,
		dbComplaint,
	)
}
