package shh

import "io"

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