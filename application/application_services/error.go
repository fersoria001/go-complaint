package application_services

import "fmt"

var ErrTokenExpired = fmt.Errorf("token expired")
var ErrInvalidToken = fmt.Errorf("invalid token")
var ErrTokenNotFound = fmt.Errorf("token not found")
var ErrUnauthorized = fmt.Errorf("unauthorized")
var ErrNotFound = fmt.Errorf("not found")
var ErrBadRequest = fmt.Errorf("bad request")
