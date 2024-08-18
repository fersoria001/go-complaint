package commands

import "fmt"

var ErrChatAlreadyExists = fmt.Errorf("chat already exists")
var ErrFeedbackAlreadyExists = fmt.Errorf("feedback with (complaintId,enterpriseId) pair already exists")
var ErrEnterpriseActivityAlreadyExists = fmt.Errorf("an activityId userId pair is unique and it already exists")
var ErrComplaintAlreadyExists = fmt.Errorf("there's already a complaint in the db with status writing that match the author and receiver id")
var ErrBadRequest = fmt.Errorf("bad request")
var ErrConfirmationNotFound = fmt.Errorf("confirmation not found")
var ErrWrongTypeAssertion = fmt.Errorf("wrong type assertion")
var ErrAlreadyVerified = fmt.Errorf("already verified")
var ErrForbidden = fmt.Errorf("forbidden request")
var ErrUnauthorized = fmt.Errorf("unauthorized request")
var ErrAlreadyHired = fmt.Errorf("already hired")
var ErrNotFound = fmt.Errorf("not found")
