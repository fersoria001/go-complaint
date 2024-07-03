package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
)

type CSRFProtectedMux struct {
	mux *http.ServeMux
}

func NewCSRFProtectedMux(options ...OptionsCSRFProtectedMuxFunc) *CSRFProtectedMux {
	csrfpm := &CSRFProtectedMux{
		mux: http.NewServeMux(),
	}
	for _, option := range options {
		err := option(csrfpm)
		if err != nil {
			log.Fatalf("Error starting server %s", err)
		}
	}
	return csrfpm
}

type OptionsCSRFProtectedMuxFunc func(c *CSRFProtectedMux) error

func CSRFPMWithHandlerFunc(pattern string, handler http.HandlerFunc) OptionsCSRFProtectedMuxFunc {
	return func(c *CSRFProtectedMux) error {
		c.mux.HandleFunc(pattern, handler)
		return nil
	}
}

func (csrfpm *CSRFProtectedMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	csrfKey := os.Getenv("CSRF-KEY")
	origin := os.Getenv("ORIGIN")
	log.Println(origin)
	csrfMiddleware := csrf.Protect(
		[]byte(csrfKey),
		csrf.Path("/"),
		csrf.HttpOnly(false),
		csrf.Secure(false),
		csrf.SameSite(csrf.SameSiteLaxMode),
		csrf.TrustedOrigins([]string{origin}),
		csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := csrf.FailureReason(r)
			log.Println(err)
		})),
		csrf.CookieName("x-csrf-token"),
	)
	csrfMiddleware(csrfpm.mux).ServeHTTP(w, r)
}
