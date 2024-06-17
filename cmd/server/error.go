package server

import (
	"fmt"
)

var ErrMessageTypeUnsupported = fmt.Errorf("message type unsupported")
var ErrWrongTypeAssertion = fmt.Errorf("wrong type assertion")
var ErrForbidden = fmt.Errorf("forbidden")
var ErrSubscriberDisconnected = fmt.Errorf("client disconnected")
var ErrAuthenticationFailed = fmt.Errorf("authentication failed")
var ErrUnmarshalFailed = fmt.Errorf("deserialization failed")
var ErrMarshalFailed = fmt.Errorf("serialization failed")
var ErrContextDone = fmt.Errorf("context done")
var ErrSubscriberNotFound = fmt.Errorf("subscriber not found")
