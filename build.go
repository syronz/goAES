package goaes

import (
	"crypto/aes"
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
	return buildModel, a.err
}
