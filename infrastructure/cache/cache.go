package cache

import (
	"net/mail"
	"slices"
	"strings"
)

type Type int

const (
	READ Type = iota
	WRITE
	DELETE
)

type Request struct {
	Type    Type
	Payload interface{}
	Key     string
	Out     chan Request
}

var RequestChannel = make(chan Request)
var RequestChannel2 = make(chan Request)

func Cache(ch chan Request) {
	var cache = make(map[string]interface{})

	for {
		req := <-ch
		switch req.Type {
		case READ:
			req.Payload = cache[req.Key]
			SendToChannel(req.Out, req)
		case WRITE:
			if req.Key == "ENTERPRISE_USER" {
				p := req.Payload.(*Pair)
				v, ok := cache[p.One]
				s := make([]string, 0)
				if ok {
					s = v.([]string)
					delete(cache, p.One)
				}
				s = append(s, p.Two)
				cache[p.One] = s
			} else {
				_, ok := cache[req.Key]
				if ok {
					delete(cache, req.Key)
				}
				cache[req.Key] = req.Payload
			}
		case DELETE:
			if req.Key == "ENTERPRISE_USER" {
				p := req.Payload.(*Pair)
				v, ok := cache[p.One]
				if ok {
					s := v.([]string)
					delete(cache, p.One)
					s = slices.DeleteFunc(s, func(i string) bool {
						return i == p.Two
					})
					cache[p.One] = s
				}
			}
			delete(cache, req.Key)
		}
	}
}

type Pair struct {
	One string `json:"one"`
	Two string `json:"two"`
}

// return a pair enterpriseID, userID
func RecognizeId(operationID string) *Pair {
	s := strings.Split(operationID, "?")
	for i := range s {
		if strings.HasPrefix(s[i], "chat") {
			_, after, found := strings.Cut(s[i], ":")
			if !found {
				continue
			}
			s1 := strings.Split(after, "=")
			if len(s1) != 2 {
				continue
			}
			enterpriseID := s1[0]
			s2 := strings.Split(s1[1], "#")
			if len(s2) != 1 {
				continue
			}
			userID := s2[0]
			if _, err := mail.ParseAddress(userID); err != nil {
				continue
			}
			return &Pair{
				One: enterpriseID,
				Two: userID,
			}
		}
	}
	return nil
}

func SendToChannel(ch chan Request, req Request) {
	ch <- req
}
