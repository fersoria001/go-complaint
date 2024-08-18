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

type ComplaintsRatedByAuthorIdQuery struct {
	AuthorId string `json:"authorId"`
	Term     string `json:"term"`
}

func NewComplaintsRatedByAuthorIdQuery(authorId, term string) *ComplaintsRatedByAuthorIdQuery {
	return &ComplaintsRatedByAuthorIdQuery{
		AuthorId: authorId,
		Term:     term,
	}
}

func (q ComplaintsRatedByAuthorIdQuery) Execute(ctx context.Context) ([]*dto.Complaint, error) {
	authorId, err := uuid.Parse(q.AuthorId)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	result := make([]*dto.Complaint, 0)
	c, err := r.FindAll(ctx, find_all_complaints.ByAuthorAndStatusIn(authorId, []string{complaint.CLOSED.String(), complaint.IN_HISTORY.String()}))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return result, nil
		}
		return nil, err
	}
	if q.Term != "" {
		trie := trie.NewTrie()
		for i := range c {
			trie.InsertText(c[i].Id().String(), c[i].Title(), " ")
			trie.InsertText(c[i].Id().String(), c[i].Description(), " ")
			trie.InsertText(c[i].Id().String(), c[i].Body(), " ")
			trie.InsertText(c[i].Id().String(), c[i].Author().SubjectName(), " ")
			trie.InsertText(c[i].Id().String(), c[i].Receiver().SubjectName(), " ")
			trie.InsertText(c[i].Id().String(), c[i].Rating().Comment(), " ")
			trie.InsertText(c[i].Id().String(), c[i].Rating().LastUpdate().UTC().Format(time.UnixDate), " ")
		}
		searchResults := trie.Search(q.Term)
		c = slices.DeleteFunc(c, func(e *complaint.Complaint) bool {
			return !searchResults.Contains(e.Id().String())
		})
	}
	for _, v := range c {
		result = append(result, dto.NewComplaint(*v))
	}
	return result, nil
}
