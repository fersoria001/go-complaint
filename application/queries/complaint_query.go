package queries

import (
	"context"
	"encoding/json"
	"fmt"
	"go-complaint/domain/model/complaint"
	"go-complaint/dto"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/persistence/counters/count_complaints_where"
	"go-complaint/infrastructure/persistence/finders/find_all_complaints"
	"go-complaint/infrastructure/persistence/finders/find_all_enterprises"
	"go-complaint/infrastructure/persistence/finders/find_all_events"
	"go-complaint/infrastructure/persistence/finders/find_all_users"
	"go-complaint/infrastructure/persistence/repositories"
	"net/mail"

	"github.com/google/uuid"
)

type ComplaintQuery struct {
	Term       string `json:"term"`
	ID         string `json:"id"`
	ReceiverID string `json:"receiver_id"`
	AuthorID   string `json:"author_id"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	UserID     string `json:"user_id"`
	Status     string `json:"status"`
}

/*
Package queries
<< Query >>
@params: context.Context, Term
@returns: []dto.ComplaintReceiver, error
*/
func (query ComplaintQuery) FindReceivers(
	ctx context.Context,
) ([]dto.ComplaintReceiver, error) {
	enterpriseReceivers, err := repositories.MapperRegistryInstance().Get(
		"Enterprise",
	).(repositories.EnterpriseRepository).FindAll(
		ctx,
		find_all_enterprises.NewByNameLike(query.Term),
	)
	if err != nil {
		return nil, err
	}
	userReceivers, err := repositories.MapperRegistryInstance().Get(
		"User",
	).(repositories.UserRepository).FindAll(
		ctx,
		find_all_users.NewByFirstNameOrLastNameLike(query.Term),
	)
	if err != nil {
		return nil, err
	}
	result := make([]dto.ComplaintReceiver, 0, enterpriseReceivers.Cardinality()+userReceivers.Cardinality())
	for enterpriseReceiver := range enterpriseReceivers.Iter() {
		result = append(result, dto.ComplaintReceiver{
			ID:        enterpriseReceiver.Name(),
			FullName:  enterpriseReceiver.Name(),
			Thumbnail: enterpriseReceiver.LogoIMG(),
		})
	}
	for userReceiver := range userReceivers.Iter() {
		result = append(result, dto.ComplaintReceiver{
			ID:        userReceiver.Email(),
			FullName:  userReceiver.FullName(),
			Thumbnail: userReceiver.ProfileIMG(),
		})
	}
	return result, nil
}

/*
Package queries
<< Query >>
@params: context.Context, ReceiverID, Limit, Offset
@returns: dto.ComplaintListDTO, error
*/
func (query ComplaintQuery) Inbox(
	ctx context.Context,
) (dto.ComplaintListDTO, error) {
	if query.ReceiverID == "" {
		return dto.ComplaintListDTO{}, ErrBadRequest
	}
	count, err := repositories.MapperRegistryInstance().Get(
		"Complaint",
	).(repositories.ComplaintRepository).Count(
		ctx,
		count_complaints_where.NewReceiverID(query.ReceiverID),
	)
	if err != nil {
		return dto.ComplaintListDTO{}, err
	}
	if count == 0 {
		return dto.ComplaintListDTO{}, nil
	}
	received, err := repositories.MapperRegistryInstance().Get(
		"Complaint",
	).(repositories.ComplaintRepository).FindAll(
		ctx,
		find_all_complaints.NewByReceiverIDWithLimitAndOffset(
			query.ReceiverID,
			query.Limit,
			query.Offset,
		),
	)
	if err != nil {
		return dto.ComplaintListDTO{}, err
	}
	complaints := make([]dto.ComplaintDTO, 0, received.Cardinality())
	for c := range received.Iter() {
		complaints = append(complaints, dto.NewComplaintDTO(c))
		if _, err := mail.ParseAddress(c.AuthorID()); err == nil {
			author, err := repositories.MapperRegistryInstance().Get(
				"User",
			).(repositories.UserRepository).Get(
				ctx,
				c.AuthorID(),
			)
			if err != nil {
				return dto.ComplaintListDTO{}, err
			}
			complaints[len(complaints)-1].AuthorFullName = author.FullName()
		}
		complaints[len(complaints)-1].AuthorFullName = c.AuthorID()
	}
	return dto.ComplaintListDTO{
		Complaints:    complaints,
		Count:         count,
		CurrentLimit:  query.Limit,
		CurrentOffset: query.Offset,
	}, nil
}

/*
Package queries
<< Query >>
@params: context.Context, AuthorID, Limit, Offset
@returns: dto.ComplaintListDTO, error
*/
func (query ComplaintQuery) Sent(
	ctx context.Context,
) (dto.ComplaintListDTO, error) {
	if query.AuthorID == "" {
		return dto.ComplaintListDTO{}, ErrBadRequest
	}
	count, err := repositories.MapperRegistryInstance().Get(
		"Complaint",
	).(repositories.ComplaintRepository).Count(
		ctx,
		count_complaints_where.NewAuthorID(query.AuthorID),
	)
	if err != nil {
		return dto.ComplaintListDTO{}, err
	}
	if count == 0 {
		return dto.ComplaintListDTO{}, nil
	}
	received, err := repositories.MapperRegistryInstance().Get(
		"Complaint",
	).(repositories.ComplaintRepository).FindAll(
		ctx,
		find_all_complaints.NewByAuthorIDWithLimitAndOffset(
			query.AuthorID,
			query.Limit,
			query.Offset,
		),
	)
	if err != nil {
		return dto.ComplaintListDTO{}, err
	}
	complaints := make([]dto.ComplaintDTO, 0, received.Cardinality())
	for c := range received.Iter() {
		complaints = append(complaints, dto.NewComplaintDTO(c))
		if _, err := mail.ParseAddress(c.ReceiverID()); err == nil {
			receiver, err := repositories.MapperRegistryInstance().Get(
				"User",
			).(repositories.UserRepository).Get(
				ctx,
				c.ReceiverID(),
			)
			if err != nil {
				return dto.ComplaintListDTO{}, err
			}
			complaints[len(complaints)-1].ReceiverFullName = receiver.FullName()
		}
		complaints[len(complaints)-1].ReceiverFullName = c.ReceiverID()
	}
	return dto.ComplaintListDTO{
		Complaints:    complaints,
		Count:         count,
		CurrentLimit:  query.Limit,
		CurrentOffset: query.Offset,
	}, nil
}

/*
Package queries
<< Query >>
@params: context.Context, AuthorID, Limit, Offset
@returns: dto.ComplaintListDTO, error
*/
func (query ComplaintQuery) History(
	ctx context.Context,
) (dto.ComplaintListDTO, error) {
	if query.Status == "" || query.UserID == "" {
		return dto.ComplaintListDTO{}, ErrBadRequest
	}
	count, err := repositories.MapperRegistryInstance().Get(
		"Complaint",
	).(repositories.ComplaintRepository).Count(
		ctx,
		count_complaints_where.StatusAndReceiverIDOrAuthorIDAre(complaint.IN_HISTORY.String(),
			query.UserID),
	)
	if err != nil {
		return dto.ComplaintListDTO{}, err
	}
	if count == 0 {
		return dto.ComplaintListDTO{}, nil
	}
	received, err := repositories.MapperRegistryInstance().Get(
		"Complaint",
	).(repositories.ComplaintRepository).FindAll(
		ctx,
		find_all_complaints.ByStatusStatusAndReceiverIDOrAuthorIDWithLimitAndOffset(
			complaint.IN_HISTORY.String(),
			query.UserID,
			query.Limit,
			query.Offset,
		),
	)
	if err != nil {
		return dto.ComplaintListDTO{}, err
	}
	complaints := make([]dto.ComplaintDTO, 0, received.Cardinality())
	for c := range received.Iter() {
		complaints = append(complaints, dto.NewComplaintDTO(c))
		complaints[len(complaints)-1].ReceiverFullName = c.ReceiverID()
		complaints[len(complaints)-1].AuthorFullName = c.AuthorID()
		if _, err := mail.ParseAddress(c.AuthorID()); err == nil {
			author, err := repositories.MapperRegistryInstance().Get(
				"User",
			).(repositories.UserRepository).Get(
				ctx,
				c.AuthorID(),
			)
			if err != nil {
				return dto.ComplaintListDTO{}, err
			}
			complaints[len(complaints)-1].ReceiverFullName = author.FullName()
		}
		if _, err := mail.ParseAddress(c.ReceiverID()); err == nil {
			receiver, err := repositories.MapperRegistryInstance().Get(
				"User",
			).(repositories.UserRepository).Get(
				ctx,
				c.ReceiverID(),
			)
			if err != nil {
				return dto.ComplaintListDTO{}, err
			}
			complaints[len(complaints)-1].ReceiverFullName = receiver.FullName()
		}
	}
	return dto.ComplaintListDTO{
		Complaints:    complaints,
		Count:         count,
		CurrentLimit:  query.Limit,
		CurrentOffset: query.Offset,
	}, nil
}

/*
Package queries
<< Query >>
@params: context.Context, ID
@returns: dto.ComplaintDTO, error
*/
func (query ComplaintQuery) Complaint(
	ctx context.Context,
) (dto.ComplaintDTO, error) {
	if query.ID == "" {
		return dto.ComplaintDTO{}, ErrBadRequest
	}
	parsedID, err := uuid.Parse(query.ID)
	if err != nil {
		return dto.ComplaintDTO{}, ErrBadRequest
	}
	complaint, err := repositories.MapperRegistryInstance().Get(
		"Complaint",
	).(repositories.ComplaintRepository).Get(
		ctx,
		parsedID,
	)
	if err != nil {
		return dto.ComplaintDTO{}, err
	}
	return dto.NewComplaintDTO(*complaint), nil
}

/*
Package queries
<< Query >>
@params: context.Context, ID
@returns: dto.ReplyDTO, error
*/
func (query ComplaintQuery) ComplaintLastReply(
	ctx context.Context,
) (dto.ReplyDTO, error) {
	if query.ID == "" {
		return dto.ReplyDTO{}, ErrBadRequest
	}
	objectID := "complaintLastReply:" + query.ID
	v, ok := infrastructure.InMemoryCacheInstance().Get(objectID)
	if !ok {
		return dto.ReplyDTO{}, fmt.Errorf("complaint not found in cache %v", objectID)
	}
	cachedComplaint, ok := v.(*complaint.Complaint)
	if !ok {
		return dto.ReplyDTO{}, fmt.Errorf("incorrect type of cached object %v", objectID)
	}
	return dto.NewReplyDTO(cachedComplaint.LastReply(), cachedComplaint.Status().String()), nil
}

/* */
func (query ComplaintQuery) PendingComplaintReviews(
	ctx context.Context,
) ([]dto.PendingComplaintReview, error) {
	if query.UserID == "" {
		return nil, ErrBadRequest
	}
	storedEvents, err := repositories.MapperRegistryInstance().Get("StoredEvent").(repositories.EventRepository).FindAll(
		ctx,
		find_all_events.By(),
	)
	if err != nil {
		return nil, err
	}
	waitingForReview := make(map[string]map[string]complaint.ComplaintSentForReview, 0)

	for storedEvent := range storedEvents.Iter() {
		if storedEvent.TypeName == "*complaint.ComplaintSentForReview" {
			var c complaint.ComplaintSentForReview
			err := json.Unmarshal(storedEvent.EventBody, &c)
			if err != nil {
				return nil, err
			}
			if c.AuthorID() == query.UserID {
				waitingForReview[c.ComplaintID().String()] = map[string]complaint.ComplaintSentForReview{
					storedEvent.EventId.String(): c,
				}
			}
		}
	}

	for storedEvent := range storedEvents.Iter() {
		switch storedEvent.TypeName {
		case "*complaint.ComplaintRated":
			var c complaint.ComplaintRated
			err := json.Unmarshal(storedEvent.EventBody, &c)
			if err != nil {
				return nil, err
			}
			delete(waitingForReview, c.ComplaintID().String())
		case "*complaint.ComplaintClosed":
			var c complaint.ComplaintClosed
			err := json.Unmarshal(storedEvent.EventBody, &c)
			if err != nil {
				return nil, err
			}
			delete(waitingForReview, c.ComplaintID().String())
		case "*complaint.ComplaintSentToHistory":
			var c complaint.ComplaintSentToHistory
			err := json.Unmarshal(storedEvent.EventBody, &c)
			if err != nil {
				return nil, err
			}
			delete(waitingForReview, c.ComplaintID().String())
		}
	}

	complaints := make([]dto.PendingComplaintReview, 0, len(waitingForReview))
	for _, c := range waitingForReview {
		for k, v := range c {
			dbComplaint, err := repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).Get(ctx, v.ComplaintID())
			if err != nil {
				return nil, err
			}
			triggeredBy, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, v.TriggeredBy())
			if err != nil {
				return nil, err
			}
			complaints = append(complaints, dto.PendingComplaintReview{
				EventID:     k,
				Complaint:   dto.NewComplaintDTO(*dbComplaint),
				TriggeredBy: dto.NewUser(*triggeredBy),
			})
		}
	}
	return complaints, nil
}
