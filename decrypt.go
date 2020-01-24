package goaes

import (
	"crypto/cipher"
	"encoding/hex"
)

// Decrypt is used for decrypt an encrypted term
func (b BuildModel) Decrypt(str string) string {

	// // TODO: should fix auto read from stream
	// encrypted, _ := hex.DecodeString(str)
	// bReader := bytes.NewReader(encrypted)
	// stream := cipher.NewOFB(b.block, b.iv)
	// reader := &cipher.StreamReader{S: stream, R: bReader}

	// buf := new(bytes.Buffer)
	// buf.ReadFrom(reader)
	// return buf.String()

	//------------------------- different method with CTR specific length
	encrypt, _ := hex.DecodeString(str)
	result := make([]byte, b.length)
	stream := cipher.NewCTR(b.block, b.iv)
	stream.XORKeyStream(result, encrypt)

	result = trimResult(result)
	return string(result)

}

func trimResult(data []byte) []byte {

	k := 0
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] == 0 {
			k++
		}
	}

	data = data[0 : len(data)-k]
	// fmt.Println(">>>>>>>>>>>>>....................", k, data)

	return data

}
