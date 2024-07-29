package complaint

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/recipient"
	"go-complaint/erros"
	"slices"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

/*
*
<<Entity>>
<<Value Object>> Status 1
<<Value Object>> Message 1
<<Value Object>> Rating 1
<<Value Object>> ReceiverID 1
Relationships:
- Complaint *..1 Author trough AuthorID
- Complaint *..1 Status
- Complaint *..1 Rating
- Complaint *..1 ReceiverID trough ReceiverID
- Complaint 1..* Reply  trough ID
*/
type Complaint struct {
	id          uuid.UUID
	author      recipient.Recipient
	receiver    recipient.Recipient
	title       string
	description string
	status      Status
	rating      Rating
	createdAt   common.Date
	updatedAt   common.Date
	replies     mapset.Set[*Reply]
}

/*
validated by the requester ID
if its an email then its an user and it must match the author
-users can only close their own complaints-
if its an employee ID then it must match the enterprise
-employees can close sent or received complaints from their enterprise-
*/
func (c *Complaint) SendToHistory(
	ctx context.Context,
	closeRequesterID uuid.UUID,
) error {
	if c.status != CLOSED {
		return &erros.ValidationError{
			Expected: "a complaint must be closed to be sent to history",
		}
	}
	err := c.setStatus(IN_HISTORY)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Publish(
		ctx,
		NewComplaintSentToHistory(
			c.id,
			closeRequesterID,
		),
	)
	return nil
}

func (complaintt *Complaint) Rate(
	ctx context.Context,
	triggeredBy uuid.UUID,
	rate int,
	comment string,
) error {
	if complaintt.Status() != IN_REVIEW {
		return &ValidationError{Message: "a complaint must be in review or closed to be rated"}
	}
	rating, err := NewRating(complaintt.id, rate, comment)
	if err != nil {
		return err
	}
	complaintt.rating = rating
	err = complaintt.setStatus(CLOSED)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Publish(
		ctx,
		NewComplaintClosed(
			complaintt.id,
			complaintt.author.Id(),
			triggeredBy,
		),
	)
	if !complaintt.receiver.IsEnterprise() {
		err = complaintt.setStatus(IN_HISTORY)
		if err != nil {
			return err
		}
		domain.DomainEventPublisherInstance().Publish(
			ctx,
			NewComplaintSentToHistory(
				complaintt.id,
				triggeredBy,
			),
		)
	}
	lastReply := complaintt.LastReply()
	domain.DomainEventPublisherInstance().Publish(
		ctx,
		NewComplaintRated(
			complaintt.Id(),
			triggeredBy,
			lastReply.sender.Id(),
			time.Now(),
		),
	)
	return nil
}

/*
If the complaint is already in review status return an error
Set the complaint status to IN_REVIEW
Publish a new event of type WaitingForReview
*/
func (complaint *Complaint) MarkAsReviewable(
	ctx context.Context,
	triggeredById uuid.UUID,
) error {
	if complaint.status >= IN_REVIEW {
		return ErrComplaintClosed
	}
	err := complaint.setStatus(IN_REVIEW)
	if err != nil {
		return err
	}
	return domain.DomainEventPublisherInstance().Publish(
		ctx,
		NewComplaintSentForReview(
			complaint.id,
			complaint.receiver.Id(),
			complaint.author.Id(),
			triggeredById,
		),
	)
}

func (c *Complaint) Reply(
	ctx context.Context,
	newReplyId uuid.UUID,
	author recipient.Recipient,
	body string,
) error {
	if c.status > IN_DISCUSSION {
		return ErrComplaintClosed
	}
	thisTime := time.Now()
	publisher := domain.DomainEventPublisherInstance()
	newReply := CreateReply(
		newReplyId,
		c.id,
		author,
		body,
	)
	switch c.replies.Cardinality() {
	case 1:
		err := c.setStatus(STARTED)
		if err != nil {
			return err
		}
		err = publisher.Publish(ctx, NewComplaintStarted(c.id, newReplyId, thisTime))
		if err != nil {
			return err
		}
	case 2:
		err := c.setStatus(IN_DISCUSSION)
		if err != nil {
			return err
		}
		err = publisher.Publish(ctx, NewDiscussionStarted(c.id, newReplyId, thisTime))
		if err != nil {
			return err
		}
	}
	added := c.replies.Add(newReply)
	if !added {
		return ErrReplyNotAdded
	}
	err := publisher.Publish(ctx, NewComplaintReplied(c.id, newReplyId, thisTime))
	if err != nil {
		return err
	}
	return nil
}

