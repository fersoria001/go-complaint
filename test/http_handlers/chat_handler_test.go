package http_handlers_test

import (
	"context"
	"encoding/json"
	"go-complaint/application/application_services"
	"go-complaint/application/commands"
	"go-complaint/application/queries"
	"go-complaint/domain"
	"go-complaint/domain/model/complaint"
	"go-complaint/dto"
	"go-complaint/http_handlers"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/finders/find_recipient"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

var complaintIds = []uuid.UUID{}
var authTokens = []string{}

func TestChatHandler_Setup(t *testing.T) {
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "5170")
	os.Setenv("ALLOWED_ORIGINS", "http://localhost:3000,localhost:3000,localhost")
	os.Setenv("CSRF_KEY", "ultrasecret")
	os.Setenv("DATABASE_URL", "postgres://postgres:sfdkwtf@localhost:5432/postgres?pool_max_conns=100&search_path=public&connect_timeout=5")
	os.Setenv("PORT", "5170")
	os.Setenv("DNS", "http://localhost:3000")
	os.Setenv("SEND_GRID_API_KEY", "Bearer mlsn.0557f4217143328c73149ad91c7455121924f188c63af0fe093b42feb3fa1de1")
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	userRepository, ok := reg.Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	recipientsRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewUsers {
		domain.DomainEventPublisherInstance().Reset()
		c := commands.NewRegisterUserCommand(
			v.UserName,
			v.Password,
			v.Person.FirstName,
			v.Person.LastName,
			v.Person.Genre,
			v.Person.Pronoun,
			v.Person.BirthDate.StringRepresentation(),
			v.Person.Phone,
			v.Person.ProfileImg,
			v.Person.Address.Country().ID(),
			v.Person.Address.CountryState().ID(),
			v.Person.Address.City().ID(),
		)
		err := c.Execute(ctx)
		assert.Nil(t, err)
	}
	for _, v := range mock_data.NewRegisterEnterprises {
		domain.DomainEventPublisherInstance().Reset()
		owner, err := userRepository.Find(ctx, find_user.ByUsername(mock_data.NewUsers["valid"].UserName))
		assert.Nil(t, err)
		c := commands.NewRegisterEnterpriseCommand(
			owner.Id().String(),
			v.Name,
			v.LogoImg,
			v.BannerImg,
			v.Website,
			v.Email,
			v.Phone,
			v.FoundationDate.StringRepresentation(),
			v.Industry.Id,
			v.Address.Country().ID(),
			v.Address.CountryState().ID(),
			v.Address.City().ID(),
		)
		err = c.Execute(ctx)
		assert.Nil(t, err)
	}

	for _, v := range mock_data.NewUsers {
		q := queries.NewSignInQuery(v.UserName, v.Password, false)
		token, err := q.Execute(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, token.Token)
		_, ok := cache.InMemoryInstance().Get(token.Token)
		assert.True(t, ok)
		q1 := queries.NewLoginQuery(token.Token, 9999999)
		loginToken, err := q1.Execute(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, loginToken.Token)
		_, ok = cache.InMemoryInstance().Get(token.Token)
		assert.False(t, ok)
		svc := application_services.AuthorizationApplicationServiceInstance()
		_, err = svc.Authorize(ctx, loginToken.Token)
		assert.Nil(t, err)
		authTokens = append(authTokens, loginToken.Token)
	}

	for _, v := range mock_data.NewComplaints {
		author, err := recipientsRepository.Find(ctx, find_recipient.ByNameAndEmail(v.Author.SubjectName, v.Author.SubjectEmail))
		assert.Nil(t, err)
		receiver, err := recipientsRepository.Find(ctx, find_recipient.ByNameAndEmail(v.Receiver.SubjectName, v.Receiver.SubjectEmail))
		assert.Nil(t, err)
		var complaintId uuid.UUID
		domain.DomainEventPublisherInstance().Subscribe(
			domain.DomainEventSubscriber{
				HandleEvent: func(event domain.DomainEvent) error {
					if ev, ok := event.(*complaint.ComplaintCreated); ok {
						complaintId = ev.Id()
					}
					return nil
				},
				SubscribedToEventType: func() reflect.Type {
					return reflect.TypeOf(&complaint.ComplaintCreated{})
				},
			},
		)
		c := commands.NewCreateNewComplaintCommand(author.Id().String(), receiver.Id().String())
		err = c.Execute(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, complaintId)
		complaintIds = append(complaintIds, complaintId)
		c1 := commands.NewDescribeComplaintCommand(complaintId.String(), v.Title, v.Description)
		err = c1.Execute(ctx)
		assert.Nil(t, err)
		assert.Greater(t, len(v.Replies), 0)
		mockBody := v.Replies[0].Body
		c2 := commands.NewSendComplaintCommand(complaintId.String(), author.Id().String(), mockBody)
		err = c2.Execute(ctx)
		assert.Nil(t, err)
	}
}

