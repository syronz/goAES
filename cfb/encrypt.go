package cfb

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

// Encrypt is used for encryption
func (b BuildModel) Encrypt(str string) string {
	b.streamEncrypt = cipher.NewCTR(b.block, b.iv)
	ciphertext := make([]byte, aes.BlockSize+len(str))

	b.streamEncrypt.XORKeyStream(ciphertext[aes.BlockSize:], []byte(str))

	return hex.EncodeToString(ciphertext[aes.BlockSize:])
}
