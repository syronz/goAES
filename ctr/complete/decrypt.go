package main

import (
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"log"
)

func Decryptor(input []byte) (output []byte, err error) {
	// The key should be 16 bytes (AES-128), 24 bytes (AES-192) or
	// 32 bytes (AES-256)
	key, err := ioutil.ReadFile("key")
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
	}

	nonce := input[:gcm.NonceSize()]
	input = input[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, input, nil)
	if err != nil {
		log.Panic(err)
	}

	output = plaintext
	return

}

func main() {
	// ciphertext, err := ioutil.ReadFile("ciphertext.bin")
	ciphertext, err := ioutil.ReadFile("icon.bin")
	if err != nil {
		log.Fatal(err)
	}

	plaintext, err := Decryptor(ciphertext)

	// err = ioutil.WriteFile("plaintext2.txt", plaintext, 0777)
	err = ioutil.WriteFile("icon2.png", plaintext, 0777)
	if err != nil {
		log.Panic(err)
	}
}
