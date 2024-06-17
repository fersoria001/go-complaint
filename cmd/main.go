package main

import (
	"go-complaint/application"
	"go-complaint/cmd/api/middleware"
	"go-complaint/cmd/server"
	"log"
)

func main() {

	application.EventProcessorInstance().ResetDomainEventPublisher()
	s, err := server.NewGoComplaintServer(
		server.WithHandlerFunc(
			"/graphql",
			middleware.CORS()(middleware.Chain(server.GraphQLHandler,
				middleware.AuthenticationMiddleware(),
			)),
		),
		server.WithHandlerFunc("/subscriptions", server.SubscriptionsHandler),
		server.WithHandlerFunc("/publish", server.PublisherHandler),
	)
	if err != nil {
		log.Fatalf("Error starting server %s", err)
	}
	s.Run()
}
