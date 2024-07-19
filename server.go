package main

import (
	"fmt"
	"go-complaint/cmd/api/middleware"
	"go-complaint/cmd/server/authentication"
	"go-complaint/graph"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func main() {
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "5170")
	os.Setenv("ALLOWED_ORIGINS", "http://localhost:3000,localhost:3000,localhost")
	os.Setenv("CSRF_KEY", "ultrasecret")
	os.Setenv("DATABASE_URL", "postgres://postgres:sfdkwtf@localhost:5432/postgres?pool_max_conns=100&search_path=public&connect_timeout=5")
	os.Setenv("PORT", "5170")
	os.Setenv("DNS", "http://localhost:3000")
	os.Setenv("SEND_GRID_API_KEY", "Bearer mlsn.0557f4217143328c73149ad91c7455121924f188c63af0fe093b42feb3fa1de1")
	r := chi.NewRouter()
	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"cookie", "content-type", "upgrade", "connection", "sec-websocket-key"},
		ExposedHeaders:   []string{"set-cookie", "upgrade", "connection", "sec-websocket-accept"},
	}).Handler)
	r.Use(middleware.AuthenticationMiddleware())

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(&transport.Websocket{})

	r.Handle("/", playground.Handler("GoComplaint GraphQL", "/graphql"))
	r.Handle("/graphql", srv)
	r.HandleFunc("/sign-in", authentication.SignInHandler)
	r.HandleFunc("/confirm-sign-in", authentication.ConfirmSignInHandler)
	err := http.ListenAndServe(port, r)
	log.Println("server started")
	if err != nil {
		panic(err)
	}
}
