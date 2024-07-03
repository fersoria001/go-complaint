package queue

import "fmt"

var ErrEmptyQueue = fmt.Errorf("queue is empty")
var ErrCacheMiss = fmt.Errorf("cache miss")
var ErrInterfaceAssertion = fmt.Errorf("error type assertion")
