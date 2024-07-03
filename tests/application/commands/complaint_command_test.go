package commands_test

import (
	"context"
	"go-complaint/application/application_services"
	"go-complaint/application/commands"
	"go-complaint/tests"
	"testing"
)

func TestComplaintCommandSendNew(t *testing.T) {
	ctx := context.Background()
	command := commands.ComplaintCommand{
		AuthorID:    tests.UserRegisterAndVerifyCommands["1"].Email,
		ReceiverID:  tests.UserRegisterAndVerifyCommands["2"].Email,
		Title:       "titletitletitletitletitletitletitletitletitle",
		Description: "descriptiondescriptiondescriptiondescriptiondescriptiondescriptiondescription",
		Content:     "contentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontent",
	}
	err := command.SendNew(ctx)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}

func TestComplaintCommandReply(t *testing.T) {
	ctx := context.Background()
	command := commands.ComplaintCommand{
		ReplyAuthorID: tests.UserRegisterAndVerifyCommands["1"].Email,
		ReplyBody:     tests.RepeatString("ja", 10),
		ID:            "8e184d43-6eae-40ad-af72-51c74fc0f2c9",
	}
	err := command.Reply(ctx)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestComplaintCommandSendForReview(t *testing.T) {
	ctx := context.Background()
	command := commands.ComplaintCommand{
		ID:     "8e184d43-6eae-40ad-af72-51c74fc0f2c9",
		UserID: tests.UserRegisterAndVerifyCommands["2"].Email,
	}
	err := command.SendForReviewing(ctx)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestComplaintCommandRate(t *testing.T) {
	ctx := context.Background()
	command := commands.ComplaintCommand{
		ID:      "8e184d43-6eae-40ad-af72-51c74fc0f2c9",
		UserID:  tests.UserRegisterAndVerifyCommands["2"].Email,
		Rating:  5,
		Comment: "tea is good",
		EventID: "94585dff-b2e3-4ae4-94f6-f7dac97d3a39",
	}
	ctx, err := application_services.AuthorizationApplicationServiceInstance().Authorize(
		ctx,
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJlcmNobzAwMUBnbWFpbC5jb20iLCJmdWxsX25hbWUiOiJGZXJuYW5kbyBBZ3VzdGluIFNvcmlhIiwicHJvZmlsZV9pbWciOiIvZGVmYXVsdC5qcGciLCJnZW5kZXIiOiJNQUxFIiwicHJvbm91biI6IkhlIiwiY2xpZW50X2RhdGEiOnsiaXAiOiIiLCJkZXZpY2UiOiIiLCJnZW9sb2NhbGl6YXRpb24iOnsibGF0aXR1ZGUiOjAsImxvbmdpdHVkZSI6MH0sImxvZ2luX2RhdGUiOiIxNzE4Mzk3NTI0MzE3In0sInJlbWVtYmVyX21lIjp0cnVlLCJhdXRob3JpdGllcyI6bnVsbCwiZXhwIjoxNzE4NDgzOTI0LCJpYXQiOjE3MTgzOTc1MjR9.oPMec11HwoN_85ZmeKeC99LcTJS0-3EBZx4uIA_yCA0",
	)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	err = command.RateComplaint(ctx)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
