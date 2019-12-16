package shh

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"
)

type Wizard struct {
	random io.Reader
}

func (w *Wizard) generatePair(bits ...int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	b := 2048
	if len(bits) > 0 {
		b = bits[0]
	}

	key, err := rsa.GenerateKey(w.random, b)
	if err != nil {
		return nil, nil, err
	}

	return key, &key.PublicKey, nil
}

// func (w *Wizard)

func (w *Wizard) CraftNewDefaultPair() (*Pair, error) {
	x, p, err := w.generatePair(2048)
	if err != nil {
		return nil, err
	}

	memPri := pem.EncodeToMemory(&pem.Block{
		Type:    "RSA PRIVATE KEY",
		Bytes:   x509.MarshalPKCS1PrivateKey(x),
	})


	memPub := pem.EncodeToMemory(&pem.Block{
		Type:    "RSA PUBLIC KEY",
		Bytes:   x509.MarshalPKCS1PublicKey(p),
	})

	return &Pair{
		PrivateKeyPem: string(memPri),
		PublicKeyPem:  string(memPub),
	}, nil
}


func NewWizard() *Wizard {
	return &Wizard{random:rand.Reader}
}