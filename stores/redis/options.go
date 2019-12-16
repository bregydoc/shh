package redis

import "time"

type Option func (*Store) error

func WithAddress(address string) Option {
	return func (s *Store) error {
		s.address = address
		return nil
	}
}


func WithExpirationTime(duration time.Duration) Option {
	return func (s *Store) error {
		s.pairExpiration = duration
		return nil
	}
}