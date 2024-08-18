package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"go-complaint/application/application_services"
	"go-complaint/application/commands"
	"go-complaint/application/queries"
	"go-complaint/dto"
	"go-complaint/infrastructure/cache"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024 * 8
)

var (
	newline = []byte{'\n'}
)

type ChatMessageType int

const (
	ConnectionInit ChatMessageType = iota
	ConnectionAcknowledged
	Data
	Complete
	UserOnline
	UserOffline
)

func (cmt ChatMessageType) String() string {
	switch cmt {
	case ConnectionInit:
		return "connection_init"
	case ConnectionAcknowledged:
		return "connection_ack"
	case Data:
		return "data"
	case Complete:
		return "complete"
	case UserOnline:
		return "user_online"
	case UserOffline:
		return "user_offline"
	default:
		return ""
	}
}

type ChatMessage struct {
	Type    string `json:"type"`
	Payload []byte `json:"payload"`
}

type Client struct {
	cas             *ChatAdapter
	conn            *websocket.Conn
	send            chan []byte
	isAuthenticated bool
	id              string
	mu              sync.Mutex
}

func NewClient(cas *ChatAdapter, conn *websocket.Conn) *Client {
	return &Client{
		cas:  cas,
		conn: conn,
		send: make(chan []byte),
	}
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) Read() {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		//log.Printf("closing connection %v", c)
		offlineMsg, err := json.Marshal(&ChatMessage{Type: UserOffline.String(), Payload: []byte(c.id)})
		if err != nil {
			log.Printf("error marshaling offline user msg connection init %v", err)
		}
		c.cas.broadcast <- offlineMsg
		c.cas.unregister <- c
		c.conn.Close()
		cancel()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error at reading message: %v", err)
			}
			log.Printf("error at reading message: %v", err)
			break
		}
		var chatMessage ChatMessage
		err = json.Unmarshal(message, &chatMessage)
		if err != nil {
			log.Printf("error at unmarshal token: %v", err)
			break
		}

		switch chatMessage.Type {
		case ConnectionInit.String():
			svc := application_services.AuthorizationApplicationServiceInstance()
			authCtx, err := svc.Authorize(ctx, string(chatMessage.Payload))
			response := ChatMessage{Type: ConnectionAcknowledged.String()}
			if err != nil {
				response.Payload = []byte("false")
			} else {
				credentials, err := svc.Credentials(authCtx)
				if err != nil {
					log.Printf("error obtaining credentials connection init %v", err)
					break
				}
				c.mu.Lock()
				c.id = credentials.Id
				c.isAuthenticated = true
				c.mu.Unlock()
				onlineMsg, err := json.Marshal(&ChatMessage{Type: UserOnline.String(), Payload: []byte(c.id)})
				if err != nil {
					log.Printf("error marshaling online user msg connection init %v", err)
					break
				}
				c.cas.broadcast <- onlineMsg
				response.Payload = []byte("true")
			}
			m, err := json.Marshal(response)
			if err != nil {
				log.Printf("error marshalling authorize response false: %v", err)
				break
			}
			c.send <- m
		case Data.String():
			if !c.isAuthenticated {
				log.Printf("error at sub protocol handler client is not authenticated: %v", err)
				break
			}
			p, err := c.HandleSubProtocol(ctx, c.conn.Subprotocol(), chatMessage.Payload)
			if err != nil {
				log.Printf("error at sub protocol handler false: %v", err)
				break
			}
			response, err := json.Marshal(&ChatMessage{Type: Data.String(), Payload: p})
			if err != nil {
				log.Printf("error marshalling data response: %v", err)
				break
			}
			c.cas.broadcast <- response
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) HandleSubProtocol(ctx context.Context, subProtocol string, payload []byte) ([]byte, error) {
	switch subProtocol {
	case "complaint":
		return complaintSubProtocol(ctx, payload)
	case "enterpriseChat":
		return enterpriseChatSubProtocol(ctx, payload)
	default:
		return nil, fmt.Errorf(("unsupported subProtocol"))
	}
}

type SubProtocolPayload struct {
	SubProtocolDataType string `json:"subProtocolDataType"`
	Command             []byte `json:"command"`
}

