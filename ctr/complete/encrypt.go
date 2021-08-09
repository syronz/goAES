package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func Encryptor(input []byte) (output []byte, err error) {
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

	// var gcm int
	var gcm cipher.AEAD
	gcm, err = cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
	}

	// Never use more than 2^32 random nonces with a given key
	// because of the risk of repeat.
	var nonce []byte
	nonce = make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, input, nil)
	fmt.Printf(">>>>>>>>>>>> %T\n", ciphertext)

	output = ciphertext

	return
}

func main() {
	log.Print("File encryption example")

	// plaintext, err := ioutil.ReadFile("plaintext.txt")
	plaintext, err := ioutil.ReadFile("icon.png")
	if err != nil {
		log.Fatal(err)
	}

	ciphertext, err := Encryptor(plaintext)
	if err != nil {
		log.Fatal(err)
	}

	// Save back to file
	// err = ioutil.WriteFile("ciphertext.bin", ciphertext, 0777)
	err = ioutil.WriteFile("icon.bin", ciphertext, 0777)
	if err != nil {
		log.Panic(err)
	}
}
