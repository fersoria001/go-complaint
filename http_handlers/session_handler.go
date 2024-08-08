package http_handlers

import (
	"encoding/json"
	"go-complaint/application/application_services"
	"net/http"
)

func SessionHandler(w http.ResponseWriter, r *http.Request) {
	svc := application_services.AuthorizationApplicationServiceInstance()
	session, err := svc.Credentials(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(session)
}
