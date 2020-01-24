package goaes

import "testing"

func TestDecrypt(t *testing.T) {
	samples := []struct {
		secret    string
		iv        string
		result    string
		encrypted string
		err       string
	}{
		{"one", "1234567890123456", "hello", "7f7775ff73", ""},
		{"one", "1234567890123456", "hell", "7f7775ff", ""},
		{"one", "123456789012345", "hello", "7f7775ff73", "length of iv should be 16"},
		{"one", "1234567890123456", "", "", ""},
		{"", "1234567890123456", "hello", "f6632cc3c4", ""},
		{"", "1234567890123456", "\n", "94", ""},
		{"", "1234567890123456", "   ", "be2660", ""},
		{"", "1234567890123456", "3.14159265359", "ad28719b9a22a6fd384f118682", ""},
		{"", "1234567890123456", "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.", "d26932cac637d6bf7d0f4f93d2be919ad2de59d7cb23a4140add0a89e7864d4a42abc1f25c51fa4af226dfe97c9ecbed4973d63eb037fab21fae3074c17c162e06293bf515dfc689361a5ae85066b37de38012e280c317bcbc788ceeb4fdbdebea856e04bc270d4139e5341c731f881425382dd473cf4ea0de59a3eb1a89b4b96513839825bb001b231c107272027cfc0859a3936006f6bf617ff2679dd90329595e3b3a8925c5779d0d7c7415cde241a81f804d8394fbd1908fe0fba410c00202d26ebbb2adbc23f2fd4919c2b005d88b3a4ac0bd536067ff4a65a75038404ec18407657036ef4d2a254806b5087e3179f4ffb70caea937c45705077377cb00e169c89546f2d16274e05f38f4ea71441242b9f4663b532535456edc713c0910aa4f606c08e13c85f57557f7786fcabb3e41a368f5b2dd7831166af9e424982477cc2f856bd2c46910469078feaa5d5ead9b21ecc4d2291fac4011b49ba511e432e3658b345318b3620f0c68c5899d3f9753b0c237a10b1e424c84040f4dd52884d77c38e2dc146135fc291ca7f4947fb5386aead48bbea56b177f8b5a13e07f2d300b025a8b6b2e0ce422c7e855d7e17d7d7a111031720a1044f0960c5792aff69c4de9cadae968bce6077a3add127bd0ebb97a2c6f9720fcf6479e1c1521caddf9ce999dd03bed29c8e1241b0b3138a308c38f53d6c64083d3cb19df55be7bead0afb7cb0835fc55af38f611e5f5b72164fb4f273d52b5c031d4ba299fae675f4c399f834a380768db5be22371134643383e42492fbeced8c904a92a1f", ""},
		// {"secret key", "kd23w[qDn.5+2/Ok", "123456", "17d48e94c980", ""},
		// {"secret key", "kd23w[qDn.5+2/Ok", ".", "08", ""},
		// {"secret key", "kd23w[qDn.5+2/Ok", "golang", "4189d1c192d1", ""},
		// {"secret key", "kd23w[qDn.5+2/Ok", "okaid----------", "498ddcc9989b13d9470762ed808aaa", ""},
		// {"secret key", "kd23w[qDn.5+2/Ok", "00000000", "16d68d90cc860ec4", ""},
		// {"secret key", "kd23w[qDn.5+2/Ok", "+", "0d", ""},
		// {"secret key", "kd23w[qDn.5+2/Ok", "\n\t\t\n", "2cefb4aa", ""},
		// {"secret key", "kd23w[qDn.5+2/Ok", "9872342", "1fde8a92cf820c", ""},
	}

	for _, v := range samples {
		aes, err := New().Key(v.secret).IV(v.iv).Build()
		var errStr string
		switch err {
		case nil:
			encryptedStr := aes.Decrypt(v.encrypted)
			if encryptedStr != v.result {
				t.Errorf("for New(%q, %q).Encrypt(%q) result suppose to be %q but it is %q !",
					v.secret, v.iv, v.encrypted, v.result, encryptedStr)

			}

		default:
			errStr = err.Error()
			if errStr != v.err {
				t.Fatalf("for New().Key(%q).IV(%q) then aes.Encrypt(%q) the error suppose to be %q but it is %q",
					v.secret, v.iv, v.encrypted, v.err, errStr)

			}

		} //switch

	}

}
