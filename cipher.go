package shh

import (
	"crypto/rsa"
	"crypto/x509"
)

func (w *Wizard) encrypt(publicKey string, message string) (string, error) {
	pub, err := x509.ParsePKCS1PublicKey([]byte(publicKey))
	if err != nil {
		return "", err
	}

	encryptedMessage, err := rsa.EncryptOAEP(
		w.hash,
		w.random,
		pub,
		[]byte(message),
		[]byte(""),
	)
	if err != nil {
		return "", err
	}

	return string(encryptedMessage), nil
}


func (w *Wizard) decryptCipher(privateKey string, cipherText string) (string, error) {
	priv, err := x509.ParsePKCS1PrivateKey([]byte(privateKey))
	if err != nil {
		return "", err
	}

	message, err := rsa.DecryptOAEP(
		w.hash,
		w.random,
		priv,
		[]byte(cipherText),
		[]byte(""),
	)
	if err != nil {
		return "", err
	}

	return string(message), nil
}

