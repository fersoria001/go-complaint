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
	updateCommand := string(`UPDATE chat_reply
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
			id        uuid.UUID = reply.ID()
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
	chatID enterprise.ChatID,
) error {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return err
	}
	one := "%" + chatID.String() + "%"
	two := "%" + chatID.Reverse().String() + "%"
	deleteCommand := string(`DELETE FROM chat_reply WHERE CHAT_ID LIKE $1 OR CHAT_ID LIKE $2`)
	_, err = conn.Exec(ctx, deleteCommand, &one, &two)
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
		INSERT INTO chat_reply
			(ID, CHAT_ID, USER_ID, CONTENT, SEEN, CREATED_AT, UPDATED_AT)
			VALUES
			($1, $2, $3, $4, $5, $6, $7)`)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for _, reply := range replies {
		var (
			id        uuid.UUID = reply.ID()
			chatID    string    = reply.ChatID().String()
			userID    string    = reply.User().Email()
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
	updateCommand := string(`UPDATE chat_reply
	SET
		CONTENT = $1,
		SEEN = $2,
		UPDATED_AT = $3
	WHERE ID = $4`)
	var (
		content   string    = reply.Content()
		seen      bool      = reply.Seen()
		updatedAt string    = common.StringDate(reply.UpdatedAt())
		id        uuid.UUID = reply.ID()
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
	INSERT INTO chat_reply
		(ID, CHAT_ID, USER_ID, CONTENT, SEEN, CREATED_AT, UPDATED_AT)
		VALUES
		($1, $2, $3, $4, $5, $6, $7)`)
	var (
		id        uuid.UUID = reply.ID()
		chatID    string    = reply.ChatID().String()
		userID    string    = reply.User().Email()
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
		USER_ID,
		CONTENT,
		SEEN,
		CREATED_AT,
		UPDATED_AT
	FROM chat_reply
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
	rows, err := conn.Query(ctx, source.Query(), source.Args()...)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return r.loadAll(ctx, rows)
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
		chatID    string
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
	user, err := MapperRegistryInstance().Get("User").(UserRepository).Get(ctx, userId)
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
	parsedChatID, err := enterprise.NewChatID(chatID)
	if err != nil {
		return nil, err
	}
	return enterprise.NewReply(
		id,
		*parsedChatID,
		*user,
		content,
		seen,
		createdAtTime,
		updatedAtTime,
	), nil
}
