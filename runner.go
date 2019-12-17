package shh

import (
	"crypto/sha256"
	"errors"
)

func (w *Wizard) Run() chan error {
	errChan := make(chan error, 1)

	if w.hash == nil {
		w.hash = sha256.New()
	}

	if w.store == nil {
		errChan <- errors.New("please, add a store module, `shh` cannot works without it")
		return errChan
	}

	if w.api != nil {
		go func () {
			errChan <- w.api.run(w)
		}()
	}

	errChan <- w.runRPCService()
	return errChan
}