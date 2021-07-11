package main

import (
	"log"
)

func init() {
	initAESEncrypter()
	initRSAEncrypter()
}

func main() {
	if err := Attack(); err != nil {
		log.Fatalln(err)
	}
}