/*
Pre:
  - title must not be empty and must have more than 10 characters and less than 80 characters
  - description must not be empty and must have more than 30 characters and less than 120 characters
  - body must not be empty and must have more than 50 characters and less than 250 characters
*/
func (c *Complaint) Send(ctx context.Context) error {
	if c.status != WRITING {
		return &ValidationError{Message: "the complaint has already sent"}
	}
	if c.title == "" {
		return &ValidationError{Message: "title can't be null"}
	}
	if c.description == "" {
		return &ValidationError{Message: "description can't be null"}
	}
	if c.Body() == "" {
		return &ValidationError{Message: "body can't be null"}
	}
	c.status = OPEN
	publisher := domain.DomainEventPublisherInstance()
	event := NewComplaintSent(c.id, c.author.Id(), c.receiver.Id(), c.updatedAt.Date())
	err := publisher.Publish(ctx, event)
	if err != nil {
		return err
	}
	return nil
}

/*
Pre:
  - body must not be empty and must have more than 50 characters and less than 250 characters
*/
func (c *Complaint) SetBody(ctx context.Context, body string) error {
	err := c.setBody(body)
	if err != nil {
		return err
	}
	c.updatedAt = common.NewDate(time.Now())
	return domain.DomainEventPublisherInstance().Publish(
		ctx,
		NewComplaintBodySet(
			c.id,
			body,
		),
	)
}

func (c *Complaint) setBody(body string) error {
	if len(body) <= 50 {
		return &ValidationError{Message: "body has to be more than 50 characters long"}
	}
	if len(body) > 250 {
		return &ValidationError{Message: "body has to be less than 251 characters long"}
	}
	createdAt := common.NewDate(time.Now())
	firstReply, err := NewReply(
		c.id,
		c.id,
		c.author,
		body,
		false,
		createdAt,
		createdAt,
		createdAt,
	)
	if err != nil {
		return err
	}
	c.AddReply(firstReply)
	return nil
}

/*
Pre:
- description must not be empty and must have more than 30 characters and less than 120 characters
*/
func (c *Complaint) SetDescription(ctx context.Context, description string) error {
	err := c.setDescription(description)
	if err != nil {
		return err
	}
	c.updatedAt = common.NewDate(time.Now())
	return domain.DomainEventPublisherInstance().Publish(
		ctx,
		NewComplaintDescriptionSet(c.id, description),
	)
}

func (c *Complaint) setDescription(description string) error {
	if len(description) <= 30 {
		return &ValidationError{Message: "description has to be more than 30 characters long"}
	}
	if len(description) > 120 {
		return &ValidationError{Message: "description has to be less than 121 characters long"}
	}
	c.description = description
	return nil
}

/*
Pre:
- title must not be empty and must have more than 10 characters and less than 80 characters
*/
func (c *Complaint) SetTitle(ctx context.Context, title string) error {
	err := c.setTitle(title)
	if err != nil {
		return err
	}
	c.updatedAt = common.NewDate(time.Now())
	return domain.DomainEventPublisherInstance().Publish(
		ctx,
		NewComplaintTitleSet(c.id, title),
	)
}

func (c *Complaint) setTitle(title string) error {
	if len(title) <= 10 {
		return &ValidationError{Message: "title has to be more than 10 characters length"}
	}
	if len(title) > 80 {
		return &ValidationError{Message: "title has to be less than 81 characters length"}
	}
	c.title = title
	return nil
}

