package commands

import (
	"context"
	"go-complaint/infrastructure/persistence/repositories"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type MarkRepliesAsReadCommand struct {
	ComplaintId string   `json:"complaintId"`
	RepliesIds  []string `json:"repliesIds"`
}

func NewMarkRepliesAsReadCommand(complaintId string, repliesIds []string) *MarkRepliesAsReadCommand {
	return &MarkRepliesAsReadCommand{
		ComplaintId: complaintId,
		RepliesIds:  repliesIds,
	}
}

func (c MarkRepliesAsReadCommand) Execute(ctx context.Context) error {
	complaintId, err := uuid.Parse(c.ComplaintId)
	if err != nil {
		return err
	}
	repliesIds := mapset.NewSet[uuid.UUID]()
	for i := range c.RepliesIds {
		id, err := uuid.Parse(c.RepliesIds[i])
		if err != nil {
			return err
		}
		repliesIds.Add(id)
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	dbC, err := r.Get(ctx, complaintId)
	if err != nil {
		return err
	}
	_ = dbC.MarkRepliesAsSeen(repliesIds)
	err = r.Update(ctx, dbC)
	if err != nil {
		return err
	}
	return nil
}
