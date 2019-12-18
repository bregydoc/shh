package shh

import (
	"io"

	"github.com/gin-gonic/gin"
)

type Option func(wizard *Wizard) error

func WithStore(s Store) Option {
	return func(wizard *Wizard) error {
		wizard.store = s
		return nil
	}
}

func WithRandomSource(reader io.Reader) Option {
	return func(wizard *Wizard) error {
		wizard.random = reader
		return nil
	}
}

func WithDefaultAPI() Option {
	return func(wizard *Wizard) error {
		wizard.api = &API{
			engine:  gin.Default(),
			baseURL: "/api",
			port:    ":8080",
		}
		return nil
	}
}

func WithConfigAPI(prefixURL string, port string, e ...*gin.Engine) Option {
	return func(wizard *Wizard) error {
		wizard.api = &API{
			baseURL: prefixURL,
			port:    port,
		}

		if len(e) > 0 {
			wizard.api.engine = e[0]
		} else {
			wizard.api.engine = gin.Default()
		}

		return nil
	}
}

func WithConfigRPC(port string) Option {
	return func(wizard *Wizard) error {
		wizard.rpcPort = port
		return nil
	}
}


func WithFullAvailableAPI() Option {
	return func(wizard *Wizard) error {
		wizard.fullAvailableAPI = true
		return nil
	}
}
