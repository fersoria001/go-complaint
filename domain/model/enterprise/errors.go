package enterprise

import "errors"

var ErrWrongTypeAssertion = errors.New("wrong type assertion")
var ErrUnknownColleagueType = errors.New("unknown colleague type")
var ErrPositionNotExists = errors.New("position not exists")
var ErrForbidden = errors.New("either you don't belong to this enterprise or you don't have the necessary permissions")
