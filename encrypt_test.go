package goaes

import "testing"

func testDiffrentInputs() {

}

func TestEncrypt(t *testing.T) {
	samples := []struct {
		secret string
		iv     string
		str    string
		result string
		err    error
	}{
		{"one", "1234567890123456", "hello", "7f7775ff73", nil},
	}

	for _, v := range samples {
		encryptedStr := New(v.secret, v.iv).Encrypt(v.str)
		if encryptedStr != v.result {
			t.Errorf("for New(%q, %q).Encrypt(%q) result suppose to be %q but it is %q !",
				v.secret, v.iv, v.str, v.result, encryptedStr)

			t.Log("the key is ", New(v.secret, v.iv).GetKey())
		}

		t.Log(v)
	}

	// aes := New("secret key hashed with sha256", "iv 16 character!")
	// pinString := "Hello"
	// pinEncrypted := aes.Encrypt(pinString)
	// stringKey := aes.GetKey()
	// t.Log("this is test encryption", pinEncrypted, stringKey)
}
