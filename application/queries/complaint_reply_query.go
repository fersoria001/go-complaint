package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type ComplaintReplyQuery struct {
	ReplyId string `json:"replyId"`
}

func NewComplaintReplyQuery(replyId string) *ComplaintReplyQuery {
	return &ComplaintReplyQuery{
		ReplyId: replyId,
	}
}

func (c ComplaintReplyQuery) Execute(ctx context.Context) (*dto.Reply, error) {
	replyId, err := uuid.Parse(c.ReplyId)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	r := reg.Get("Reply").(repositories.ComplaintRepliesRepository)
	reply, err := r.Get(ctx, replyId)
	if err != nil {
		return nil, err
	}
	return dto.NewReply(*reply), nil
}
