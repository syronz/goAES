package goaes

import (
	"encoding/hex"
)

// GetKey returns back the hashed key in string
func (a Model) GetKey() string {
	return string(hex.EncodeToString(a.key))
}
