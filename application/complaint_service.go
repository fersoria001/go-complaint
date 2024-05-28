package application

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/complaint"
	"go-complaint/dto"
	"go-complaint/erros"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"

	mapset "github.com/deckarep/golang-set/v2"
)

type ComplaintService struct {
	complaintRepository *repositories.ComplaintRepository
	repliesRepository   *repositories.ReplyRepository
}

func NewComplaintService(
	complaintRepository *repositories.ComplaintRepository,
	repliesRepository *repositories.ReplyRepository) *ComplaintService {
	return &ComplaintService{
		complaintRepository: complaintRepository,
		repliesRepository:   repliesRepository,
	}
}

// It must be closed first, how it ends up closed?
// func (cs *ComplaintService) ProvideComplaintAsResource() {}
func (cs *ComplaintService) RateComplaint(
	ctx context.Context,
	authorID,
	id string,
	rating int,
	comment string,
) error {
	var (
		complaint *complaint.Complaint
		err       error
	)
	complaint, err = cs.complaintRepository.Get(ctx, id)
	if err != nil {
		return err
	}
	if complaint.AuthorID() != authorID {
		return &erros.UnauthorizedError{}
	}
	err = complaint.Rate(ctx, authorID, rating, comment)
	if err != nil {
		return err
	}
	err = cs.complaintRepository.Update(ctx, complaint)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ComplaintService) Close(
	ctx context.Context,
	id string,
	closeRequesterID string) error {
	var (
		complaint *complaint.Complaint
		err       error
	)
	complaint, err = cs.complaintRepository.Get(ctx, id)
	if err != nil {
		return err
	}
	err = complaint.Close(ctx, closeRequesterID)
	if err != nil {
		return err
	}
	return cs.complaintRepository.Update(ctx, complaint)
}

func (cs *ComplaintService) SendForReviewing(
	ctx context.Context,
	complaintID,
	assistantID string) error {
	var (
		complaint *complaint.Complaint
		err       error
	)
	complaint, err = cs.complaintRepository.Get(ctx, complaintID)
	if err != nil {
		return err
	}
	err = complaint.MarkAsReviewable(
		ctx,
		assistantID,
	)
	if err != nil {
		return err
	}
	return cs.complaintRepository.Update(ctx, complaint)
}

func (cs *ComplaintService) ReplyFromBroadCast(ctx context.Context, id string, msg dto.Message) error {
	return cs.ReplyComplaint(ctx,
		id,
		msg.Reply.SenderIMG,
		msg.Reply.SenderName,
		msg.Reply.SenderID,
		msg.Reply.Body)
}

// Reply
func (cs *ComplaintService) ReplyComplaint(
	ctx context.Context,
	id string,
	profileIMG string,
	name string,
	senderID string,
	body string,
) error {
	var (
		complaintt *complaint.Complaint
		err        error
	)
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*complaint.ComplaintReplied); ok {
				fmt.Println("Reply Complaint Handler on" + event.OccurredOn().String())
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&complaint.ComplaintReplied{})
		},
	})

	complaintt, err = cs.complaintRepository.Get(ctx, id)
	if err != nil {
		return err
	}
	count, err := cs.repliesRepository.Count(ctx, id)
	if err != nil {
		return err
	}
	reply, err := complaintt.ReplyComplaint(ctx, count, senderID, profileIMG, name, body)
	if err != nil {
		return err
	}
	err = cs.repliesRepository.Save(ctx, reply)
	if err != nil {
		return err
	}
	err = cs.complaintRepository.Update(ctx, complaintt)
	if err != nil {
		return err
	}
	return nil
}

