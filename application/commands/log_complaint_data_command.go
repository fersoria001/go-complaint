package commands

import (
	"context"
	"fmt"
	"go-complaint/application"
	"go-complaint/domain/model/complaint"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"
	"time"

	"github.com/google/uuid"
)

type LogComplaintDataCommand struct {
	OwnerId     string `json:"ownerId"`
	AuthorId    string `json:"authorId"`
	ReceiverId  string `json:"receiverId"`
	ComplaintId string `json:"complaintId"`
	DataType    string `json:"dataType"`
}

func NewLogComplaintDataCommand(ownerId, authorId, receiverId, complaintId, dataType string) *LogComplaintDataCommand {
	return &LogComplaintDataCommand{
		OwnerId:     ownerId,
		AuthorId:    authorId,
		ReceiverId:  receiverId,
		ComplaintId: complaintId,
		DataType:    dataType,
	}
}

func (c LogComplaintDataCommand) Execute(ctx context.Context) error {
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("ComplaintData").(repositories.ComplaintDataRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	ownerId, err := uuid.Parse(c.OwnerId)
	if err != nil {
		return err
	}
	authorId, err := uuid.Parse(c.AuthorId)
	if err != nil {
		return err
	}
	receiverId, err := uuid.Parse(c.ReceiverId)
	if err != nil {
		return err
	}
	complaintId, err := uuid.Parse(c.ComplaintId)
	if err != nil {
		return err
	}
	dataType := complaint.ParseComplaintDataType(c.DataType)
	if dataType < 0 {
		return fmt.Errorf("unexpected data type at log complaint data command")
	}
	complaintData := complaint.NewComplaintData(uuid.New(),
		ownerId, authorId, receiverId, complaintId, time.Now(), dataType)
	err = r.Save(ctx, *complaintData)
	if err != nil {
		return err
	}
	svc := application.ApplicationMessagePublisherInstance()
	svc.Publish(application.NewApplicationMessage(
		c.OwnerId,
		"complaintData",
		*dto.NewComplaintData(*complaintData),
	))
	return nil
}
