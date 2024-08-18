package find_all_chat_replies

import "github.com/google/uuid"

type ChatID struct {
	query string
	args  []interface{}
}

func ByChatID(chatId uuid.UUID) *ChatID {
	return &ChatID{
		query: string(`
		SELECT 
			ID,
			CHAT_ID,
			SENDER_ID,
			CONTENT,
			SEEN,
			CREATED_AT,
			UPDATED_AT
		FROM chat_replies
		WHERE chat_id = $1
		`),
		args: []interface{}{chatId},
	}
}

func (e *ChatID) Query() string {
	return e.query
}

func (e *ChatID) Args() []interface{} {
	return e.args
}
