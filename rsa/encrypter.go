package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
)

type Encrypter struct {
	PublicKey *rsa.PublicKey
}

func (e Encrypter) EncryptRSA(plaintext []byte) ([]byte, error) {
	ciphertext, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, e.PublicKey, plaintext, nil)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}
