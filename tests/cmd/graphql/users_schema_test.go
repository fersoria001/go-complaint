package graphql_test

import (
	"encoding/json"
	"go-complaint/cmd/api/graphql_"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUsersSchema(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(graphql_.ExecuteAndEncodeBody))
	queries := map[string]string{
		"createUserMutation": `
		{
			"query": "mutation { createUser(profileIMG: \"google.com\", 
			email: \"email@gmail.com\",password: \"Password1\", 
			firstName: \"firstName\", lastName: \"firstName\",
			 birthDate: \"1713205999629\", phone:\"01234567890\",
			  country: \"country\", county: \"county\",
			   city: \"city\")}"
		}`,
		"loginUserQuery": `
		{
			"query": "{ Login(email: \"email@gmail.com\", password: \"Password1\") { token } }"
		}`,
	}
	deleteLineBreaks := strings.NewReplacer("\n", "", "\t", "")
	t.Run("a user should be created", func(t *testing.T) {
		query := deleteLineBreaks.Replace(queries["createUserMutation"])
		payload := strings.NewReader(query)
		resp, err := http.Post(testServer.URL, "application/json", payload)
		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status OK, got %v", resp.Status)
		}
	})
	t.Run("a user should be logged in", func(t *testing.T) {
		query := deleteLineBreaks.Replace(queries["loginUserQuery"])
		payload := strings.NewReader(query)
		resp, err := http.Post(testServer.URL, "application/json", payload)
		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status OK, got %v", resp.Status)
		}
		deserializeIntoMap := struct {
			Token string `json:"token"`
		}{}
		err = json.NewDecoder(resp.Body).Decode(&deserializeIntoMap)
		if err != nil {
			t.Error(err)
		}
	})
}
