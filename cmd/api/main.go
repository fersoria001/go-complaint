package main

import (
	"go-complaint/cmd"
	"go-complaint/cmd/api/chat"
	"go-complaint/cmd/api/graphql_"
	"go-complaint/cmd/api/middleware"
	"go-complaint/cmd/api/notifications"
	"log"
)

func main() {
	s, err := cmd.NewServer(
		cmd.WithHandler("/graphql",
			middleware.CORS()(
				middleware.Chain(graphql_.ExecuteAndEncodeBody, middleware.AuthenticationMiddleware()),
			),
		),
		cmd.WithHandler("/chat",
			middleware.Chain(chat.ChatHandler, middleware.WebsocketAuthenticationMiddleware()),
		),
		cmd.WithHandler("/notifications",
			middleware.CORS()(
				middleware.Chain(notifications.ProvideNotifications, middleware.AuthenticationMiddleware()),
			),
		),
		cmd.WithHandler("/notification",
			middleware.CORS()(
				middleware.Chain(notifications.Notification, middleware.AuthenticationMiddleware()),
			),
		),
	)

	if err != nil {
		log.Fatalf("Error starting server %s", err)
	}

	s.Run()
}