func TestChatHandler_ServeWS(t *testing.T) {
	TestChatHandler_Setup(t)
	s := httptest.NewServer(http.HandlerFunc(http_handlers.ServeWS))
	defer s.Close()

	assert.GreaterOrEqual(t, len(complaintIds), 1)
	id := complaintIds[len(complaintIds)-1]
	// Convert http://127.0.0.1 to ws://127.0.0.1
	u := "ws" + strings.TrimPrefix(s.URL, "http") + "?id=" + id.String()

	d := websocket.DefaultDialer
	d.Subprotocols = []string{"complaint"}
	// Connect to the server
	ws, _, err := d.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()
	assert.GreaterOrEqual(t, len(authTokens), 1)
	authToken := authTokens[len(authTokens)-1]
	tokenMsg := map[string]string{"token": authToken}
	authTokenB, err := json.Marshal(tokenMsg)
	assert.Nil(t, err)
	if err := ws.WriteMessage(websocket.TextMessage, authTokenB); err != nil {
		t.Fatalf("%v", err)
	}
	_, p, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("%v", err)
	}
	var expected map[string]bool
	if err := json.Unmarshal(p, &expected); err != nil {
		t.Fatalf("bad message")
	}
	assert.NotNil(t, expected["authenticated"])
	if !expected["authenticated"] {
		t.Fatalf("authentication failed: %v", expected["authenticated"])
	}
	assert.GreaterOrEqual(t, len(complaintIds), 1)
	complaintId := complaintIds[len(complaintIds)-1]
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	recipientsRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, repliesMock := range mock_data.NewReplies {
		for _, replyMock := range repliesMock {
			author, err := recipientsRepository.Find(ctx, find_recipient.ByNameAndEmail(replyMock.Sender.SubjectName,
				replyMock.Sender.SubjectEmail))
			assert.Nil(t, err)
			c := commands.NewReplyComplaintCommand(author.Id().String(), replyMock.Sender.Id.String(), complaintId.String(), replyMock.Body)
			b, err := json.Marshal(*c)
			assert.Nil(t, err)
			if err := ws.WriteMessage(websocket.TextMessage, b); err != nil {
				t.Fatalf("%v", err)
			}
			_, p, err := ws.ReadMessage()
			if err != nil {
				t.Fatalf("%v", err)
			}
			var expected dto.Reply
			if err := json.Unmarshal(p, &expected); err != nil {
				t.Fatalf("bad message")
			}
		}
	}

	t.Cleanup(func() {
		ctx := context.Background()
		reg := repositories.MapperRegistryInstance()
		userRepository, ok := reg.Get("User").(repositories.UserRepository)
		assert.True(t, ok)
		enterpriseRepository, ok := reg.Get("Enterprise").(repositories.EnterpriseRepository)
		assert.True(t, ok)
		complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
		assert.True(t, ok)
		recipientsRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for _, v := range mock_data.NewComplaints {
			err := complaintRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewRegisterEnterprises {
			err := enterpriseRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
			err = recipientsRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewUsers {
			err := userRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
			err = recipientsRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestChatHandler_ServeWS_2Users(t *testing.T) {
	TestChatHandler_Setup(t)
	s := httptest.NewServer(http.HandlerFunc(http_handlers.ServeWS))
	defer s.Close()

	assert.GreaterOrEqual(t, len(complaintIds), 1)
	id := complaintIds[len(complaintIds)-1]
	// Convert http://127.0.0.1 to ws://127.0.0.1
	u := "ws" + strings.TrimPrefix(s.URL, "http") + "?id=" + id.String()
	wsMap := map[string]*websocket.Conn{}
	for _, v := range authTokens {
		d := websocket.DefaultDialer
		d.Subprotocols = []string{"complaint"}
		// Connect to the server
		ws, _, err := d.Dial(u, nil)
		if err != nil {
			t.Fatalf("%v", err)
		}
		defer ws.Close()
		wsMap[v] = ws
	}
	for k, v := range wsMap {
		tokenMsg := map[string]string{"token": k}
		authTokenB, err := json.Marshal(tokenMsg)
		assert.Nil(t, err)
		if err := v.WriteMessage(websocket.TextMessage, authTokenB); err != nil {
			t.Fatalf("%v", err)
		}
		_, p, err := v.ReadMessage()
		if err != nil {
			t.Fatalf("%v", err)
		}
		var expected map[string]bool
		if err := json.Unmarshal(p, &expected); err != nil {
			t.Fatalf("bad message")
		}
		assert.NotNil(t, expected["authenticated"])
		if !expected["authenticated"] {
			t.Fatalf("authentication failed: %v", expected["authenticated"])
		}
	}
	assert.GreaterOrEqual(t, len(complaintIds), 1)
	complaintId := complaintIds[len(complaintIds)-1]
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	recipientsRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	msgs := [][]byte{}

	for _, repliesMock := range mock_data.NewReplies {
		for _, replyMock := range repliesMock {
			author, err := recipientsRepository.Find(ctx, find_recipient.ByNameAndEmail(replyMock.Sender.SubjectName,
				replyMock.Sender.SubjectEmail))
			assert.Nil(t, err)
			c := commands.NewReplyComplaintCommand(author.Id().String(), replyMock.Sender.Id.String(), complaintId.String(), replyMock.Body)
			b, err := json.Marshal(*c)
			assert.Nil(t, err)
			msgs = append(msgs, b)
		}
	}

	var pickOne *websocket.Conn
	for _, ws := range wsMap {
		pickOne = ws
	}

	for _, v := range msgs {
		if err := pickOne.WriteMessage(websocket.TextMessage, v); err != nil {
			t.Fatalf("%v", err)
		}
	}

	for _, ws := range wsMap {
		for range len(msgs) {
			_, p, err := ws.ReadMessage()
			if err != nil {
				t.Fatalf("%v", err)
			}
			var expected dto.Reply
			if err := json.Unmarshal(p, &expected); err != nil {
				t.Fatalf("bad message")
			}
		}
	}

	t.Cleanup(func() {
		ctx := context.Background()
		reg := repositories.MapperRegistryInstance()
		userRepository, ok := reg.Get("User").(repositories.UserRepository)
		assert.True(t, ok)
		enterpriseRepository, ok := reg.Get("Enterprise").(repositories.EnterpriseRepository)
		assert.True(t, ok)
		complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
		assert.True(t, ok)
		recipientsRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for _, v := range mock_data.NewComplaints {
			err := complaintRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewRegisterEnterprises {
			err := enterpriseRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
			err = recipientsRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
		for _, v := range mock_data.NewUsers {
			err := userRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
			err = recipientsRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}
