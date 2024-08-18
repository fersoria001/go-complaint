package repositories

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/datasource"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ChatRepliesRepository struct {
	db datasource.Schema
}

func NewChatRepliesRepository(db datasource.Schema) ChatRepliesRepository {
	return ChatRepliesRepository{db: db}
}

func (r ChatRepliesRepository) UpdateAll(
	ctx context.Context,
	replies []*enterprise.Reply,
) error {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return err
	}
	updateCommand := string(`UPDATE chat_replies
	SET
		CONTENT = $1,
		SEEN = $2,
		UPDATED_AT = $3
	WHERE ID = $4`)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for _, reply := range replies {
		var (
			content   string    = reply.Content()
			seen      bool      = reply.Seen()
			updatedAt string    = common.StringDate(reply.UpdatedAt())
			id        uuid.UUID = reply.Id()
		)
		_, err = tx.Exec(
			ctx,
			updateCommand,
			&content,
			&seen,
			&updatedAt,
			&id,
		)
		if err != nil {
			return err
		}
	}
	defer conn.Release()
	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	return nil
}

func (r ChatRepliesRepository) DeleteAll(
	ctx context.Context,
	chatId uuid.UUID,
) error {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return err
	}
	deleteCommand := string(`DELETE FROM chat_replies WHERE CHAT_ID = $1`)
	_, err = conn.Exec(ctx, deleteCommand, &chatId)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (r ChatRepliesRepository) SaveAll(
	ctx context.Context,
	replies []*enterprise.Reply,
) error {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return err
	}
	insertCommand := string(`
		INSERT INTO chat_replies
			(ID, CHAT_ID, SENDER_ID, CONTENT, SEEN, CREATED_AT, UPDATED_AT)
			VALUES
			($1, $2, $3, $4, $5, $6, $7)`)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for _, reply := range replies {
		var (
			id        uuid.UUID = reply.Id()
			chatID    uuid.UUID = reply.ChatId()
			userID    uuid.UUID = reply.Sender().Id()
			content   string    = reply.Content()
			seen      bool      = reply.Seen()
			createdAt string    = common.StringDate(reply.CreatedAt())
			updatedAt string    = common.StringDate(reply.UpdatedAt())
		)
		_, err = tx.Exec(
			ctx,
			insertCommand,
			&id,
			&chatID,
			&userID,
			&content,
			&seen,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	defer conn.Release()
	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	return nil
}
func (r ChatRepliesRepository) Update(
	ctx context.Context,
	reply *enterprise.Reply,
) error {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return err
	}
	updateCommand := string(`UPDATE chat_replies
	SET
		CONTENT = $1,
		SEEN = $2,
		UPDATED_AT = $3
	WHERE ID = $4`)
	var (
		content   string    = reply.Content()
		seen      bool      = reply.Seen()
		updatedAt string    = common.StringDate(reply.UpdatedAt())
		id        uuid.UUID = reply.Id()
	)
	_, err = conn.Exec(
		ctx,
		updateCommand,
		&content,
		&seen,
		&updatedAt,
		&id,
	)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (r ChatRepliesRepository) Save(
	ctx context.Context,
	reply *enterprise.Reply,
) error {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return err
	}
	insertCommand := string(`
	INSERT INTO chat_replies
		(ID, CHAT_ID, SENDER_ID, CONTENT, SEEN, CREATED_AT, UPDATED_AT)
		VALUES
		($1, $2, $3, $4, $5, $6, $7)`)
	var (
		id        uuid.UUID = reply.Id()
		chatID    uuid.UUID = reply.ChatId()
		userID    uuid.UUID = reply.Sender().Id()
		content   string    = reply.Content()
		seen      bool      = reply.Seen()
		createdAt string    = common.StringDate(reply.CreatedAt())
		updatedAt string    = common.StringDate(reply.UpdatedAt())
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&id,
		&chatID,
		&userID,
		&content,
		&seen,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (r ChatRepliesRepository) Get(
	ctx context.Context,
	id string,
) (*enterprise.Reply, error) {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	selectQuery := string(`
	SELECT 
		ID,
		CHAT_ID,
		SENDER_ID,
		CONTENT,
		SEEN,
		CREATED_AT,
		UPDATED_AT
	FROM chat_replies
	WHERE ID = $1
	`)
	row := conn.QueryRow(ctx, selectQuery, &id)
	reply, err := r.load(ctx, row)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (r ChatRepliesRepository) FindAll(
	ctx context.Context,
	source StatementSource,
) ([]*enterprise.Reply, error) {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, source.Query(), source.Args()...)
	if err != nil {
		return nil, err
	}
	result, err := r.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	rows.Close()
	return result, nil
}

func (r ChatRepliesRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) ([]*enterprise.Reply, error) {
	var replies []*enterprise.Reply
	for rows.Next() {
		reply, err := r.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		replies = append(replies, reply)
	}
	return replies, nil
}

func (r ChatRepliesRepository) load(ctx context.Context, row pgx.Row) (*enterprise.Reply, error) {
	var (
		id        uuid.UUID
		chatID    uuid.UUID
		userId    uuid.UUID
		content   string
		seen      bool
		createdAt string
		updatedAt string
	)
	err := row.Scan(
		&id,
		&chatID,
		&userId,
		&content,
		&seen,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}
	reg := MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(RecipientRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	user, err := recipientRepository.Get(ctx, userId)
	if err != nil {
		return nil, err
	}
	createdAtTime, err := common.ParseDate(createdAt)
	if err != nil {
		return nil, err
	}
	updatedAtTime, err := common.ParseDate(updatedAt)
	if err != nil {
		return nil, err
	}
	return enterprise.NewReply(
		id,
		chatID,
		*user,
		content,
		seen,
		createdAtTime,
		updatedAtTime,
	), nil
}
