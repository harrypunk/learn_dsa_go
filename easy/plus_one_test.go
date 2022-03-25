package easy

import (
	"testing"
)

func TestHello(t *testing.T) {
	digits := []int{3, 4}
	var result []int = PlusOne(digits)
	expected := []int{3, 5}
	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Fail()
		}
	}
}
