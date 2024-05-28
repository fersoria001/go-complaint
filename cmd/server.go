package cmd

import (
	"log"
	"net/http"
)

type Server struct {
	configuration *Configuration
	mux           *http.ServeMux
}

func NewServer(options ...OptionsServerFunc) (*Server, error) {
	config, err := NewConfiguration()
	if err != nil {
		return nil, err
	}
	s := &Server{
		configuration: config,
		mux:           http.NewServeMux(),
	}
	for _, option := range options {
		err := option(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (s *Server) Run() error {
	log.Printf("Server is running on %s", s.configuration.String())
	return http.ListenAndServe(s.configuration.String(), s.mux)
}

type OptionsServerFunc func(c *Server) error

func WithConfiguration(configuration *Configuration) OptionsServerFunc {
	return func(c *Server) error {
		c.configuration = configuration
		return nil
	}
}

func WithHandler(path string, handler http.Handler) OptionsServerFunc {
	return func(c *Server) error {
		c.mux.Handle(path, handler)
		return nil
	}
}
