package application_test

import (
	"context"
	"go-complaint/application"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/tests"
	"testing"
)

func TestComplaintService(t *testing.T) {
	ctx := context.Background()
	complaintRepository := repositories.NewComplaintRepository(datasource.ComplaintSchema())
	repliesRepository := repositories.NewReplyRepository(datasource.ComplaintSchema())
	complaintService := application.NewComplaintService(complaintRepository, repliesRepository)

	t.Run("Send a complaint", func(t *testing.T) {
		for _, info := range complaintInfo {
			err := complaintService.CreateComplaint(
				ctx, info.authorID, info.receiverID, info.title, info.description, info.content)
			if err != nil {
				t.Errorf("error: %v", err)
			}
		}
	})
	t.Run("Get complaints from author", func(t *testing.T) {
		for _, info := range complaintInfo {
			complaints, err := complaintService.GetComplaintsFrom(ctx, info.authorID, "", 10, 0)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if len(complaints.Complaints) == 0 {
				t.Errorf("expected at least one complaint")
			}
		}
	})

	t.Run("Get complaints from receiver", func(t *testing.T) {
		for _, info := range complaintInfo {
			complaints, err := complaintService.GetComplaintsTo(ctx, info.receiverID, "", 10, 0)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if len(complaints.Complaints) == 0 {
				t.Errorf("expected at least one complaint")
			}
		}
	})

	t.Run(`Return the draft get an ID, reply the complaint
	and return the single complaint with its replies
	if its the first reply the complaint status will change to STARTED`,
		func(t *testing.T) {
			anUseCase := complaintInfo["userEnterprise"]
			complaintsDtos, err := complaintService.GetComplaintsTo(ctx, anUseCase.receiverID, "", 10, 0)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if len(complaintsDtos.Complaints) == 0 {
				t.Errorf("expected at least one complaint")
			}

			if len(complaintsDtos.Complaints) > 0 {
				dtoWithoutReplies := complaintsDtos.Complaints[0]
				err := complaintService.ReplyComplaint(
					ctx,
					dtoWithoutReplies.ID.String(),
					"ReceiverPROFILEIMG.jpg",
					"Receiver FullName",
					dtoWithoutReplies.ReceiverID,
					"This is the body of the reply",
				)
				if err != nil {
					t.Errorf("error: %v", err)
				}
				complaintDB, err := complaintService.Complaint(
					ctx,
					dtoWithoutReplies.ID.String(),
				)

				if err != nil {
					t.Errorf("error: %v", err)
				}
				if complaintDB == nil {
					t.Errorf("expected complaint")
				}
				if complaintDB != nil && complaintDB.Replies == nil {
					t.Errorf("expected replies")
				}
				if complaintDB != nil && complaintDB.Replies != nil && len(complaintDB.Replies) != 1 {
					t.Errorf("expected one reply")
				}
				if complaintDB.Status != complaint.STARTED.String() {
					t.Errorf("expected complaint status to be STARTED, got %v", complaintDB.Status)
				}
			}
		})
	t.Run(`A STARTED complaint can only be replied by the author`, func(t *testing.T) {
		anUseCase := complaintInfo["userEnterprise"]
		complaintsDtos, err := complaintService.GetComplaintsTo(ctx, anUseCase.receiverID, "", 10, 0)
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if len(complaintsDtos.Complaints) == 0 {
			t.Errorf("expected at least one complaint")
		}

		if len(complaintsDtos.Complaints) > 0 {
			dtoWithoutReplies := complaintsDtos.Complaints[0]
			if dtoWithoutReplies.Status != complaint.STARTED.String() {
				t.Errorf("expected complaint status to be STARTED, got %v", dtoWithoutReplies.Status)
			}
			err := complaintService.ReplyComplaint(
				ctx,
				dtoWithoutReplies.ID.String(),
				"ReceiverPROFILEIMG.jpg",
				"Receiver FullName",
				dtoWithoutReplies.ReceiverID,
				"This is the body of the reply",
			)
			if err == nil {
				t.Errorf("error: %v", err)
			}
			complaintDB, err := complaintService.Complaint(
				ctx,
				dtoWithoutReplies.ID.String(),
			)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if complaintDB == nil {
				t.Errorf("expected complaint")
			}
			if complaintDB != nil && complaintDB.Replies == nil {
				t.Errorf("expected replies")
			}
			if complaintDB != nil && complaintDB.Replies != nil && len(complaintDB.Replies) != 1 {
				t.Errorf("expected one reply")
			}
			if complaintDB.Status != complaint.STARTED.String() {
				t.Errorf("expected complaint status to be STARTED, got %v", complaintDB.Status)
			}
		}
	})

	t.Run(`Get the sent complaint that has STARTED
		and reply it, the complaint status should change to IN_DISCUSSION`, func(t *testing.T) {
		anUseCase := complaintInfo["userEnterprise"]
		complaintsDtos, err := complaintService.GetComplaintsFrom(ctx, anUseCase.authorID, "STARTED", 10, 0)
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if len(complaintsDtos.Complaints) == 0 {
			t.Errorf("expected at least one complaint")
		}
		if len(complaintsDtos.Complaints) > 0 {

			dtoWithoutReplies := complaintsDtos.Complaints[0]
			if dtoWithoutReplies.Status != complaint.STARTED.String() {
				t.Errorf("expected complaint status to be STARTED, got %v", dtoWithoutReplies.Status)
			}
			err := complaintService.ReplyComplaint(
				ctx,
				dtoWithoutReplies.ID.String(),
				"AuthorPROFILEIMG.jpg",
				"Author FullName",
				dtoWithoutReplies.AuthorID,
				"This is the body of the reply",
			)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			complaintDB, err := complaintService.Complaint(
				ctx,
				dtoWithoutReplies.ID.String(),
			)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if complaintDB == nil {
				t.Errorf("expected complaint")
			}
			if complaintDB != nil && complaintDB.Replies == nil {
				t.Errorf("expected replies")
			}
			if complaintDB != nil && complaintDB.Replies != nil && len(complaintDB.Replies) != 2 {
				t.Errorf("expected two reply")
			}
			if complaintDB.Status != complaint.IN_DISCUSSION.String() {
				t.Errorf("expected complaint status to be IN_DISCUSSION, got %v", complaintDB.Status)
			}

		}
	})

	t.Run(`An OPEN complaint can only be replied by the receiver, we has taken another use case here`,
		func(t *testing.T) {
			anUseCase := complaintInfo["userUser"]
			complaintsDtos, err := complaintService.GetComplaintsTo(ctx, anUseCase.receiverID, "OPEN", 10, 0)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if len(complaintsDtos.Complaints) == 0 {
				t.Errorf("expected at least one complaint")
			}
			if len(complaintsDtos.Complaints) > 0 {
				dtoWithoutReplies := complaintsDtos.Complaints[0]
				if dtoWithoutReplies.Status != complaint.OPEN.String() {
					t.Errorf("expected complaint status to be OPEN, got %v", dtoWithoutReplies.Status)
				}
				err := complaintService.ReplyComplaint(
					ctx,
					dtoWithoutReplies.ID.String(),
					"AuthorPROFILEIMG.jpg",
					"Author FullName",
					dtoWithoutReplies.AuthorID,
					"This is the body of the reply",
				)
				if err == nil {
					t.Errorf("expected error")
				}
				complaintDB, err := complaintService.Complaint(
					ctx,
					dtoWithoutReplies.ID.String(),
				)
				if err != nil {
					t.Errorf("error: %v", err)
				}
				if complaintDB == nil {
					t.Errorf("expected complaint")
				}
				if complaintDB != nil && len(complaintDB.Replies) > 0 {
					t.Errorf("expected  no replies")
				}
				if complaintDB.Status != complaint.OPEN.String() {
					t.Errorf("expected complaint status to be OPEN, got %v", complaintDB.Status)
				}
			}
		})

	t.Run(`A IN_DISCUSSION complaint can be answered by author OR receiver`, func(t *testing.T) {
		anUseCase := complaintInfo["userEnterprise"]
		complaintsDtos, err := complaintService.GetComplaintsTo(ctx, anUseCase.receiverID, "", 10, 0)
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if len(complaintsDtos.Complaints) == 0 {
			t.Errorf("expected at least one complaint")
		}
		if len(complaintsDtos.Complaints) > 0 {
			dtoWithoutReplies := complaintsDtos.Complaints[0]
			if dtoWithoutReplies.Status != complaint.IN_DISCUSSION.String() {
				t.Errorf("expected complaint status to be IN_DISCUSSION, got %v", dtoWithoutReplies.Status)
			}
			err := complaintService.ReplyComplaint(
				ctx,
				dtoWithoutReplies.ID.String(),
				"ReceiverPROFILEIMG.jpg",
				"Receiver FullName",
				dtoWithoutReplies.ReceiverID,
				"This is the body of the reply",
			)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			complaintDB, err := complaintService.Complaint(
				ctx,
				dtoWithoutReplies.ID.String(),
			)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if complaintDB == nil {
				t.Errorf("expected complaint")
			}
			if complaintDB != nil && complaintDB.Replies == nil {
				t.Errorf("expected replies")
			}
			if complaintDB != nil && complaintDB.Replies != nil && len(complaintDB.Replies) != 3 {
				t.Errorf("expected three replies")
			}
			if complaintDB.Status != complaint.IN_DISCUSSION.String() {
				t.Errorf("expected complaint status to be IN_DISCUSSION, got %v", complaintDB.Status)
			}
			err = complaintService.ReplyComplaint(
				ctx,
				dtoWithoutReplies.ID.String(),
				"ReceiverPROFILEIMG.jpg",
				"Receiver FullName",
				dtoWithoutReplies.ReceiverID,
				"This is the body of the reply",
			)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			err = complaintService.ReplyComplaint(
				ctx,
				dtoWithoutReplies.ID.String(),
				"AUTHORPROFILEIMG.jpg",
				"AUTHOR FullName",
				dtoWithoutReplies.AuthorID,
				"This is the body of the reply",
			)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			complaintDB, err = complaintService.Complaint(
				ctx,
				dtoWithoutReplies.ID.String(),
			)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if complaintDB == nil {
				t.Errorf("expected complaint")
			}
			if complaintDB != nil && complaintDB.Replies == nil {
				t.Errorf("expected replies")
			}
			if complaintDB != nil && complaintDB.Replies != nil && len(complaintDB.Replies) != 5 {
				t.Errorf("expected five replies")
			}
			if complaintDB.Status != complaint.IN_DISCUSSION.String() {
				t.Errorf("expected complaint status to be IN_DISCUSSION, got %v", complaintDB.Status)
			}
		}
	})

	t.Run(`An IN_DISCUSSION enterprise complaint can't be closed directly by anyone`,
		func(t *testing.T) {
			anUseCase := complaintInfo["userEnterprise"]
			complaintsDtos, err := complaintService.GetComplaintsTo(ctx, anUseCase.receiverID, "IN_DISCUSSION", 10, 0)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if len(complaintsDtos.Complaints) == 0 {
				t.Errorf("expected at least one complaint")
			}
			if len(complaintsDtos.Complaints) > 0 {
				dtoWithoutReplies := complaintsDtos.Complaints[0]
				err := complaintService.Close(
					ctx,
					dtoWithoutReplies.ID.String(),
					tests.NewEmployeeID(anUseCase.authorID, "managerid@gmail.com", "thisEmployeeUserID@gmail.com"),
				)
				if err == nil {
					t.Errorf("expected error")
				}
				complaintDB, err := complaintService.Complaint(
					ctx,
					dtoWithoutReplies.ID.String(),
				)
				if err != nil {
					t.Errorf("error: %v", err)
				}
				if complaintDB == nil {
					t.Errorf("expected complaint")
				}
				if complaintDB != nil && complaintDB.Status != complaint.IN_DISCUSSION.String() {
					t.Errorf("expected complaint status to be IN_DISCUSSION, got %v", complaintDB.Status)
				}
			}
		})

	t.Run(`An IN_DISCUSSION complaint can be send to be review by the enterprise ONLY by an employee of the enterprise`,
		func(t *testing.T) {
			indiscussionUseCase := complaintInfo["userEnterprise"]
			complaintsDtos, err := complaintService.GetComplaintsTo(
				ctx,
				indiscussionUseCase.receiverID,
				"IN_DISCUSSION",
				10, 0,
			)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if len(complaintsDtos.Complaints) == 0 {
				t.Errorf("expected at least one complaint")
			}
			if len(complaintsDtos.Complaints) > 0 {
				dtoWithoutReplies := complaintsDtos.Complaints[0]
				if dtoWithoutReplies.Status != complaint.IN_DISCUSSION.String() {
					t.Errorf("expected complaint status to be IN_DISCUSSION, got %v", dtoWithoutReplies.Status)
				}
				err := complaintService.SendForReviewing(
					ctx,
					dtoWithoutReplies.ID.String(),
					"enterprise",
				)
				if err == nil {
					t.Errorf("expected error")
				}
				aValidEmployeeID := tests.NewEmployeeID(indiscussionUseCase.receiverID, "manager@gmail.com", "thisEmployeeUserID@gmail.com")
				err = complaintService.SendForReviewing(
					ctx,
					dtoWithoutReplies.ID.String(),
					aValidEmployeeID,
				)
				if err != nil {
					t.Errorf("error: %v", err)
					t.Errorf("ArgumentID: %v, complaint authorID: %v, complaint receiverID: %v",
						aValidEmployeeID, indiscussionUseCase.authorID, indiscussionUseCase.receiverID)
				}
				complaintDB, err := complaintService.Complaint(
					ctx,
					dtoWithoutReplies.ID.String(),
				)
				if err != nil {
					t.Errorf("error: %v", err)
				}
				if complaintDB == nil {
					t.Errorf("expected complaint")
				}
				if complaintDB != nil && complaintDB.Status != complaint.IN_REVIEW.String() {
					t.Errorf("expected complaint status to be IN_REVIEW, got %v", complaintDB.Status)
				}
			}
		})

	t.Run(`An User complaint can't be send to review`,
		func(t *testing.T) {
			anUseCase := complaintInfo["userUser"]
			complaintsDtos, err := complaintService.GetComplaintsTo(ctx, anUseCase.receiverID, "", 10, 0)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if len(complaintsDtos.Complaints) == 0 {
				t.Errorf("expected at least one complaint")
			}
			if len(complaintsDtos.Complaints) > 0 {
				dtoWithoutReplies := complaintsDtos.Complaints[0]
				err := complaintService.ReplyComplaint(
					ctx,
					dtoWithoutReplies.ID.String(),
					"ReceiverPROFILEIMG.jpg",
					"Receiver FullName",
					dtoWithoutReplies.ReceiverID,
					"This is the body of the reply",
				)
				if err != nil {
					t.Errorf("error: %v", err)
				}
				err = complaintService.ReplyComplaint(
					ctx,
					dtoWithoutReplies.ID.String(),
					"Author.jpg",
					"Author FullName",
					dtoWithoutReplies.AuthorID,
					"This is the body of the reply",
				)
				if err != nil {
					t.Errorf("error: %v", err)
				}
				err = complaintService.SendForReviewing(
					ctx,
					dtoWithoutReplies.ID.String(),
					anUseCase.authorID,
				)
				if err == nil {
					t.Errorf("expected error")
				}
				complaintDB, err := complaintService.Complaint(
					ctx,
					dtoWithoutReplies.ID.String(),
				)
				if err != nil {
					t.Errorf("error: %v", err)
				}
				if complaintDB == nil {
					t.Errorf("expected complaint")
				}
				if complaintDB != nil && complaintDB.Status != complaint.IN_DISCUSSION.String() {
					t.Errorf("expected complaint status to be IN_DISCUSSION, got %v", complaintDB.Status)
				}
			}

		})

	t.Run(`An user complaint can be closed ONLY by the author, without going trough review`,
		func(t *testing.T) {
			anUseCase := complaintInfo["userUser"]
			complaintsDtos, err := complaintService.GetComplaintsTo(ctx, anUseCase.receiverID, "IN_DISCUSSION", 10, 0)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if len(complaintsDtos.Complaints) == 0 {
				t.Errorf("expected at least one complaint")
			}
			if len(complaintsDtos.Complaints) > 0 {
				dtoWithoutReplies := complaintsDtos.Complaints[0]
				err := complaintService.Close(
					ctx,
					dtoWithoutReplies.ID.String(),
					"anyotherID@gmail.com",
				)
				if err == nil {
					t.Errorf("expected error")
				}
				err = complaintService.Close(
					ctx,
					dtoWithoutReplies.ID.String(),
					anUseCase.authorID,
				)
				if err != nil {
					t.Errorf("error: %v", err)
				}
				complaintDB, err := complaintService.Complaint(
					ctx,
					dtoWithoutReplies.ID.String(),
				)
				if err != nil {
					t.Errorf("error: %v", err)
				}
				if complaintDB == nil {
					t.Errorf("expected complaint")
				}
				if complaintDB != nil && complaintDB.Status != complaint.CLOSED.String() {
					t.Errorf("expected complaint status to be CLOSED, got %v", complaintDB.Status)
				}
			}
		})

	t.Cleanup(func() {
		cs, err := complaintRepository.GetAll(ctx)
		if err != nil {
			t.Errorf("error: %v", err)
		}
		for iterObj := range cs.Iter() {
			c, err := complaintRepository.Get(ctx, iterObj.ID().String())
			if err != nil {
				t.Errorf("error: %v", err)
			}
			rs, err := repliesRepository.FindByComplaintID(ctx, c.ID().String())
			if err != nil {
				t.Errorf("error: %v", err)
			}
			for _, r := range rs {
				err = repliesRepository.Remove(ctx, r.ID().String())
				if err != nil {
					t.Errorf("error: %v", err)
				}
			}
			err = complaintRepository.Remove(ctx, c.ID().String())
			if err != nil {
				t.Errorf("error: %v", err)
			}
		}

	})
}
