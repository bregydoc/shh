package main

import (
	"crypto/rand"
	"errors"
	"os"
	"time"

	"github.com/bregydoc/shh"
	"github.com/bregydoc/shh/stores/redis"
)

func NewDefaultWizard() (*shh.Wizard, error) {
	configFile := os.Getenv("CONFIG_FILE")
	var conf *shh.Config
	var err error
	if configFile != "" {
		if conf, err = shh.LoadConfig(configFile); err != nil {
			return nil,  err
		}
	}

	if conf, err = shh.LoadConfig(); err != nil {
		return nil,  err
	}

	if conf.StoreBackend.Type != "redis" {
		return nil,  errors.New("invalid store backend, currently SHH only have a redis store backend")
	}

	store, err := redis.NewStore(
		redis.WithAddress(conf.StoreBackend.Address),
		redis.WithExpirationTime(5*time.Minute),
	)
	if err != nil {
		return nil,  err
	}

	wizard, err := shh.NewWizard(
		shh.WithConfigAPI("/api", conf.APIPort),
		shh.WithConfigRPC(conf.RPCPort),
		shh.WithRandomSource(rand.Reader),
		shh.WithStore(store),
	)

	if err != nil {
		return nil,  err
	}

	return wizard, nil
}
