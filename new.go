package goaes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	// "errors"
	"fmt"
	"log"
)

// New create the instance of Model and setup block and iv
func New(secretKey string, ivKey string) Model {
	a := Model{}
	if len(ivKey) != aes.BlockSize {
		a.err = fmt.Errorf("length of iv should be %v", aes.BlockSize)
		return a
	}

	var err error
	key := sha256.Sum256([]byte(secretKey))

	a.key = key[:]

	a.block, err = aes.NewCipher(a.key)
	if err != nil {
		log.Fatal(err.Error())
	}
	// if a.err == nil && err != nil {
	// 	err
	a.err = err

	a.iv = []byte(ivKey)

	a.streamEncrypt = cipher.NewCTR(a.block, a.iv)
	return a
}

// Build return generated model and error if exist
func (a Model) Build() (Model, error) {
	return a, a.err
}
