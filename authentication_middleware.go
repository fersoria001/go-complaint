package main

import (
	"fmt"
	"go-complaint/application/application_services"
	"net/http"
	"strings"
)

// it checks the cookie "jwt"
func tokenFromRequest(r *http.Request) (string, error) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(cookie.Value, "Bearer ") {
		return cookie.Value[7:], nil
	}
	if strings.HasPrefix(cookie.Value, "Bearer%20") {
		return cookie.Value[9:], nil
	}
	return "", fmt.Errorf("cookie has a value that is not the expected token format")
}

func AuthenticationMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtToken, err := tokenFromRequest(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			svc := application_services.AuthorizationApplicationServiceInstance()
			authorized, err := svc.Authorize(r.Context(), jwtToken)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			r = r.WithContext(authorized)
			next.ServeHTTP(w, r)
		})
	}
}
