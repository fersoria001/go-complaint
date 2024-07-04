package infrastructure_test

import (
	"context"
	"go-complaint/application/commands"
	"testing"
)

func TestEmailServiceSendBackground(t *testing.T) {
	ctx := context.Background()
	c := commands.SendEmailCommand{
		ToEmail:           "bercho001@gmail.com",
		ToName:            "fer",
		ConfirmationToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImJlcmNobzAwMUBnbWFpbC5jb20iLCJDb2RlIjozOTU3MTQ5LCJleHAiOjE3MjAyMTkzNzEsImlhdCI6MTcyMDEzMjk3MX0.kXQRqBi98bMmhm02gmpp7iEOXZXDuXgneCZSlMw0DPA",
		ConfirmationCode:  3463748,
	}
	c.Welcome(ctx)
}
