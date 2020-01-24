package goaes

import (
	"testing"
)

func TestNew(t *testing.T) {
	samples := []struct {
		secret string
		iv     string
		err    string
	}{
		{"one", "1234567890123456", ""},
		{"one", "123456789012345", "length of iv should be 16"},
		{"", "1234567890123456", ""},
		{"", "k3dl29d+2k3jf832", ""},
		{"2345678dfghj3249283uyr2840298341;lk098134jj098s7fdq1kj09", "k3dl29d+2k3jf832", ""},
		{"2345678dfghj3249283uyr2840298341;lk098134jj098s7fdq1kj09", "k3dl29d+2k3jf832", ""},
	}

	for _, v := range samples {
		_, err := New(v.secret, v.iv).Build()
		var errStr string
		if err != nil {
			errStr = err.Error()
		}
		if errStr != v.err {
			t.Errorf("for New(%q, %q).Build() err suppose to be %q but it is %q !",
				v.secret, v.iv, v.err, errStr)

		}

	}

}
