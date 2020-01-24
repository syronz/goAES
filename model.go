package goaes

import "crypto/cipher"

// Model is the core of builder design pattern
type Model struct {
	key           []byte
	iv            []byte
	block         cipher.Block
	streamEncrypt cipher.Stream
	// stream *cipher.ctr
}
