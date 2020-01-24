package goaes

import (
	"bytes"
	"crypto/cipher"
	"encoding/hex"
)

// Decrypt is used for decrypt an encrypted term
func (b BuildModel) Decrypt(str string) string {

	encrypted, _ := hex.DecodeString(str)
	bReader := bytes.NewReader(encrypted)

	stream := cipher.NewOFB(b.block, b.iv)

	reader := &cipher.StreamReader{S: stream, R: bReader}

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	return buf.String()

	//------------------------- different method with CTR specific length
	// encryptO, _ := hex.DecodeString(str)
	// length := 5
	// result := make([]byte, length)
	// stream := cipher.NewCTR(b.block, b.iv)
	// stream.XORKeyStream(result, encryptO)
	// return string(result)

}
