package goaes

import (
	"crypto/aes"
	"crypto/cipher"
)

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
