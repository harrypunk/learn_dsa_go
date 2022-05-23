package easy

import (
	"math"
)

func ReverseInteger(x int) int {
	rev := 0
	var max10 = math.MaxInt32 / 10
	var min10 = math.MinInt32 / 10
	for x != 0 {
		var digit = x % 10
		if rev > max10 || (rev == max10 && digit > 7) {
			return 0
		}
		if rev < min10 || (rev == min10 && digit < -8) {
			return 0
		}
		rev = rev*10 + digit
		x /= 10
	}

	return rev
}
