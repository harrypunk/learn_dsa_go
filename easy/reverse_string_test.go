package easy

import "testing"

func TestReverseString(t *testing.T) {
	var s1 = []byte("hello")
	ReverseString(s1)
	var result = []byte{'o', 'l', 'l', 'e', 'h'}
	for i := 0; i < len(s1); i++ {
		if result[i] != s1[i] {
			t.Fail()
		}
	}

}
