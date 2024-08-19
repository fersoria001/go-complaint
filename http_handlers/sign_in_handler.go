package http_handlers

import (
	"encoding/json"
	"fmt"
	"go-complaint/application/queries"
	"net/http"
	"net/mail"
	"os"
	"time"
)

type AuthenticationRequest struct {
	Username   string
	Password   string
	RememberMe bool
}

func (ar *AuthenticationRequest) valid() error {
	if _, err := mail.ParseAddress(ar.Username); err != nil {
		return fmt.Errorf("invalid username, it must be an email address")
	}
	if len(ar.Password) < 8 {
		return fmt.Errorf("invalid password, it must have at least 8 characters length")
	}
	return nil
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var p AuthenticationRequest
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	r.Body.Close()
	if err := p.valid(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	q := queries.NewSignInQuery(p.Username, p.Password, p.RememberMe)
	token, err := q.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	value := token.Token
	raw := fmt.Sprintf("jwt=%s", value)
	environment := os.Getenv("ENVIRONMENT")
	secure := false
	httpOnly := true
	domain := ""
	if environment == "PROD" {
		domain = "go-complaint.com"
		secure = true
		httpOnly = false
	}
	maxAge := time.Second * 60 * 60 * 24 * 7
	cookie := &http.Cookie{
		Name:       "jwt",
		Value:      value,
		Path:       "/",
		Domain:     domain,
		Expires:    time.Now().Add(time.Hour * 24 * 7),
		RawExpires: "",
		MaxAge:     int(maxAge.Seconds()),
		Secure:     secure,
		HttpOnly:   httpOnly,
		SameSite:   http.SameSiteLaxMode,
		Raw:        raw,
		Unparsed:   []string{raw},
	}
	if err = cookie.Valid(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(200)
}
