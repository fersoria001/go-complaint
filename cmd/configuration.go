package cmd

import (
	"fmt"
	"go-complaint/erros"
)

type Configuration struct {
	host string
	port int
}

func NewConfiguration(options ...OptionsConfigurationFunc) (*Configuration, error) {
	c := &Configuration{
		host: "localhost",
		port: 8080,
	}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

type OptionsConfigurationFunc func(co *Configuration) error

func WithHost(host string) OptionsConfigurationFunc {
	return func(co *Configuration) error {
		co.host = host
		return nil
	}
}

func WithPort(port int) OptionsConfigurationFunc {
	return func(co *Configuration) error {
		if port >= 65535 || port <= 0 {
			return &erros.InvalidPortError{Port: port}
		}
		co.port = port
		return nil
	}
}

func (c *Configuration) String() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}
