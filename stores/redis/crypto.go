package redis

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
)

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	k := make([]byte, 32)
	hsh := sha256.Sum256(key)

	copy(k[:], hsh[:])

	c, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(cipherText []byte, key []byte) ([]byte, error) {
	k := make([]byte, 32)
	hsh := sha256.Sum256(key)

	copy(k[:], hsh[:])

	c, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("cipherText too short")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	return gcm.Open(nil, nonce, cipherText, nil)
}
