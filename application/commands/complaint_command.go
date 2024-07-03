package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"go-complaint/application/application_services"
	"go-complaint/domain"
	"go-complaint/domain/model/complaint"
	"go-complaint/dto"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/repositories"
	"net/mail"
	"reflect"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type ComplaintCommand struct {
	ID                 string   `json:"id"`
	IDS                []string `json:"ids"`
	UserID             string   `json:"user_id"`
	AuthorID           string   `json:"author_id"`
	ReceiverID         string   `json:"receiver_id"`
	ReceiverFullName   string   `json:"receiver_full_name"`
	ReceiverProfileIMG string   `json:"receiver_profile_img"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Content            string   `json:"content"`
	ReplyAuthorID      string   `json:"reply_author_id"`
	ReplyBody          string   `json:"reply_body"`
	ReplyEnterpriseID  string   `json:"reply_enterprise_id"`
	Rating             int      `json:"rating"`
	Comment            string   `json:"comment"`
	EventID            string   `json:"event_id"`
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
		command.Title == "" || command.Description == "" ||
		command.Content == "" || command.ReceiverFullName == "" ||
		command.ReceiverProfileIMG == "" {
		return ErrBadRequest
	}
	newID := uuid.New()

	var fullName string
	var profileIMG string
	var link string
	if _, err := mail.ParseAddress(command.AuthorID); err != nil {
		author, err := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Get(ctx, command.AuthorID)
		if err != nil {
			return err
		}
		fullName = author.Name()
		profileIMG = author.LogoIMG()

	} else {
		author, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, command.AuthorID)
		if err != nil {
			return err
		}
		fullName = author.FullName()
		profileIMG = author.ProfileIMG()
	}
	if _, err := mail.ParseAddress(command.ReceiverID); err != nil {
		link = fmt.Sprintf("/%s/inbox/%s", command.ReceiverID, newID.String())
	} else {
		link = fmt.Sprintf("/inbox/%s", newID.String())
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*complaint.ComplaintSent); ok {

					NotificationCommand{
						OwnerID:     command.ReceiverID,
						ThumbnailID: command.AuthorID,
						Thumbnail:   profileIMG,
						Title:       fmt.Sprintf("New Complaint from %s", fullName),
						Content:     fmt.Sprintf("You have a new complaint from %s", fullName),
						Link:        link,
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
		fullName,
		profileIMG,
		command.ReceiverID,
		command.ReceiverFullName,
		command.ReceiverProfileIMG,
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
	author, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, command.ReplyAuthorID)
	if err != nil {
		return err
	}
	_, err = dbComplaint.Reply(
		ctx,
		uuid.New(),
		command.ReplyAuthorID,
		author.ProfileIMG(),
		author.FullName(),
		command.ReplyBody,
		command.ReplyEnterpriseID,
	)
	if err != nil {
		return err
	}

	cache.RequestChannel <- cache.Request{
		Key:     fmt.Sprintf("complaintLastReply:%s", dbComplaint.ID().String()),
		Payload: dto.NewReplyDTO(dbComplaint.LastReply(), dbComplaint.Status().String()),
	}
	err = repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).Update(
		ctx, dbComplaint)
	if err != nil {
		return err
	}
	return nil
}

func (command ComplaintCommand) MarkAsSeen(ctx context.Context) error {
	if command.IDS == nil || len(command.IDS) == 0 || command.ID == "" {
		return ErrBadRequest
	}
	parseIds := make([]uuid.UUID, len(command.IDS))
	for _, replyID := range command.IDS {
		parsedReplyID, err := uuid.Parse(replyID)
		if err != nil {
			return ErrBadRequest
		}
		parseIds = append(parseIds, parsedReplyID)
	}
	parsedIds := mapset.NewSet(parseIds...)
	complaintID, err := uuid.Parse(command.ID)
	if err != nil {
		return ErrBadRequest
	}
	r := repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository)
	dbComplaint, err := r.Get(ctx, complaintID)
	if err != nil {
		return err
	}
	_ = dbComplaint.MarkRepliesAsSeen(parsedIds)
	err = r.Update(ctx, dbComplaint)
	if err != nil {
		return err
	}
	replies := dbComplaint.Replies().ToSlice()
	dtos := make([]dto.ReplyDTO, 0, len(replies))
	for _, reply := range replies {
		dtos = append(dtos, dto.NewReplyDTO(reply, dbComplaint.Status().String()))
	}
	cache.RequestChannel <- cache.Request{
		Key:     fmt.Sprintf("complaintLastReply:%s", dbComplaint.ID().String()),
		Payload: dtos,
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
	link := "/reviews"
	if _, err := mail.ParseAddress(dbComplaint.AuthorID()); err != nil {
		notificationContent = fmt.Sprintf("%s has been asked for you to review %s attention to your complaint %s you  made to %s",
			triggeredBy.FullName(), triggeredBy.Pronoun(), dbComplaint.Message().Title(), dbComplaint.AuthorID())
		link = fmt.Sprintf("/%s/reviews", dbComplaint.AuthorID())
	} else {
		notificationContent = fmt.Sprintf("%s has been asked for you to review %s attention to your complaint %s",
			triggeredBy.FullName(), triggeredBy.Pronoun(), dbComplaint.Message().Title())
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*complaint.ComplaintSentForReview); ok {
					NotificationCommand{
						OwnerID:     dbComplaint.AuthorID(),
						ThumbnailID: dbComplaint.ReceiverID(),
						Thumbnail:   dbComplaint.ReceiverProfileIMG(),
						Title: fmt.Sprintf("%s has been ask you to review %s attention",
							triggeredBy.FullName(), triggeredBy.Pronoun()),
						Content: notificationContent,
						Link:    link,
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
		*triggeredBy,
	)
	if err != nil {
		return err
	}
	cache.RequestChannel <- cache.Request{
		Key:     fmt.Sprintf("complaintLastReply:%s", dbComplaint.ID().String()),
		Payload: dto.NewReplyDTO(dbComplaint.LastReply(), complaint.IN_REVIEW.String()),
	}
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
	link := "/reviews"
	notificationContent := fmt.Sprintf("%s reviewed the attention %s received from %s",
		credentials.FullName, credentials.Pronoun, triggeredBy.FullName())
	if _, err := mail.ParseAddress(dbComplaint.ReceiverID()); err != nil {
		notificationContent = fmt.Sprintf("%s reviewed the attention %s received from %s at %s",
			credentials.FullName, credentials.Pronoun, triggeredBy.FullName(), dbComplaint.ReceiverFullName())
		link = fmt.Sprintf("/%s/reviews", dbComplaint.ReceiverFullName())
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*complaint.ComplaintRated); ok {
					err = NotificationCommand{
						OwnerID:     dbComplaint.ReceiverID(),
						ThumbnailID: dbComplaint.ReceiverID(),
						Thumbnail:   dbComplaint.ReceiverProfileIMG(),
						Title:       fmt.Sprintf("%s reviewed your attention", credentials.FullName),
						Content:     notificationContent,
						Link:        link,
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
		triggeredBy.Email(),
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
