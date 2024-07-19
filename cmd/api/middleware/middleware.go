package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Middleware1 func(http.Handler) http.Handler

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
