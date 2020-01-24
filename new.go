package goaes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"log"
)

// New create the instance of Model and setup block and iv
func New(secretKey string, ivKey string) Model {
	a := Model{}
	key := sha256.Sum256([]byte(secretKey))

	a.key = key[:]

	block, err := aes.NewCipher(key[:])
	if err != nil {
		log.Fatal(err.Error())
	}

	a.block = block

	iv := []byte(ivKey)
	a.iv = iv

	a.streamEncrypt = cipher.NewCTR(block, iv)
	return a
}
