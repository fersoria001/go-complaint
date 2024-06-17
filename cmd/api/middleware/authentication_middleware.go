package middleware

import (
	"errors"
	"go-complaint/application/application_services"
	"net/http"
)

func GetTokenFromRequest(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		header, err := r.Cookie("Authorization")
		if err != nil {
			return "", err
		}

		return checkAndSliceHeader(header.Value)
	}
	return checkAndSliceHeader(header)
}

func GetTokenFromWebsocketRequest(r *http.Request) (string, error) {
	header := r.Header.Get("Sec-WebSocket-Protocol")
	if header == "" {
		header, err := r.Cookie("Sec-WebSocket-Protocol")
		if err != nil {
			return "", err
		}

		return checkAndSliceHeader(header.Value)
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
func AuthenticationMiddleware() Middleware {

	middleware := func(next http.HandlerFunc) http.HandlerFunc {

		handler := func(w http.ResponseWriter, r *http.Request) {

			jwtToken, err := GetTokenFromRequest(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
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
			next(w, r)
		}

		return handler
	}
	return middleware
}
