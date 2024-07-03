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
	"os"

	"github.com/gorilla/csrf"
)

func main() {
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "5170")
	os.Setenv("ORIGIN", "http://localhost:5173")
	os.Setenv("CSRF-KEY", "ultrasecret")
	os.Setenv("JWT_SECRET", "supersecret123$%^&*")
	os.Setenv("DATABASE_URL", "postgres://postgres:sfdkwtf@localhost:5432/postgres?pool_max_conns=100&search_path=public&connect_timeout=5")
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
