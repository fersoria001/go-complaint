package commands

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/feedback"
	"go-complaint/infrastructure/persistence/repositories"
	"log"
	"reflect"

	"github.com/google/uuid"
)

type EndFeedbackCommand struct {
	FeedbackId string `json:"feedbackId"`
	ReviewerId string `json:"reviewerId"`
}

func NewEndFeedbackCommand(feedbackId, reviewerId string) *EndFeedbackCommand {
	return &EndFeedbackCommand{
		FeedbackId: feedbackId,
		ReviewerId: reviewerId,
	}
}

func (c EndFeedbackCommand) Execute(ctx context.Context) error {
	feedbackId, err := uuid.Parse(c.FeedbackId)
	if err != nil {
		return err
	}
	reviewerId, err := uuid.Parse(c.ReviewerId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	feedbackRepository, ok := reg.Get("Feedback").(repositories.FeedbackRepository)
	if !ok {

		return ErrWrongTypeAssertion
	}
	complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	if !ok {

		return ErrWrongTypeAssertion
	}
	f, err := feedbackRepository.Get(ctx, feedbackId)
	if err != nil {
		log.Println("feedback not found")
		return err
	}
	dbc, err := complaintRepository.Get(ctx, f.ComplaintId())
	if err != nil {
		log.Println("complaint not found")
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if e, ok := event.(*feedback.FeedbackDone); ok {
					reg := repositories.MapperRegistryInstance()
					complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
					if !ok {
						return ErrWrongTypeAssertion
					}
					recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
					if !ok {
						return ErrWrongTypeAssertion
					}
					reviewerRecipient, err := recipientRepository.Get(ctx, e.ReviewedId())
					if err != nil {
						log.Println("recipient not found")
						return err
					}
					dbc, err := complaintRepository.Get(ctx, e.ComplaintId())
					if err != nil {
						log.Println("complaint not found event")
						return err
					}
					command := NewSendNotificationCommand(
						e.ReviewedId().String(),
						dbc.Receiver().Id().String(),
						fmt.Sprintf("%s has made a feedback on your attention", reviewerRecipient.SubjectName()),
						fmt.Sprintf("You have received a feedback from %s on your attention", reviewerRecipient.SubjectName()),
						fmt.Sprintf("/%s/feedbacks?id=%s", dbc.Receiver().Id(), e.FeedbackId()),
					)
					err = command.Execute(ctx)
					if err != nil {
						return err
					}
					return nil
				}
				return ErrWrongTypeAssertion
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&feedback.FeedbackDone{})
			},
		})
	err = f.EndFeedback(ctx)
	if err != nil {
		return err
	}
	err = dbc.SendToHistory(ctx, reviewerId)
	if err != nil {
		return err
	}
	err = feedbackRepository.Update(ctx, f)
	if err != nil {
		return err
	}
	err = complaintRepository.Update(ctx, dbc)
	if err != nil {
		return err
	}
	return nil

}