// HUGE DOMAIN ISSUE
func (cs *ComplaintService) ProvideDomainComplaintAndReplies(ctx context.Context, id string) (*complaint.Complaint, []*complaint.Reply, error) {
	complaint, err := cs.complaintRepository.Get(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	replies, err := cs.repliesRepository.FindByComplaintID(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	return complaint, replies, nil
}

func (cs *ComplaintService) Complaint(ctx context.Context, id string) (*dto.ComplaintDTO, error) {
	complaint, err := cs.complaintRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	replies, err := cs.repliesRepository.FindByComplaintID(ctx, id)
	if err != nil {
		return nil, err
	}
	complaintDTO := dto.NewComplaintDTO(complaint, replies)
	return complaintDTO, nil
}

func (cs *ComplaintService) GetComplaintsTo(ctx context.Context, receiverID string, status string, limit, offset int) (dto.ComplaintListDTO, error) {
	var (
		complaintList = dto.ComplaintListDTO{}
		complaints    mapset.Set[*complaint.Complaint]
		parsedStatus  complaint.Status = -1
		count         int
		err           error
	)
	if status != "" {
		parsedStatus, err = complaint.ParseStatus(status)
		if err != nil {
			return complaintList, err
		}
	}
	complaints, count, err = cs.complaintRepository.FindByReceiver(ctx, receiverID, limit, offset)
	if err != nil {
		return complaintList, err
	}
	dtoSlice := []*dto.ComplaintDTO{}
	//filter
	for _, c := range complaints.ToSlice() {
		if parsedStatus == -1 && (c.Status() == complaint.IN_HISTORY || c.Status() == complaint.CLOSED) {
			complaints.Remove(c)
		}
	}
	if complaints.Cardinality() > 0 {
		for c := range complaints.Iter() {
			complaintt := dto.NewComplaintDTO(c, []*complaint.Reply{})
			dtoSlice = append(dtoSlice, complaintt)
		}
		complaintList.Complaints = dtoSlice
		complaintList.Count = count
		complaintList.CurrentLimit = limit
		complaintList.CurrentOffset = offset
	} else {
		return complaintList, &erros.ValueNotFoundError{}
	}
	return complaintList, nil
}

// without replies
func (cs *ComplaintService) GetComplaintsFrom(
	ctx context.Context,
	authorID string,
	status string,
	limit,
	offset int,
) (dto.ComplaintListDTO, error) {
	var (
		complaintList = dto.ComplaintListDTO{}
		complaints    mapset.Set[*complaint.Complaint]
		parsedStatus  complaint.Status = -1
		count         int
		err           error
	)
	if status != "" {
		parsedStatus, err = complaint.ParseStatus(status)
		if err != nil {
			return complaintList, err
		}
	}
	complaints, count, err = cs.complaintRepository.FindByAuthor(ctx, authorID, limit, offset)
	if err != nil {

		return complaintList, err
	}
	complaintSlic := []*dto.ComplaintDTO{}
	for _, c := range complaints.ToSlice() {
		if parsedStatus > -1 && c.Status() != parsedStatus {
			complaints.Remove(c)
		}
	}
	if complaints.Cardinality() > 0 {
		for c := range complaints.Iter() {
			complaintDTO := dto.NewComplaintDTO(c, []*complaint.Reply{})
			complaintSlic = append(complaintSlic, complaintDTO)
		}
		complaintList.Complaints = complaintSlic
		complaintList.Count = count
		complaintList.CurrentLimit = limit
		complaintList.CurrentOffset = offset
	} else {
		return complaintList, &erros.ValueNotFoundError{}
	}
	return complaintList, nil
}

// PRE: authorID and receiverID are valid ids from user(email) OR enterprise(name)
func (cs *ComplaintService) CreateComplaint(
	ctx context.Context,
	authorID,
	receiverID,
	title,
	description,
	content string) error {
	var (
		newComplaint *complaint.Complaint
		err          error
	)
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*complaint.ComplaintSent); ok {
					fmt.Println("Complaint Created Handler", event.OccurredOn().Format("2006-01-02 15:04:05"))
					return nil
				}
				return &erros.ValueNotFoundError{}
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&complaint.ComplaintSent{})
			},
		})

	newComplaint, err = complaint.SendComplaint(
		ctx,
		authorID,
		receiverID,
		title,
		description,
		content)
	if err != nil {
		return err
	}
	err = cs.complaintRepository.Save(ctx, newComplaint)
	if err != nil {
		return err
	}
	return nil
}
