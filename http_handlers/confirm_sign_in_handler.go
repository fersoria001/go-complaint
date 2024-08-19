package http_handlers

import (
	"encoding/json"
	"fmt"
	"go-complaint/application/application_services"
	"go-complaint/application/queries"
	"net/http"
	"os"
	"time"
)

type ConfirmSignInRequest struct {
	ConfirmationCode int
}

func (csr *ConfirmSignInRequest) valid() error {
	if csr.ConfirmationCode < 1000000 ||
		csr.ConfirmationCode > 9999999 {
		return fmt.Errorf("the code provided is not a valid confirmation code")
	}
	return nil
}

func ConfirmSignInHandler(w http.ResponseWriter, r *http.Request) {
	signInToken, err := application_services.AuthorizationApplicationServiceInstance().JWTToken(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	var p ConfirmSignInRequest
	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	r.Body.Close()
	if err = p.valid(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	q := queries.NewLoginQuery(signInToken, p.ConfirmationCode)
	sessionToken, err := q.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	value := sessionToken.Token
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
