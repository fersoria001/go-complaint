package main

import (
	"go-complaint/application/commands"
	"go-complaint/cmd/api/files"
	"go-complaint/cmd/api/middleware"
	"go-complaint/cmd/server"
	"go-complaint/cmd/server/authentication"
	"go-complaint/cmd/server/graphql_subscriptions"
	"go-complaint/graph"
	"go-complaint/infrastructure/cache"
	projectpath "go-complaint/project_path"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gorilla/csrf"
)

func main() {
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "5170")
	os.Setenv("ORIGIN", "http://localhost:3000,localhost:3000,localhost")
	os.Setenv("CSRF_KEY", "ultrasecret")
	os.Setenv("DATABASE_URL", "postgres://postgres:sfdkwtf@localhost:5432/postgres?pool_max_conns=100&search_path=public&connect_timeout=5")
	os.Setenv("PORT", "5170")
	os.Setenv("DNS", "http//localhost:3000")
	os.Setenv("SEND_GRID_API_KEY", "Bearer mlsn.0557f4217143328c73149ad91c7455121924f188c63af0fe093b42feb3fa1de1")
	profileImgHandler := http.StripPrefix("/profile_img/", http.FileServer(http.Dir(projectpath.ProfileImgsPath)))
	logoImgsHandler := http.StripPrefix("/logo_img/", http.FileServer(http.Dir(projectpath.LogoImgsPath)))
	bannerImgsHandler := http.StripPrefix("/banner_img/", http.FileServer(http.Dir(projectpath.BannerImgsPath)))
	graphQLHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	csrfProtectedMux := server.NewCSRFProtectedMux(
		server.CSRFPMWithHandlerFunc(
			"/upload",
			middleware.CORS()(middleware.Chain(files.UploadFileHandler,
				middleware.AuthenticationMiddleware(),
			)),
		),
		server.CSRFPMWithHandlerFunc(
			"/sign-in",
			middleware.CORS()(middleware.Chain(authentication.SignInHandler, middleware.AuthenticationMiddleware())),
		),
		server.CSRFPMWithHandlerFunc(
			"/confirm-sign-in",
			middleware.CORS()(middleware.Chain(authentication.ConfirmSignInHandler, middleware.AuthenticationMiddleware())),
		),
	)
	s, err := server.NewGoComplaintServer(
		server.WithHandler("/", csrfProtectedMux),
		server.WithHandler("/graphql", middleware.CORS1()(graphQLHandler)),
		server.WithHandler("/profile_img/", profileImgHandler),
		server.WithHandler("/logo_img/", logoImgsHandler),
		server.WithHandler("/banner_img/", bannerImgsHandler),
		server.WithHandlerFunc("/subscriptions", server.SubscriptionsHandler),
		server.WithHandlerFunc("/confirmation-link", ValidateEmail),
		server.WithHandlerFunc("/csrf", middleware.CORS()(CSRF)),
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

func ValidateEmail(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	token := queries.Get("token")
	userCommand := commands.UserCommand{
		EmailVerificationToken: token,
	}
	err := userCommand.VerifyEmail(r.Context())
	if err != nil {
		newUrl := "https://www.go-complaint.com/"
		log.Println("errorLoginIn")
		http.Redirect(w, r, newUrl, http.StatusSeeOther)
		return
	}
	newUrl := "https://www.go-complaint.com/"
	http.Redirect(w, r, newUrl, http.StatusSeeOther)
}
