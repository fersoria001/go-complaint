package complaint

import (
	"go-complaint/erros"
)


type Message struct {
	title       string
	description string
	body        string
}


func NewMessage(title string, description string, body string) (Message, error) {
	var m Message = Message{}
	err := m.setTitle(title)
	if err != nil {
		return m, err
	}
	err = m.setDescription(description)

	if err != nil {
		return m, err
	}
	err = m.setBody(body)
	if err != nil {
		return m, err
	}
	return m, nil
}

func (m *Message) setTitle(title string) error {
	if title == "" {
		return &erros.NullValueError{}
	}
	if len(title) < 10 {
		return &erros.ValidationError{

			Expected: "more than 10 characters",
		}
	}
	if len(title) > 80 {
		return &erros.ValidationError{

			Expected: "less than 80 characters",
		}
	}
	m.title = title
	return nil
}

func (m *Message) setDescription(description string) error {
	if description == "" {
		return &erros.NullValueError{}
	}
	if len(description) < 30 {
		return &erros.InvalidLengthError{
			AttributeName: "description",
			MinLength:     30,
			MaxLength:     120,
			CurrentLength: len(description),
		}
	}
	if len(description) > 120 {
		return &erros.InvalidLengthError{
			AttributeName: "description",
			MinLength:     30,
			MaxLength:     120,
			CurrentLength: len(description),
		}
	}
	m.description = description
	return nil
}

func (m *Message) setBody(body string) error {
	if body == "" {
		return &erros.NullValueError{}
	}
	if len(body) < 50 {
		return &erros.InvalidLengthError{
			AttributeName: "body",
			MinLength:     50,
			MaxLength:     250,
			CurrentLength: len(body),
		}
	}
	if len(body) > 250 {
		return &erros.InvalidLengthError{
			AttributeName: "body",
			MinLength:     50,
			MaxLength:     250,
			CurrentLength: len(body),
		}
	}
	m.body = body
	return nil
}

func (m Message) Title() string {
	return m.title
}

func (m Message) Description() string {
	return m.description
}

func (m Message) Body() string {
	return m.body
}
