package repositories

import (
	"context"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/finders/find_all_chat_replies"

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
	insertCommand := `INSERT INTO chat (ID) VALUES ($1)`
	var id string = chat.ID().String()
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&id,
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
		chat.ID(),
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

func (r ChatRepository) Get(
	ctx context.Context,
	chatID enterprise.ChatID,
) (*enterprise.Chat, error) {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	one := "%" + chatID.String() + "%"
	two := "%" + chatID.Reverse().String() + "%"
	selectQuery := `
		SELECT id
		FROM chat
			WHERE id LIKE $1 
 		OR id LIKE $2
	`
	row := conn.QueryRow(
		ctx,
		selectQuery,
		&one,
		&two,
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
	rows, err := conn.Query(ctx, source.Query(), source.Args()...)
	if err != nil {
		return nil, err
	}
	chats, err := r.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
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
	var id string
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}
	chatID, err := enterprise.NewChatID(id)
	if err != nil {
		return nil, err
	}
	replies, err := MapperRegistryInstance().Get("enterprise.Reply").(ChatRepliesRepository).FindAll(
		ctx,
		find_all_chat_replies.ByChatID(*chatID),
	)
	if err != nil {
		return nil, err
	}
	return enterprise.NewChat(
		*chatID,
		replies,
	), nil
}
