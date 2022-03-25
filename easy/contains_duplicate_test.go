package easy

import "testing"

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	if !containsDuplicate(nums) {
		t.Fail()
	}
}
