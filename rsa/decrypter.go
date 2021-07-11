package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
)

type Decrypter struct {
	PrivateKey *rsa.PrivateKey
}

func (d Decrypter) Decrypt(ciphertext []byte) ([]byte, error) {
	plaintext, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, d.PrivateKey, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
