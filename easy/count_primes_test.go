package easy

import "testing"

func TestCountPrimes(t *testing.T) {
	n := 10
	expected := 4
	t.Log(CountPrimes(n))
	if CountPrimes(n) != expected {
		t.Fail()
	}

}

func TestCountPrimes2(t *testing.T) {
	n := 5000000
	t.Log(CountPrimes(n))
}

func TestCountPrimes3(t *testing.T) {
	n := 5000000
	t.Log(CountPrimes2(n))
}
