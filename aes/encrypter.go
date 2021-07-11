package aes

import (
	"crypto/aes"
	"crypto/cipher"
)

type Encrypter struct {
	Key   Key
	Nonce []byte
}

func (e Encrypter) EncryptAES(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(e.Key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nil, e.Nonce, plaintext, nil)

	return ciphertext, nil
}
