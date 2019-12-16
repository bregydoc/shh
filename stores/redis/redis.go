package redis

import (
	"time"

	"github.com/bregydoc/shh"
	"github.com/go-redis/redis"
)

type Store struct {
	address string
	password string
	db *redis.Client

	pairExpiration time.Duration

	registeredCallbacks map[string][]Callback
}

func (s *Store) RegisterNewPair(token string, pair *shh.Pair) error {
	private, err := encrypt([]byte(pair.PrivateKeyPem), []byte(token))
	if err != nil {
		return err
	}
	return s.db.Set(pair.PublicKeyPem, private, s.pairExpiration).Err()
}

func (s *Store) ObservePublic(token, publicKey string, callback func(e *shh.PairEvent)) error {
	if s.registeredCallbacks == nil {
		s.registeredCallbacks = map[string][]Callback{}
	}

	if s.registeredCallbacks[publicKey] == nil {
		s.registeredCallbacks[publicKey] = make([]Callback, 0)
	}

	s.registeredCallbacks[publicKey] = append(s.registeredCallbacks[publicKey], callback)

	return nil
}

func (s *Store) GetPair(token, publicKey string) (*shh.Pair, error) {
	private, err := s.db.Get(publicKey).Result()
	if err != nil {
		return nil, err
	}

	x, err := decrypt([]byte(private), []byte(token))
	if err != nil {
		return nil, err
	}

	return &shh.Pair{
		PrivateKeyPem: x,
		PublicKeyPem:  publicKey,
	}, nil
}


