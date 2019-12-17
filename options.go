package shh

import (
	"io"

	"github.com/gin-gonic/gin"
)

type Option func(wizard *Wizard) error


func WithStore(s Store) Option {
	return func (wizard *Wizard) error {
		wizard.store = s
		return nil
	}
}

func WithRandomSource(reader io.Reader) Option {
	return func (wizard *Wizard) error {
		wizard.random = reader
		return nil
	}
}

func WithDefaultAPI() Option {
	return func (wizard *Wizard) error {
		wizard.api = &API{
			engine:  gin.Default(),
			baseURL: "/api",
			port:    ":8080",
		}
		return nil
	}
}