type SubProtocolResult struct {
	SubProtocolDataType string `json:"subProtocolDataType"`
	Result              []byte `json:"result"`
}

type MarkAsSeenData struct {
	Id      string `json:"id"`
	ReplyId string `json:"replyId"`
}

type EnterpriseChatSubProtocolDataType int

const (
	MarkAsSeenCommand EnterpriseChatSubProtocolDataType = iota
	ReplyCommand
	ChatReply
	Chat
)

func (ecspdt EnterpriseChatSubProtocolDataType) String() string {
	switch ecspdt {
	case MarkAsSeenCommand:
		return "mark_as_seen_command"
	case ReplyCommand:
		return "reply_command"
	case ChatReply:
		return "chat_reply"
	case Chat:
		return "chat"
	default:
		return ""
	}
}

func enterpriseChatSubProtocol(ctx context.Context, payload []byte) ([]byte, error) {
	var p SubProtocolPayload
	err := json.Unmarshal(payload, &p)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal enterprise chat subprotocol payload: %v", err)
	}
	switch p.SubProtocolDataType {
	case ReplyCommand.String():
		var c commands.ReplyEnterpriseChatCommand
		err := json.Unmarshal(p.Command, &c)
		if err != nil {
			return nil, fmt.Errorf("error unmarshal reply enterprise chat command %v", err)
		}

		err = c.Execute(ctx)
		if err != nil {
			return nil, fmt.Errorf("error execute reply chat command: %v", err)
		}
		replyMsg, ok := cache.InMemoryInstance().Get(c.ChatId)
		if !ok {
			return nil, fmt.Errorf("error receiving replyMsg in replyEnterpriseChat subProtocol")
		}
		reply, ok := replyMsg.(dto.ChatReply)
		if !ok {
			return nil, fmt.Errorf("error casting replyMsg in replyEnterpriseChat subProtocol")
		}
		b, err := json.Marshal(reply)
		if err != nil {
			return nil, fmt.Errorf("err marshal reply at replyEnterpriseChat sub protocol, %v", err)
		}
		responsePayload, err := json.Marshal(SubProtocolResult{
			SubProtocolDataType: ChatReply.String(),
			Result:              b,
		})
		if err != nil {
			return nil, fmt.Errorf("error marshaling response payload at enterpriseChatReply subProtocol %v", err)
		}
		return responsePayload, nil
	case MarkAsSeenCommand.String():
		var d []MarkAsSeenData
		err := json.Unmarshal(p.Command, &d)
		if err != nil {
			return nil, fmt.Errorf(" error at unmarshall markAsSeen data in enterpriseChat subProtocol %v", err)
		}
		if len(d) <= 0 {
			return nil, fmt.Errorf(" error at unmarshall markAsSeen data in enterpriseChat subProtocol, the length is zero %v", err)
		}
		chatId := d[len(d)-1].Id
		ids := make([]string, 0, len(d))
		for i := range d {
			ids = append(ids, d[i].ReplyId)
		}
		c := commands.NewMarkEnterpriseChatReplyAsSeenCommand(chatId, ids)
		err = c.Execute(ctx)
		if err != nil {
			return nil, fmt.Errorf("error execute markEnterpriseChatReplyAsSeen in enterpriseChatSubProtocol %v", err)
		}
		q := queries.NewEnterpriseChatByIdQuery(chatId)
		dbC, err := q.Execute(ctx)
		if err != nil {
			return nil, fmt.Errorf("error execute query after mark replies as seen enterpriseChatSubProtocol %v", err)
		}
		b, err := json.Marshal(dbC)
		if err != nil {
			return nil, fmt.Errorf("error marshal query result after mark replies as seen enterpriseChatSubProtocol %v", err)
		}
		responsePayload, err := json.Marshal(SubProtocolResult{
			SubProtocolDataType: Chat.String(),
			Result:              b,
		})
		if err != nil {
			return nil, fmt.Errorf("error marshal response payload result enterpriseChatSubProtocol %v", err)
		}
		return responsePayload, nil
	default:
		return nil, fmt.Errorf("not implemented yet")
	}
}

type ComplaintSubProtocolDataType int

const (
	ReplyComplaint ComplaintSubProtocolDataType = iota
	MarkAsRead
	ComplaintReply
	Complaint
	SendToReview
)

