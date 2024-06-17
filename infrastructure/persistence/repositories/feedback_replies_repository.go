package repositories

import (
	"context"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/persistence/datasource"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type FeedbackRepliesRepository struct {
	schema datasource.Schema
}

func NewFeedbackRepliesRepository(
	feedbackSchema datasource.Schema,
) FeedbackRepliesRepository {
	return FeedbackRepliesRepository{schema: feedbackSchema}
}

func (fr FeedbackRepliesRepository) Save(
	ctx context.Context,
	replyReviewID uuid.UUID,
	replies mapset.Set[complaint.Reply],
) error {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return nil
	}
	insertCommand := string(`
		INSERT INTO feedback_replies
		(
			id,
			reply_review_id,
			reply_id
		)
		VALUES ($1, $2, $3)
		`)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for reply := range replies.Iter() {
		var (
			feedbackReplyID = uuid.New()
			replyReviewID   = replyReviewID
			replyID         = reply.ID()
		)
		_, err = tx.Exec(
			ctx,
			insertCommand,
			&feedbackReplyID,
			&replyReviewID,
			&replyID,
		)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (fr FeedbackRepliesRepository) FindAll(
	ctx context.Context,
	source StatementSource,
) (mapset.Set[*complaint.Reply], error) {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(ctx, source.Query(), source.Args()...)
	if err != nil {
		return nil, err
	}
	replies, err := fr.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return replies, nil
}

func (fr FeedbackRepliesRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[*complaint.Reply], error) {
	var replies = mapset.NewSet[*complaint.Reply]()
	for rows.Next() {
		reply, err := fr.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		replies.Add(reply)
	}
	return replies, nil
}

func (fr FeedbackRepliesRepository) load(
	ctx context.Context,
	row pgx.Row,
) (*complaint.Reply, error) {
	mapper := MapperRegistryInstance().Get("Reply")
	if mapper == nil {
		return nil, ErrMapperNotRegistered
	}
	replyRepository, ok := mapper.(ComplaintRepliesRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	var (
		ID            uuid.UUID
		ReplyReviewId uuid.UUID
		ReplyID       uuid.UUID
	)
	err := row.Scan(&ID, &ReplyReviewId, &ReplyID)
	if err != nil {
		return nil, err
	}
	reply, err := replyRepository.Get(ctx, ReplyID)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
