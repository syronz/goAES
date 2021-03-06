package main

import (
	"fmt"
	"log"

	goaes "github.com/syronz/goAES"
)

func main() {
	aes, err := goaes.New().Key("secret key").IV("iv is 16 char!!!").Build()
	if err != nil {
		log.Fatal(err)
	}

	term := "Hello"
	encrypted := aes.Encrypt(term)

	fmt.Printf("Encryption of %q is %q \n", term, encrypted)
}
