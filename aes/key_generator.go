package aes

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Key []byte

// Hex returns the key as hexadecimal.
func (k Key) Hex() string {
	return hex.EncodeToString(k)
}

type KeySize int

// Get validates the AES key size and then returns the key size as an integer.
func (k KeySize) Get() (int, error) {
	if k != 16 && k != 32 {
		return 0, fmt.Errorf("expected key size to be 16 or 32, got %d instead", k)
	}

	return int(k), nil
}

func GenerateAESKey(size KeySize) (Key, error) {
	// First, we need to know the key size.
	keySize, err := size.Get()
	if err != nil {
		return nil, err
	}

	// Now we need to generate a pseudorandom key based on the given key size.
	key := make([]byte, keySize)
	_, err = rand.Read(key)
	if err != nil {
		return nil, err
	}

	return Key(key), nil
}
