package chat

import (
	"context"
	"encoding/json"
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
		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		if !c.isAuthenticated {
			var tokenMsg map[string]string
			err = json.Unmarshal(message, &tokenMsg)
			if err != nil {
				log.Printf("error at unmarshal token: %v", err)
				break
			}
			r := map[string]bool{"authenticated": false}
			svc := application_services.AuthorizationApplicationServiceInstance()
			_, err = svc.Authorize(context.Background(), tokenMsg["token"])
			if err != nil {
				log.Printf("error at authorize: %v", err)
				m, err := json.Marshal(r)
				if err != nil {
					log.Printf("error marshalling authorize response false: %v", err)
					break
				}
				c.send <- m
				break
			}
			c.isAuthenticated = true
			r["authenticated"] = true
			m, err := json.Marshal(r)
			if err != nil {
				log.Printf("error marshaling authorize response true: %v", err)
				break
			}
			c.send <- m
		} else {
			subProtocol := c.conn.Subprotocol()
			switch subProtocol {
			case "complaint":
				var command commands.ReplyComplaintCommand
				err = json.Unmarshal(message, &command)
				if err != nil {
					log.Printf("error unmarshal command: %v", err)
					break
				}
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				err = command.Execute(ctx)
				if err != nil {
					cancel()
					log.Printf("error execute command: %v", err)
					break
				}
				replyId, ok := cache.InMemoryInstance().Get(command.ComplaintId)
				if !ok {
					cancel()
					log.Printf("replyId not cached with []byte key")
					break
				}
				q := queries.NewComplaintReplyQuery(replyId.(string))
				reply, err := q.Execute(ctx)
				if err != nil {
					cancel()
					log.Printf("err execute query: %v", err)
					break
				}
				b, err := json.Marshal(reply)
				if err != nil {
					cancel()
					log.Printf("err marshaling query result: %v", err)
					break
				}
				c.cas.broadcast <- b
				cancel()
			case "enterprise":
			default:
			}
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
