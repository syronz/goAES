package cfb

import "crypto/cipher"

// Model is the core of builder design pattern
type Model struct {
	key    []byte
	iv     []byte
	err    error
	length int
}

// BuildModel return back finilize and executable struct
type BuildModel struct {
	block         cipher.Block
	streamEncrypt cipher.Stream
	iv            []byte
	length        int
}
