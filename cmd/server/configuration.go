package server

import (
	"fmt"
	"go-complaint/erros"
	"os"
	"strconv"
)

type Configuration struct {
	host string
	port int
}

func NewConfiguration(options ...OptionsConfigurationFunc) (*Configuration, error) {
	portv := os.Getenv("PORT")
	port, err := strconv.Atoi(portv)
	if err != nil {
		return nil, err
	}
	c := &Configuration{
		host: os.Getenv("HOST"),
		port: port,
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

func (c *Configuration) Address() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}

func (c *Configuration) URL() string {
	return fmt.Sprintf("http://%s", c.Address())
}