func CreateNew(
	ctx context.Context,
	id uuid.UUID,
	author,
	receiver recipient.Recipient,
) (*Complaint, error) {
	t := time.Now()
	c := &Complaint{
		id:        id,
		author:    author,
		receiver:  receiver,
		status:    WRITING,
		createdAt: common.NewDate(t),
		updatedAt: common.NewDate(t),
		replies:   mapset.NewSet[*Reply](),
	}
	p := domain.DomainEventPublisherInstance()
	err := p.Publish(
		ctx,
		NewComplaintCreated(
			id,
			author.Id(),
			author.SubjectName(),
			author.SubjectThumbnail(),
			author.IsEnterprise(),
			receiver.Id(),
			receiver.SubjectName(),
			receiver.SubjectThumbnail(),
			receiver.IsEnterprise(),
			WRITING,
			t,
		),
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewComplaint(
	id uuid.UUID,
	author,
	receiver recipient.Recipient,
	status Status,
	title,
	description string,
	createdAt,
	updatedAt common.Date,
	rating Rating,
	replies mapset.Set[*Reply],
) (*Complaint, error) {
	var c *Complaint = new(Complaint)
	err := c.setID(id)
	if err != nil {
		return nil, err
	}
	err = c.setStatus(status)
	if err != nil {
		return nil, err
	}
	c.title = title
	c.description = description
	c.author = author
	c.receiver = receiver
	c.setRating(rating)
	err = c.setCreatedAt(createdAt)
	if err != nil {
		return nil, err
	}
	err = c.setUpdatedAt(updatedAt)
	if err != nil {
		return nil, err
	}
	err = c.setReplies(replies)
	if err != nil {
		return nil, err
	}
	c.replies = replies
	return c, nil
}

func (c *Complaint) AddReply(reply *Reply) {
	ok := c.replies.Add(reply)
	if ok {
		c.updatedAt = common.NewDate(time.Now())
	}
}

// nullable
func (c *Complaint) setRating(rating Rating) {
	if rating == (Rating{}) {
		rating = Rating{
			rate:    0,
			comment: "",
		}
	}
	c.rating = rating
}

func (c *Complaint) setID(id uuid.UUID) error {
	if id == uuid.Nil {
		return &erros.NullValueError{}
	}
	c.id = id
	return nil
}

func (c *Complaint) setStatus(status Status) error {
	if status < 0 || status > 5 {
		return &erros.ValidationError{
			Expected: "a value between 0 and 5",
		}
	}
	c.status = status
	return nil
}

func (c *Complaint) setCreatedAt(createdAt common.Date) error {
	if createdAt == (common.Date{}) {
		return &erros.EmptyStructError{}
	}
	c.createdAt = createdAt
	return nil
}

func (c *Complaint) setUpdatedAt(updatedAt common.Date) error {
	if updatedAt == (common.Date{}) {
		return &erros.EmptyStructError{}
	}
	c.updatedAt = updatedAt
	return nil
}

func (c *Complaint) setReplies(replies mapset.Set[*Reply]) error {
	if replies == nil {
		return &erros.NullValueError{}
	}
	c.replies = replies
	return nil
}

func (c Complaint) Id() uuid.UUID {
	return c.id
}

func (c Complaint) Author() recipient.Recipient {
	return c.author
}

func (c Complaint) Receiver() recipient.Recipient {
	return c.receiver
}

func (c Complaint) Status() Status {
	return c.status
}

func (c Complaint) Title() string {
	return c.title
}

func (c Complaint) Description() string {
	return c.description
}

func (c Complaint) Rating() Rating {
	return c.rating
}

func (c Complaint) CreatedAt() common.Date {
	return c.createdAt
}

func (c Complaint) UpdatedAt() common.Date {
	return c.updatedAt
}

func (c Complaint) Replies() mapset.Set[Reply] {
	valueCopy := mapset.NewSet[Reply]()
	for reply := range c.replies.Iter() {
		valueCopy.Add(*reply)
	}
	return valueCopy
}

func (c *Complaint) MarkRepliesAsSeen(parsedIds mapset.Set[uuid.UUID]) int {
	replies := c.replies.ToSlice()
	count := 0
	for _, reply := range replies {
		if parsedIds.Contains(reply.ID()) {
			reply.MarkAsRead()
			count++
		}
	}
	c.replies = mapset.NewSet(replies...)
	return count
}

func (c Complaint) RepliesDifference(replies mapset.Set[*Reply]) mapset.Set[Reply] {
	valueCopy := mapset.NewSet[Reply]()
	difference := c.replies.Difference(replies)
	for reply := range difference.Iter() {
		valueCopy.Add(*reply)
	}
	return valueCopy
}

func (c Complaint) LastReply() Reply {
	var lastReply Reply
	if c.replies.Cardinality() == 0 {
		return lastReply
	}
	sliceCopy := c.replies.ToSlice()
	slices.SortStableFunc(sliceCopy, func(i, j *Reply) int {
		if i.createdAt.Date().Before(j.createdAt.Date()) {
			return -1
		}
		if i.createdAt.Date().After(j.createdAt.Date()) {
			return 1
		}
		return 0
	})
	return *sliceCopy[len(sliceCopy)-1]
}

func (c Complaint) Body() string {
	if c.replies.Cardinality() < 1 {
		return ""
	}
	s := c.replies.ToSlice()
	slices.SortStableFunc(s, func(i, j *Reply) int {
		if i.createdAt.Date().Before(j.createdAt.Date()) {
			return -1
		}
		if i.createdAt.Date().After(j.createdAt.Date()) {
			return 1
		}
		return 0
	})
	return s[0].body
}
