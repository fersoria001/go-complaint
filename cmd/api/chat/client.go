package chat

import (
	"context"
	"go-complaint/dto"
	"log"
	"sync"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Room *Room
	Send chan dto.Message
	m    sync.Mutex
}

func (client *Client) ReadPump(ctx context.Context) {
	for {
		var msg dto.Message
		err := wsjson.Read(ctx, client.Conn, &msg)
		log.Printf("dto.Message read: %v", msg)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			log.Println("Client disconnected normally")
			break
		}
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		client.m.Lock()
		replyy, err := client.Room.Complaint.ReplyComplaint(
			ctx,
			client.Room.Count,
			client.ID,
			msg.Reply.SenderIMG,
			msg.Reply.SenderName,
			msg.Reply.Body,
		)
		if err != nil {
			client.Room.Errors <- err
		}
		msg.Reply = *dto.NewReplyDTO(replyy)
		client.Room.NewReplies = append(client.Room.NewReplies, replyy)
		client.Room.Count++
		client.m.Unlock()
		client.Room.Broadcast <- msg
		log.Printf("dto.Message read: %s", msg.Content)
	}
}

func (client *Client) WritePump(ctx context.Context) {
	for {
		select {
		case msg := <-client.Send:
			if ctx.Err() != nil {
				return
			}
			err := wsjson.Write(ctx, client.Conn, msg)
			if err != nil {
				log.Printf("Error writing message: %v", err)
				return
			}
			log.Printf("dto.Message written: %s", msg.Content)
		case <-ctx.Done():
			log.Println("Context done in WritePump")
			return
		}
	}
}
