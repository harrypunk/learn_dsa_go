package easy

func PlusOne(digits []int) []int {
	var len1 int = len(digits)

	t := 1
	for i := len1 - 1; i >= 0; i-- {
		n := digits[i] + t
		t = n / 10
		digits[i] = n % 10
	}

	if t > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
