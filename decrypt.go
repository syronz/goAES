package goaes

import (
	"crypto/cipher"
	"encoding/hex"
)

// Decrypt is used for decrypt an encrypted term
func (a Model) Decrypt(str string, length int) string {
	encryptO, _ := hex.DecodeString(str)
	result := make([]byte, length)
	stream := cipher.NewCTR(a.block, a.iv)
	stream.XORKeyStream(result, encryptO)
	return string(result)

}
