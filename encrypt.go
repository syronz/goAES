package goaes

import (
	"crypto/aes"
	"encoding/hex"
)

// Encrypt is used for encryption
func (a Model) Encrypt(str string) string {
	ciphertext := make([]byte, aes.BlockSize+len(str))
	a.streamEncrypt.XORKeyStream(ciphertext[aes.BlockSize:], []byte(str))
	return hex.EncodeToString(ciphertext[aes.BlockSize:])
}
