package queries

import (
	"context"
	"errors"
	"go-complaint/domain/model/complaint"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_all_complaints"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/infrastructure/trie"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PendingReviewsByAuthorIdQuery struct {
	AuthorId string `json:"authorId"`
	Term     string `json:"term"`
}

func NewPendingReviewsByAuthorIdQuery(authorId, term string) *PendingReviewsByAuthorIdQuery {
	return &PendingReviewsByAuthorIdQuery{
		AuthorId: authorId,
		Term:     term,
	}
}

func (q PendingReviewsByAuthorIdQuery) Execute(ctx context.Context) ([]*dto.Complaint, error) {
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	authorId, err := uuid.Parse(q.AuthorId)
	if err != nil {
		return nil, err
	}
	result := make([]*dto.Complaint, 0)
	p, err := r.FindAll(ctx, find_all_complaints.ByAuthorAndStatusIn(authorId, []string{
		complaint.IN_REVIEW.String(),
	}))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return result, nil
		}
		return nil, err
	}
	if q.Term != "" {
		trie := trie.NewTrie()
		for i := range p {
			trie.InsertText(p[i].Id().String(), p[i].Title(), " ")
			trie.InsertText(p[i].Id().String(), p[i].Description(), " ")
			trie.InsertText(p[i].Id().String(), p[i].Body(), " ")
			trie.InsertText(p[i].Id().String(), p[i].Author().SubjectName(), " ")
			trie.InsertText(p[i].Id().String(), p[i].Receiver().SubjectName(), " ")
			trie.InsertText(p[i].Id().String(), p[i].Rating().Comment(), " ")
			trie.InsertText(p[i].Id().String(), p[i].Rating().LastUpdate().UTC().Format(time.UnixDate), " ")
		}
		searchResults := trie.Search(q.Term)
		p = slices.DeleteFunc(p, func(e *complaint.Complaint) bool {
			return !searchResults.Contains(e.Id().String())
		})
	}
	for _, v := range p {
		result = append(result, dto.NewComplaint(*v))
	}
	return result, nil
}