func (cspdt ComplaintSubProtocolDataType) String() string {
	switch cspdt {
	case ReplyComplaint:
		return "reply_complaint"
	case MarkAsRead:
		return "mark_as_read"
	case ComplaintReply:
		return "complaint_reply"
	case Complaint:
		return "complaint"
	case SendToReview:
		return "send_to_review"
	default:
		return ""
	}
}

func complaintSubProtocol(ctx context.Context, payload []byte) ([]byte, error) {
	var p SubProtocolPayload
	err := json.Unmarshal(payload, &p)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal complaint subprotocol payload: %v", err)
	}
	switch p.SubProtocolDataType {
	case ReplyComplaint.String():
		var c commands.ReplyComplaintCommand
		err := json.Unmarshal(p.Command, &c)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling reply complaint command: %v", err)
		}
		err = c.Execute(ctx)
		if err != nil {
			return nil, fmt.Errorf("error execute command: %v", err)
		}
		replyId, ok := cache.InMemoryInstance().Get(c.ComplaintId)
		if !ok {
			return nil, fmt.Errorf("replyId not cached with []byte key")
		}
		q := queries.NewComplaintReplyQuery(replyId.(string))
		reply, err := q.Execute(ctx)
		if err != nil {
			return nil, fmt.Errorf("err execute query: %v", err)
		}
		b, err := json.Marshal(reply)
		if err != nil {
			return nil, fmt.Errorf("err marshaling query result: %v", err)
		}
		responsePayload, err := json.Marshal(SubProtocolResult{
			SubProtocolDataType: ComplaintReply.String(),
			Result:              b,
		})
		if err != nil {
			return nil, fmt.Errorf("err marshaling response payload result: %v", err)
		}
		return responsePayload, nil
	case MarkAsRead.String():
		var d []MarkAsSeenData
		err = json.Unmarshal(p.Command, &d)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling mark as seen data: %v", err)
		}
		if len(d) <= 0 {
			return nil, fmt.Errorf("error unmarshalling mark as seen data the length is zero: %v", err)
		}
		complaintId := d[len(d)-1].Id
		ids := make([]string, 0, len(d))
		for i := range d {
			ids = append(ids, d[i].ReplyId)
		}
		c := commands.NewMarkRepliesAsReadCommand(complaintId, ids)
		err := c.Execute(ctx)
		if err != nil {
			return nil, fmt.Errorf("error executing mark replies as read command: %v", err)
		}
		q := queries.NewComplaintByIdQuery(complaintId)
		dbC, err := q.Execute(ctx)
		if err != nil {
			return nil, fmt.Errorf("error executing query after mark replies as read: %v", err)
		}
		b, err := json.Marshal(dbC)
		if err != nil {
			return nil, fmt.Errorf("error marshalling query  result after mark replies as read: %v", err)
		}
		responsePayload, err := json.Marshal(SubProtocolResult{
			SubProtocolDataType: Complaint.String(),
			Result:              b,
		})
		if err != nil {
			return nil, fmt.Errorf("err marshaling response payload result: %v", err)
		}
		return responsePayload, nil
	case SendToReview.String():
		var c commands.SendComplaintToReviewCommand
		err := json.Unmarshal(p.Command, &c)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling send complaint to review command data: %v", err)
		}
		err = c.Execute(ctx)
		if err != nil {
			return nil, fmt.Errorf("error executing send complaint to review command: %v", err)
		}
		q := queries.NewComplaintByIdQuery(c.ComplaintId)
		dbC, err := q.Execute(ctx)
		if err != nil {
			return nil, fmt.Errorf("error executing query after send complaint to review: %v", err)
		}
		b, err := json.Marshal(dbC)
		if err != nil {
			return nil, fmt.Errorf("error marshalling query  result after send complaint to review: %v", err)
		}
		responsePayload, err := json.Marshal(SubProtocolResult{
			SubProtocolDataType: Complaint.String(),
			Result:              b,
		})
		if err != nil {
			return nil, fmt.Errorf("err marshaling response payload result: %v", err)
		}
		return responsePayload, nil
	default:
		return nil, fmt.Errorf("default case not implemented")
	}
}
