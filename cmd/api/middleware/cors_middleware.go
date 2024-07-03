package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

func CORS() Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			origin := os.Getenv("ORIGIN")
			origins := strings.Split(origin, ",")
			c := cors.New(cors.Options{
				AllowedOrigins:   origins,
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"authorization", "content-type", "x-csrf-token", "subscription-id"},
				ExposedHeaders:   []string{"x-csrf-token"},
				AllowCredentials: true,
			})
			handler := c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				f(w, r)
			}))
			handler.ServeHTTP(w, r)
		}
	}
}
