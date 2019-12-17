package shh

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

func (w *Wizard) encrypt(publicKey string, message string) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKey))
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	encryptedMessage, err := rsa.EncryptOAEP(
		w.hash,
		w.random,
		pub,
		[]byte(message),
		[]byte(""),
	)
	if err != nil {
		return nil, err
	}

	return encryptedMessage, nil
}


func (w *Wizard) decryptCipher(privateKey string, cipherText string) ([]byte, error) {
	block, _ := pem.Decode([]byte(privateKey))
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}

	message, err := rsa.DecryptOAEP(
		w.hash,
		w.random,
		priv,
		data,
		[]byte(""),
	)
	if err != nil {
		return nil, err
	}

	return message, nil
}

