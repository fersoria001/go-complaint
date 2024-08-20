package http_handlers

import (
	"fmt"
	"go-complaint/application/commands"
	"net/http"
	"os"
)

func EmailValidationHandler(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("token") {
		http.Error(w, fmt.Errorf("error: token is not present in the get request query").Error(), http.StatusUnauthorized)
	}
	token := r.URL.Query().Get("token")
	c := commands.NewVerifyUserEmailCommand(token)
	err := c.Execute(r.Context())
	if err != nil {
		http.Error(w, fmt.Errorf("error at verifyUserEmailCommand %w", err).Error(), http.StatusInternalServerError)
	}
	rootUrl := os.Getenv("FRONT_END_URL")
	signInUrl := fmt.Sprintf("%s/%s", rootUrl, "sign-in")
	http.Redirect(w, r, signInUrl, http.StatusSeeOther)
}
