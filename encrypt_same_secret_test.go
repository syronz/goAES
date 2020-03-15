package goaes

import "testing"

func TestEncryptSame(t *testing.T) {
	samples := []struct {
		secret string
		iv     string
		str    string
		result string
		err    string
	}{
		{"secret key", "kd23w[qDn.5+2/Ok", "123456", "17d48e94c980", ""},
		{"secret key", "kd23w[qDn.5+2/Ok", ".", "08", ""},
		{"secret key", "kd23w[qDn.5+2/Ok", "golang", "4189d1c192d1", ""},
		{"secret key", "kd23w[qDn.5+2/Ok", "okaid----------", "498ddcc9989b13d9470762ed808aaa", ""},
		{"secret key", "kd23w[qDn.5+2/Ok", "00000000", "16d68d90cc860ec4", ""},
		{"secret key", "kd23w[qDn.5+2/Ok", "+", "0d", ""},
		{"secret key", "kd23w[qDn.5+2/Ok", "\n\t\t\n", "2cefb4aa", ""},
		{"secret key", "kd23w[qDn.5+2/Ok", "9872342", "1fde8a92cf820c", ""},
	}

	aes, err := New().Key(samples[0].secret).IV(samples[0].iv).Build()
	for _, v := range samples {
		var errStr string
		switch err {
		case nil:
			encryptedStr := aes.Encrypt(v.str)
			if encryptedStr != v.result {
				t.Errorf("for New(%q, %q).Encrypt(%q) result suppose to be %q but it is %q !",
					v.secret, v.iv, v.str, v.result, encryptedStr)

			}

		default:
			errStr = err.Error()
			if errStr != v.err {
				t.Fatalf("for New().Key(%q).IV(%q) then aes.Encrypt(%q) the error suppose to be %q but it is %q",
					v.secret, v.iv, v.str, v.err, errStr)

			}

		} //switch

	}

}
