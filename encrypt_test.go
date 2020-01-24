package goaes

import "testing"

func TestEncrypt(t *testing.T) {
	aes, err := New("secret key hashed with sha256", "iv 16 character!")
	pinString := "Hello"
	pinEncrypted := aes.Encrypt(pinString)
	stringKey := aes.GetKey()
	t.Log("this is test encryption", pinEncrypted, stringKey)
}
