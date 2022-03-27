package easy

import (
	"math"
)

func CountPrimes(n int) int {
	if n < 2 {
		return 0
	}

	count := 0
	for i := 1; i <= n; i++ {
		if IsPrime(i) {
			count += 1
		}
	}

	return count
}

func IsPrime(n int) bool {
	var root2 int = int(math.Sqrt(float64(n)))
	for i := 1; i <= root2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
