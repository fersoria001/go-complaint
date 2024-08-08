package main

import (
	"fmt"
	"go-complaint/application"
	"go-complaint/application/application_services"
	"go-complaint/graph"
	"go-complaint/http_handlers"
	"slices"
	"time"

	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

func main() {
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "5170")
	os.Setenv("ALLOWED_ORIGINS", "http://localhost:5170,http://localhost:3000,localhost:3000,localhost")
	os.Setenv("CSRF_KEY", "ultrasecret")
	os.Setenv("DATABASE_URL", "postgres://postgres:sfdkwtf@localhost:5432/postgres?pool_max_conns=100&search_path=public&connect_timeout=5")
	os.Setenv("PORT", "5170")
	os.Setenv("DNS", "http://localhost:3000")
	os.Setenv("SEND_GRID_API_KEY", "Bearer mlsn.0557f4217143328c73149ad91c7455121924f188c63af0fe093b42feb3fa1de1")
	r := chi.NewRouter()
	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r.Use(middleware.Logger)
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"cookie", "content-type", "upgrade", "connection", "sec-websocket-key"},
		ExposedHeaders:   []string{"set-cookie", "upgrade", "connection", "sec-websocket-accept"},
	}).Handler)
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("jwt")
			if err != nil {
				h.ServeHTTP(w, r)
				return
			}
			jwtToken := cookie.Value
			svc := application_services.AuthorizationApplicationServiceInstance()
			authorized, err := svc.Authorize(r.Context(), jwtToken)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			r = r.WithContext(authorized)
			h.ServeHTTP(w, r)
		})
	})
	publisher := application.ApplicationMessagePublisherInstance()
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Publisher: publisher}}))
	srv.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				origin := r.Header.Get("Origin")
				return slices.Contains(allowedOrigins, origin)
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	r.Handle("/", playground.Handler("GoComplaint GraphQL", "/graphql"))
	r.Handle("/graphql", srv)
	r.HandleFunc("/sign-in", http_handlers.SignInHandler)
	r.HandleFunc("/confirm-sign-in", http_handlers.ConfirmSignInHandler)
	r.HandleFunc("/chat", http_handlers.ServeWS)
	r.HandleFunc("/session", http_handlers.SessionHandler)
	err := http.ListenAndServe(port, r)
	if err != nil {
		panic(err)
	}
}
