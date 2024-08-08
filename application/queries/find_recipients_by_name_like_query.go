package queries

import (
	"context"
	"errors"
	"go-complaint/domain/model/recipient"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_all_recipients"
	"go-complaint/infrastructure/persistence/repositories"
	"slices"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type FindRecipientsByNameLikeQuery struct {
	UserId string `json:"userId"`
	Term   string `json:"term"`
}

func NewFindRecipientsByNameLikeQuery(userId, term string) *FindRecipientsByNameLikeQuery {
	return &FindRecipientsByNameLikeQuery{
		UserId: userId,
		Term:   term,
	}
}

func (frbnl FindRecipientsByNameLikeQuery) Execute(ctx context.Context) ([]*dto.Recipient, error) {
	userId, err := uuid.Parse(frbnl.UserId)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	results := make([]*dto.Recipient, 0)
	dbR, err := repository.FindAll(ctx, find_all_recipients.ByNameLike(frbnl.Term))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return results, nil
		}
		return nil, err
	}
	f := slices.DeleteFunc(dbR, func(e *recipient.Recipient) bool {
		return e.Id() == userId
	})
	for _, v := range f {
		results = append(results, dto.NewRecipient(*v))
	}
	return results, nil
}
