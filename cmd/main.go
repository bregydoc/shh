package main

import (
	"crypto/rand"
	"time"

	"github.com/bregydoc/shh"
	"github.com/bregydoc/shh/stores/redis"
)

func main() {
	store, err := redis.NewStore(
		redis.WithAddress("127.0.0.1:6379"),
		redis.WithExpirationTime(5*time.Minute),
	)
	if err != nil {
		panic(err)
	}

	wizard, err := shh.NewWizard(
		shh.WithDefaultAPI(),
		shh.WithRandomSource(rand.Reader),
		shh.WithStore(store),
	)
	if err != nil {
		panic(err)
	}

	errChan := wizard.Run()

	<- errChan
}
