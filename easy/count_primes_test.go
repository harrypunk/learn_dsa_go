package easy

import "testing"

func TestCountPrimes(t *testing.T) {
	n := 10
	expected := 4
	t.Log(CountPrimes(10))
	if CountPrimes(n) != expected {
		t.Fail()
	}

}
