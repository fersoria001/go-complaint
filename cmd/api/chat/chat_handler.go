package chat

import (
	"context"
	"go-complaint/application"
	"go-complaint/domain/model/complaint"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"log"
	"net/http"
	"sync"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// spaghetti monster
// refactor after release, it will require the in memory object change detection weekends impl
var rooms = make(map[string]*Room)
var roomsMutex = sync.Mutex{}

type AuthMessage struct {
	Content      string `json:"content"`
	JWTToken     string `json:"jwt_token"`
	EnterpriseID string `json:"enterprise_id"`
}

type ConfirmationMessage struct {
	Content string `json:"content"`
	Success bool   `json:"success"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	complaintsRepository := repositories.NewComplaintRepository(datasource.ComplaintSchema())
	repliesRepository := repositories.NewReplyRepository(datasource.ComplaintSchema())
	enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
	service := application.NewComplaintService(complaintsRepository, repliesRepository)
	roomID := r.URL.Query().Get("room")
	if roomID == "" {
		http.Error(w, "Room ID is required", http.StatusBadRequest)
		return
	}

	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"localhost:5173"},
	})
	if err != nil {
		log.Printf("Error accepting websocket connection: %v", err)
		return
	}
	log.Println("WebSocket connection opened, authenticating...")
	var authMessage AuthMessage
	authCtx, cancelAuthCtx := context.WithCancel(context.Background())
	defer cancelAuthCtx()

	var authWG sync.WaitGroup
	authWG.Add(1) // Adding one goroutine to the WaitGroup for authentication

	var clientID string
	go func() {
		defer authWG.Done() // Notify WaitGroup that authentication processing is done
		authSuccessMsg := ConfirmationMessage{
			Content: "auth",
			Success: false,
		}
		err = wsjson.Read(authCtx, conn, &authMessage)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			log.Println("Client disconnected normally at auth")
			cancelAuthCtx()
			return
		}
		if err != nil {
			log.Printf("Error reading authentication message: %v", err)
			cancelAuthCtx()
			conn.Close(websocket.StatusPolicyViolation, "Invalid authentication message")
			return
		}
		claims, err := application.NewJWTService().ParseUserDescriptor(authMessage.JWTToken)
		if err != nil {
			log.Printf("Error parsing JWT token, authentication failed: %v", err)
			err = wsjson.Write(authCtx, conn, authSuccessMsg)
			if err != nil {
				log.Printf("Error writing auth failed message: %v", err)
			}
			cancelAuthCtx()
			conn.Close(websocket.StatusPolicyViolation, "Invalid JWT token")
			return
		}
		clientID = claims.Email
		if authMessage.EnterpriseID != "" {
			ep, err := enterpriseRepository.Get(authCtx, authMessage.EnterpriseID)
			if err != nil {
				log.Printf("Error getting enterprise: %v", err)
				err = wsjson.Write(authCtx, conn, authSuccessMsg)
				if err != nil {
					log.Printf("Error writing auth failed message: %v", err)
				}
				cancelAuthCtx()
				conn.Close(websocket.StatusPolicyViolation, "Invalid enterprise ID")
				return
			}
			clientID = ep.Name()
		}
		authSuccessMsg.Success = true
		err = wsjson.Write(authCtx, conn, authSuccessMsg)
		if err != nil {
			log.Printf("Error writing message: %v", err)
			return
		}
		log.Printf("User %s authenticated", clientID)
	}()
	// Wait for the authentication process to complete before proceeding
	authWG.Wait()
	if clientID == "" {
		// Authentication failed, do not proceed
		return
	}

	roomsMutex.Lock()
	room, ok := rooms[roomID]
	if !ok {
		complaintt, prevReplies, err := service.ProvideDomainComplaintAndReplies(context.Background(), roomID)
		if err != nil {
			if complaintt == nil {
				log.Printf("Error getting complaint: %v", err)
				http.Error(w, "Room not found", http.StatusNotFound)
				return
			}
			prevReplies = []*complaint.Reply{}
		}
		count, err := repliesRepository.Count(context.Background(), roomID)
		if err != nil {
			http.Error(w, "Error counting replies", http.StatusInternalServerError)
			return
		}
		room = &Room{
			ID:          roomID,
			Clients:     make(map[*Client]bool),
			Broadcast:   make(chan dto.Message),
			Register:    make(chan *Client),
			Unregister:  make(chan *Client),
			Complaint:   complaintt,
			PrevReplies: prevReplies,
			NewReplies:  []*complaint.Reply{},
			Count:       count,
		}
		rooms[roomID] = room
		go room.Run(
			func(ct *complaint.Complaint) error {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				return complaintsRepository.Update(ctx, ct)
			},
			func(reply *complaint.Reply) error {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				return repliesRepository.Update(ctx, reply)
			},
			func(reply *complaint.Reply) error {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				return repliesRepository.Save(ctx, reply)
			},
		)
		log.Printf("Room %s created", roomID)
	}
	roomsMutex.Unlock()

	client := &Client{ID: clientID, Conn: conn, Room: room, Send: make(chan dto.Message)}
	room.Register <- client

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(2) // Add two goroutines (ReadPump and WritePump)

	go func() {
		defer cancel()
		client.ReadPump(ctx)
		wg.Done() // Notify WaitGroup that ReadPump has finished
	}()
	go func() {
		defer cancel()
		client.WritePump(ctx)
		wg.Done() // Notify WaitGroup that WritePump has finished
	}()

	// Wait for all client goroutines to finish
	go func() {
		wg.Wait()
		log.Println("All client goroutines have finished")
		client.Room.Unregister <- client
		log.Println("Client unregistered")
	}()

	// Keep the handler alive until the WebSocket connection is closed
	<-ctx.Done()
	log.Println("Context done in ChatHandler")
}
