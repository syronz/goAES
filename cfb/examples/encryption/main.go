package main

import (
	"fmt"
	"log"

	"github.com/syronz/goAES/cfb"
)

func main() {
	aes, err := cfb.New().Key("secret key").IV("iv is 16 char!!!").Build()
	if err != nil {
		log.Fatal(err)
	}

	term := "Hello"
	encrypted := aes.Encrypt(term)

	fmt.Printf("Encryption of %q is %q \n", term, encrypted)
}
