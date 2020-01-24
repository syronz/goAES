package goaes

import "testing"

func TestGetKey(t *testing.T) {
	samples := []struct {
		secret string
		iv     string
		key    string
		err    error
	}{
		{"one",
			"1234567890123456",
			"7692c3ad3540bb803c020b3aee66cd8887123234ea0c6e7143c0add73ff431ed",
			nil,
		},
	}

	for _, v := range samples {
		result := New(v.secret, v.iv).GetKey()
		if result != v.key {
			t.Errorf("for New(%q, %q).GetKey() result suppose to be %q but it is %q !",
				v.secret, v.iv, v.key, result)
		}

	}

}
