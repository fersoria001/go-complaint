package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func CORS() Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			c := cors.New(cors.Options{
				AllowedOrigins:   []string{"http://localhost:5173"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"Authorization", "Content-Type"},
				AllowCredentials: true,
			})
			handler := c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				f(w, r)
			}))

			handler.ServeHTTP(w, r)
		}
	}
}
