package repositories

import (
	"context"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/finders/find_all_chat_replies"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ChatRepository struct {
	db datasource.Schema
}

func NewChatRepository(db datasource.Schema) ChatRepository {
	return ChatRepository{db: db}
}

func (r ChatRepository) Save(
	ctx context.Context,
	chat *enterprise.Chat,
) error {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	insertCommand := `INSERT INTO chats (ID, ENTERPRISE_ID, RECIPIENT_ONE_ID, RECIPIENT_TWO_ID) VALUES ($1, $2, $3, $4)`
	var (
		id             uuid.UUID = chat.Id()
		enterpriseId   uuid.UUID = chat.EnterpriseId()
		recipientOneId uuid.UUID = chat.RecipientOne().Id()
		recipientTwoId uuid.UUID = chat.RecipientTwo().Id()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&id,
		&enterpriseId,
		&recipientOneId,
		&recipientTwoId,
	)
	if err != nil {
		return err
	}
	err = MapperRegistryInstance().Get("enterprise.Reply").(ChatRepliesRepository).SaveAll(
		ctx,
		chat.Replies(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (r ChatRepository) Update(
	ctx context.Context,
	chat *enterprise.Chat,
) error {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return err
	}
	err = MapperRegistryInstance().Get("enterprise.Reply").(ChatRepliesRepository).DeleteAll(
		ctx,
		chat.Id(),
	)
	if err != nil {
		return err
	}
	err = MapperRegistryInstance().Get("enterprise.Reply").(ChatRepliesRepository).SaveAll(
		ctx,
		chat.Replies(),
	)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (r ChatRepository) Find(
	ctx context.Context,
	src StatementSource,
) (*enterprise.Chat, error) {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	row := conn.QueryRow(
		ctx,
		src.Query(),
		src.Args()...,
	)
	chat, err := r.load(ctx, row)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (r ChatRepository) Get(
	ctx context.Context,
	chatId uuid.UUID,
) (*enterprise.Chat, error) {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	selectQuery := string(`
		SELECT ID,
		ENTERPRISE_ID,
		RECIPIENT_ONE_ID,
		RECIPIENT_TWO_ID
		FROM chats
			WHERE id = $1
	`)
	row := conn.QueryRow(
		ctx,
		selectQuery,
		&chatId,
	)
	chat, err := r.load(ctx, row)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	return chat, nil
}

func (r ChatRepository) FindAll(
	ctx context.Context,
	source StatementSource,
) ([]*enterprise.Chat, error) {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	conn.Release()
	rows, err := conn.Query(ctx, source.Query(), source.Args()...)
	if err != nil {
		return nil, err
	}
	chats, err := r.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	rows.Close()
	return chats, nil
}

func (r ChatRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) ([]*enterprise.Chat, error) {
	var chats []*enterprise.Chat
	for rows.Next() {
		chat, err := r.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	return chats, nil
}

func (r ChatRepository) load(
	ctx context.Context,
	row pgx.Row,
) (*enterprise.Chat, error) {
	var (
		id             uuid.UUID
		enterpriseId   uuid.UUID
		recipientOneId uuid.UUID
		recipientTwoId uuid.UUID
	)
	err := row.Scan(&id, &enterpriseId, &recipientOneId, &recipientTwoId)
	if err != nil {
		return nil, err
	}
	reg := MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(RecipientRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	repliesRepository, ok := reg.Get("enterprise.Reply").(ChatRepliesRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}

	recipientOne, err := recipientRepository.Get(ctx, recipientOneId)
	if err != nil {
		return nil, err
	}
	recipientTwo, err := recipientRepository.Get(ctx, recipientTwoId)
	if err != nil {
		return nil, err
	}
	replies, err := repliesRepository.FindAll(
		ctx,
		find_all_chat_replies.ByChatID(id),
	)
	if err != nil {
		return nil, err
	}
	return enterprise.NewChat(
		id,
		enterpriseId,
		*recipientOne,
		*recipientTwo,
		replies,
	), nil
}
