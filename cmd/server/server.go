package server

import (
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

type GoComplaintServer struct {
	configuration *Configuration
	mux           *http.ServeMux
}

func NewGoComplaintServer(options ...OptionsServerFunc) (*GoComplaintServer, error) {
	config, err := NewConfiguration()
	if err != nil {
		return nil, err
	}
	mux := http.NewServeMux()
	s := &GoComplaintServer{
		configuration: config,
		mux:           mux,
	}
	for _, option := range options {
		err := option(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (s *GoComplaintServer) URL() string {
	return s.configuration.URL()
}

func (s *GoComplaintServer) Run() {
	log.Printf("Server is running on %s", s.configuration.URL())
	http.ListenAndServe("localhost:8080", csrf.Protect([]byte("32-byte-long-auth-key"))(s.mux))
}

type OptionsServerFunc func(c *GoComplaintServer) error

func WithConfiguration(configuration *Configuration) OptionsServerFunc {
	return func(c *GoComplaintServer) error {
		c.configuration = configuration
		return nil
	}
}

func WithHandlerFunc(path string, handler func(http.ResponseWriter, *http.Request)) OptionsServerFunc {
	return func(c *GoComplaintServer) error {
		c.mux.HandleFunc(path, handler)
		return nil
	}
}

func WithHandler(path string, handler http.Handler) OptionsServerFunc {
	return func(c *GoComplaintServer) error {
		c.mux.Handle(path, handler)
		return nil
	}
}
