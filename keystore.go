package shh

import "errors"

func (w *Wizard) verifyToken(token string) error {
	if token == "public:access" {
		return nil
	}

	return errors.New("invalid token")
}
