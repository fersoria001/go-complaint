package middleware

import (
	"errors"
	"go-complaint/application/application_services"
	"net/http"
)

func GetTokenFromRequest(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			return "", err
		}
		return checkAndSliceHeader(cookie.Value)
	}
	return checkAndSliceHeader(header)
}

func checkAndSliceHeader(header string) (string, error) {
	if startWith(header, "Bearer ") {
		return header[7:], nil
	} else if startWith(header, "Bearer%20") {
		return header[9:], nil
	} else {
		return "", errors.New("invalid bearer token format")
	}
}

func startWith(s string, prefix string) bool {
	if len(s) == 0 || len(prefix) == 0 || len(s) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}
func AuthenticationMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtToken, err := GetTokenFromRequest(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			authorized, err := application_services.AuthorizationApplicationServiceInstance().Authorize(
				r.Context(),
				jwtToken,
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			r = r.WithContext(authorized)
			next.ServeHTTP(w, r)
		})
	}
}
