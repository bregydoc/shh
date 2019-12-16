package redis

import "github.com/go-redis/redis"

func NewStore(opts ...Option) (*Store, error) {
	s := new(Store)
	s.address = "localhost:6379"

	for _, o := range opts {
		if err := o(s); err != nil {
			return nil, err
		}
	}

	s.db = redis.NewClient(&redis.Options{
		Addr:               s.address,
		Password:           s.password,
	})

	return s, nil
}