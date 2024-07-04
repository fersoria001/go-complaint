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
		ConfirmationToken: "89175",
	}
	c.Welcome(ctx)
}
