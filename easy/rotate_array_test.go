package easy

import "testing"

func TestRotate(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}
	Rotate(nums, 3)
	exp := []int{5, 6, 7, 1, 2, 3, 4}
	for i := 0; i < len(nums); i++ {
		if exp[i] != nums[i] {
			t.Fail()
		}
	}
}

func TestRotateEx2(t *testing.T) {
	var nums = []int{-1, -100, 3, 99}
	Rotate(nums, 2)
	exp := []int{3, 99, -1, -100}
	for i := 0; i < len(nums); i++ {
		if exp[i] != nums[i] {
			t.Fail()
		}
	}
}
