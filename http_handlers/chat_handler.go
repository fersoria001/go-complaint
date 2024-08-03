package http_handlers

import (
	"go-complaint/chat"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/gorilla/websocket"
)

var allowedOrigins = strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return slices.Contains(allowedOrigins, r.Host)
	},
	ReadBufferSize:  8192,
	WriteBufferSize: 8192,
	Subprotocols:    []string{"complaint", "enterprise"},
}

var chats = map[string]*chat.ChatAdapter{}

func ServeWS(w http.ResponseWriter, r *http.Request) {
	chatId := r.URL.Query().Get("id")
	if chatId == "" {
		log.Println("query is empty")
		w.WriteHeader(http.StatusBadRequest)
	}
	log.Println("ws headers", r.Header)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error at upgrade", err)
		return
	}
	svc, ok := chats[chatId]
	if !ok {
		newInstance := chat.NewChatAdapter()
		go newInstance.Run()
		chats[chatId] = newInstance
		svc = newInstance
	}
	client := chat.NewClient(svc, conn)
	svc.Register(client)

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.Write()
	go client.Read()
}
