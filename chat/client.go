package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"go-complaint/application/application_services"
	"go-complaint/application/commands"
	"go-complaint/application/queries"
	"go-complaint/infrastructure/cache"
	"log"
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
	maxMessageSize = 4096
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
	defer func() {
		c.cas.unregister <- c
		c.conn.Close()
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
			_, err = svc.Authorize(context.Background(), string(chatMessage.Payload))
			response := ChatMessage{Type: ConnectionAcknowledged.String()}
			if err != nil {
				response.Payload = []byte("false")
			} else {
				response.Payload = []byte("true")
				c.isAuthenticated = true
			}
			m, err := json.Marshal(response)
			if err != nil {
				log.Printf("error marshalling authorize response false: %v", err)
				break
			}
			c.send <- m
		case Data.String():
			if !c.isAuthenticated {
				break
			}
			p, err := c.HandleSubProtocol(c.conn.Subprotocol(), chatMessage.Payload)
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

type ComplaintSubProtocolDataType int

func (c *Client) HandleSubProtocol(subProtocol string, payload []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	switch subProtocol {
	case "complaint":
		return complaintSubProtocol(ctx, payload)
	case "enterprise":
		return nil, fmt.Errorf(("not implemented yet"))
	default:
		return nil, fmt.Errorf(("unsupported subProtocol"))
	}
}

const (
	ReplyComplaint ComplaintSubProtocolDataType = iota
	MarkAsRead
	ComplaintReply
	Complaint
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
	default:
		return ""
	}
}

type ComplaintSubProtocolPayload struct {
	SubProtocolDataType string `json:"subProtocolDataType"`
	Command             []byte `json:"command"`
}

type ComplaintSubProtocolResult struct {
	SubProtocolDataType string `json:"subProtocolDataType"`
	Result              []byte `json:"result"`
}

type markAsSeenData struct {
	ComplaintId string `json:"complaintId"`
	ReplyId     string `json:"replyId"`
}

func complaintSubProtocol(ctx context.Context, payload []byte) ([]byte, error) {
	var p ComplaintSubProtocolPayload
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
		responsePayload, err := json.Marshal(ComplaintSubProtocolResult{
			SubProtocolDataType: ComplaintReply.String(),
			Result:              b,
		})
		if err != nil {
			return nil, fmt.Errorf("err marshaling response payload result: %v", err)
		}
		return responsePayload, nil
	case MarkAsRead.String():
		var d []markAsSeenData
		err = json.Unmarshal(p.Command, &d)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling mark as seen data: %v", err)
		}
		if len(d) <= 0 {
			return nil, fmt.Errorf("error unmarshalling mark as seen data the length is zero: %v", err)
		}
		complaintId := d[len(d)-1].ComplaintId
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
		responsePayload, err := json.Marshal(ComplaintSubProtocolResult{
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
