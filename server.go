package main

import (
	"fmt"
	"go-complaint/application"
	"go-complaint/application/application_services"
	"go-complaint/domain"
	"go-complaint/graph"
	"go-complaint/http_handlers"
	projectpath "go-complaint/project_path"
	"slices"
	"time"

	"crypto/tls"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

func main() {
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "5170")
	os.Setenv("ALLOWED_ORIGINS", "http://localhost:5170,http://localhost:3000,localhost:3000,localhost")
	os.Setenv("CSRF_KEY", "")
	os.Setenv("JWT_SECRET", "")
	os.Setenv("DATABASE_URL", "")
	os.Setenv("PORT", "5170")
	os.Setenv("DNS", "http://localhost:5170")
	os.Setenv("SEND_GRID_API_KEY", "")
	os.Setenv("ENVIRONMENT", "DEV")
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
			apiKey := r.Header.Get("api_key")
			if apiKey == "" {
				apiKey := r.URL.Query().Get("api_key")
				if apiKey == "" {
					http.Error(w, fmt.Errorf("api not found in request").Error(), http.StatusForbidden)
					return
				}
			}
			svc := application_services.JWTApplicationServiceInstance()
			err := svc.ParseApiKey(apiKey)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			h.ServeHTTP(w, r)
		})
	})
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
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pub := domain.DomainEventPublisherInstance()
			pub.Reset()
			eventProcessor := application.EventProcessor{}
			pub.Subscribe(domain.DomainEventSubscriber{
				HandleEvent:           eventProcessor.HandleEvent,
				SubscribedToEventType: eventProcessor.SubscribedToEventType,
			})
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
	profileImgHandler := http.StripPrefix("/profile_img/", http.FileServer(http.Dir(projectpath.ProfileImgsPath)))
	logoImgsHandler := http.StripPrefix("/logo_img/", http.FileServer(http.Dir(projectpath.LogoImgsPath)))
	bannerImgsHandler := http.StripPrefix("/banner_img/", http.FileServer(http.Dir(projectpath.BannerImgsPath)))
	//r.Handle("/", playground.Handler("GoComplaint GraphQL", "/graphql"))
	r.Handle("/graphql", srv)
	r.HandleFunc("/sign-in", http_handlers.SignInHandler)
	r.HandleFunc("/confirm-sign-in", http_handlers.ConfirmSignInHandler)
	r.HandleFunc("/chat", http_handlers.ServeWS)
	r.HandleFunc("/session", http_handlers.SessionHandler)
	r.Handle("/profile_img/*", profileImgHandler)
	r.Handle("/logo_img/*", logoImgsHandler)
	r.Handle("/banner_img/*", bannerImgsHandler)

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	mainserver := &http.Server{
		Addr:         port,
		Handler:      r,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	err := mainserver.ListenAndServeTLS(projectpath.CertPath, projectpath.KeyPath)

	//err := http.ListenAndServe(port, r)
	if err != nil {
		log.Printf("error at ListenAndServeTLS %v", err)
		panic(err)
	}
}
