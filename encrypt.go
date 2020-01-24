package goaes

import (
	"crypto/aes"
	"encoding/hex"
)

// Encrypt is used for encryption
func (b BuildModel) Encrypt(str string) string {
	ciphertext := make([]byte, aes.BlockSize+len(str))
	b.streamEncrypt.XORKeyStream(ciphertext[aes.BlockSize:], []byte(str))
	return hex.EncodeToString(ciphertext[aes.BlockSize:])
}
