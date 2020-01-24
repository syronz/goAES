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

	encrypted := "257d85d55b"
	decrypted := aes.Decrypt(encrypted)

	fmt.Printf("Decryption of %q is %q \n", encrypted, decrypted)
}
