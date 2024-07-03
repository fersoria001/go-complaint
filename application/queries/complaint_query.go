package queries

import (
	"context"
	"encoding/json"
	"fmt"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/dto"
	"math"

	"go-complaint/infrastructure/persistence/counters/count_complaints"
	"go-complaint/infrastructure/persistence/finders/find_all_complaints"
	"go-complaint/infrastructure/persistence/finders/find_all_enterprises"
	"go-complaint/infrastructure/persistence/finders/find_all_events"
	"go-complaint/infrastructure/persistence/finders/find_all_users"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/infrastructure/trie"
	"net/mail"
	"slices"

	"github.com/google/uuid"
)

type ComplaintQuery struct {
	Term         string `json:"term"`
	ID           string `json:"id"`
	ReceiverID   string `json:"receiver_id"`
	ReceiverName string `json:"receiver_name"`
	AuthorID     string `json:"author_id"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
	UserID       string `json:"user_id"`
	Status       string `json:"status"`
	AfterDate    string `json:"after_date"`
	BeforeDate   string `json:"before_date"`
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
	if query.ID == "" {
		return nil, ErrBadRequest
	}
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
		find_all_users.ByFullNameLike(query.Term),
	)
	if err != nil {
		return nil, err
	}
	result := make([]dto.ComplaintReceiver, 0, enterpriseReceivers.Cardinality()+userReceivers.Cardinality())
	for enterpriseReceiver := range enterpriseReceivers.Iter() {
		if enterpriseReceiver.Name() != query.ID {
			result = append(result, dto.ComplaintReceiver{
				ID:        enterpriseReceiver.Name(),
				FullName:  enterpriseReceiver.Name(),
				Thumbnail: enterpriseReceiver.LogoIMG(),
			})
		}
	}
	for userReceiver := range userReceivers.Iter() {
		if userReceiver.Email() != query.ID {
			result = append(result, dto.ComplaintReceiver{
				ID:        userReceiver.Email(),
				FullName:  userReceiver.FullName(),
				Thumbnail: userReceiver.ProfileIMG(),
			})
		}
	}
	return result, nil
}

func (query ComplaintQuery) FindAuthor(
	ctx context.Context,
) (dto.ComplaintReceiver, error) {
	if query.AuthorID == "" {
		return dto.ComplaintReceiver{}, ErrBadRequest
	}
	enterpriseReceiver, err := repositories.MapperRegistryInstance().Get(
		"Enterprise",
	).(repositories.EnterpriseRepository).Get(ctx, query.AuthorID)
	if err != nil {
		userReceiver, err := repositories.MapperRegistryInstance().Get(
			"User",
		).(repositories.UserRepository).Get(ctx, query.AuthorID)
		if err != nil {
			return dto.ComplaintReceiver{}, err
		}
		return dto.ComplaintReceiver{
			ID:        userReceiver.Email(),
			FullName:  userReceiver.FullName(),
			Thumbnail: userReceiver.ProfileIMG(),
		}, nil
	}
	return dto.ComplaintReceiver{
		ID:        enterpriseReceiver.Name(),
		FullName:  enterpriseReceiver.Name(),
		Thumbnail: enterpriseReceiver.LogoIMG(),
	}, nil
}

func (query ComplaintQuery) IsValidComplaintReceiver(
	ctx context.Context,
) bool {
	if query.ReceiverID == "" {
		return false
	}
	_, err := repositories.MapperRegistryInstance().Get(
		"Enterprise",
	).(repositories.EnterpriseRepository).Get(
		ctx,
		query.ReceiverID,
	)
	if err != nil {
		_, err = repositories.MapperRegistryInstance().Get(
			"User",
		).(repositories.UserRepository).Get(
			ctx,
			query.ReceiverID,
		)
		if err != nil {
			return false
		}
	}
	return true
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
		count_complaints.WhereReceiverID(query.ReceiverID),
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
	}
	return dto.ComplaintListDTO{
		Complaints:    complaints,
		Count:         count,
		CurrentLimit:  query.Limit,
		CurrentOffset: query.Offset,
	}, nil
}

func (query ComplaintQuery) InboxSearch(
	ctx context.Context,
) (dto.ComplaintListDTO, error) {
	if query.ReceiverID == "" {
		return dto.ComplaintListDTO{}, ErrBadRequest
	}
	received, err := repositories.MapperRegistryInstance().Get(
		"Complaint",
	).(repositories.ComplaintRepository).FindAll(
		ctx,
		find_all_complaints.ByReceiverAndStatusIn(
			query.ReceiverID,
			[]string{
				complaint.OPEN.String(),
				complaint.STARTED.String(),
				complaint.IN_DISCUSSION.String(),
			},
		),
	)
	if err != nil {
		return dto.ComplaintListDTO{}, err
	}

	if query.AfterDate != "" && query.BeforeDate != "" {
		after, err := common.NewDateFromString(query.AfterDate)
		if err != nil {
			return dto.ComplaintListDTO{}, err
		}
		before, err := common.NewDateFromString(query.BeforeDate)
		if err != nil {
			return dto.ComplaintListDTO{}, err
		}
		for c := range received.Iter() {
			if !c.CreatedAt().Date().After(after.Date()) && !c.CreatedAt().Date().Before(before.Date()) {
				received.Remove(c)
			}
		}
	}

	if query.Term != "" {
		trie := trie.NewTrie()
		for c := range received.Iter() {
			trie.InsertText(c.ID().String(), c.AuthorFullName(), " ")
			trie.InsertText(c.ID().String(), c.Message().Title(), " ")
			trie.InsertText(c.ID().String(), c.Message().Description(), " ")
			trie.InsertText(c.ID().String(), c.Message().Body(), " ")
			trie.InsertText(c.ID().String(), c.Status().String(), "_")
		}
		ids := trie.Search(query.Term)
		if ids == nil {
			return dto.ComplaintListDTO{
				Complaints:    []dto.ComplaintDTO{},
				Count:         0,
				CurrentLimit:  query.Limit,
				CurrentOffset: query.Offset,
			}, nil
		}
		slice := received.ToSlice()
		for c := range slice {
			if !ids.Contains(slice[c].ID().String()) {
				received.Remove(slice[c])
			}
		}
	}

	count := received.Cardinality()
	complaints := make([]dto.ComplaintDTO, 0, count)

	receivedSlice := received.ToSlice()

	slices.SortStableFunc(receivedSlice, func(i, j complaint.Complaint) int {
		if i.CreatedAt().Date().After(j.CreatedAt().Date()) {
			return -1
		}
		if i.CreatedAt().Date().Before(j.CreatedAt().Date()) {
			return 1
		}
		return 0
	})

	for _, c := range receivedSlice {
		complaints = append(complaints, dto.NewComplaintDTO(c))
	}
	if query.Limit != 0 {
		offset := query.Offset
		limit := query.Limit
		length := len(complaints)
		offsetLimit := offset + limit
		//offset: 0 | < len | > len
		//limit: 10 | < len | > len
		if offset > length {
			return dto.ComplaintListDTO{}, fmt.Errorf("offset is greater than the length of the list")
		}
		if offset+limit > length {
			offsetLimit = offset + (length - offset)
		}
		complaints = complaints[offset:offsetLimit]
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
		count_complaints.WhereAuthorID(query.AuthorID),
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

func (query ComplaintQuery) SentSearch(
	ctx context.Context,
) (dto.ComplaintListDTO, error) {
	if query.AuthorID == "" {
		return dto.ComplaintListDTO{}, ErrBadRequest
	}
	received, err := repositories.MapperRegistryInstance().Get(
		"Complaint",
	).(repositories.ComplaintRepository).FindAll(
		ctx,
		find_all_complaints.ByAuthorAndStatusIn(
			query.AuthorID,
			[]string{
				complaint.OPEN.String(),
				complaint.STARTED.String(),
				complaint.IN_DISCUSSION.String(),
			},
		),
	)
	if err != nil {
		return dto.ComplaintListDTO{}, err
	}

	if query.AfterDate != "" && query.BeforeDate != "" {

		after, err := common.NewDateFromString(query.AfterDate)
		if err != nil {
			return dto.ComplaintListDTO{}, err
		}
		before, err := common.NewDateFromString(query.BeforeDate)
		if err != nil {
			return dto.ComplaintListDTO{}, err
		}
		for c := range received.Iter() {
			if !c.CreatedAt().Date().After(after.Date()) && !c.CreatedAt().Date().Before(before.Date()) {
				received.Remove(c)
			}
		}
	}

	if query.Term != "" {

		trie := trie.NewTrie()
		for c := range received.Iter() {
			trie.InsertText(c.ID().String(), c.ReceiverFullName(), " ")
			trie.InsertText(c.ID().String(), c.Message().Title(), " ")
			trie.InsertText(c.ID().String(), c.Message().Description(), " ")
			trie.InsertText(c.ID().String(), c.Message().Body(), " ")
			trie.InsertText(c.ID().String(), c.Status().String(), "_")
		}
		ids := trie.Search(query.Term)
		if ids == nil {
			return dto.ComplaintListDTO{
				Complaints:    []dto.ComplaintDTO{},
				Count:         0,
				CurrentLimit:  query.Limit,
				CurrentOffset: query.Offset,
			}, nil
		}
		slice := received.ToSlice()
		for c := range slice {
			if !ids.Contains(slice[c].ID().String()) {
				received.Remove(slice[c])
			}
		}
	}

	count := received.Cardinality()
	complaints := make([]dto.ComplaintDTO, 0, count)

	receivedSlice := received.ToSlice()

	slices.SortStableFunc(receivedSlice, func(i, j complaint.Complaint) int {
		if i.CreatedAt().Date().After(j.CreatedAt().Date()) {
			return -1
		}
		if i.CreatedAt().Date().Before(j.CreatedAt().Date()) {
			return 1
		}
		return 0
	})

	for _, c := range receivedSlice {
		complaints = append(complaints, dto.NewComplaintDTO(c))
	}
	if query.Limit != 0 {
		offset := query.Offset
		limit := query.Limit
		length := len(complaints)
		offsetLimit := offset + limit
		//offset: 0 | < len | > len
		//limit: 10 | < len | > len
		if offset > length {
			return dto.ComplaintListDTO{}, fmt.Errorf("offset is greater than the length of the list")
		}
		if offset+limit > length {
			offsetLimit = offset + (length - offset)
		}
		complaints = complaints[offset:offsetLimit]
	}
	return dto.ComplaintListDTO{
		Complaints:    complaints,
		Count:         count,
		CurrentLimit:  query.Limit,
		CurrentOffset: query.Offset,
	}, nil
}

func (query ComplaintQuery) History(
	ctx context.Context,
) (dto.ComplaintListDTO, error) {
	if query.ReceiverID == "" {
		return dto.ComplaintListDTO{}, ErrBadRequest
	}
	received, err := repositories.MapperRegistryInstance().Get(
		"Complaint",
	).(repositories.ComplaintRepository).FindAll(
		ctx,
		find_all_complaints.ByReceiverAndStatus(
			query.ReceiverID,
			complaint.IN_HISTORY.String(),
		),
	)
	if err != nil {
		return dto.ComplaintListDTO{}, err
	}

	if query.AfterDate != "" && query.BeforeDate != "" {
		after, err := common.NewDateFromString(query.AfterDate)
		if err != nil {
			return dto.ComplaintListDTO{}, err
		}
		before, err := common.NewDateFromString(query.BeforeDate)
		if err != nil {
			return dto.ComplaintListDTO{}, err
		}
		for c := range received.Iter() {
			if !c.CreatedAt().Date().After(after.Date()) && !c.CreatedAt().Date().Before(before.Date()) {
				received.Remove(c)
			}
		}
	}

	if query.Term != "" {
		trie := trie.NewTrie()
		for c := range received.Iter() {
			trie.InsertText(c.ID().String(), c.AuthorFullName(), " ")
			trie.InsertText(c.ID().String(), c.Message().Title(), " ")
			trie.InsertText(c.ID().String(), c.Message().Description(), " ")
			trie.InsertText(c.ID().String(), c.Message().Body(), " ")
			trie.InsertText(c.ID().String(), c.Status().String(), "_")
		}
		ids := trie.Search(query.Term)
		if ids == nil {
			return dto.ComplaintListDTO{
				Complaints:    []dto.ComplaintDTO{},
				Count:         0,
				CurrentLimit:  query.Limit,
				CurrentOffset: query.Offset,
			}, nil
		}
		slice := received.ToSlice()
		for c := range slice {
			if !ids.Contains(slice[c].ID().String()) {
				received.Remove(slice[c])
			}
		}
	}

	count := received.Cardinality()
	complaints := make([]dto.ComplaintDTO, 0, count)

	receivedSlice := received.ToSlice()

	slices.SortStableFunc(receivedSlice, func(i, j complaint.Complaint) int {
		if i.CreatedAt().Date().After(j.CreatedAt().Date()) {
			return -1
		}
		if i.CreatedAt().Date().Before(j.CreatedAt().Date()) {
			return 1
		}
		return 0
	})

	for _, c := range receivedSlice {
		complaints = append(complaints, dto.NewComplaintDTO(c))
	}
	offset := query.Offset
	limit := query.Limit
	length := len(complaints)
	offsetLimit := offset + limit
	//offset: 0 | < len | > len
	//limit: 10 | < len | > len
	if offset > length {
		return dto.ComplaintListDTO{}, fmt.Errorf("offset is greater than the length of the list")
	}
	if offset+limit > length {
		offsetLimit = offset + (length - offset)
	}
	complaints = complaints[offset:offsetLimit]
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

/* */
func (query ComplaintQuery) PendingComplaintReviews(
	ctx context.Context,
) ([]dto.PendingComplaintReview, error) {
	if query.UserID == "" {
		return nil, ErrBadRequest
	}
	storedEvents, err := repositories.MapperRegistryInstance().Get("Event").(repositories.EventRepository).FindAll(
		ctx,
		find_all_events.By(),
	)
	if err != nil {

		return nil, err
	}
	waitingForReview := make(map[string]*dto.PendingComplaintReview, 0)

	for storedEvent := range storedEvents.Iter() {
		if storedEvent.TypeName == "*complaint.ComplaintSentForReview" {
			var c complaint.ComplaintSentForReview
			err := json.Unmarshal(storedEvent.EventBody, &c)
			if err != nil {
				return nil, err
			}
			if c.AuthorID() == query.UserID || c.TriggeredBy() == query.UserID {
				triggeredBy, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, c.TriggeredBy())
				if err != nil {

					return nil, err
				}
				dbComplaint, err := repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).Get(ctx,
					c.ComplaintID())
				if err != nil {

					return nil, err
				}
				waitingForReview[c.ComplaintID().String()] = &dto.PendingComplaintReview{
					EventID:     storedEvent.EventId.String(),
					TriggeredBy: dto.NewUser(*triggeredBy),
					Status:      dto.PENDING.String(),
					Complaint:   dto.NewComplaintDTO(*dbComplaint),
					OccurredOn:  common.StringDate(c.OccurredOn()),
				}
			}
			if c.ReceiverID() == query.UserID {
				triggeredBy, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, c.TriggeredBy())
				if err != nil {

					return nil, err
				}
				dbComplaint, err := repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).Get(ctx,
					c.ComplaintID())
				if err != nil {

					return nil, err
				}
				waitingForReview[c.ComplaintID().String()] = &dto.PendingComplaintReview{
					EventID:     storedEvent.EventId.String(),
					TriggeredBy: dto.NewUser(*triggeredBy),
					Status:      dto.WAITING.String(),
					Complaint:   dto.NewComplaintDTO(*dbComplaint),
					OccurredOn:  common.StringDate(c.OccurredOn()),
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
			if v, ok := waitingForReview[c.ComplaintID().String()]; ok {
				ratedBy, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, c.RatedBy())
				if err != nil {

					return nil, err
				}
				v.SetStatus(dto.RATED.String())
				v.SetRatedBy(dto.NewUser(*ratedBy))
				v.SetOccurredOn(common.StringDate(c.OccurredOn()))
			}
		}
	}

	result := make([]dto.PendingComplaintReview, 0, len(waitingForReview))
	for _, v := range waitingForReview {
		result = append(result, *v)
	}
	slices.SortStableFunc(result, func(i, j dto.PendingComplaintReview) int {
		dateI, _ := common.ParseDate(i.OccurredOn)
		dateJ, _ := common.ParseDate(j.OccurredOn)
		if dateI.After(dateJ) {
			return -1
		}
		if dateI.Before(dateJ) {
			return 1
		}
		return 0
	})
	return result, nil
}

func (query ComplaintQuery) ComplaintsReceivedInfo(
	ctx context.Context,
) (dto.ComplaintInfo, error) {
	if query.ID == "" {
		return dto.ComplaintInfo{}, ErrBadRequest
	}
	received, err := repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).Count(
		ctx,
		count_complaints.WhereReceiverID(
			query.ID,
		),
	)
	if err != nil {
		return dto.ComplaintInfo{}, err
	}
	resolved, err := repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).Count(
		ctx,
		count_complaints.WhereReceiverIDAndStatusIN(
			query.ID,
			[]string{
				complaint.IN_REVIEW.String(),
				complaint.CLOSED.String(),
				complaint.IN_HISTORY.String(),
			},
		),
	)
	if err != nil {
		return dto.ComplaintInfo{}, err
	}
	reviewed, err := repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).Count(
		ctx,
		count_complaints.WhereReceiverIDAndStatusIN(
			query.ID,
			[]string{
				complaint.CLOSED.String(),
				complaint.IN_HISTORY.String(),
			},
		),
	)
	if err != nil {
		return dto.ComplaintInfo{}, err
	}

	cPending, err := repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).Count(
		ctx,
		count_complaints.WhereReceiverIDAndStatusIN(
			query.ID,
			[]string{
				complaint.OPEN.String(),
				complaint.STARTED.String(),
				complaint.IN_DISCUSSION.String(),
			},
		),
	)
	if err != nil {
		return dto.ComplaintInfo{}, err
	}
	cs, err := repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).FindAll(
		ctx,
		find_all_complaints.ByReceiver(query.ID),
	)
	if err != nil {
		return dto.ComplaintInfo{}, err
	}
	sum := 0
	for c := range cs.Iter() {
		sum += c.Rating().Rate()
	}
	average := float64(sum) / float64(reviewed)
	if math.IsNaN(average) {
		average = 0
	}
	return dto.ComplaintInfo{
		ComplaintsReceived: received,
		ComplaintsResolved: resolved,
		ComplaintsReviewed: reviewed,
		ComplaintsPending:  cPending,
		AverageRating:      average,
	}, nil
}
