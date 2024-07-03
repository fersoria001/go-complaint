package main

import (
	"go-complaint/cmd/api/files"
	"go-complaint/cmd/api/middleware"
	"go-complaint/cmd/server"
	"go-complaint/cmd/server/graphql_subscriptions"
	"go-complaint/infrastructure/cache"
	projectpath "go-complaint/project_path"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

func main() {
	profileImgHandler := http.StripPrefix("/profile_img/", http.FileServer(http.Dir(projectpath.ProfileImgsPath)))
	logoImgsHandler := http.StripPrefix("/logo_img/", http.FileServer(http.Dir(projectpath.LogoImgsPath)))
	bannerImgsHandler := http.StripPrefix("/banner_img/", http.FileServer(http.Dir(projectpath.BannerImgsPath)))
	csrfProtectedMux := server.NewCSRFProtectedMux(
		server.CSRFPMWithHandlerFunc(
			"/graphql",
			middleware.CORS()(middleware.Chain(server.GraphQLHandler,
				middleware.AuthenticationMiddleware(),
			)),
		),
		server.CSRFPMWithHandlerFunc(
			"/publish",
			middleware.CORS()(middleware.Chain(server.PublisherHandler,
				middleware.AuthenticationMiddleware(),
			)),
		),
		server.CSRFPMWithHandlerFunc(
			"/csrf",
			middleware.CORS()(CSRF),
		),
		server.CSRFPMWithHandlerFunc(
			"/upload",
			middleware.CORS()(middleware.Chain(files.UploadFileHandler,
				middleware.AuthenticationMiddleware(),
			)),
		),
	)
	s, err := server.NewGoComplaintServer(
		server.WithHandler("/", csrfProtectedMux),
		server.WithHandler("/profile_img/", profileImgHandler),
		server.WithHandler("/logo_img/", logoImgsHandler),
		server.WithHandler("/banner_img/", bannerImgsHandler),
		server.WithHandlerFunc("/subscriptions", server.SubscriptionsHandler),
	)
	if err != nil {
		log.Fatalf("Error starting server %s", err)
	}
	go graphql_subscriptions.SubscriptionsPublisherInstance().Background(cache.RequestChannel)
	go cache.Cache(cache.RequestChannel2)
	s.Run()
}

func CSRF(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
}
