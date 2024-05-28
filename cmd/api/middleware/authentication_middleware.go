package middleware

import (
	"context"
	"errors"
	"go-complaint/application"
	"go-complaint/erros"
	"net/http"
)

type AuthContextKey struct {
	name string
}

var AuthCtxKey = AuthContextKey{"user_email"}

type TokenContextKey struct {
	token string
}

var TokenCtxKey = TokenContextKey{"jwt_token"}

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
			tokenStr, err := GetTokenFromRequest(r)
			if err != nil {
				next(w, r)
				return
			}

			claims, err := application.NewJWTService().ParseUserDescriptor(tokenStr)
			//this can lead to errors if the token expired at serverside and client still has it
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			ctxKey := AuthCtxKey
			ctx := context.WithValue(r.Context(), ctxKey, &claims.Email)
			ctx = context.WithValue(ctx, TokenCtxKey, tokenStr)
			r = r.WithContext(ctx)
			next(w, r)
		}

		return handler
	}

	return middleware
}

func WebsocketAuthenticationMiddleware() Middleware {

	middleware := func(next http.HandlerFunc) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			tokenStr, err := GetTokenFromWebsocketRequest(r)
			if err != nil {
				next(w, r)
				return
			}

			claims, err := application.NewJWTService().ParseUserDescriptor(tokenStr)
			//this can lead to errors if the token expired at serverside and client still has it
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			ctxKey := AuthCtxKey
			ctx := context.WithValue(r.Context(), ctxKey, &claims.Email)
			ctx = context.WithValue(ctx, TokenCtxKey, tokenStr)
			r = r.WithContext(ctx)
			next(w, r)
		}

		return handler
	}

	return middleware
}

func GetContextPersonID(ctx context.Context) (string, error) {
	key := AuthCtxKey
	raw, found := ctx.Value(key).(*string)
	if !found {
		return "", &erros.UnauthorizedError{}
	}
	return *raw, nil
}

func GetContextToken(ctx context.Context) (string, error) {
	key := TokenCtxKey
	raw, found := ctx.Value(key).(string)
	if !found {
		return "", &erros.UnauthorizedError{}
	}
	return raw, nil
}
