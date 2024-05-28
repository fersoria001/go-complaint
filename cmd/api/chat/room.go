package chat

import (
	"go-complaint/domain/model/complaint"
	"go-complaint/dto"
	"log"
)

type Room struct {
	ID               string
	Clients          map[*Client]bool
	Broadcast        chan dto.Message
	Register         chan *Client
	Unregister       chan *Client
	Errors           chan error
	Complaint        *complaint.Complaint
	PrevReplies      []*complaint.Reply
	NewReplies       []*complaint.Reply
	Count            int
	PreviousMessages []dto.Message
}

func GenerateMessagesFromReplies(replies []*complaint.Reply) []dto.Message {
	var messages []dto.Message = []dto.Message{}
	for _, reply := range replies {
		messages = append(messages, dto.Message{
			Content: "reply",
			Reply: dto.ReplyDTO{
				ID:          reply.ID().String(),
				ComplaintID: reply.ComplaintID().String(),
				SenderID:    reply.SenderID(),
				SenderName:  reply.SenderName(),
				SenderIMG:   reply.SenderIMG(),
				Body:        reply.Body(),
				Read:        reply.Read(),
				CreatedAt:   reply.CreatedAt().StringRepresentation(),
				ReadAt:      reply.ReadAt().StringRepresentation(),
			},
		})
	}
	return messages
}

func (room *Room) Run(
	updateComplaintFn func(complaint *complaint.Complaint) error,
	updateReplyFn func(reply *complaint.Reply) error,
	saveReplyFn func(reply *complaint.Reply) error) {
	log.Printf("Room %s is running", room.ID)
	room.PreviousMessages = GenerateMessagesFromReplies(room.PrevReplies)
	for {
		select {
		case client := <-room.Register:
			room.Clients[client] = true
			log.Printf("Client with ID %s registered in room %s", client.ID, room.ID)
			for _, msg := range room.PreviousMessages {
				client.Send <- msg
			}
		case client := <-room.Unregister:
			if _, ok := room.Clients[client]; ok {
				delete(room.Clients, client)
				close(client.Send)
				log.Printf("Client with ID %s unregistered from room %s", client.ID, room.ID)
				if len(room.Clients) == 0 {
					log.Printf("Room %s is empty", room.ID)
					for _, reply := range room.NewReplies {
						err := saveReplyFn(reply)
						if err != nil {
							room.Errors <- err
						}
					}
					for _, reply := range room.PrevReplies {
						err := updateReplyFn(reply)
						if err != nil {
							room.Errors <- err
						}
					}
					err := updateComplaintFn(room.Complaint)
					if err != nil {
						room.Errors <- err
					}
					roomsMutex.Lock()
					delete(rooms, room.ID)
					roomsMutex.Unlock()
					log.Printf("Room %s is empty,db has been updated and has been deleted", room.ID)
					return
				}
				log.Printf("Room %s has %d clients", room.ID, len(room.Clients))
			}
		case message := <-room.Broadcast:
			log.Printf("dto.Message received in room %s: %s", room.ID, message.Content)
			room.PreviousMessages = append(room.PreviousMessages, message)
			for client := range room.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(room.Clients, client)
					log.Printf("Failed to send message to a client in room %s, client removed", room.ID)
				}
			}
		case err := <-room.Errors:
			log.Printf("Loggin error from Errors channel from Room %s: %v", room.ID, err)
		}
	}
}
