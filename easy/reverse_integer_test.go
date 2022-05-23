package easy

import "testing"

func TestReverseInteger(t *testing.T) {
	var a = 1234
	var a1 = ReverseInteger(a)
	if a1 != 4321 {
		t.Fail()
	}
}
