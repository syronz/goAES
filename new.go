package goaes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
)

// New create the instance of Model and setup block and iv
func New() Model {
	a := Model{
		length: 100,
	}

	return a
}

// Key is used for getting a secret-key and hash it with sha256
func (a Model) Key(secretKey string) Model {
	key := sha256.Sum256([]byte(secretKey))
	a.key = key[:]

	return a
}

// IV get a string and use it as iv
func (a Model) IV(iv string) Model {
	if len(iv) != aes.BlockSize {
		a.err = fmt.Errorf("length of iv should be %v", aes.BlockSize)
		return a
	}

	a.iv = []byte(iv)

	return a
}

// Length is used in case the result is more than 100 character
func (a Model) Length(length int) Model {
	a.length = length
	return a
}

// Build return generated model and error if exist
func (a Model) Build() (BuildModel, error) {
	buildModel := BuildModel{
		iv:     a.iv,
		length: a.length,
	}

	if a.err != nil {
		return buildModel, a.err
	}
	buildModel.block, _ = aes.NewCipher(a.key)
	buildModel.streamEncrypt = cipher.NewCTR(buildModel.block, a.iv)
	return buildModel, a.err
}
