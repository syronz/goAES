package cfb

import "testing"

func TestGetKey(t *testing.T) {
	samples := []struct {
		secret string
		key    string
	}{
		{"one", "7692c3ad3540bb803c020b3aee66cd8887123234ea0c6e7143c0add73ff431ed"},
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"\n", "01ba4719c80b6fe911b091a7c05124b64eeece964e09c058ef8f9805daca546b"},
	}

	for _, v := range samples {
		result := New().Key(v.secret).GetKey()
		if result != v.key {
			t.Errorf("for New().Key(%q).GetKey() result suppose to be %q but it is %q !",
				v.secret, v.key, result)
		}

	}

}
